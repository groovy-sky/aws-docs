package crawl

import (
	"testing"

	"github.com/groovy-sky/aws-docs/internal/config"
)

func TestIsAllowedURLAcceptsNormalizedLatestPathWithoutTrailingSlash(t *testing.T) {
	cfg := config.Default()
	cfg.IncludePathPatterns = []string{}
	cfg.ExcludePathPatterns = []string{}

	if !IsAllowedURL("https://docs.aws.amazon.com/vpc/latest", cfg) {
		t.Fatalf("IsAllowedURL should allow normalized latest path without trailing slash")
	}
}

func TestIsAllowedURLAcceptsNormalizedUserGuidePathWithoutTrailingSlash(t *testing.T) {
	cfg := config.Default()
	cfg.IncludePathPatterns = []string{}
	cfg.ExcludePathPatterns = []string{}

	if !IsAllowedURL("https://docs.aws.amazon.com/vpc/latest/userguide", cfg) {
		t.Fatalf("IsAllowedURL should allow normalized userguide path without trailing slash")
	}
}

func TestResolveURLRewritesRootLatestToSectionLatest(t *testing.T) {
	cfg := config.Default()

	got, err := ResolveURL("https://docs.aws.amazon.com/vpc/latest/userguide/", "/latest/peering/", cfg)
	if err != nil {
		t.Fatalf("ResolveURL returned error: %v", err)
	}

	want := "https://docs.aws.amazon.com/vpc/latest/peering"
	if got != want {
		t.Fatalf("ResolveURL = %q, want %q", got, want)
	}
}

func TestResolveURLRewritesRootLatestWithoutTrailingSlash(t *testing.T) {
	cfg := config.Default()

	got, err := ResolveURL("https://docs.aws.amazon.com/vpc/latest/userguide/", "/latest", cfg)
	if err != nil {
		t.Fatalf("ResolveURL returned error: %v", err)
	}

	want := "https://docs.aws.amazon.com/vpc/latest"
	if got != want {
		t.Fatalf("ResolveURL = %q, want %q", got, want)
	}
}

func TestResolveURLDoesNotRewriteRootLatestWhenBaseHasNoSection(t *testing.T) {
	cfg := config.Default()

	got, err := ResolveURL("https://docs.aws.amazon.com/", "/latest/userguide/", cfg)
	if err != nil {
		t.Fatalf("ResolveURL returned error: %v", err)
	}

	want := "https://docs.aws.amazon.com/latest/userguide"
	if got != want {
		t.Fatalf("ResolveURL = %q, want %q", got, want)
	}
}
