package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/groovy-sky/aws-docs/internal/config"
	"github.com/groovy-sky/aws-docs/internal/crawl"
	"github.com/groovy-sky/aws-docs/internal/store"
	"github.com/groovy-sky/aws-docs/internal/write"
)

type Options struct {
	ConfigPath      string
	Mode            string
	URL             string
	MaxSections     int
	DetailedLogging bool
}

func Run(ctx context.Context, options Options) (runErr error) {
	if options.Mode == "" || options.Mode == "partial" {
		options.Mode = "incremental"
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	writer := write.New(workingDir)
	outputDir := "docs"
	defer func() {
		status := "success"
		errorText := ""
		if runErr != nil {
			status = "error"
			errorText = runErr.Error()
		}

		indexErr := writer.WriteServicesIndex(outputDir, write.ServicesIndexRunInfo{
			Mode:        options.Mode,
			Status:      status,
			Error:       errorText,
			GeneratedAt: time.Now().UTC(),
		})
		if indexErr != nil && runErr == nil {
			runErr = fmt.Errorf("write services index: %w", indexErr)
		}
	}()

	cfg, err := config.Load(options.ConfigPath)
	if err != nil {
		return err
	}
	cfg.DetailedLogging = options.DetailedLogging
	outputDir = cfg.OutputDir
	metadataPath := resolvePath(workingDir, cfg.MetadataDB)
	if options.MaxSections > 0 {
		cfg.MaxSections = options.MaxSections
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
	converter := crawl.NewConverter(cfg, mapper, writer.Exists)

	runner := crawl.NewCrawler(cfg, database, fetcher, extractor, converter, mapper, writer, robots)
	runErr = runner.Run(ctx, crawl.RunOptions{
		Mode:        options.Mode,
		URL:         options.URL,
		MaxSections: cfg.MaxSections,
	})
	if runErr != nil {
		return runErr
	}

	return nil
}

func resolvePath(baseDir string, value string) string {
	if filepath.IsAbs(value) {
		return filepath.Clean(value)
	}
	return filepath.Join(baseDir, filepath.Clean(value))
}
