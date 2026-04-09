package crawl

import (
	"testing"
)

func TestRepoPathGeneralURLNotDoubled(t *testing.T) {
	mapper := NewMapper("docs")

	got := mapper.RepoPath("https://docs.aws.amazon.com/general/latest/gr/docconventions.html")
	want := "docs/general/latest/gr/docconventions.md"
	if got != want {
		t.Fatalf("RepoPath = %q, want %q", got, want)
	}
}

func TestRepoPathServiceURL(t *testing.T) {
	mapper := NewMapper("docs")

	got := mapper.RepoPath("https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html")
	want := "docs/services/vpc/latest/userguide/what-is-amazon-vpc.md"
	if got != want {
		t.Fatalf("RepoPath = %q, want %q", got, want)
	}
}

func TestRepoPathReferenceURLNotDoubled(t *testing.T) {
	mapper := NewMapper("docs")

	got := mapper.RepoPath("https://docs.aws.amazon.com/acm/latest/APIReference/Welcome.html")
	want := "docs/reference/acm/latest/apireference/welcome.md"
	if got != want {
		t.Fatalf("RepoPath = %q, want %q", got, want)
	}
}

func TestRepoPathGeneralRootURL(t *testing.T) {
	mapper := NewMapper("docs")

	// Root URL maps to general index.
	got := mapper.RepoPath("https://docs.aws.amazon.com/")
	want := "docs/general/index.md"
	if got != want {
		t.Fatalf("RepoPath = %q, want %q", got, want)
	}
}

func TestRepoPathStripsQueryParameters(t *testing.T) {
	mapper := NewMapper("docs")

	// Query parameters should be stripped from the path mapping
	got := mapper.RepoPath("https://docs.aws.amazon.com/ec2/latest/userguide/index.html?icmpid=docs_homepage_featuredsvcs")
	want := "docs/services/ec2/latest/userguide/index.md"
	if got != want {
		t.Fatalf("RepoPath = %q, want %q", got, want)
	}
}

func TestRepoPathStripsFragments(t *testing.T) {
	mapper := NewMapper("docs")

	// Fragments should also be stripped from the path mapping
	got := mapper.RepoPath("https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html#section")
	want := "docs/services/vpc/latest/userguide/what-is-amazon-vpc.md"
	if got != want {
		t.Fatalf("RepoPath = %q, want %q", got, want)
	}
}
