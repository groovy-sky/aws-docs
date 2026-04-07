package crawl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/groovy-sky/aws-docs/internal/config"
	"golang.org/x/time/rate"
)

type Fetcher struct {
	client     *http.Client
	limiter    *rate.Limiter
	userAgent  string
	maxRetries int
}

func NewFetcher(cfg config.Config) *Fetcher {
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          cfg.Concurrency * 4,
		MaxIdleConnsPerHost:   cfg.Concurrency * 2,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: time.Duration(cfg.RequestTimeoutSeconds) * time.Second,
	}

	return &Fetcher{
		client: &http.Client{
			Timeout:   time.Duration(cfg.RequestTimeoutSeconds) * time.Second,
			Transport: transport,
		},
		limiter:    rate.NewLimiter(rate.Limit(cfg.RequestsPerSecond), cfg.Burst),
		userAgent:  cfg.UserAgent,
		maxRetries: cfg.MaxRetries,
	}
}

func (f *Fetcher) HTTPClient() *http.Client {
	return f.client
}

func (f *Fetcher) Fetch(ctx context.Context, rawURL string) (FetchResult, error) {
	var result FetchResult
	operation := func() error {
		if err := f.limiter.Wait(ctx); err != nil {
			return backoff.Permanent(err)
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("build request: %w", err))
		}
		request.Header.Set("User-Agent", f.userAgent)
		request.Header.Set("Accept", "text/html,application/xhtml+xml")

		response, err := f.client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		result.StatusCode = response.StatusCode
		result.FinalURL = response.Request.URL.String()
		result.ContentType = response.Header.Get("Content-Type")
		result.ETag = response.Header.Get("ETag")
		result.LastModified = response.Header.Get("Last-Modified")

		if response.StatusCode == http.StatusNotModified {
			result.NotModified = true
			return nil
		}

		if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
			if response.StatusCode >= http.StatusInternalServerError || response.StatusCode == http.StatusTooManyRequests {
				return fmt.Errorf("transient HTTP status %d", response.StatusCode)
			}
			return backoff.Permanent(fmt.Errorf("unexpected HTTP status %d", response.StatusCode))
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("read response body: %w", err)
		}
		result.Body = body
		return nil
	}

	exponentialBackoff := backoff.NewExponentialBackOff()
	exponentialBackoff.MaxElapsedTime = time.Duration(f.maxRetries) * 2 * time.Second
	exponentialBackoff.Reset()
	retryPolicy := backoff.WithContext(exponentialBackoff, ctx)

	err := backoff.Retry(operation, backoff.WithMaxRetries(retryPolicy, uint64(f.maxRetries)))
	if err != nil {
		result.RequestFailed = err
		return result, err
	}

	if !strings.Contains(strings.ToLower(result.ContentType), "html") && !result.NotModified {
		return result, fmt.Errorf("unsupported content type %q", result.ContentType)
	}

	return result, nil
}
