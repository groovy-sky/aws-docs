package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/groovy-sky/aws-docs/internal/config"
	"github.com/groovy-sky/aws-docs/internal/crawl"
	"github.com/groovy-sky/aws-docs/internal/store"
	"github.com/groovy-sky/aws-docs/internal/write"
)

type Options struct {
	ConfigPath string
	Mode       string
	URL        string
	MaxPages   int
}

func Run(ctx context.Context, options Options) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	cfg, err := config.Load(options.ConfigPath)
	if err != nil {
		return err
	}
	metadataPath := resolvePath(workingDir, cfg.MetadataDB)
	if options.MaxPages > 0 {
		cfg.MaxPages = options.MaxPages
	}

	database, err := store.Open(metadataPath)
	if err != nil {
		return err
	}
	defer database.Close()

	fetcher := crawl.NewFetcher(cfg)
	robots, err := crawl.LoadRobots(ctx, fetcher.HTTPClient(), "https://docs.aws.amazon.com", cfg.UserAgent)
	if err != nil {
		return err
	}

	mapper := crawl.NewMapper(cfg.OutputDir)
	extractor := crawl.NewExtractor(cfg)
	converter := crawl.NewConverter(cfg, mapper)
	writer := write.New(workingDir)

	runner := crawl.NewCrawler(cfg, database, fetcher, extractor, converter, mapper, writer, robots)
	return runner.Run(ctx, crawl.RunOptions{
		Mode:     options.Mode,
		URL:      options.URL,
		MaxPages: cfg.MaxPages,
	})
}

func resolvePath(baseDir string, value string) string {
	if filepath.IsAbs(value) {
		return filepath.Clean(value)
	}
	return filepath.Join(baseDir, filepath.Clean(value))
}
