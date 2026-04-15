package crawl

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sync"
	"testing"

	"github.com/groovy-sky/aws-docs/internal/config"
	"github.com/groovy-sky/aws-docs/internal/store"
	"github.com/groovy-sky/aws-docs/internal/write"
)

func TestProcessURLUsesConditionalFetchMetadata(t *testing.T) {
	var (
		mu               sync.Mutex
		ifNoneMatchValue string
		requestCount     int
	)

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		mu.Lock()
		requestCount++
		if requestCount == 2 {
			ifNoneMatchValue = request.Header.Get("If-None-Match")
		}
		mu.Unlock()

		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Header().Set("ETag", `"etag-v1"`)
		if request.Header.Get("If-None-Match") == `"etag-v1"` {
			writer.WriteHeader(http.StatusNotModified)
			return
		}

		_, _ = writer.Write([]byte(`<!doctype html><html><head><title>Guide</title></head><body><main><h1>Guide</h1><p>Hello world.</p></main></body></html>`))
	}))
	defer server.Close()

	metadataPath := filepath.Join(t.TempDir(), "crawl.db")
	database, err := store.Open(metadataPath)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer database.Close()

	root := t.TempDir()
	writer := write.New(root)
	cfg := config.Default()
	cfg.AllowedHosts = []string{serverURLHost(t, server.URL)}
	cfg.OutputDir = "docs"
	fetcher := NewFetcher(cfg)
	mapper := NewMapper(cfg.OutputDir)
	extractor := NewExtractor(cfg)
	converter := NewConverter(cfg, mapper, writer.Exists)
	crawler := NewCrawler(cfg, database, fetcher, extractor, converter, mapper, writer, nil)

	pageURL := server.URL + "/service/latest/userguide/index.html"
	if _, err := crawler.processURL(context.Background(), pageURL); err != nil {
		t.Fatalf("first processURL: %v", err)
	}

	record, found, err := database.GetPage(pageURL)
	if err != nil {
		t.Fatalf("get page record after first crawl: %v", err)
	}
	if !found {
		t.Fatal("expected page record after first crawl")
	}
	if record.ETag != `"etag-v1"` {
		t.Fatalf("stored ETag = %q, want %q", record.ETag, `"etag-v1"`)
	}
	if record.ContentHash == "" {
		t.Fatal("expected content hash to be stored")
	}

	if _, err := crawler.processURL(context.Background(), pageURL); err != nil {
		t.Fatalf("second processURL: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if requestCount != 2 {
		t.Fatalf("request count = %d, want 2", requestCount)
	}
	if ifNoneMatchValue != `"etag-v1"` {
		t.Fatalf("If-None-Match = %q, want %q", ifNoneMatchValue, `"etag-v1"`)
	}

	record, found, err = database.GetPage(pageURL)
	if err != nil {
		t.Fatalf("get page record after second crawl: %v", err)
	}
	if !found {
		t.Fatal("expected page record after second crawl")
	}
	if record.StatusCode != http.StatusNotModified {
		t.Fatalf("stored status code = %d, want %d", record.StatusCode, http.StatusNotModified)
	}
	if record.LastError != "" {
		t.Fatalf("last error = %q, want empty", record.LastError)
	}
}
