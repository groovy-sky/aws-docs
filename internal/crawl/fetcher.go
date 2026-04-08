package crawl

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/groovy-sky/aws-docs/internal/config"
	"golang.org/x/time/rate"
)

type Fetcher struct {
	client          *http.Client
	limiter         *rate.Limiter
	userAgent       string
	maxRetries      int
	minDelay        time.Duration
	maxDelay        time.Duration
	random          *rand.Rand
	detailedLogging bool
}

func NewFetcher(cfg config.Config) *Fetcher {
	jar, _ := cookiejar.New(nil)

	requestTimeout := time.Duration(cfg.RequestTimeoutSeconds) * time.Second
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          cfg.Concurrency * 4,
		MaxIdleConnsPerHost:   cfg.Concurrency * 2,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: requestTimeout,
		ForceAttemptHTTP2:     true,
	}

	minDelay := time.Duration(cfg.MinRequestDelayMS) * time.Millisecond
	maxDelay := time.Duration(cfg.MaxRequestDelayMS) * time.Millisecond

	return &Fetcher{
		client: &http.Client{
			Timeout:   requestTimeout,
			Transport: transport,
			Jar:       jar,
			CheckRedirect: func(request *http.Request, via []*http.Request) error {
				if len(via) >= 10 {
					return fmt.Errorf("stopped after %d redirects", len(via))
				}
				if cfg.DetailedLogging && len(via) > 0 {
					log.Printf("fetch redirect: from=%s to=%s hop=%d", via[len(via)-1].URL.String(), request.URL.String(), len(via))
				}

				request.Header.Set("User-Agent", cfg.UserAgent)
				request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
				request.Header.Set("Accept-Language", "en-US,en;q=0.9")
				request.Header.Set("Upgrade-Insecure-Requests", "1")
				return nil
			},
		},
		limiter:         rate.NewLimiter(rate.Limit(cfg.RequestsPerSecond), cfg.Burst),
		userAgent:       cfg.UserAgent,
		maxRetries:      cfg.MaxRetries,
		minDelay:        minDelay,
		maxDelay:        maxDelay,
		random:          rand.New(rand.NewSource(time.Now().UnixNano())),
		detailedLogging: cfg.DetailedLogging,
	}
}

func (f *Fetcher) HTTPClient() *http.Client {
	return f.client
}

func (f *Fetcher) Fetch(ctx context.Context, rawURL string) (FetchResult, error) {
	var result FetchResult
	attempt := 0
	operation := func() error {
		attempt++
		if err := f.limiter.Wait(ctx); err != nil {
			return backoff.Permanent(err)
		}
		if err := f.waitWithJitter(ctx); err != nil {
			return backoff.Permanent(err)
		}
		f.logf("fetch start: url=%s attempt=%d", rawURL, attempt)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("build request: %w", err))
		}
		request.Header.Set("User-Agent", f.userAgent)
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Set("Accept-Language", "en-US,en;q=0.9")
		request.Header.Set("Upgrade-Insecure-Requests", "1")

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
		f.logf("fetch response: url=%s final-url=%s status=%d content-type=%q", rawURL, result.FinalURL, response.StatusCode, result.ContentType)

		if response.StatusCode == http.StatusNotModified {
			result.NotModified = true
			f.logf("fetch not-modified: url=%s", rawURL)
			return nil
		}

		if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
			if shouldRetryHTTPStatus(response.StatusCode) {
				return fmt.Errorf("transient HTTP status %d", response.StatusCode)
			}
			return backoff.Permanent(fmt.Errorf("unexpected HTTP status %d", response.StatusCode))
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("read response body: %w", err)
		}

		if isChallenge, reason := detectBotChallenge(response.StatusCode, response.Header.Get("Content-Type"), body); isChallenge {
			result.PotentialBotChallenge = true
			result.BotChallengeReason = reason
			f.logf("fetch challenge-detected: url=%s final-url=%s reason=%s", rawURL, result.FinalURL, reason)
			return fmt.Errorf("potential bot challenge page detected: %s", reason)
		}

		result.Body = body
		f.logf("fetch success: url=%s final-url=%s bytes=%d", rawURL, result.FinalURL, len(body))
		return nil
	}

	exponentialBackoff := backoff.NewExponentialBackOff()
	exponentialBackoff.MaxElapsedTime = time.Duration(f.maxRetries) * 2 * time.Second
	exponentialBackoff.Reset()
	retryPolicy := backoff.WithContext(exponentialBackoff, ctx)

	err := backoff.RetryNotify(operation, backoff.WithMaxRetries(retryPolicy, uint64(f.maxRetries)), func(err error, next time.Duration) {
		f.logf("fetch retry: url=%s attempt=%d error=%v next-backoff=%s", rawURL, attempt, err, next)
	})
	if err != nil {
		result.RequestFailed = err
		f.logf("fetch failed: url=%s attempts=%d error=%v", rawURL, attempt, err)
		return result, err
	}

	if !strings.Contains(strings.ToLower(result.ContentType), "html") && !result.NotModified {
		return result, fmt.Errorf("unsupported content type %q", result.ContentType)
	}

	return result, nil
}

func (f *Fetcher) logf(format string, args ...any) {
	if !f.detailedLogging {
		return
	}
	log.Printf(format, args...)
}

func (f *Fetcher) waitWithJitter(ctx context.Context) error {
	if f.maxDelay <= 0 {
		return nil
	}

	delay := f.maxDelay
	if f.maxDelay > f.minDelay {
		delay = f.minDelay + time.Duration(f.random.Int63n(int64(f.maxDelay-f.minDelay)+1))
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func shouldRetryHTTPStatus(statusCode int) bool {
	if statusCode == http.StatusForbidden || statusCode == http.StatusTooManyRequests || statusCode == http.StatusServiceUnavailable {
		return true
	}
	return statusCode >= http.StatusInternalServerError
}

func detectBotChallenge(statusCode int, contentType string, body []byte) (bool, string) {
	if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
		return false, ""
	}
	if !strings.Contains(strings.ToLower(contentType), "html") {
		return false, ""
	}

	lowerBody := strings.ToLower(string(body))
	challengeSignals := []string{
		"captcha",
		"verify you are human",
		"attention required",
		"are you a robot",
		"bot detection",
		"challenge-platform",
		"cf-challenge",
		"security check",
	}

	for _, signal := range challengeSignals {
		if strings.Contains(lowerBody, signal) {
			return true, signal
		}
	}

	return false, ""
}
