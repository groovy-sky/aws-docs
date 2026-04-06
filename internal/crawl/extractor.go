package crawl

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/groovy-sky/aws-docs/internal/config"
	"golang.org/x/net/html"
)

type Extractor struct {
	config config.Config
}

func NewExtractor(cfg config.Config) *Extractor {
	return &Extractor{config: cfg}
}

func (e *Extractor) Extract(rawURL string, body []byte) (ExtractedDocument, error) {
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return ExtractedDocument{}, fmt.Errorf("parse HTML: %w", err)
	}

	root := e.findMainContent(document)
	if root.Length() == 0 {
		return ExtractedDocument{}, fmt.Errorf("could not find main content")
	}

	cleaned := root.Clone()
	for _, selector := range e.config.ExcludeSelectors {
		if selector == "" {
			continue
		}
		cleaned.Find(selector).Each(func(_ int, selection *goquery.Selection) {
			selection.Remove()
		})
	}

	links := make([]string, 0)
	cleaned.Find("a[href]").Each(func(_ int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")
		if exists {
			links = append(links, strings.TrimSpace(href))
		}
	})

	title := strings.TrimSpace(cleaned.Find("h1").First().Text())
	if title == "" {
		title = strings.TrimSpace(document.Find("title").First().Text())
	}

	canonicalURL := rawURL
	if href, exists := document.Find("link[rel='canonical']").Attr("href"); exists && strings.TrimSpace(href) != "" {
		canonicalURL = strings.TrimSpace(href)
	}

	htmlValue, err := renderSelectionHTML(cleaned)
	if err != nil {
		return ExtractedDocument{}, fmt.Errorf("render cleaned HTML: %w", err)
	}

	return ExtractedDocument{
		CanonicalURL: canonicalURL,
		Title:        title,
		HTML:         htmlValue,
		Links:        links,
	}, nil
}

func (e *Extractor) findMainContent(document *goquery.Document) *goquery.Selection {
	for _, selector := range e.config.Selectors {
		selection := document.Find(selector).First()
		if selection.Length() > 0 && strings.TrimSpace(selection.Text()) != "" {
			return selection
		}
	}

	best := document.Find("body").First()
	bestSize := 0
	document.Find("main, article, section, div").Each(func(_ int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())
		if text == "" {
			return
		}
		score := len(text)
		if selection.Find("h1").Length() > 0 {
			score += 500
		}
		if score > bestSize {
			best = selection
			bestSize = score
		}
	})
	return best
}

func renderSelectionHTML(selection *goquery.Selection) (string, error) {
	var builder strings.Builder
	for _, node := range selection.Nodes {
		if err := html.Render(&builder, node); err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}
