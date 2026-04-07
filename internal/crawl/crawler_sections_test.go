package crawl

import "testing"

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
