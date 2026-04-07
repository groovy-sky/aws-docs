package crawl

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/groovy-sky/aws-docs/internal/config"
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
		discovered, err := c.processURL(ctx, current)
		if err != nil {
			log.Printf("skip %s: %v", current, err)
			continue
		}

		if options.Mode == "refresh-url" {
			continue
		}

		for _, link := range discovered {
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
			normalized, normalizeErr := canonicalizeSeedURL(seed, c.config)
			if normalizeErr != nil {
				changed = true
				continue
			}
			if normalized == "" {
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

		filteredSeeds := c.filterReachableSeeds(ctx, storedSeeds)
		if len(filteredSeeds) != len(storedSeeds) {
			if err := c.store.PutSeeds(filteredSeeds); err != nil {
				return nil, err
			}
			storedSeeds = filteredSeeds
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
		normalized, err := canonicalizeSeedURL(seed, c.config)
		if err != nil {
			return nil, fmt.Errorf("normalize seed %q: %w", seed, err)
		}
		if normalized == "" {
			continue
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
			normalized, err := canonicalizeSeedURL(seed, c.config)
			if err != nil {
				return nil, fmt.Errorf("normalize fallback seed %q: %w", seed, err)
			}
			if normalized == "" {
				continue
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
		normalized, err := canonicalizeSeedURL(rawURL, c.config)
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

	seeds = c.filterReachableSeeds(ctx, seeds)

	if err := c.store.PutSeeds(seeds); err != nil {
		return nil, err
	}
	log.Printf("stored %d robots-derived seeds in metadata db", len(seeds))
	return seeds, nil
}

func (c *Crawler) processURL(ctx context.Context, pageURL string) ([]string, error) {
	if !c.allowedByRobots(pageURL) {
		return nil, fmt.Errorf("blocked by robots.txt")
	}

	result, err := c.fetcher.Fetch(ctx, pageURL)
	if err != nil {
		return nil, err
	}

	extracted, err := c.extractor.Extract(result.FinalURL, result.Body)
	if err != nil {
		return nil, err
	}

	canonicalURL, err := NormalizeURL(extracted.CanonicalURL, c.config)
	if err != nil {
		canonicalURL = pageURL
	}
	markdownDocument, err := c.converter.Convert(extracted, canonicalURL)
	if err != nil {
		return nil, fmt.Errorf("convert %s: %w", canonicalURL, err)
	}

	repoPath := c.mapper.RepoPath(canonicalURL)
	if err := c.writer.Write(repoPath, markdownDocument.Markdown); err != nil {
		return nil, fmt.Errorf("write %s: %w", repoPath, err)
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

	return discovered, nil
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

func canonicalizeSeedURL(rawURL string, cfg config.Config) (string, error) {
	normalized, err := NormalizeURL(rawURL, cfg)
	if err != nil {
		return "", err
	}

	parsed, err := url.Parse(normalized)
	if err != nil {
		return "", err
	}

	trimmed := strings.Trim(parsed.Path, "/")
	if trimmed == "" {
		return normalized, nil
	}
	segments := strings.Split(trimmed, "/")
	if len(segments) == 1 {
		return "", nil
	}
	if strings.Contains(segments[0], ".") {
		return normalized, nil
	}

	seedPath := deriveSectionSeedPath(parsed.Path)
	if seedPath == "" {
		return "", nil
	}
	parsed.Path = seedPath
	parsed.RawPath = ""
	parsed.RawQuery = ""
	parsed.Fragment = ""

	return NormalizeURL(parsed.String(), cfg)
}

func (c *Crawler) filterReachableSeeds(ctx context.Context, seeds []string) []string {
	filtered := make([]string, 0, len(seeds))
	for _, seed := range seeds {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, seed, nil)
		if err != nil {
			continue
		}
		request.Header.Set("User-Agent", c.config.UserAgent)
		request.Header.Set("Accept", "text/html,application/xhtml+xml")

		response, err := c.fetcher.HTTPClient().Do(request)
		if err != nil {
			continue
		}
		_, _ = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()

		if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusBadRequest {
			filtered = append(filtered, seed)
		}
	}
	return filtered
}
