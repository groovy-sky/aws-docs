package crawl

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/groovy-sky/aws-docs/internal/config"
	"github.com/groovy-sky/aws-docs/internal/model"
	"github.com/groovy-sky/aws-docs/internal/store"
	"github.com/groovy-sky/aws-docs/internal/write"
)

type Crawler struct {
	config    config.Config
	store     *store.Store
	fetcher   *Fetcher
	extractor *Extractor
	converter *Converter
	mapper    *Mapper
	writer    *write.FileWriter
	robots    *Robots
}

func NewCrawler(cfg config.Config, db *store.Store, fetcher *Fetcher, extractor *Extractor, converter *Converter, mapper *Mapper, writer *write.FileWriter, robots *Robots) *Crawler {
	return &Crawler{
		config:    cfg,
		store:     db,
		fetcher:   fetcher,
		extractor: extractor,
		converter: converter,
		mapper:    mapper,
		writer:    writer,
		robots:    robots,
	}
}

func (c *Crawler) Run(ctx context.Context, options RunOptions) error {
	queue, err := c.seedQueue(ctx, options)
	if err != nil {
		return err
	}

	sectionLimitedMode := options.Mode == "incremental" || options.Mode == "partial"
	allowedSections := map[string]struct{}{}
	if sectionLimitedMode && options.MaxSections > 0 {
		for _, item := range queue {
			section := sectionKey(item)
			if _, exists := allowedSections[section]; exists {
				continue
			}
			allowedSections[section] = struct{}{}
			if len(allowedSections) >= options.MaxSections {
				break
			}
		}
	}

	seen := make(map[string]struct{}, len(queue))
	filteredQueue := make([]string, 0, len(queue))
	for _, item := range queue {
		if len(allowedSections) > 0 {
			if _, ok := allowedSections[sectionKey(item)]; !ok {
				continue
			}
		}
		seen[item] = struct{}{}
		filteredQueue = append(filteredQueue, item)
	}
	queue = filteredQueue

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if err := c.processURL(ctx, current, options.Mode); err != nil {
			log.Printf("skip %s: %v", current, err)
		}

		if options.Mode == "refresh-url" {
			continue
		}

		record, err := c.store.GetPage(current)
		if err != nil || record == nil || record.LastError != "" {
			continue
		}

		links, err := c.store.GetLinks(current)
		if err != nil {
			continue
		}
		for _, link := range links {
			if _, exists := seen[link]; exists {
				continue
			}
			if !IsAllowedURL(link, c.config) || !c.allowedByRobots(link) {
				continue
			}
			if len(allowedSections) > 0 {
				if _, ok := allowedSections[sectionKey(link)]; !ok {
					continue
				}
			}
			seen[link] = struct{}{}
			queue = append(queue, link)
		}
	}

	return nil
}

func (c *Crawler) seedQueue(ctx context.Context, options RunOptions) ([]string, error) {
	if options.Mode == "refresh-url" {
		if strings.TrimSpace(options.URL) == "" {
			return nil, fmt.Errorf("refresh-url mode requires -url")
		}
		normalized, err := NormalizeURL(options.URL, c.config)
		if err != nil {
			return nil, fmt.Errorf("normalize refresh URL: %w", err)
		}
		return []string{normalized}, nil
	}

	storedSeeds, err := c.store.GetSeeds()
	if err != nil {
		return nil, err
	}
	if len(storedSeeds) > 0 {
		normalizedSeeds := make([]string, 0, len(storedSeeds))
		seenNormalized := map[string]struct{}{}
		changed := false
		for _, seed := range storedSeeds {
			normalized, normalizeErr := NormalizeURL(seed, c.config)
			if normalizeErr != nil {
				changed = true
				continue
			}
			if normalized != seed {
				changed = true
			}
			if _, exists := seenNormalized[normalized]; exists {
				changed = true
				continue
			}
			seenNormalized[normalized] = struct{}{}
			normalizedSeeds = append(normalizedSeeds, normalized)
		}
		if changed {
			if err := c.store.PutSeeds(normalizedSeeds); err != nil {
				return nil, err
			}
			storedSeeds = normalizedSeeds
		}
	}
	if len(storedSeeds) == 0 {
		storedSeeds, err = c.discoverAndPersistSeeds(ctx)
		if err != nil {
			return nil, err
		}
	}

	queue := make([]string, 0, len(storedSeeds))
	seen := map[string]struct{}{}
	for _, seed := range storedSeeds {
		normalized, err := NormalizeURL(seed, c.config)
		if err != nil {
			return nil, fmt.Errorf("normalize seed %q: %w", seed, err)
		}
		if IsAllowedURL(normalized, c.config) {
			if _, exists := seen[normalized]; !exists {
				seen[normalized] = struct{}{}
				queue = append(queue, normalized)
			}
		}
	}

	if len(queue) == 0 {
		for _, seed := range c.config.Seeds {
			normalized, err := NormalizeURL(seed, c.config)
			if err != nil {
				return nil, fmt.Errorf("normalize fallback seed %q: %w", seed, err)
			}
			if _, exists := seen[normalized]; exists || !IsAllowedURL(normalized, c.config) {
				continue
			}
			seen[normalized] = struct{}{}
			queue = append(queue, normalized)
		}
	}

	if options.Mode == "full" && c.config.EnableRobotsSitemaps {
		sitemapURLs, err := DiscoverURLsFromRobotsSitemaps(ctx, c.fetcher.HTTPClient(), c.config)
		if err != nil {
			log.Printf("sitemap discovery failed: %v", err)
		} else {
			for _, discoveredURL := range sitemapURLs {
				if _, exists := seen[discoveredURL]; exists {
					continue
				}
				seen[discoveredURL] = struct{}{}
				queue = append(queue, discoveredURL)
			}

			log.Printf("seeded %d sitemap URLs for full crawl", len(sitemapURLs))
		}
	}

	return queue, nil
}

func (c *Crawler) discoverAndPersistSeeds(ctx context.Context) ([]string, error) {
	seeds := make([]string, 0)
	seen := map[string]struct{}{}

	addSeed := func(rawURL string) {
		normalized, err := NormalizeURL(rawURL, c.config)
		if err != nil || !IsAllowedURL(normalized, c.config) {
			return
		}
		if _, exists := seen[normalized]; exists {
			return
		}
		seen[normalized] = struct{}{}
		seeds = append(seeds, normalized)
	}

	for _, host := range c.config.AllowedHosts {
		addSeed("https://" + host + "/")
	}

	if c.config.EnableRobotsStructure {
		for _, host := range c.config.AllowedHosts {
			structure, err := DiscoverRobotsStructure(ctx, c.fetcher.HTTPClient(), host, c.config)
			if err != nil {
				log.Printf("robots structure discovery failed for %s: %v", host, err)
				continue
			}
			for _, rootURL := range structure.SectionRoots {
				addSeed(rootURL)
			}
		}
	}

	if c.config.EnableRobotsSitemaps {
		sitemapURLs, err := DiscoverURLsFromRobotsSitemaps(ctx, c.fetcher.HTTPClient(), c.config)
		if err != nil {
			log.Printf("sitemap discovery failed: %v", err)
		} else {
			for _, rootURL := range DeriveServiceRootSeeds(sitemapURLs, c.config) {
				addSeed(rootURL)
			}
		}
	}

	if len(seeds) == 0 {
		for _, seed := range c.config.Seeds {
			addSeed(seed)
		}
	}

	if err := c.store.PutSeeds(seeds); err != nil {
		return nil, err
	}
	log.Printf("stored %d robots-derived seeds in metadata db", len(seeds))
	return seeds, nil
}

func (c *Crawler) processURL(ctx context.Context, pageURL string, mode string) error {
	if !c.allowedByRobots(pageURL) {
		return fmt.Errorf("blocked by robots.txt")
	}

	previous, err := c.store.GetPage(pageURL)
	if err != nil {
		return fmt.Errorf("read metadata: %w", err)
	}
	expectedPath := c.mapper.RepoPath(pageURL)
	if mode == "full" {
		previous = nil
	} else if mode == "refresh-url" {
		// Refresh should always re-fetch and regenerate content for the requested URL.
		previous = nil
	} else if previous != nil && (previous.RepoPath != expectedPath || !c.writer.Exists(expectedPath)) {
		previous = nil
	}

	result, err := c.fetcher.Fetch(ctx, pageURL, previous)
	if err != nil {
		return c.store.PutPage(model.PageRecord{
			URL:         pageURL,
			RepoPath:    c.mapper.RepoPath(pageURL),
			LastFetched: time.Now().UTC(),
			StatusCode:  result.StatusCode,
			LastError:   err.Error(),
		})
	}

	if result.NotModified && previous != nil {
		previous.LastFetched = time.Now().UTC()
		previous.StatusCode = result.StatusCode
		previous.LastError = ""
		return c.store.PutPage(*previous)
	}

	extracted, err := c.extractor.Extract(result.FinalURL, result.Body)
	if err != nil {
		return c.store.PutPage(model.PageRecord{
			URL:          pageURL,
			RepoPath:     c.mapper.RepoPath(pageURL),
			ETag:         result.ETag,
			LastModified: result.LastModified,
			LastFetched:  time.Now().UTC(),
			StatusCode:   result.StatusCode,
			LastError:    err.Error(),
		})
	}

	canonicalURL, err := NormalizeURL(extracted.CanonicalURL, c.config)
	if err != nil {
		canonicalURL = pageURL
	}
	markdownDocument, err := c.converter.Convert(extracted, canonicalURL)
	if err != nil {
		return fmt.Errorf("convert %s: %w", canonicalURL, err)
	}

	repoPath := c.mapper.RepoPath(canonicalURL)
	if previous == nil || previous.ContentHash != markdownDocument.ContentHash || !c.writer.Exists(repoPath) {
		if err := c.writer.Write(repoPath, markdownDocument.Markdown); err != nil {
			return fmt.Errorf("write %s: %w", repoPath, err)
		}
	}

	discovered := make([]string, 0, len(extracted.Links))
	for _, link := range extracted.Links {
		resolved, err := ResolveURL(canonicalURL, link, c.config)
		if err != nil {
			continue
		}
		if IsAllowedURL(resolved, c.config) {
			discovered = append(discovered, resolved)
		}
	}

	if err := c.store.PutLinks(canonicalURL, discovered); err != nil {
		return err
	}
	if canonicalURL != pageURL {
		if err := c.store.PutLinks(pageURL, discovered); err != nil {
			return err
		}
	}

	record := model.PageRecord{
		URL:          canonicalURL,
		RepoPath:     repoPath,
		ETag:         result.ETag,
		LastModified: result.LastModified,
		ContentHash:  markdownDocument.ContentHash,
		LastFetched:  time.Now().UTC(),
		StatusCode:   result.StatusCode,
	}
	if err := c.store.PutPage(record); err != nil {
		return err
	}

	if canonicalURL != pageURL {
		record.URL = pageURL
		if err := c.store.PutPage(record); err != nil {
			return err
		}
	}

	return nil
}

func (c *Crawler) allowedByRobots(rawURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return c.robots == nil || c.robots.Allowed(parsed.Path)
}

func sectionKey(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	host := strings.ToLower(parsed.Host)
	trimmedPath := strings.Trim(parsed.EscapedPath(), "/")
	if trimmedPath == "" {
		return host + "/"
	}
	firstSegment := strings.SplitN(trimmedPath, "/", 2)[0]
	return host + "/" + firstSegment
}
