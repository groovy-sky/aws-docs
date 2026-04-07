package crawl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/groovy-sky/aws-docs/internal/config"
	"github.com/temoto/robotstxt"
)

type Robots struct {
	group *robotstxt.Group
}

type RobotsStructure struct {
	Sitemaps      []string
	AllowPaths    []string
	DisallowPaths []string
	SectionRoots  []string
}

func LoadRobots(ctx context.Context, client *http.Client, baseURL string, userAgent string) (*Robots, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/robots.txt", nil)
	if err != nil {
		return nil, fmt.Errorf("build robots request: %w", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch robots.txt: %w", err)
	}
	defer response.Body.Close()

	robotsData, err := robotstxt.FromResponse(response)
	if err != nil {
		return nil, fmt.Errorf("parse robots.txt: %w", err)
	}

	return &Robots{group: robotsData.FindGroup(userAgent)}, nil
}

func (r *Robots) Allowed(path string) bool {
	if r == nil || r.group == nil {
		return true
	}
	return r.group.Test(path)
}

func DiscoverRobotsStructure(ctx context.Context, client *http.Client, host string, cfg config.Config) (RobotsStructure, error) {
	robotsURL := "https://" + strings.TrimSpace(host) + "/robots.txt"
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, robotsURL, nil)
	if err != nil {
		return RobotsStructure{}, fmt.Errorf("build robots request: %w", err)
	}
	request.Header.Set("User-Agent", cfg.UserAgent)

	response, err := client.Do(request)
	if err != nil {
		return RobotsStructure{}, fmt.Errorf("fetch robots.txt: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return RobotsStructure{}, fmt.Errorf("robots status %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return RobotsStructure{}, fmt.Errorf("read robots.txt: %w", err)
	}

	structure := parseRobotsLines(string(body))
	structure.SectionRoots = deriveSectionRootsFromRobots(host, structure, cfg)
	return structure, nil
}

func parseRobotsLines(content string) RobotsStructure {
	allowSet := map[string]struct{}{}
	disallowSet := map[string]struct{}{}
	sitemapSet := map[string]struct{}{}

	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if idx := strings.Index(trimmed, "#"); idx >= 0 {
			trimmed = strings.TrimSpace(trimmed[:idx])
			if trimmed == "" {
				continue
			}
		}

		parts := strings.SplitN(trimmed, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(parts[0]))
		value := strings.TrimSpace(parts[1])
		if value == "" {
			continue
		}

		switch key {
		case "allow":
			allowSet[value] = struct{}{}
		case "disallow":
			disallowSet[value] = struct{}{}
		case "sitemap":
			sitemapSet[value] = struct{}{}
		}
	}

	allow := setToSortedSlice(allowSet)
	disallow := setToSortedSlice(disallowSet)
	sitemaps := setToSortedSlice(sitemapSet)

	return RobotsStructure{
		Sitemaps:      sitemaps,
		AllowPaths:    allow,
		DisallowPaths: disallow,
	}
}

func deriveSectionRootsFromRobots(host string, structure RobotsStructure, cfg config.Config) []string {
	seen := map[string]struct{}{}
	roots := make([]string, 0, cfg.MaxRobotsSections)

	addFromPath := func(rawPath string) {
		value := strings.TrimSpace(rawPath)
		if value == "" || value == "/" || strings.HasPrefix(value, "*") {
			return
		}

		seedPath := deriveSectionSeedPath(value)
		if seedPath == "" {
			return
		}
		candidate := "https://" + host + seedPath
		normalized, err := NormalizeURL(candidate, cfg)
		if err != nil || !IsAllowedURL(normalized, cfg) {
			return
		}
		if _, exists := seen[normalized]; exists {
			return
		}
		seen[normalized] = struct{}{}
		roots = append(roots, normalized)
	}

	for _, value := range structure.AllowPaths {
		if len(roots) >= cfg.MaxRobotsSections {
			break
		}
		addFromPath(value)
	}
	for _, value := range structure.DisallowPaths {
		if len(roots) >= cfg.MaxRobotsSections {
			break
		}
		addFromPath(value)
	}

	return roots
}

func setToSortedSlice(values map[string]struct{}) []string {
	result := make([]string, 0, len(values))
	for value := range values {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
