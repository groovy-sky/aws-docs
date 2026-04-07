package write

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type serviceEntry struct {
	Name string
	Link string
}

type ServicesIndexRunInfo struct {
	Mode        string
	Status      string
	Error       string
	GeneratedAt time.Time
}

func (w *FileWriter) WriteServicesIndex(outputDir string, runInfo ServicesIndexRunInfo) error {
	outputRoot := filepath.Join(w.root, filepath.Clean(outputDir))
	servicesRoot := filepath.Join(outputRoot, "services")
	entries, err := collectServiceEntries(servicesRoot, w.root)
	if err != nil {
		return err
	}

	content := buildServicesIndexMarkdown(entries, runInfo)
	if err := w.Write("SERVICES.md", content); err != nil {
		return fmt.Errorf("write root services index: %w", err)
	}
	return nil
}

func collectServiceEntries(servicesRoot string, repoRoot string) ([]serviceEntry, error) {
	items, err := os.ReadDir(servicesRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return []serviceEntry{}, nil
		}
		return nil, fmt.Errorf("read services directory: %w", err)
	}

	entries := make([]serviceEntry, 0, len(items))
	for _, item := range items {
		if !item.IsDir() {
			continue
		}

		serviceKey := item.Name()
		servicePath := filepath.Join(servicesRoot, serviceKey)
		link, found := pickServiceLink(servicePath, repoRoot)
		if !found {
			continue
		}

		entries = append(entries, serviceEntry{
			Name: humanizeServiceName(serviceKey),
			Link: link,
		})
	}

	sort.Slice(entries, func(i int, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	return entries, nil
}

func pickServiceLink(servicePath string, repoRoot string) (string, bool) {
	candidates := []string{
		filepath.Join(servicePath, "index.md"),
		filepath.Join(servicePath, "latest", "index.md"),
	}

	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			rel, relErr := filepath.Rel(repoRoot, candidate)
			if relErr != nil {
				continue
			}
			return slashify(rel), true
		}
	}

	allMarkdown := make([]string, 0)
	_ = filepath.WalkDir(servicePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if strings.EqualFold(filepath.Ext(path), ".md") {
			allMarkdown = append(allMarkdown, path)
		}
		return nil
	})

	if len(allMarkdown) == 0 {
		return "", false
	}
	sort.Strings(allMarkdown)

	rel, err := filepath.Rel(repoRoot, allMarkdown[0])
	if err != nil {
		return "", false
	}

	return slashify(rel), true
}

func humanizeServiceName(value string) string {
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == '-' || r == '_'
	})
	if len(parts) == 0 {
		return strings.ToUpper(value)
	}

	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}

func buildServicesIndexMarkdown(entries []serviceEntry, runInfo ServicesIndexRunInfo) string {
	var builder strings.Builder
	builder.WriteString("# Parsed Services\n\n")
	builder.WriteString("This file is auto-generated after crawl runs and lists currently parsed services.\n\n")
	builder.WriteString("## Last run\n\n")
	builder.WriteString("- Mode: `")
	builder.WriteString(strings.TrimSpace(runInfo.Mode))
	builder.WriteString("`\n")
	builder.WriteString("- Status: `")
	builder.WriteString(strings.TrimSpace(runInfo.Status))
	builder.WriteString("`\n")
	builder.WriteString("- Updated at (UTC): `")
	builder.WriteString(runInfo.GeneratedAt.UTC().Format(time.RFC3339))
	builder.WriteString("`\n")
	if strings.TrimSpace(runInfo.Error) != "" {
		builder.WriteString("- Error: `")
		builder.WriteString(runInfo.Error)
		builder.WriteString("`\n")
	}
	builder.WriteString("\n")

	if len(entries) == 0 {
		builder.WriteString("No parsed services found yet.\n")
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("Total services: %d\n\n", len(entries)))
	for _, entry := range entries {
		builder.WriteString("- [")
		builder.WriteString(entry.Name)
		builder.WriteString("](")
		builder.WriteString(entry.Link)
		builder.WriteString(")\n")
	}

	return builder.String()
}

func slashify(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
