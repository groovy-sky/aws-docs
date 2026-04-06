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

	flag.StringVar(&options.ConfigPath, "config", "config.example.json", "path to crawler configuration")
	flag.StringVar(&options.Mode, "mode", "incremental", "run mode: incremental, full, refresh-url")
	flag.StringVar(&options.URL, "url", "", "single URL to refresh when mode=refresh-url")
	flag.IntVar(&options.MaxPages, "max-pages", 0, "maximum pages to process in this run (0 = config default)")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx, options); err != nil {
		log.Fatal(err)
	}
}
