package crawl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/temoto/robotstxt"
)

type Robots struct {
	group *robotstxt.Group
}

func LoadRobots(ctx context.Context, client *http.Client, baseURL string, userAgent string) (*Robots, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/robots.txt", nil)
	if err != nil {
		return nil, fmt.Errorf("build robots request: %w", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch robots.txt: %w", err)
	}
	defer response.Body.Close()

	robotsData, err := robotstxt.FromResponse(response)
	if err != nil {
		return nil, fmt.Errorf("parse robots.txt: %w", err)
	}

	return &Robots{group: robotsData.FindGroup(userAgent)}, nil
}

func (r *Robots) Allowed(path string) bool {
	if r == nil || r.group == nil {
		return true
	}
	return r.group.Test(path)
}
