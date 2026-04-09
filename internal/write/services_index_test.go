package write

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestPickWelcomeOrFirstPrefersWelcomeMd(t *testing.T) {
	root := t.TempDir()
	servicePath := filepath.Join(root, "docs", "reference", "myservice")

	files := []string{
		filepath.Join(servicePath, "latest", "apireference", "api-zzz.md"),
		filepath.Join(servicePath, "latest", "apireference", "api-aaa.md"),
		filepath.Join(servicePath, "latest", "apireference", "welcome.md"),
	}
	for _, f := range files {
		if err := os.MkdirAll(filepath.Dir(f), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(f, []byte("# content"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	link, found := pickServiceLink(servicePath, root)
	if !found {
		t.Fatal("expected a link to be found")
	}
	if !strings.HasSuffix(link, "welcome.md") {
		t.Fatalf("expected welcome.md to be preferred, got %q", link)
	}
}

func TestPickWelcomeOrFirstPrefersShallowestWelcomeMd(t *testing.T) {
	root := t.TempDir()
	servicePath := filepath.Join(root, "docs", "reference", "myservice")

	files := []string{
		filepath.Join(servicePath, "latest", "apireference", "nested", "deep", "welcome.md"),
		filepath.Join(servicePath, "latest", "apireference", "welcome.md"),
	}
	for _, f := range files {
		if err := os.MkdirAll(filepath.Dir(f), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(f, []byte("# content"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	link, found := pickServiceLink(servicePath, root)
	if !found {
		t.Fatal("expected a link to be found")
	}
	if !strings.Contains(link, "apireference/welcome.md") {
		t.Fatalf("expected shallowest welcome.md, got %q", link)
	}
}

func TestPickWelcomeOrFirstFallsBackToAlphabeticalWhenNoWelcomeMd(t *testing.T) {
	root := t.TempDir()
	servicePath := filepath.Join(root, "docs", "reference", "myservice")

	files := []string{
		filepath.Join(servicePath, "latest", "apireference", "api-zzz.md"),
		filepath.Join(servicePath, "latest", "apireference", "api-aaa.md"),
	}
	for _, f := range files {
		if err := os.MkdirAll(filepath.Dir(f), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(f, []byte("# content"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	link, found := pickServiceLink(servicePath, root)
	if !found {
		t.Fatal("expected a link to be found")
	}
	if !strings.HasSuffix(link, "api-aaa.md") {
		t.Fatalf("expected first alphabetical file, got %q", link)
	}
}

func TestBuildServicesIndexMarkdownOrdersServicesBeforeLastRun(t *testing.T) {
	buckets := []docBucket{
		{
			SectionName: "Services",
			Entries: []serviceEntry{
				{Name: "Ec2", Link: "docs/services/ec2/index.md"},
				{Name: "Vpc", Link: "docs/services/vpc/index.md"},
			},
		},
	}
	runInfo := ServicesIndexRunInfo{
		Mode:        "incremental",
		Status:      "success",
		GeneratedAt: time.Date(2026, time.April, 8, 9, 55, 40, 0, time.UTC),
	}

	content := buildServicesIndexMarkdown(buckets, runInfo)

	servicesPos := strings.Index(content, "## Services (2)")
	lastRunPos := strings.Index(content, "## Last run")
	if servicesPos == -1 || lastRunPos == -1 {
		t.Fatalf("expected Services and Last run sections to be present, got:\n%s", content)
	}
	if servicesPos > lastRunPos {
		t.Fatalf("expected services section before last run summary, got:\n%s", content)
	}

	if !strings.Contains(content, "Last run: mode=incremental, status=success, updated_at_utc=2026-04-08T09:55:40Z") {
		t.Fatalf("expected one-line last run summary, got:\n%s", content)
	}
	if strings.Contains(content, "- Mode:") || strings.Contains(content, "- Status:") || strings.Contains(content, "- Updated at (UTC):") {
		t.Fatalf("expected last run summary to be one line, got:\n%s", content)
	}
}

func TestBuildServicesIndexMarkdownRendersAllBuckets(t *testing.T) {
	buckets := []docBucket{
		{SectionName: "Services", Entries: []serviceEntry{{Name: "Vpc", Link: "docs/services/vpc/index.md"}}},
		{SectionName: "Reference", Entries: []serviceEntry{{Name: "Acm", Link: "docs/reference/acm/latest/apireference.md"}}},
		{SectionName: "General", Entries: []serviceEntry{{Name: "General", Link: "docs/general/index.md"}}},
	}
	runInfo := ServicesIndexRunInfo{
		Mode:        "full",
		Status:      "success",
		GeneratedAt: time.Date(2026, time.April, 8, 9, 55, 40, 0, time.UTC),
	}

	content := buildServicesIndexMarkdown(buckets, runInfo)

	for _, heading := range []string{"## Services (1)", "## Reference (1)", "## General (1)"} {
		if !strings.Contains(content, heading) {
			t.Fatalf("expected heading %q in output:\n%s", heading, content)
		}
	}

	svcPos := strings.Index(content, "## Services")
	refPos := strings.Index(content, "## Reference")
	genPos := strings.Index(content, "## General")
	lastRunPos := strings.Index(content, "## Last run")

	if !(svcPos < refPos && refPos < genPos && genPos < lastRunPos) {
		t.Fatalf("expected sections in order Services > Reference > General > Last run, got:\n%s", content)
	}
}
