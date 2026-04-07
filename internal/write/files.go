package write

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileWriter struct {
	root string
}

func New(root string) *FileWriter {
	return &FileWriter{root: root}
}

func (w *FileWriter) Exists(relativePath string) bool {
	cleanRelative := filepath.Clean(relativePath)
	fullPath := filepath.Join(w.root, cleanRelative)
	info, err := os.Stat(fullPath)
	return err == nil && !info.IsDir()
}

func (w *FileWriter) Write(relativePath string, content string) error {
	cleanRelative := filepath.Clean(relativePath)
	fullPath := filepath.Join(w.root, cleanRelative)
	if !strings.HasPrefix(fullPath, w.root) {
		return fmt.Errorf("refusing to write outside root: %s", relativePath)
	}
	content = sanitizeSensitiveContent(content)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
