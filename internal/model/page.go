package model

import "time"

type PageRecord struct {
	URL          string    `json:"url"`
	RepoPath     string    `json:"repo_path"`
	ETag         string    `json:"etag,omitempty"`
	LastModified string    `json:"last_modified,omitempty"`
	ContentHash  string    `json:"content_hash,omitempty"`
	LastFetched  time.Time `json:"last_fetched"`
	StatusCode   int       `json:"status_code"`
	LastError    string    `json:"last_error,omitempty"`
}
