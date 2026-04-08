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

type docBucket struct {
	SectionName string
	Entries     []serviceEntry
}

type ServicesIndexRunInfo struct {
	Mode        string
	Status      string
	Error       string
	GeneratedAt time.Time
}

func (w *FileWriter) WriteServicesIndex(outputDir string, runInfo ServicesIndexRunInfo) error {
	outputRoot := filepath.Join(w.root, filepath.Clean(outputDir))
	buckets := make([]docBucket, 0, 3)

	serviceEntries, err := collectBucketEntries(filepath.Join(outputRoot, "services"), w.root)
	if err != nil {
		return err
	}
	if len(serviceEntries) > 0 {
		buckets = append(buckets, docBucket{SectionName: "Services", Entries: serviceEntries})
	}

	referenceEntries, err := collectBucketEntries(filepath.Join(outputRoot, "reference"), w.root)
	if err != nil {
		return err
	}
	if len(referenceEntries) > 0 {
		buckets = append(buckets, docBucket{SectionName: "Reference", Entries: referenceEntries})
	}

	generalEntries, err := collectBucketEntries(filepath.Join(outputRoot, "general"), w.root)
	if err != nil {
		return err
	}
	if len(generalEntries) == 0 {
		link, found := pickServiceLink(filepath.Join(outputRoot, "general"), w.root)
		if found {
			generalEntries = []serviceEntry{{Name: "General", Link: link}}
		}
	}
	if len(generalEntries) > 0 {
		buckets = append(buckets, docBucket{SectionName: "General", Entries: generalEntries})
	}

	content := buildServicesIndexMarkdown(buckets, runInfo)
	if err := w.Write("SERVICES.md", content); err != nil {
		return fmt.Errorf("write root services index: %w", err)
	}
	return nil
}

func collectBucketEntries(bucketRoot string, repoRoot string) ([]serviceEntry, error) {
	items, err := os.ReadDir(bucketRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return []serviceEntry{}, nil
		}
		return nil, fmt.Errorf("read bucket directory: %w", err)
	}

	entries := make([]serviceEntry, 0, len(items))
	for _, item := range items {
		if !item.IsDir() {
			continue
		}

		serviceKey := item.Name()
		servicePath := filepath.Join(bucketRoot, serviceKey)
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

func buildServicesIndexMarkdown(buckets []docBucket, runInfo ServicesIndexRunInfo) string {
	var builder strings.Builder
	builder.WriteString("# Parsed Services\n\n")
	builder.WriteString("This file is auto-generated after crawl runs and lists currently mirrored content.\n\n")

	hasContent := false
	for _, bucket := range buckets {
		if len(bucket.Entries) == 0 {
			continue
		}
		hasContent = true
		builder.WriteString(fmt.Sprintf("## %s (%d)\n\n", bucket.SectionName, len(bucket.Entries)))
		for _, entry := range bucket.Entries {
			builder.WriteString("- [")
			builder.WriteString(entry.Name)
			builder.WriteString("](")
			builder.WriteString(entry.Link)
			builder.WriteString(")\n")
		}
		builder.WriteString("\n")
	}

	if !hasContent {
		builder.WriteString("No content found yet.\n\n")
	}

	builder.WriteString("## Last run\n\n")
	builder.WriteString("Last run: mode=")
	builder.WriteString(strings.TrimSpace(runInfo.Mode))
	builder.WriteString(", status=")
	builder.WriteString(strings.TrimSpace(runInfo.Status))
	builder.WriteString(", updated_at_utc=")
	builder.WriteString(runInfo.GeneratedAt.UTC().Format(time.RFC3339))
	if strings.TrimSpace(runInfo.Error) != "" {
		builder.WriteString(", error=")
		builder.WriteString(runInfo.Error)
	}
	builder.WriteString("\n")

	return builder.String()
}

func slashify(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
