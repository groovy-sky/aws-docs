package write

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteIfChangedSkipsIdenticalContent(t *testing.T) {
	root := t.TempDir()
	writer := New(root)

	wrote, err := writer.WriteIfChanged(filepath.Join("docs", "sample.md"), "hello\n")
	if err != nil {
		t.Fatalf("first write: %v", err)
	}
	if !wrote {
		t.Fatal("first write reported unchanged")
	}

	wrote, err = writer.WriteIfChanged(filepath.Join("docs", "sample.md"), "hello\n")
	if err != nil {
		t.Fatalf("second write: %v", err)
	}
	if wrote {
		t.Fatal("second write reported file changed")
	}

	content, err := os.ReadFile(filepath.Join(root, "docs", "sample.md"))
	if err != nil {
		t.Fatalf("read written file: %v", err)
	}
	if string(content) != "hello\n" {
		t.Fatalf("file content = %q, want %q", string(content), "hello\n")
	}

	wrote, err = writer.WriteIfChanged(filepath.Join("docs", "sample.md"), "updated\n")
	if err != nil {
		t.Fatalf("third write: %v", err)
	}
	if !wrote {
		t.Fatal("third write reported unchanged")
	}
}
