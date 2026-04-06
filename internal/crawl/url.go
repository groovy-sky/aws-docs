package crawl

import (
	"net/url"
	"path"
	"sort"
	"strings"

	"github.com/groovy-sky/aws-docs/internal/config"
)

func NormalizeURL(rawURL string, cfg config.Config) (string, error) {
	parsed, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		return "", err
	}

	parsed.Scheme = strings.ToLower(parsed.Scheme)
	parsed.Host = strings.ToLower(parsed.Host)
	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}
	if parsed.Scheme == "http" && hostAllowed(parsed.Host, cfg) {
		parsed.Scheme = "https"
	}
	parsed.Fragment = ""
	if parsed.Path == "" {
		parsed.Path = "/"
	}
	parsed.Path = path.Clean(parsed.Path)
	if !strings.HasPrefix(parsed.Path, "/") {
		parsed.Path = "/" + parsed.Path
	}

	query := parsed.Query()
	for key := range query {
		lowerKey := strings.ToLower(key)
		if shouldDropQueryParam(lowerKey, cfg.IgnoredQueryParams) {
			query.Del(key)
		}
	}
	if len(query) == 0 {
		parsed.RawQuery = ""
	} else {
		keys := make([]string, 0, len(query))
		for key := range query {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		encoded := url.Values{}
		for _, key := range keys {
			values := query[key]
			sort.Strings(values)
			for _, value := range values {
				encoded.Add(key, value)
			}
		}
		parsed.RawQuery = encoded.Encode()
	}

	return parsed.String(), nil
}

func ResolveURL(baseURL, href string, cfg config.Config) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	reference, err := url.Parse(strings.TrimSpace(href))
	if err != nil {
		return "", err
	}
	return NormalizeURL(base.ResolveReference(reference).String(), cfg)
}

func IsAllowedURL(rawURL string, cfg config.Config) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	pathValue := strings.ToLower(parsed.Path)
	host := strings.ToLower(parsed.Host)
	if !hostAllowed(host, cfg) {
		return false
	}

	for _, pattern := range cfg.ExcludePathPatterns {
		if pattern != "" && strings.Contains(pathValue, strings.ToLower(pattern)) {
			return false
		}
	}
	if len(cfg.IncludePathPatterns) == 0 {
		return isLikelyDocPath(pathValue)
	}
	for _, pattern := range cfg.IncludePathPatterns {
		if pattern != "" && strings.Contains(pathValue, pattern) {
			return true
		}
	}
	return pathValue == "/" || isLikelyDocPath(pathValue)
}

func isLikelyDocPath(pathValue string) bool {
	if pathValue == "/" {
		return true
	}
	trimmed := strings.Trim(pathValue, "/")
	if trimmed == "" {
		return true
	}
	segments := strings.Split(trimmed, "/")
	if len(segments) == 1 {
		// Many AWS docs landings are top-level slugs, e.g. /next-generation-sagemaker/.
		return true
	}
	docSegments := []string{
		"/latest/",
		"/userguide/",
		"/developerguide/",
		"/api-reference/",
		"/apireference/",
		"/reference/",
		"/getting-started/",
		"/installation/",
		"/security/",
		"/gr/",
	}
	for _, segment := range docSegments {
		if strings.Contains(pathValue, segment) {
			return true
		}
	}
	return strings.HasSuffix(pathValue, ".html") && strings.Count(strings.Trim(pathValue, "/"), "/") >= 2
}

func shouldDropQueryParam(key string, ignored []string) bool {
	for _, prefix := range ignored {
		if prefix != "" && strings.HasPrefix(key, prefix) {
			return true
		}
	}
	return false
}

func hostAllowed(host string, cfg config.Config) bool {
	host = strings.ToLower(strings.TrimSpace(host))
	for _, value := range cfg.AllowedHosts {
		if host == strings.ToLower(strings.TrimSpace(value)) {
			return true
		}
	}
	return false
}
