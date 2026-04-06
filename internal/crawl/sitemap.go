package crawl

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/groovy-sky/aws-docs/internal/config"
)

type sitemapURLSet struct {
	URLs []struct {
		Loc string `xml:"loc"`
	} `xml:"url"`
}

type sitemapIndex struct {
	Sitemaps []struct {
		Loc string `xml:"loc"`
	} `xml:"sitemap"`
}

func DiscoverURLsFromRobotsSitemaps(ctx context.Context, client *http.Client, cfg config.Config) ([]string, error) {
	if len(cfg.AllowedHosts) == 0 {
		return nil, nil
	}

	sitemapRoots := make([]string, 0)
	for _, host := range cfg.AllowedHosts {
		robotsURL := "https://" + host + "/robots.txt"
		urls, err := parseSitemapsFromRobots(ctx, client, robotsURL, cfg.UserAgent)
		if err != nil {
			continue
		}
		sitemapRoots = append(sitemapRoots, urls...)
	}

	if len(sitemapRoots) == 0 {
		return nil, nil
	}

	seenSitemaps := map[string]struct{}{}
	seenPages := map[string]struct{}{}
	queue := make([]string, 0, len(sitemapRoots))
	for _, value := range sitemapRoots {
		normalized, err := NormalizeURL(value, cfg)
		if err != nil {
			continue
		}
		if _, exists := seenSitemaps[normalized]; exists {
			continue
		}
		seenSitemaps[normalized] = struct{}{}
		queue = append(queue, normalized)
	}

	results := make([]string, 0, cfg.MaxSitemapURLs)
	for len(queue) > 0 {
		if len(results) >= cfg.MaxSitemapURLs || len(seenSitemaps) > cfg.MaxSitemapFiles {
			break
		}

		current := queue[0]
		queue = queue[1:]
		pageURLs, childSitemaps, err := parseSitemap(ctx, client, current, cfg.UserAgent)
		if err != nil {
			continue
		}

		for _, pageURL := range pageURLs {
			normalized, err := NormalizeURL(pageURL, cfg)
			if err != nil || !IsAllowedURL(normalized, cfg) {
				continue
			}
			if _, exists := seenPages[normalized]; exists {
				continue
			}
			seenPages[normalized] = struct{}{}
			results = append(results, normalized)
			if len(results) >= cfg.MaxSitemapURLs {
				break
			}
		}
		if len(results) >= cfg.MaxSitemapURLs {
			break
		}

		for _, sitemapURL := range childSitemaps {
			normalized, err := NormalizeURL(sitemapURL, cfg)
			if err != nil {
				continue
			}
			if !sitemapHostAllowed(normalized, cfg) {
				continue
			}
			if _, exists := seenSitemaps[normalized]; exists {
				continue
			}
			if len(seenSitemaps) >= cfg.MaxSitemapFiles {
				break
			}
			seenSitemaps[normalized] = struct{}{}
			queue = append(queue, normalized)
		}
	}

	return results, nil
}

func DeriveServiceRootSeeds(pageURLs []string, cfg config.Config) []string {
	seen := map[string]struct{}{}
	roots := make([]string, 0)

	for _, raw := range pageURLs {
		parsed, err := url.Parse(raw)
		if err != nil {
			continue
		}
		if parsed.Host == "" || !sitemapHostAllowed(parsed.String(), cfg) {
			continue
		}

		cleanPath := strings.Trim(path.Clean(parsed.Path), "/")
		if cleanPath == "" {
			continue
		}

		firstSegment := strings.Split(cleanPath, "/")[0]
		if firstSegment == "" {
			continue
		}

		rootURL := parsed.Scheme + "://" + parsed.Host + "/" + firstSegment + "/"
		normalized, err := NormalizeURL(rootURL, cfg)
		if err != nil || !IsAllowedURL(normalized, cfg) {
			continue
		}
		if _, exists := seen[normalized]; exists {
			continue
		}
		seen[normalized] = struct{}{}
		roots = append(roots, normalized)
	}

	return roots
}

func parseSitemapsFromRobots(ctx context.Context, client *http.Client, robotsURL string, userAgent string) ([]string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, robotsURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", userAgent)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("robots status %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for _, line := range strings.Split(string(body), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if !strings.HasPrefix(strings.ToLower(trimmed), "sitemap:") {
			continue
		}
		value := strings.TrimSpace(trimmed[len("sitemap:"):])
		if value != "" {
			result = append(result, value)
		}
	}
	return result, nil
}

func parseSitemap(ctx context.Context, client *http.Client, sitemapURL string, userAgent string) ([]string, []string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, sitemapURL, nil)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("User-Agent", userAgent)

	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return nil, nil, fmt.Errorf("sitemap status %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	if strings.HasSuffix(strings.ToLower(sitemapURL), ".gz") {
		gzipReader, err := gzip.NewReader(bytes.NewReader(body))
		if err == nil {
			decompressed, readErr := io.ReadAll(gzipReader)
			_ = gzipReader.Close()
			if readErr == nil {
				body = decompressed
			}
		}
	}

	urlSet := sitemapURLSet{}
	if err := xml.Unmarshal(body, &urlSet); err == nil && len(urlSet.URLs) > 0 {
		urls := make([]string, 0, len(urlSet.URLs))
		for _, entry := range urlSet.URLs {
			if strings.TrimSpace(entry.Loc) != "" {
				urls = append(urls, strings.TrimSpace(entry.Loc))
			}
		}
		return urls, nil, nil
	}

	index := sitemapIndex{}
	if err := xml.Unmarshal(body, &index); err == nil && len(index.Sitemaps) > 0 {
		sitemaps := make([]string, 0, len(index.Sitemaps))
		for _, entry := range index.Sitemaps {
			if strings.TrimSpace(entry.Loc) != "" {
				sitemaps = append(sitemaps, strings.TrimSpace(entry.Loc))
			}
		}
		return nil, sitemaps, nil
	}

	return nil, nil, fmt.Errorf("unsupported sitemap format")
}

func sitemapHostAllowed(rawURL string, cfg config.Config) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	host := strings.ToLower(parsed.Host)
	for _, allowed := range cfg.AllowedHosts {
		if host == strings.ToLower(allowed) {
			return true
		}
	}
	return false
}
