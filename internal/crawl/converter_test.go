package crawl

import (
	"strings"
	"testing"

	"github.com/groovy-sky/aws-docs/internal/config"
)

func TestRewriteHrefRewritesInDomainLinkEvenWhenExcludedByFilters(t *testing.T) {
	cfg := config.Default()
	cfg.ExcludePathPatterns = []string{"/aws-sdk-php/v3/"}
	converter := NewConverter(cfg, NewMapper("docs"), nil)

	source := "https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html"
	href := "https://docs.aws.amazon.com/aws-sdk-php/v3/api-accessanalyzer-2019-11-01.html"

	rewritten := converter.rewriteHref(source, href)
	if strings.HasPrefix(rewritten, "https://docs.aws.amazon.com/") {
		t.Fatalf("rewriteHref returned absolute docs URL %q, want repository-relative path", rewritten)
	}
	if !strings.HasSuffix(rewritten, "reference/aws-sdk-php/v3/api-accessanalyzer-2019-11-01.md") {
		t.Fatalf("rewriteHref = %q, want suffix %q", rewritten, "reference/aws-sdk-php/v3/api-accessanalyzer-2019-11-01.md")
	}
}

func TestRewriteHrefKeepsExternalLinkAbsolute(t *testing.T) {
	converter := NewConverter(config.Default(), NewMapper("docs"), nil)

	source := "https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html"
	href := "https://example.com/docs/page.html"

	rewritten := converter.rewriteHref(source, href)
	if rewritten != "https://example.com/docs/page.html" {
		t.Fatalf("rewriteHref = %q, want external absolute URL", rewritten)
	}
}

func TestRewriteHrefKeepsImageLinkAbsolute(t *testing.T) {
	converter := NewConverter(config.Default(), NewMapper("docs"), nil)

	source := "https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html"
	href := "https://docs.aws.amazon.com/images/vpc/latest/userguide/images/security-group-overview.png"

	rewritten := converter.rewriteHref(source, href)
	if rewritten != href {
		t.Fatalf("rewriteHref = %q, want original image URL %q", rewritten, href)
	}
}
