package write

import (
	"bytes"
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
	_, err := w.WriteIfChanged(relativePath, content)
	return err
}

func (w *FileWriter) WriteIfChanged(relativePath string, content string) (bool, error) {
	cleanRelative := filepath.Clean(relativePath)
	fullPath := filepath.Join(w.root, cleanRelative)
	if !strings.HasPrefix(fullPath, w.root) {
		return false, fmt.Errorf("refusing to write outside root: %s", relativePath)
	}
	content = sanitizeSensitiveContent(content)
	newContent := []byte(content)
	if existingContent, err := os.ReadFile(fullPath); err == nil {
		if bytes.Equal(existingContent, newContent) {
			return false, nil
		}
	} else if !os.IsNotExist(err) {
		return false, fmt.Errorf("read existing file: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return false, fmt.Errorf("create output directory: %w", err)
	}
	if err := os.WriteFile(fullPath, newContent, 0o644); err != nil {
		return false, fmt.Errorf("write file: %w", err)
	}
	return true, nil
}

func (w *FileWriter) Read(relativePath string) ([]byte, error) {
	cleanRelative := filepath.Clean(relativePath)
	fullPath := filepath.Join(w.root, cleanRelative)
	if !strings.HasPrefix(fullPath, w.root) {
		return nil, fmt.Errorf("refusing to read outside root: %s", relativePath)
	}
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
