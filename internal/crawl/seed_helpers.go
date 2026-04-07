package crawl

import (
	"path"
	"strings"
)

func deriveSectionSeedPath(rawPath string) string {
	trimmed := strings.TrimSpace(rawPath)
	if trimmed == "" {
		return ""
	}
	trimmed = strings.SplitN(trimmed, "?", 2)[0]
	trimmed = strings.SplitN(trimmed, "#", 2)[0]
	if !strings.HasPrefix(trimmed, "/") {
		trimmed = "/" + trimmed
	}

	cleaned := strings.Trim(path.Clean(trimmed), "/")
	if cleaned == "" {
		return ""
	}

	segments := strings.Split(cleaned, "/")
	if len(segments) == 0 || strings.Contains(segments[0], ".") {
		return ""
	}

	if len(segments) > 1 && segments[1] != "" && !strings.Contains(segments[1], ".") {
		if strings.EqualFold(segments[1], "latest") && len(segments) > 2 && segments[2] != "" && !strings.Contains(segments[2], ".") {
			return "/" + segments[0] + "/" + segments[1] + "/" + segments[2]
		}
		return "/" + segments[0] + "/" + segments[1]
	}

	return "/" + segments[0]
}
