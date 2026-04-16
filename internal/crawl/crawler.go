package crawl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/groovy-sky/aws-docs/internal/config"
	"github.com/groovy-sky/aws-docs/internal/model"
	"github.com/groovy-sky/aws-docs/internal/store"
	"github.com/groovy-sky/aws-docs/internal/write"
)

const searchIndexPath = "static/search-index.json"

type searchIndexEntry struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
	Section string `json:"section"`
}

type Crawler struct {
	config    config.Config
	store     *store.Store
	fetcher   *Fetcher
	extractor *Extractor
	converter *Converter
	mapper    *Mapper
	writer    *write.FileWriter
	robots    *Robots
	search    map[string]searchIndexEntry
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
		search:    make(map[string]searchIndexEntry),
	}
}

func (c *Crawler) Run(ctx context.Context, options RunOptions) error {
	if err := c.loadSearchIndex(); err != nil {
		log.Printf("search index load: %v", err)
	}

	queue, err := c.seedQueue(ctx, options)
	if err != nil {
		return err
	}

	sectionLimitedMode := options.Mode == "incremental" || options.Mode == "partial"
	allowedSections := map[string]struct{}{}
	selectedSectionName := ""
	if sectionLimitedMode && options.MaxSections > 0 {
		cursor, err := c.store.GetSectionCursor()
		if err != nil {
			return err
		}

		var nextCursor int
		allowedSections, nextCursor = selectSectionsForRun(queue, options.MaxSections, cursor)
		if err := c.store.PutSectionCursor(nextCursor); err != nil {
			return err
		}
	}

	if options.Mode == "section" {
		sectionName, err := normalizeSectionName(options.Name)
		if err != nil {
			return err
		}
		selectedSectionName = sectionName
		allowedSections = selectSectionsByName(queue, sectionName)
		if len(allowedSections) == 0 {
			return fmt.Errorf("section mode: no seed URLs found for section %q", sectionName)
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
	if options.Mode == "section" {
		log.Printf("crawl run: mode=%s section=%s selected-sections=%d seed-urls=%d", options.Mode, selectedSectionName, len(allowedSections), len(queue))
	} else if len(allowedSections) > 0 {
		log.Printf("crawl run: mode=%s max-sections=%d selected-sections=%d seed-urls=%d", options.Mode, options.MaxSections, len(allowedSections), len(queue))
	} else {
		log.Printf("crawl run: mode=%s seed-urls=%d", options.Mode, len(queue))
	}

	processedCount := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		c.logf("crawl dequeue: url=%s remaining=%d", current, len(queue))
		discovered, err := c.processURL(ctx, current)
		if err != nil {
			log.Printf("skip %s: %v", current, err)
			continue
		}
		processedCount++
		c.logf("crawl processed: url=%s discovered=%d", current, len(discovered))

		if options.Mode == "refresh-url" {
			continue
		}

		for _, link := range discovered {
			if _, exists := seen[link]; exists {
				c.logf("crawl skip-seen: url=%s", link)
				continue
			}
			if !IsAllowedURL(link, c.config) || !c.allowedByRobots(link) {
				c.logf("crawl skip-filtered: url=%s", link)
				continue
			}
			if len(allowedSections) > 0 {
				if _, ok := allowedSections[sectionKey(link)]; !ok {
					continue
				}
			}
			seen[link] = struct{}{}
			queue = append(queue, link)
			c.logf("crawl enqueue: url=%s queue-size=%d", link, len(queue))
		}
	}

	log.Printf("crawl run completed: processed-urls=%d", processedCount)

	if options.Mode != "refresh-url" {
		c.persistDiscoveredSeeds(seen)
	}

	if err := c.saveSearchIndex(); err != nil {
		return fmt.Errorf("write search index: %w", err)
	}

	return nil
}

// persistDiscoveredSeeds canonicalizes every URL that was seen during the BFS
// run to its section root and stores any roots not already in the seed store.
// This ensures that sections discovered via link traversal (e.g. reachability
// found in a VPC landing page) survive across runs and are not silently dropped
// when the current run ends before crawling them.
func (c *Crawler) persistDiscoveredSeeds(seen map[string]struct{}) {
	if c.store == nil || len(seen) == 0 {
		return
	}

	storedSeeds, err := c.store.GetSeeds()
	if err != nil {
		return
	}

	storedSet := make(map[string]struct{}, len(storedSeeds))
	for _, s := range storedSeeds {
		storedSet[s] = struct{}{}
	}

	var newSeeds []string
	for rawURL := range seen {
		canonical, err := canonicalizeSeedURL(rawURL, c.config)
		if err != nil || canonical == "" {
			continue
		}
		if _, exists := storedSet[canonical]; !exists {
			storedSet[canonical] = struct{}{}
			newSeeds = append(newSeeds, canonical)
		}
	}

	if len(newSeeds) > 0 {
		if err := c.store.PutSeeds(append(storedSeeds, newSeeds...)); err != nil {
			log.Printf("persist discovered seeds: %v", err)
			return
		}
		log.Printf("crawl seeds: persisted %d newly discovered section roots", len(newSeeds))
	}
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

	for _, seed := range c.discoverRobotsDerivedSeeds(ctx) {
		if _, exists := seen[seed]; exists {
			continue
		}
		seen[seed] = struct{}{}
		queue = append(queue, seed)
	}

	if (options.Mode == "incremental" || options.Mode == "partial") && options.MaxSections > 0 {
		if uniqueSectionsCount(queue) < options.MaxSections {
			for _, seed := range c.config.Seeds {
				normalized, err := canonicalizeSeedURL(seed, c.config)
				if err != nil || normalized == "" {
					continue
				}
				if _, exists := seen[normalized]; exists || !IsAllowedURL(normalized, c.config) {
					continue
				}
				seen[normalized] = struct{}{}
				queue = append(queue, normalized)
				if uniqueSectionsCount(queue) >= options.MaxSections {
					break
				}
			}
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

	if options.Mode == "section" {
		sectionName, err := normalizeSectionName(options.Name)
		if err != nil {
			return nil, err
		}

		for _, host := range c.config.AllowedHosts {
			candidate := "https://" + strings.TrimSpace(host) + "/" + sectionName + "/latest/"
			normalized, err := canonicalizeSeedURL(candidate, c.config)
			if err != nil || normalized == "" {
				continue
			}
			if _, exists := seen[normalized]; exists || !IsAllowedURL(normalized, c.config) || !c.allowedByRobots(normalized) {
				continue
			}
			seen[normalized] = struct{}{}
			queue = append(queue, normalized)
		}
	}

	return queue, nil
}

func normalizeSectionName(value string) (string, error) {
	name := strings.Trim(strings.TrimSpace(value), "/")
	if name == "" {
		return "", fmt.Errorf("section mode requires -name")
	}
	if strings.Contains(name, "/") {
		return "", fmt.Errorf("section mode requires a top-level section name, got %q", value)
	}
	return name, nil
}

func selectSectionsByName(queue []string, sectionName string) map[string]struct{} {
	selected := map[string]struct{}{}
	for _, item := range queue {
		if !sectionNameMatches(item, sectionName) {
			continue
		}
		key := sectionKey(item)
		if key == "" {
			continue
		}
		selected[key] = struct{}{}
	}
	return selected
}

func sectionNameMatches(rawURL string, sectionName string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	trimmedPath := strings.Trim(parsed.EscapedPath(), "/")
	if trimmedPath == "" {
		return false
	}
	firstSegment := strings.SplitN(trimmedPath, "/", 2)[0]
	return strings.EqualFold(firstSegment, sectionName)
}

func uniqueSectionsCount(urls []string) int {
	seen := map[string]struct{}{}
	for _, item := range urls {
		key := sectionKey(item)
		if key == "" {
			continue
		}
		seen[key] = struct{}{}
	}
	return len(seen)
}

func (c *Crawler) discoverAndPersistSeeds(ctx context.Context) ([]string, error) {
	seeds := c.discoverRobotsDerivedSeeds(ctx)
	seen := map[string]struct{}{}
	for _, seed := range seeds {
		seen[seed] = struct{}{}
	}

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

func (c *Crawler) discoverRobotsDerivedSeeds(ctx context.Context) []string {
	seeds := make([]string, 0)
	seen := map[string]struct{}{}

	addSeed := func(rawURL string) {
		normalized, err := canonicalizeSeedURL(rawURL, c.config)
		if err != nil || !IsAllowedURL(normalized, c.config) || !c.allowedByRobots(normalized) {
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

	return seeds
}

func (c *Crawler) processURL(ctx context.Context, pageURL string) ([]string, error) {
	if !c.allowedByRobots(pageURL) {
		c.logf("crawl blocked-by-robots: url=%s", pageURL)
		return nil, fmt.Errorf("blocked by robots.txt")
	}
	c.logf("crawl process-start: url=%s", pageURL)

	pageRecord, hasPageRecord, err := c.loadPageRecord(pageURL)
	if err != nil {
		return nil, err
	}

	result, err := c.fetcher.Fetch(ctx, pageURL, FetchOptions{
		IfNoneMatch:     pageRecord.ETag,
		IfModifiedSince: pageRecord.LastModified,
	})
	if err != nil {
		_ = c.storeFetchFailure(pageURL, pageRecord, hasPageRecord, result, err)
		if result.PotentialBotChallenge {
			log.Printf("potential bot challenge: url=%s reason=%s", pageURL, result.BotChallengeReason)
		}
		return nil, err
	}
	if !c.allowedByRobots(result.FinalURL) {
		c.logf("crawl final-url-blocked-by-robots: source=%s final=%s", pageURL, result.FinalURL)
		_ = c.storeFetchFailure(pageURL, pageRecord, hasPageRecord, result, fmt.Errorf("final URL blocked by robots.txt"))
		return nil, fmt.Errorf("final URL blocked by robots.txt")
	}
	if result.NotModified {
		if err := c.storeNotModifiedPage(pageURL, result, pageRecord, hasPageRecord); err != nil {
			return nil, err
		}
		c.logf("crawl not-modified: url=%s final-url=%s", pageURL, result.FinalURL)
		return nil, nil
	}

	extracted, err := c.extractor.Extract(result.FinalURL, result.Body)
	if err != nil {
		_ = c.storeFetchFailure(pageURL, pageRecord, hasPageRecord, result, err)
		return nil, err
	}
	if strings.TrimSpace(extracted.RedirectURL) != "" {
		resolved, err := ResolveURL(result.FinalURL, extracted.RedirectURL, c.config)
		if err != nil {
			_ = c.storeFetchFailure(pageURL, pageRecord, hasPageRecord, result, err)
			return nil, fmt.Errorf("resolve redirect target %q: %w", extracted.RedirectURL, err)
		}
		if err := c.storePageRecordAliases(model.PageRecord{
			URL:          pageURL,
			RepoPath:     c.mapper.RepoPath(resolved),
			ETag:         result.ETag,
			LastModified: result.LastModified,
			LastFetched:  time.Now().UTC(),
			StatusCode:   result.StatusCode,
		}, pageURL, result.FinalURL, resolved); err != nil {
			return nil, err
		}
		c.logf("crawl meta-refresh: source=%s target=%s", result.FinalURL, resolved)
		return []string{resolved}, nil
	}
	c.logf("crawl extracted: url=%s title=%q links=%d", result.FinalURL, extracted.Title, len(extracted.Links))

	canonicalURL, err := NormalizeURL(extracted.CanonicalURL, c.config)
	if err != nil {
		canonicalURL = pageURL
	}
	markdownDocument, err := c.converter.Convert(extracted, canonicalURL)
	if err != nil {
		return nil, fmt.Errorf("convert %s: %w", canonicalURL, err)
	}

	repoPath := c.mapper.RepoPath(canonicalURL)
	contentWithFrontMatter := addTitleFrontMatter(extracted.Title, markdownDocument.Markdown)
	wroteFile, err := c.writer.WriteIfChanged(repoPath, contentWithFrontMatter)
	if err != nil {
		_ = c.storeFetchFailure(pageURL, pageRecord, hasPageRecord, result, err)
		return nil, fmt.Errorf("write %s: %w", repoPath, err)
	}
	if err := c.storePageRecordAliases(model.PageRecord{
		URL:          pageURL,
		RepoPath:     repoPath,
		ETag:         result.ETag,
		LastModified: result.LastModified,
		ContentHash:  markdownDocument.ContentHash,
		LastFetched:  time.Now().UTC(),
		StatusCode:   result.StatusCode,
		LastError:    "",
	}, pageURL, result.FinalURL, canonicalURL); err != nil {
		return nil, err
	}
	c.recordSearchEntry(repoPath, extracted.Title, markdownDocument.Markdown)
	if wroteFile {
		c.logf("crawl wrote: canonical-url=%s path=%s markdown-bytes=%d", canonicalURL, repoPath, len(markdownDocument.Markdown))
	} else {
		c.logf("crawl skip-write-unchanged: canonical-url=%s path=%s", canonicalURL, repoPath)
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

func (c *Crawler) loadPageRecord(pageURL string) (model.PageRecord, bool, error) {
	if c.store == nil {
		return model.PageRecord{}, false, nil
	}
	record, found, err := c.store.GetPage(pageURL)
	if err != nil || found {
		return record, found, err
	}

	normalizedURL, err := NormalizeURL(pageURL, c.config)
	if err != nil || normalizedURL == pageURL {
		return model.PageRecord{}, false, nil
	}

	return c.store.GetPage(normalizedURL)
}

func (c *Crawler) storeNotModifiedPage(pageURL string, result FetchResult, existing model.PageRecord, found bool) error {
	if c.store == nil {
		return nil
	}

	record := existing
	if !found {
		record = model.PageRecord{}
	}
	record.URL = pageURL
	if record.RepoPath == "" {
		record.RepoPath = c.mapper.RepoPath(result.FinalURL)
	}
	if strings.TrimSpace(result.ETag) != "" {
		record.ETag = result.ETag
	}
	if strings.TrimSpace(result.LastModified) != "" {
		record.LastModified = result.LastModified
	}
	record.LastFetched = time.Now().UTC()
	record.StatusCode = result.StatusCode
	record.LastError = ""

	return c.storePageRecordAliases(record, pageURL, result.FinalURL)
}

func (c *Crawler) storeFetchFailure(pageURL string, existing model.PageRecord, found bool, result FetchResult, fetchErr error) error {
	if c.store == nil {
		return nil
	}

	record := existing
	if !found {
		record = model.PageRecord{}
	}
	record.URL = pageURL
	if record.RepoPath == "" {
		record.RepoPath = c.mapper.RepoPath(pageURL)
	}
	if strings.TrimSpace(result.FinalURL) != "" && record.RepoPath == "" {
		record.RepoPath = c.mapper.RepoPath(result.FinalURL)
	}
	if strings.TrimSpace(result.ETag) != "" {
		record.ETag = result.ETag
	}
	if strings.TrimSpace(result.LastModified) != "" {
		record.LastModified = result.LastModified
	}
	record.LastFetched = time.Now().UTC()
	record.StatusCode = result.StatusCode
	record.LastError = fetchErr.Error()

	return c.storePageRecordAliases(record, pageURL, result.FinalURL)
}

func (c *Crawler) storePageRecordAliases(record model.PageRecord, rawURLs ...string) error {
	if c.store == nil {
		return nil
	}

	seen := make(map[string]struct{}, len(rawURLs))
	for _, rawURL := range rawURLs {
		trimmed := strings.TrimSpace(rawURL)
		if trimmed == "" {
			continue
		}

		candidates := []string{trimmed}
		if normalized, err := NormalizeURL(trimmed, c.config); err == nil && normalized != trimmed {
			candidates = append(candidates, normalized)
		}

		for _, candidate := range candidates {
			if _, exists := seen[candidate]; exists {
				continue
			}
			seen[candidate] = struct{}{}

			aliasRecord := record
			aliasRecord.URL = candidate
			if aliasRecord.RepoPath == "" {
				aliasRecord.RepoPath = c.mapper.RepoPath(candidate)
			}
			if err := c.store.PutPage(aliasRecord); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Crawler) logf(format string, args ...any) {
	if !c.config.DetailedLogging {
		return
	}
	log.Printf(format, args...)
}

func (c *Crawler) loadSearchIndex() error {
	content, err := c.writer.Read(searchIndexPath)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "no such file") {
			return nil
		}
		return err
	}

	var entries []searchIndexEntry
	if err := json.Unmarshal(content, &entries); err != nil {
		return fmt.Errorf("unmarshal %s: %w", searchIndexPath, err)
	}

	for _, entry := range entries {
		if strings.TrimSpace(entry.URL) == "" {
			continue
		}
		c.search[entry.URL] = entry
	}

	return nil
}

func (c *Crawler) saveSearchIndex() error {
	entries := make([]searchIndexEntry, 0, len(c.search))
	for _, entry := range c.search {
		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].URL < entries[j].URL
	})

	content, err := json.Marshal(entries)
	if err != nil {
		return fmt.Errorf("marshal search index: %w", err)
	}

	if err := c.writer.Write(searchIndexPath, string(content)); err != nil {
		return err
	}

	return nil
}

func (c *Crawler) recordSearchEntry(repoPath string, title string, markdown string) {
	permalink := repoPathToPermalink(repoPath)
	if permalink == "" {
		return
	}

	cleanTitle := strings.TrimSpace(title)
	if cleanTitle == "" {
		cleanTitle = humanizeSlug(filepath.Base(strings.TrimSuffix(repoPath, filepath.Ext(repoPath))))
	}

	c.search[permalink] = searchIndexEntry{
		Title:   cleanTitle,
		URL:     permalink,
		Content: searchSnippet(markdown),
		Section: sectionLabelFromRepoPath(repoPath),
	}
}

func addTitleFrontMatter(title string, markdown string) string {
	cleanTitle := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(title, "\r", " "), "\n", " "))
	if cleanTitle == "" {
		return markdown
	}

	var builder strings.Builder
	builder.WriteString("---\n")
	builder.WriteString("title: \"")
	builder.WriteString(escapeYAMLDoubleQuoted(cleanTitle))
	builder.WriteString("\"\n")
	builder.WriteString("---\n\n")
	builder.WriteString(strings.TrimLeft(markdown, "\n"))

	return builder.String()
}

func escapeYAMLDoubleQuoted(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "\"", "\\\"")
	return value
}

func repoPathToPermalink(repoPath string) string {
	pathValue := strings.TrimSpace(filepath.ToSlash(filepath.Clean(repoPath)))
	if pathValue == "" || !strings.HasSuffix(pathValue, ".md") {
		return ""
	}

	pathValue = strings.TrimPrefix(pathValue, "docs/")
	pathValue = strings.TrimSuffix(pathValue, ".md")
	if strings.HasSuffix(pathValue, "/index") {
		pathValue = strings.TrimSuffix(pathValue, "/index")
	}

	pathValue = strings.Trim(pathValue, "/")
	if pathValue == "" {
		return "/"
	}

	return "/" + pathValue + "/"
}

func sectionLabelFromRepoPath(repoPath string) string {
	pathValue := strings.TrimSpace(filepath.ToSlash(filepath.Clean(repoPath)))
	pathValue = strings.TrimPrefix(pathValue, "docs/")
	pathValue = strings.TrimSuffix(pathValue, ".md")
	parts := strings.Split(strings.Trim(pathValue, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return "General"
	}

	if len(parts) == 1 {
		return humanizeSlug(parts[0])
	}

	if parts[1] == "index" {
		return humanizeSlug(parts[0])
	}

	return humanizeSlug(parts[1])
}

func humanizeSlug(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}

	value = strings.ReplaceAll(value, "-", " ")
	words := strings.Fields(value)
	for i, word := range words {
		if word == "" {
			continue
		}
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}

func searchSnippet(markdown string) string {
	normalized := strings.Join(strings.Fields(markdown), " ")
	const maxLen = 500
	if len(normalized) <= maxLen {
		return normalized
	}
	return normalized[:maxLen]
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
	firstSegment := strings.ToLower(strings.SplitN(trimmedPath, "/", 2)[0])
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

func selectSectionsForRun(queue []string, limit int, cursor int) (map[string]struct{}, int) {
	orderedSections := make([]string, 0, len(queue))
	seen := map[string]struct{}{}
	for _, item := range queue {
		section := sectionKey(item)
		if section == "" {
			continue
		}
		if _, exists := seen[section]; exists {
			continue
		}
		seen[section] = struct{}{}
		orderedSections = append(orderedSections, section)
	}

	if len(orderedSections) == 0 {
		return map[string]struct{}{}, 0
	}
	if limit <= 0 || limit >= len(orderedSections) {
		selected := make(map[string]struct{}, len(orderedSections))
		for _, section := range orderedSections {
			selected[section] = struct{}{}
		}
		return selected, 0
	}

	start := 0
	if cursor > 0 {
		start = cursor % len(orderedSections)
	}

	selected := make(map[string]struct{}, limit)
	for index := 0; index < limit; index++ {
		section := orderedSections[(start+index)%len(orderedSections)]
		selected[section] = struct{}{}
	}

	next := (start + limit) % len(orderedSections)
	return selected, next
}

func (c *Crawler) filterReachableSeeds(ctx context.Context, seeds []string) []string {
	filtered := make([]string, 0, len(seeds))
	seen := make(map[string]struct{}, len(seeds))
	for _, seed := range seeds {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, seed, nil)
		if err != nil {
			continue
		}
		request.Header.Set("User-Agent", c.config.UserAgent)
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Set("Accept-Language", "en-US,en;q=0.9")

		response, err := c.fetcher.HTTPClient().Do(request)
		if err != nil {
			continue
		}
		_, _ = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()

		if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
			continue
		}

		finalURL, err := NormalizeURL(response.Request.URL.String(), c.config)
		if err != nil {
			continue
		}
		if !IsAllowedURL(finalURL, c.config) || !c.allowedByRobots(finalURL) {
			c.logf("seed pruned: seed=%s final-url=%s", seed, finalURL)
			continue
		}
		if _, exists := seen[finalURL]; exists {
			continue
		}

		seen[finalURL] = struct{}{}
		filtered = append(filtered, finalURL)
	}
	return filtered
}
