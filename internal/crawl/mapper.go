package crawl

import (
	"path/filepath"
	"regexp"
	"strings"
)

var nonSlugPattern = regexp.MustCompile(`[^a-z0-9]+`)

type Mapper struct {
	outputDir string
}

func NewMapper(outputDir string) *Mapper {
	return &Mapper{outputDir: outputDir}
}

func (m *Mapper) RepoPath(rawURL string) string {
	cleaned := strings.TrimPrefix(rawURL, "https://docs.aws.amazon.com/")
	cleaned = strings.TrimPrefix(cleaned, "http://docs.aws.amazon.com/")
	segments := strings.Split(cleaned, "/")
	if len(segments) == 0 || segments[0] == "" {
		return filepath.Join(m.outputDir, "general", "index.md")
	}

	section := "misc"
	service := ""
	first := segments[0]
	switch {
	case strings.EqualFold(first, "general") || strings.Contains(cleaned, "/general/"):
		section = "general"
		// Consume the leading "general" segment so it is not appended again
		// when building the path (section already contributes "general/").
		if strings.EqualFold(first, "general") {
			segments = segments[1:]
		}
	case strings.Contains(strings.ToLower(cleaned), "/cli/"):
		section = filepath.Join("reference", "cli")
	case strings.Contains(strings.ToLower(cleaned), "sdk") || strings.Contains(strings.ToLower(cleaned), "apireference"):
		section = "reference"
	default:
		section = "services"
		service = mapServiceCode(first)
	}

	pathSegments := make([]string, 0, len(segments)+2)
	pathSegments = append(pathSegments, m.outputDir)
	pathSegments = append(pathSegments, filepath.SplitList(section)...)
	if strings.Contains(section, string(filepath.Separator)) {
		pathSegments = []string{m.outputDir}
		pathSegments = append(pathSegments, strings.Split(section, string(filepath.Separator))...)
	}
	if service != "" {
		pathSegments = append(pathSegments, service)
		segments = segments[1:]
	}

	if len(segments) == 0 {
		segments = []string{"index"}
	}

	for _, segment := range segments {
		if segment == "" {
			continue
		}
		pathSegments = append(pathSegments, slugify(segment))
	}

	last := pathSegments[len(pathSegments)-1]
	last = strings.TrimSuffix(last, ".html")
	last = strings.TrimSuffix(last, ".htm")
	if last == "" || last == "latest" {
		last = filepath.Join(last, "index")
	}
	if !strings.HasSuffix(last, ".md") {
		last += ".md"
	}
	pathSegments[len(pathSegments)-1] = last

	return filepath.Join(pathSegments...)
}

func (m *Mapper) RelativeLink(fromURL, toURL string) string {
	fromPath := m.RepoPath(fromURL)
	toPath := m.RepoPath(toURL)
	relative, err := filepath.Rel(filepath.Dir(fromPath), toPath)
	if err != nil {
		return toPath
	}
	return filepath.ToSlash(relative)
}

func mapServiceCode(value string) string {
	mappings := map[string]string{
		"AWSEC2":            "ec2",
		"AmazonS3":          "s3",
		"IAM":               "iam",
		"AWSCloudFormation": "cloudformation",
		"amazondynamodb":    "dynamodb",
		"eks":               "eks",
	}
	for key, mapped := range mappings {
		if strings.EqualFold(key, value) {
			return mapped
		}
	}
	return slugify(value)
}

func slugify(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.TrimSuffix(value, "/")
	value = strings.ReplaceAll(value, ".html", "")
	value = strings.ReplaceAll(value, ".htm", "")
	value = strings.ReplaceAll(value, "_", "-")
	value = nonSlugPattern.ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	if value == "" {
		return "index"
	}
	return value
}
