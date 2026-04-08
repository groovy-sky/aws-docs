package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPreservesExplicitEmptyPathPatterns(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")
	configJSON := `{
		"include_path_patterns": [],
		"exclude_path_patterns": []
	}`

	if err := os.WriteFile(configPath, []byte(configJSON), 0o644); err != nil {
		t.Fatalf("write config file: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if len(cfg.IncludePathPatterns) != 0 {
		t.Fatalf("include_path_patterns length = %d, want 0", len(cfg.IncludePathPatterns))
	}
	if len(cfg.ExcludePathPatterns) != 0 {
		t.Fatalf("exclude_path_patterns length = %d, want 0", len(cfg.ExcludePathPatterns))
	}
}

func TestLoadDefaultsPathPatternsWhenOmitted(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	if err := os.WriteFile(configPath, []byte(`{"seeds": ["https://docs.aws.amazon.com/"]}`), 0o644); err != nil {
		t.Fatalf("write config file: %v", err)
	}

	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if len(cfg.IncludePathPatterns) == 0 {
		t.Fatal("include_path_patterns should default when omitted")
	}
	if len(cfg.ExcludePathPatterns) == 0 {
		t.Fatal("exclude_path_patterns should default when omitted")
	}
}
