package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Config struct {
	Seeds                 []string `json:"seeds"`
	AllowedHosts          []string `json:"allowed_hosts"`
	OutputDir             string   `json:"output_dir"`
	MetadataDB            string   `json:"metadata_db"`
	EnableRobotsStructure bool     `json:"enable_robots_structure"`
	EnableRobotsSitemaps  bool     `json:"enable_robots_sitemaps"`
	MaxRobotsSections     int      `json:"max_robots_sections"`
	MaxSitemapURLs        int      `json:"max_sitemap_urls"`
	MaxSitemapFiles       int      `json:"max_sitemap_files"`
	UserAgent             string   `json:"user_agent"`
	Concurrency           int      `json:"concurrency"`
	RequestsPerSecond     float64  `json:"requests_per_second"`
	Burst                 int      `json:"burst"`
	MaxRetries            int      `json:"max_retries"`
	RequestTimeoutSeconds int      `json:"request_timeout_seconds"`
	MaxSections           int      `json:"max_sections"`
	Selectors             []string `json:"selectors"`
	ExcludeSelectors      []string `json:"exclude_selectors"`
	IgnoredQueryParams    []string `json:"ignored_query_params"`
	IncludePathPatterns   []string `json:"include_path_patterns"`
	ExcludePathPatterns   []string `json:"exclude_path_patterns"`
}

func Default() Config {
	return Config{
		Seeds: []string{
			"https://docs.aws.amazon.com/",
			"https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html",
			"https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html",
			"https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction.html",
		},
		AllowedHosts:          []string{"docs.aws.amazon.com"},
		OutputDir:             "docs",
		MetadataDB:            "metadata/crawl.db",
		EnableRobotsStructure: true,
		EnableRobotsSitemaps:  true,
		MaxRobotsSections:     500,
		MaxSitemapURLs:        50000,
		MaxSitemapFiles:       1000,
		UserAgent:             "aws-docs-mirror/0.1 (+https://github.com/groovy-sky/aws-docs)",
		Concurrency:           2,
		RequestsPerSecond:     1.0,
		Burst:                 1,
		MaxRetries:            3,
		RequestTimeoutSeconds: 20,
		MaxSections:           25,
		Selectors: []string{
			"#main-content",
			"main",
			"article",
			".awsdocs-content",
			".awsui-context-content-header + div",
		},
		ExcludeSelectors: []string{
			"script",
			"style",
			"noscript",
			"iframe",
			"header",
			"footer",
			"nav",
			"aside",
			".breadcrumbs",
			".awsdocs-toc",
			".awsui-util-hide",
			".feedback-section",
		},
		IgnoredQueryParams: []string{
			"utm_source",
			"utm_medium",
			"utm_campaign",
			"utm_term",
			"utm_content",
			"ref_",
		},
		IncludePathPatterns: []string{
			"/aws",
			"/amazon",
			"/iam/",
			"/general/",
			"/cli/",
			"/sdkfor",
			"/whitepapers/",
			"/latest/",
			"/userguide/",
			"/developerguide/",
			"/api-reference/",
			"/apireference/",
			"/reference/",
		},
		ExcludePathPatterns: []string{
			"/console/",
			"/marketplace/",
			"/training/",
			"/premiumsupport/",
			"/solutions/",
		},
	}
}

func Load(path string) (Config, error) {
	if strings.TrimSpace(path) == "" {
		cfg := Default()
		cfg.normalize()
		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	cfg := Default()
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	cfg.normalize()
	return cfg, nil
}

func (c *Config) normalize() {
	if c.OutputDir == "" {
		c.OutputDir = "docs"
	}
	if c.MetadataDB == "" {
		c.MetadataDB = "metadata/crawl.db"
	}
	if c.UserAgent == "" {
		c.UserAgent = Default().UserAgent
	}
	if c.MaxRobotsSections < 1 {
		c.MaxRobotsSections = 500
	}
	if c.MaxSitemapURLs < 1 {
		c.MaxSitemapURLs = 50000
	}
	if c.MaxSitemapFiles < 1 {
		c.MaxSitemapFiles = 1000
	}
	if c.Concurrency < 1 {
		c.Concurrency = 1
	}
	if c.RequestsPerSecond <= 0 {
		c.RequestsPerSecond = 1.0
	}
	if c.Burst < 1 {
		c.Burst = 1
	}
	if c.MaxRetries < 1 {
		c.MaxRetries = 1
	}
	if c.RequestTimeoutSeconds < 1 {
		c.RequestTimeoutSeconds = 20
	}
	if c.MaxSections < 0 {
		c.MaxSections = 0
	}
	if len(c.AllowedHosts) == 0 {
		c.AllowedHosts = []string{"docs.aws.amazon.com"}
	}
	for index, host := range c.AllowedHosts {
		c.AllowedHosts[index] = strings.ToLower(strings.TrimSpace(host))
	}
	if len(c.Seeds) == 0 {
		c.Seeds = slices.Clone(Default().Seeds)
	}
	if len(c.Selectors) == 0 {
		c.Selectors = slices.Clone(Default().Selectors)
	}
	if len(c.ExcludeSelectors) == 0 {
		c.ExcludeSelectors = slices.Clone(Default().ExcludeSelectors)
	}
	if len(c.IgnoredQueryParams) == 0 {
		c.IgnoredQueryParams = slices.Clone(Default().IgnoredQueryParams)
	}
	if len(c.IncludePathPatterns) == 0 {
		c.IncludePathPatterns = slices.Clone(Default().IncludePathPatterns)
	}
	if len(c.ExcludePathPatterns) == 0 {
		c.ExcludePathPatterns = slices.Clone(Default().ExcludePathPatterns)
	}
	for index, value := range c.Seeds {
		c.Seeds[index] = strings.TrimSpace(value)
	}
	for index, value := range c.Selectors {
		c.Selectors[index] = strings.TrimSpace(value)
	}
	for index, value := range c.ExcludeSelectors {
		c.ExcludeSelectors[index] = strings.TrimSpace(value)
	}
	for index, value := range c.IgnoredQueryParams {
		c.IgnoredQueryParams[index] = strings.ToLower(strings.TrimSpace(value))
	}
	for index, value := range c.IncludePathPatterns {
		c.IncludePathPatterns[index] = strings.ToLower(strings.TrimSpace(value))
	}
	for index, value := range c.ExcludePathPatterns {
		c.ExcludePathPatterns[index] = strings.TrimSpace(value)
	}
	if !filepath.IsAbs(c.OutputDir) {
		c.OutputDir = filepath.Clean(c.OutputDir)
	}
	if !filepath.IsAbs(c.MetadataDB) {
		c.MetadataDB = filepath.Clean(c.MetadataDB)
	}
}
