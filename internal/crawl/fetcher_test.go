package crawl

import "testing"

func TestShouldRetryHTTPStatus(t *testing.T) {
	cases := []struct {
		status int
		want   bool
	}{
		{status: 200, want: false},
		{status: 400, want: false},
		{status: 403, want: true},
		{status: 429, want: true},
		{status: 503, want: true},
		{status: 500, want: true},
	}

	for _, tc := range cases {
		got := shouldRetryHTTPStatus(tc.status)
		if got != tc.want {
			t.Fatalf("shouldRetryHTTPStatus(%d) = %v, want %v", tc.status, got, tc.want)
		}
	}
}

func TestDetectBotChallenge(t *testing.T) {
	body := []byte("<html><body><h1>Verify you are human</h1></body></html>")
	isChallenge, reason := detectBotChallenge(200, "text/html; charset=utf-8", body)
	if !isChallenge {
		t.Fatal("expected challenge detection")
	}
	if reason == "" {
		t.Fatal("expected non-empty challenge reason")
	}
}

func TestDetectBotChallengeSkipsNonHTML(t *testing.T) {
	body := []byte("verify you are human")
	isChallenge, _ := detectBotChallenge(200, "application/json", body)
	if isChallenge {
		t.Fatal("expected non-html body to skip challenge detection")
	}
}
