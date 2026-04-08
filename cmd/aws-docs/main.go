package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/groovy-sky/aws-docs/internal/app"
)

func main() {
	var options app.Options

	flag.StringVar(&options.ConfigPath, "config", "", "optional path to crawler configuration")
	flag.StringVar(&options.Mode, "mode", "partial", "run mode: partial, incremental, full, refresh-url")
	flag.StringVar(&options.URL, "url", "", "single URL to refresh when mode=refresh-url")
	flag.IntVar(&options.MaxSections, "max-sections", 0, "maximum sections to process in partial/incremental mode (0 = config default)")
	flag.BoolVar(&options.DetailedLogging, "detailed-logging", false, "emit per-request and per-page crawl logs")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx, options); err != nil {
		log.Fatal(err)
	}
}
