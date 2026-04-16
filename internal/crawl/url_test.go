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
