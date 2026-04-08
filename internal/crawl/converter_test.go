package crawl

import (
	"strings"
	"testing"

	"github.com/groovy-sky/aws-docs/internal/config"
)

func TestConvertAppendsSourceAttribution(t *testing.T) {
	converter := NewConverter(config.Default(), NewMapper("docs"), nil)

	document := ExtractedDocument{
		CanonicalURL: "https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html",
		HTML:         "<h1>VPC</h1><p>Example content.</p>",
	}

	markdownDocument, err := converter.Convert(document, document.CanonicalURL)
	if err != nil {
		t.Fatalf("Convert returned error: %v", err)
	}

	if !strings.Contains(markdownDocument.Markdown, "All content copied from https://docs.aws.amazon.com/.") {
		t.Fatalf("Convert markdown missing attribution footer:\n%s", markdownDocument.Markdown)
	}
	if !strings.HasSuffix(markdownDocument.Markdown, "All content copied from https://docs.aws.amazon.com/.\n") {
		t.Fatalf("Convert markdown should end with attribution footer:\n%s", markdownDocument.Markdown)
	}
}

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

func TestRewriteHrefKeepsOriginalImageURLWithQueryAndFragment(t *testing.T) {
	converter := NewConverter(config.Default(), NewMapper("docs"), nil)

	source := "https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html"
	href := "http://docs.aws.amazon.com/images/vpc/diagram.png?utm_source=mail#zoom"

	rewritten := converter.rewriteHref(source, href)
	if rewritten != href {
		t.Fatalf("rewriteHref = %q, want unchanged original image URL %q", rewritten, href)
	}
}

func TestRewriteAssetURLResolvesRelativeImageToAbsoluteWithoutDroppingParts(t *testing.T) {
	converter := NewConverter(config.Default(), NewMapper("docs"), nil)

	source := "https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html"
	src := "../images/security-group-overview.png?ref=sidebar#full"

	rewritten := converter.rewriteAssetURL(source, src)
	want := "https://docs.aws.amazon.com/vpc/latest/images/security-group-overview.png?ref=sidebar#full"
	if rewritten != want {
		t.Fatalf("rewriteAssetURL = %q, want %q", rewritten, want)
	}
}
