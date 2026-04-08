package crawl

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/groovy-sky/aws-docs/internal/config"
)

func TestFilterReachableSeedsPrunesRedirectsToExcludedPaths(t *testing.T) {
	var serverURL string
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/sdk":
			http.Redirect(writer, request, serverURL+"/aws-sdk-php/v3/api/", http.StatusFound)
		case "/aws-sdk-php/v3/api/":
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("ok"))
		default:
			writer.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	serverURL = server.URL

	cfg := config.Default()
	cfg.AllowedHosts = []string{serverURLHost(t, server.URL)}
	cfg.ExcludePathPatterns = []string{"/aws-sdk-php/v3/"}
	fetcher := NewFetcher(cfg)
	crawler := NewCrawler(cfg, nil, fetcher, nil, nil, nil, nil, nil)

	filtered := crawler.filterReachableSeeds(context.Background(), []string{server.URL + "/sdk"})
	if len(filtered) != 0 {
		t.Fatalf("filtered seeds = %v, want none", filtered)
	}
}

func TestFilterReachableSeedsNormalizesToFinalAllowedURL(t *testing.T) {
	var serverURL string
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/guide":
			http.Redirect(writer, request, serverURL+"/service/latest/userguide/", http.StatusFound)
		case "/service/latest/userguide/":
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("ok"))
		default:
			writer.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	serverURL = server.URL

	cfg := config.Default()
	cfg.AllowedHosts = []string{serverURLHost(t, server.URL)}
	cfg.ExcludePathPatterns = nil
	fetcher := NewFetcher(cfg)
	crawler := NewCrawler(cfg, nil, fetcher, nil, nil, nil, nil, nil)

	filtered := crawler.filterReachableSeeds(context.Background(), []string{server.URL + "/guide"})
	if len(filtered) != 1 {
		t.Fatalf("filtered seeds length = %d, want 1", len(filtered))
	}
	if !strings.HasSuffix(filtered[0], "/service/latest/userguide") {
		t.Fatalf("filtered seed = %q, want redirected final path suffix %q", filtered[0], "/service/latest/userguide")
	}
}

func serverURLHost(t *testing.T, rawURL string) string {
	t.Helper()

	request, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	return request.URL.Host
}
