package crawl

import (
	"testing"

	"github.com/groovy-sky/aws-docs/internal/config"
)

func TestExtractorExtractsMetaRefreshRedirect(t *testing.T) {
	extractor := NewExtractor(config.Default())
	body := []byte(`<!DOCTYPE html>
<html>
  <head>
    <title>Amazon Virtual Private Cloud</title>
    <meta http-equiv="refresh" content="10;URL=what-is-amazon-vpc.html">
  </head>
  <body></body>
</html>`)

	document, err := extractor.Extract("https://docs.aws.amazon.com/vpc/latest/userguide/", body)
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	if document.RedirectURL != "what-is-amazon-vpc.html" {
		t.Fatalf("RedirectURL = %q, want %q", document.RedirectURL, "what-is-amazon-vpc.html")
	}
	if document.HTML != "" {
		t.Fatalf("HTML = %q, want empty", document.HTML)
	}
	if len(document.Links) != 0 {
		t.Fatalf("Links length = %d, want 0", len(document.Links))
	}
}

func TestParseMetaRefreshTarget(t *testing.T) {
	target, ok := parseMetaRefreshTarget("0; URL='guide.html'")
	if !ok {
		t.Fatal("parseMetaRefreshTarget returned ok = false, want true")
	}
	if target != "guide.html" {
		t.Fatalf("target = %q, want %q", target, "guide.html")
	}
}
