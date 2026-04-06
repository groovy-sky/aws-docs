package crawl

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	markdown "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/groovy-sky/aws-docs/internal/config"
)

var blankLinePattern = regexp.MustCompile(`\n{3,}`)

type Converter struct {
	config config.Config
	mapper *Mapper
	exists func(string) bool
}

func NewConverter(cfg config.Config, mapper *Mapper, exists func(string) bool) *Converter {
	return &Converter{config: cfg, mapper: mapper, exists: exists}
}

func (c *Converter) Convert(document ExtractedDocument, sourceURL string) (MarkdownDocument, error) {
	rewrittenHTML, err := c.rewriteHTML(document.HTML, sourceURL)
	if err != nil {
		return MarkdownDocument{}, err
	}

	converter := markdown.NewConverter("", true, nil)
	markdownText, err := converter.ConvertString(rewrittenHTML)
	if err != nil {
		return MarkdownDocument{}, fmt.Errorf("convert HTML to Markdown: %w", err)
	}

	normalized := normalizeMarkdown(markdownText)
	hash := sha256.Sum256([]byte(normalized))

	return MarkdownDocument{
		Markdown:    normalized,
		ContentHash: hex.EncodeToString(hash[:]),
	}, nil
}

func (c *Converter) rewriteHTML(fragment string, sourceURL string) (string, error) {
	document, err := goquery.NewDocumentFromReader(bytes.NewBufferString("<div id='root'>" + fragment + "</div>"))
	if err != nil {
		return "", fmt.Errorf("parse HTML fragment: %w", err)
	}

	document.Find("a[href]").Each(func(_ int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		rewritten := c.rewriteHref(sourceURL, strings.TrimSpace(href))
		selection.SetAttr("href", rewritten)
	})

	document.Find("img[src]").Each(func(_ int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		rewritten := c.rewriteAssetURL(sourceURL, strings.TrimSpace(src))
		selection.SetAttr("src", rewritten)
	})

	root := document.Find("#root")
	htmlValue, err := root.Html()
	if err != nil {
		return "", fmt.Errorf("render rewritten HTML: %w", err)
	}
	return htmlValue, nil
}

func (c *Converter) rewriteHref(sourceURL string, href string) string {
	if href == "" || strings.HasPrefix(href, "#") {
		return href
	}

	resolved, err := ResolveURL(sourceURL, href, c.config)
	if err != nil {
		return href
	}
	if !IsAllowedURL(resolved, c.config) {
		return resolved
	}

	anchor := ""
	if parsed, err := url.Parse(href); err == nil && parsed.Fragment != "" {
		anchor = "#" + parsed.Fragment
	}

	targetPath := c.mapper.RepoPath(resolved)
	if c.exists == nil || !c.exists(targetPath) {
		return resolved + anchor
	}

	return c.mapper.RelativeLink(sourceURL, resolved) + anchor
}

func (c *Converter) rewriteAssetURL(sourceURL string, value string) string {
	if value == "" {
		return value
	}
	resolved, err := ResolveURL(sourceURL, value, c.config)
	if err != nil {
		return value
	}
	return resolved
}

func normalizeMarkdown(value string) string {
	lines := strings.Split(strings.ReplaceAll(value, "\r\n", "\n"), "\n")
	for index, line := range lines {
		lines[index] = strings.TrimRight(line, " \t")
	}
	value = strings.Join(lines, "\n")
	value = blankLinePattern.ReplaceAllString(value, "\n\n")
	value = strings.TrimSpace(value) + "\n"
	return value
}
