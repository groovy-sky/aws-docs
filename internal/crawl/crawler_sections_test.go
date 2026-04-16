package crawl

import (
	"testing"
	"time"
)

func TestSelectSectionsForRunWrapsAndAdvancesCursor(t *testing.T) {
	queue := []string{
		"https://docs.aws.amazon.com/vpc/latest/userguide/",
		"https://docs.aws.amazon.com/vpc/latest/peering/",
		"https://docs.aws.amazon.com/iam/latest/userguide/",
		"https://docs.aws.amazon.com/lambda/latest/dg/",
		"https://docs.aws.amazon.com/s3/latest/userguide/",
	}

	first, cursor := selectSectionsForRun(queue, 2, 0)
	if len(first) != 2 {
		t.Fatalf("first selection size = %d, want 2", len(first))
	}
	if cursor != 2 {
		t.Fatalf("first cursor = %d, want 2", cursor)
	}

	second, cursor := selectSectionsForRun(queue, 2, cursor)
	if len(second) != 2 {
		t.Fatalf("second selection size = %d, want 2", len(second))
	}
	if cursor != 0 {
		t.Fatalf("second cursor = %d, want 0", cursor)
	}

	if _, ok := first["docs.aws.amazon.com/vpc"]; !ok {
		t.Fatalf("first selection missing vpc section")
	}
	if _, ok := first["docs.aws.amazon.com/iam"]; !ok {
		t.Fatalf("first selection missing iam section")
	}
	if _, ok := second["docs.aws.amazon.com/lambda"]; !ok {
		t.Fatalf("second selection missing lambda section")
	}
	if _, ok := second["docs.aws.amazon.com/s3"]; !ok {
		t.Fatalf("second selection missing s3 section")
	}
}

func TestSelectSectionsForRunLimitAboveSectionCountSelectsAll(t *testing.T) {
	queue := []string{
		"https://docs.aws.amazon.com/vpc/latest/userguide/",
		"https://docs.aws.amazon.com/iam/latest/userguide/",
	}

	selected, cursor := selectSectionsForRun(queue, 5, 3)
	if len(selected) != 2 {
		t.Fatalf("selection size = %d, want 2", len(selected))
	}
	if cursor != 0 {
		t.Fatalf("cursor = %d, want 0", cursor)
	}
}

func TestScheduledSlotAt(t *testing.T) {
	now := time.Date(2026, time.April, 16, 13, 0, 0, 0, time.UTC)
	slot := scheduledSlotAt(now)

	if slot != 373 {
		t.Fatalf("scheduledSlotAt() = %d, want 373", slot)
	}
}

func TestSelectSectionForScheduledRun(t *testing.T) {
	queue := []string{
		"https://docs.aws.amazon.com/vpc/latest/userguide/",
		"https://docs.aws.amazon.com/s3/latest/userguide/",
		"https://docs.aws.amazon.com/iam/latest/userguide/",
		"https://docs.aws.amazon.com/vpc/latest/peering/",
	}

	now := time.Date(2026, time.April, 16, 13, 0, 0, 0, time.UTC)
	section, slot, index, sectionCount, err := selectSectionForScheduledRun(queue, now)
	if err != nil {
		t.Fatalf("selectSectionForScheduledRun() error = %v", err)
	}

	if slot != 373 {
		t.Fatalf("slot = %d, want 373", slot)
	}
	if sectionCount != 3 {
		t.Fatalf("sectionCount = %d, want 3", sectionCount)
	}
	if index != 1 {
		t.Fatalf("index = %d, want 1", index)
	}
	if section != "docs.aws.amazon.com/s3" {
		t.Fatalf("section = %q, want %q", section, "docs.aws.amazon.com/s3")
	}
}

func TestSelectSectionForScheduledRunNoSections(t *testing.T) {
	_, _, _, _, err := selectSectionForScheduledRun(nil, time.Now().UTC())
	if err == nil {
		t.Fatalf("selectSectionForScheduledRun(nil) expected error")
	}
}
func TestNormalizeSectionName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      string
		wantError bool
	}{
		{name: "valid", input: "vpc", want: "vpc"},
		{name: "trim slashes", input: "/vpc/", want: "vpc"},
		{name: "missing", input: "", wantError: true},
		{name: "nested", input: "vpc/latest", wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeSectionName(tt.input)
			if tt.wantError {
				if err == nil {
					t.Fatalf("normalizeSectionName(%q) expected error", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("normalizeSectionName(%q) error = %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("normalizeSectionName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSelectSectionsByName(t *testing.T) {
	queue := []string{
		"https://docs.aws.amazon.com/vpc/latest/userguide/",
		"https://docs.aws.amazon.com/VPC/latest/tgw/",
		"https://docs.aws.amazon.com/iam/latest/userguide/",
	}

	selected := selectSectionsByName(queue, "vpc")
	if len(selected) != 1 {
		t.Fatalf("selection size = %d, want 1", len(selected))
	}
	if _, ok := selected["docs.aws.amazon.com/vpc"]; !ok {
		t.Fatalf("selection missing docs.aws.amazon.com/vpc")
	}

	none := selectSectionsByName(queue, "rds")
	if len(none) != 0 {
		t.Fatalf("selection size = %d, want 0", len(none))
	}
}
