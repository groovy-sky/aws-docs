package write

import "regexp"

var (
	accessKeyEnvPattern   = regexp.MustCompile(`(?mi)(AWS_ACCESS_KEY_ID\s*=\s*["']?)([A-Z0-9]{16,})(["']?)`)
	secretKeyEnvPattern   = regexp.MustCompile(`(?mi)(AWS_SECRET_ACCESS_KEY\s*=\s*["']?)([A-Za-z0-9/+=]{16,})(["']?)`)
	accessKeyLiteral      = regexp.MustCompile(`\b(A3T[A-Z0-9]|AKIA|ASIA|ANPA|AROA|AIDA|ABIA)[A-Z0-9]{12,}\b`)
	secretKeyValuePattern = regexp.MustCompile(`(?mi)(AWS_SECRET_ACCESS_KEY\s*:\s*["']?)([A-Za-z0-9/+=]{16,})(["']?)`)
)

func sanitizeSensitiveContent(content string) string {
	content = accessKeyEnvPattern.ReplaceAllString(content, `${1}AWS_ACCESS_KEY_ID_REDACTED${3}`)
	content = secretKeyEnvPattern.ReplaceAllString(content, `${1}AWS_SECRET_ACCESS_KEY_REDACTED${3}`)
	content = secretKeyValuePattern.ReplaceAllString(content, `${1}AWS_SECRET_ACCESS_KEY_REDACTED${3}`)
	content = accessKeyLiteral.ReplaceAllString(content, `AWS_ACCESS_KEY_ID_REDACTED`)
	return content
}
