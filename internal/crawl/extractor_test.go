package crawl

import (
	"net/url"
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

func TestExtractorCollectsLinksFromExcludedTOC(t *testing.T) {
	extractor := NewExtractor(config.Default())
	body := []byte(`<!DOCTYPE html>
<html>
  <head><title>EBS encryption</title></head>
  <body>
    <main id="main-content">
      <div class="awsdocs-toc">
        <a href="how-ebs-encryption-works.html">How EBS encryption works</a>
        <a href="ebs-encryption-requirements.html">Requirements</a>
      </div>
      <h1>EBS encryption</h1>
      <p>Overview content.</p>
    </main>
  </body>
</html>`)

	document, err := extractor.Extract("https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html", body)
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}

	if len(document.Links) != 2 {
		t.Fatalf("Links length = %d, want 2", len(document.Links))
	}
	if document.Links[0] != "how-ebs-encryption-works.html" {
		t.Fatalf("first link = %q, want %q", document.Links[0], "how-ebs-encryption-works.html")
	}
	if document.Links[1] != "ebs-encryption-requirements.html" {
		t.Fatalf("second link = %q, want %q", document.Links[1], "ebs-encryption-requirements.html")
	}
}

func TestExtractorExtractsLandingPageXMLListCardLinks(t *testing.T) {
	extractor := NewExtractor(config.Default())
	landingXML := `<main-landing-page>
	<title>Welcome to AWS Documentation</title>
	<abstract>Find user guides and references.</abstract>
	<sections>
		<section>
			<cards>
				<list-card>
					<list-card-items>
						<list-card-item href="/vpc/?icmpid=docs_homepage_networking">
							<title>Amazon VPC</title>
						</list-card-item>
					</list-card-items>
				</list-card>
				<simple-card>
					<footer>
						<footer-item href="/cli/latest/userguide/cli-chap-welcome.html">AWS CLI User Guide</footer-item>
					</footer>
				</simple-card>
			</cards>
		</section>
	</sections>
</main-landing-page>`
	body := []byte(`<!DOCTYPE html>
<html>
	<head><title>Welcome to AWS Documentation</title></head>
	<body>
		<input id="landing-page-xml" value="` + url.QueryEscape(landingXML) + `" />
	</body>
</html>`)

	document, err := extractor.Extract("https://docs.aws.amazon.com/", body)
	if err != nil {
		t.Fatalf("Extract returned error: %v", err)
	}
	if len(document.Links) != 2 {
		t.Fatalf("Links length = %d, want 2", len(document.Links))
	}
	if document.Links[0] != "/vpc/?icmpid=docs_homepage_networking" {
		t.Fatalf("first link = %q, want %q", document.Links[0], "/vpc/?icmpid=docs_homepage_networking")
	}
	if document.Links[1] != "/cli/latest/userguide/cli-chap-welcome.html" {
		t.Fatalf("second link = %q, want %q", document.Links[1], "/cli/latest/userguide/cli-chap-welcome.html")
	}
}
