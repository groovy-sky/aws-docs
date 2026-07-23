package crawl

import "testing"

func TestAddTitleFrontMatter(t *testing.T) {
	markdown := "# Heading\n\nBody"
	got := addTitleFrontMatter("My \"Title\"", markdown)
	wantPrefix := "---\ntitle: \"My \\\"Title\\\"\"\n---\n\n"
	if len(got) <= len(wantPrefix) || got[:len(wantPrefix)] != wantPrefix {
		t.Fatalf("front matter prefix mismatch:\n%s", got)
	}
}

func TestRepoPathToPermalink(t *testing.T) {
	cases := map[string]string{
		"docs/services/acm/latest/index.md": "/services/acm/latest/",
		"docs/general/index.md":             "/general/",
		"docs/reference/cli/index.md":       "/reference/cli/",
	}

	for input, want := range cases {
		if got := repoPathToPermalink(input); got != want {
			t.Fatalf("repoPathToPermalink(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestSectionLabelFromRepoPath(t *testing.T) {
	cases := map[string]string{
		"docs/services/amazoncloudwatch/latest/index.md": "Amazoncloudwatch",
		"docs/reference/cli/index.md":                    "Cli",
		"docs/general/index.md":                          "General",
	}

	for input, want := range cases {
		if got := sectionLabelFromRepoPath(input); got != want {
			t.Fatalf("sectionLabelFromRepoPath(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestExtractMarkdownTitle(t *testing.T) {
	markdown := "Intro paragraph\n\n# What is Amazon EC2?\n\nDetails."
	if got := extractMarkdownTitle(markdown); got != "What is Amazon EC2?" {
		t.Fatalf("extractMarkdownTitle() = %q, want %q", got, "What is Amazon EC2?")
	}
}

func TestExtractMarkdownLinks(t *testing.T) {
	markdown := "- [Concepts](concepts.md)\n- [Guide](https://docs.aws.amazon.com/vpc/latest/userguide/)\n<https://docs.aws.amazon.com/ec2/latest/userguide/>"
	got := extractMarkdownLinks(markdown)
	if len(got) != 3 {
		t.Fatalf("extractMarkdownLinks() length = %d, want 3", len(got))
	}
}
