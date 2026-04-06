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

	seen := make(map[string]struct{}, len(queue))
	for _, item := range queue {
		seen[item] = struct{}{}
	}

	processed := 0
	for len(queue) > 0 {
		if options.MaxPages > 0 && processed >= options.MaxPages {
			break
		}

		current := queue[0]
		queue = queue[1:]
		if err := c.processURL(ctx, current, options.Mode); err != nil {
			log.Printf("skip %s: %v", current, err)
		}
		processed++

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

	queue := make([]string, 0, len(c.config.Seeds))
	seen := map[string]struct{}{}
	for _, seed := range c.config.Seeds {
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

	if c.config.EnableRobotsSitemaps {
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

			serviceRootSeeds := DeriveServiceRootSeeds(sitemapURLs, c.config)
			for _, rootURL := range serviceRootSeeds {
				if _, exists := seen[rootURL]; exists {
					continue
				}
				seen[rootURL] = struct{}{}
				queue = append(queue, rootURL)
			}

			log.Printf("seeded %d URLs and %d service roots from robots sitemaps", len(sitemapURLs), len(serviceRootSeeds))
		}
	}

	return queue, nil
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

	return c.store.PutPage(model.PageRecord{
		URL:          canonicalURL,
		RepoPath:     repoPath,
		ETag:         result.ETag,
		LastModified: result.LastModified,
		ContentHash:  markdownDocument.ContentHash,
		LastFetched:  time.Now().UTC(),
		StatusCode:   result.StatusCode,
	})
}

func (c *Crawler) allowedByRobots(rawURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return c.robots == nil || c.robots.Allowed(parsed.Path)
}
