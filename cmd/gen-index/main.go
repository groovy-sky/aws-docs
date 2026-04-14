package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	docsDir := "docs"
	outputPath := "public/docs-index.json"

	var paths []string
	err := filepath.Walk(docsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			rel, _ := filepath.Rel(docsDir, path)
			rel = filepath.ToSlash(rel) // normalize to forward slashes
			paths = append(paths, rel)
		}
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking docs: %v\n", err)
		os.Exit(1)
	}

	sort.Strings(paths)

	os.MkdirAll("public", 0755)
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating index: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	data, _ := json.MarshalIndent(paths, "", "  ")
	f.Write(data)

	fmt.Printf("Generated %s with %d entries\n", outputPath, len(paths))
}
