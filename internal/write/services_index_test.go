package write

import (
	"strings"
	"testing"
	"time"
)

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
