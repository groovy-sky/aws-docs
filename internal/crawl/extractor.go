package crawl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html"
	urlpkg "net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/groovy-sky/aws-docs/internal/config"
	xhtml "golang.org/x/net/html"
)

type Extractor struct {
	config config.Config
}

func NewExtractor(cfg config.Config) *Extractor {
	return &Extractor{config: cfg}
}

func (e *Extractor) Extract(rawURL string, body []byte) (ExtractedDocument, error) {
	if landing, ok := extractLandingPageFromRawHTML(rawURL, string(body)); ok {
		return landing, nil
	}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return ExtractedDocument{}, fmt.Errorf("parse HTML: %w", err)
	}

	if redirectURL, ok := extractDocumentRedirectURL(document); ok {
		return ExtractedDocument{RedirectURL: redirectURL}, nil
	}

	if landing, ok := extractLandingPageDocument(document, rawURL); ok {
		return landing, nil
	}

	root := e.findMainContent(document)
	if root.Length() == 0 {
		return ExtractedDocument{}, fmt.Errorf("could not find main content")
	}

	cleaned := root.Clone()
	links := extractSelectionLinks(cleaned)
	for _, selector := range e.config.ExcludeSelectors {
		if selector == "" {
			continue
		}
		cleaned.Find(selector).Each(func(_ int, selection *goquery.Selection) {
			selection.Remove()
		})
	}

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

func extractSelectionLinks(selection *goquery.Selection) []string {
	links := make([]string, 0)
	seen := make(map[string]struct{})
	selection.Find("a[href]").Each(func(_ int, anchor *goquery.Selection) {
		href, exists := anchor.Attr("href")
		if !exists {
			return
		}
		href = strings.TrimSpace(href)
		if href == "" {
			return
		}
		if _, ok := seen[href]; ok {
			return
		}
		seen[href] = struct{}{}
		links = append(links, href)
	})
	return links
}

func extractDocumentRedirectURL(document *goquery.Document) (string, bool) {
	var redirectURL string
	document.Find("meta[http-equiv]").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		httpEquiv, exists := selection.Attr("http-equiv")
		if !exists || !strings.EqualFold(strings.TrimSpace(httpEquiv), "refresh") {
			return true
		}

		content, exists := selection.Attr("content")
		if !exists {
			return true
		}

		if target, ok := parseMetaRefreshTarget(content); ok {
			redirectURL = target
			return false
		}

		return true
	})

	if redirectURL == "" {
		return "", false
	}

	return redirectURL, true
}

func parseMetaRefreshTarget(content string) (string, bool) {
	for _, part := range strings.Split(content, ";") {
		trimmed := strings.TrimSpace(part)
		if len(trimmed) < 4 {
			continue
		}
		if !strings.EqualFold(trimmed[:4], "url=") {
			continue
		}

		target := strings.TrimSpace(trimmed[4:])
		target = strings.Trim(target, `"'`)
		if target == "" {
			return "", false
		}
		return target, true
	}

	return "", false
}

var (
	landingPageValuePattern = regexp.MustCompile(`id=["']landing-page-xml["'][^>]*value=["']([^"']+)["']`)
	landingTitlePattern     = regexp.MustCompile(`(?s)<title>\s*(.*?)\s*</title>`)
	landingAbstractPattern  = regexp.MustCompile(`(?s)<abstract>\s*(.*?)\s*</abstract>`)
	landingHrefPattern      = regexp.MustCompile(`href\s*=\s*"([^"]+)"`)
)

type landingPage struct {
	Title    string `xml:"title"`
	Abstract string `xml:"abstract"`
	Sections struct {
		Section []struct {
			Cards struct {
				SimpleCard []struct {
					Href string `xml:"href,attr"`
				} `xml:"simple-card"`
			} `xml:"cards"`
		} `xml:"section"`
	} `xml:"sections"`
	SideNav struct {
		Sections []struct {
			Title string `xml:"title"`
			Links struct {
				SimpleLink []struct {
					Href string `xml:"href,attr"`
					Text string `xml:",chardata"`
				} `xml:"simple-link"`
			} `xml:"links"`
		} `xml:"side-nav-section"`
	} `xml:"side-nav"`
}

func extractLandingPageDocument(document *goquery.Document, rawURL string) (ExtractedDocument, bool) {
	rawValue, exists := document.Find("#landing-page-xml").Attr("value")
	if !exists || strings.TrimSpace(rawValue) == "" {
		return ExtractedDocument{}, false
	}

	decoded, err := urlpkg.QueryUnescape(rawValue)
	if err != nil || strings.TrimSpace(decoded) == "" {
		return ExtractedDocument{}, false
	}

	parsed := landingPage{}
	if err := xml.Unmarshal([]byte(decoded), &parsed); err != nil {
		return ExtractedDocument{}, false
	}

	title := strings.TrimSpace(parsed.Title)
	if title == "" {
		title = strings.TrimSpace(document.Find("title").First().Text())
	}

	canonicalURL := rawURL
	if href, exists := document.Find("link[rel='canonical']").Attr("href"); exists && strings.TrimSpace(href) != "" {
		canonicalURL = strings.TrimSpace(href)
	}

	var builder strings.Builder
	if title != "" {
		builder.WriteString("<h1>")
		builder.WriteString(html.EscapeString(title))
		builder.WriteString("</h1>\n")
	}

	abstract := strings.TrimSpace(parsed.Abstract)
	if abstract != "" {
		builder.WriteString("<p>")
		builder.WriteString(html.EscapeString(abstract))
		builder.WriteString("</p>\n")
	}

	links := make([]string, 0)

	guideLinks := make([]string, 0)
	for _, section := range parsed.Sections.Section {
		for _, card := range section.Cards.SimpleCard {
			href := strings.TrimSpace(card.Href)
			if href != "" {
				guideLinks = append(guideLinks, href)
				links = append(links, href)
			}
		}
	}
	if len(guideLinks) > 0 {
		builder.WriteString("<h2>Guides</h2>\n<ul>\n")
		for _, href := range guideLinks {
			builder.WriteString("<li><a href=\"")
			builder.WriteString(html.EscapeString(href))
			builder.WriteString("\">")
			builder.WriteString(html.EscapeString(href))
			builder.WriteString("</a></li>\n")
		}
		builder.WriteString("</ul>\n")
	}

	for _, section := range parsed.SideNav.Sections {
		sectionTitle := strings.TrimSpace(section.Title)
		if sectionTitle == "" || len(section.Links.SimpleLink) == 0 {
			continue
		}
		builder.WriteString("<h2>")
		builder.WriteString(html.EscapeString(sectionTitle))
		builder.WriteString("</h2>\n<ul>\n")
		for _, link := range section.Links.SimpleLink {
			href := strings.TrimSpace(link.Href)
			if href == "" {
				continue
			}
			text := strings.TrimSpace(link.Text)
			if text == "" {
				text = href
			}
			builder.WriteString("<li><a href=\"")
			builder.WriteString(html.EscapeString(href))
			builder.WriteString("\">")
			builder.WriteString(html.EscapeString(text))
			builder.WriteString("</a></li>\n")
			links = append(links, href)
		}
		builder.WriteString("</ul>\n")
	}

	htmlValue := strings.TrimSpace(builder.String())
	if htmlValue == "" {
		return ExtractedDocument{}, false
	}

	return ExtractedDocument{
		CanonicalURL: canonicalURL,
		Title:        title,
		HTML:         htmlValue,
		Links:        links,
	}, true
}

func extractLandingPageFromRawHTML(rawURL string, htmlBody string) (ExtractedDocument, bool) {
	matches := landingPageValuePattern.FindStringSubmatch(htmlBody)
	if len(matches) < 2 || strings.TrimSpace(matches[1]) == "" {
		return ExtractedDocument{}, false
	}

	decoded, err := urlpkg.QueryUnescape(matches[1])
	if err != nil || strings.TrimSpace(decoded) == "" {
		return ExtractedDocument{}, false
	}

	parsed := landingPage{}
	if err := xml.Unmarshal([]byte(decoded), &parsed); err == nil {
		title := strings.TrimSpace(parsed.Title)
		abstract := strings.TrimSpace(parsed.Abstract)

		links := make([]string, 0)
		for _, section := range parsed.Sections.Section {
			for _, card := range section.Cards.SimpleCard {
				href := strings.TrimSpace(card.Href)
				if href != "" {
					links = append(links, href)
				}
			}
		}
		for _, section := range parsed.SideNav.Sections {
			for _, link := range section.Links.SimpleLink {
				href := strings.TrimSpace(link.Href)
				if href != "" {
					links = append(links, href)
				}
			}
		}

		htmlValue := buildLandingHTML(title, abstract, links)
		if strings.TrimSpace(htmlValue) != "" {
			return ExtractedDocument{
				CanonicalURL: rawURL,
				Title:        title,
				HTML:         htmlValue,
				Links:        links,
			}, true
		}
	}

	title := ""
	if m := landingTitlePattern.FindStringSubmatch(decoded); len(m) > 1 {
		title = strings.TrimSpace(html.UnescapeString(m[1]))
	}
	abstract := ""
	if m := landingAbstractPattern.FindStringSubmatch(decoded); len(m) > 1 {
		abstract = strings.TrimSpace(html.UnescapeString(m[1]))
	}

	linkMatches := landingHrefPattern.FindAllStringSubmatch(decoded, -1)
	links := make([]string, 0, len(linkMatches))
	for _, m := range linkMatches {
		if len(m) < 2 {
			continue
		}
		href := strings.TrimSpace(html.UnescapeString(m[1]))
		if href != "" {
			links = append(links, href)
		}
	}

	htmlValue := buildLandingHTML(title, abstract, links)
	if strings.TrimSpace(htmlValue) == "" {
		return ExtractedDocument{}, false
	}

	return ExtractedDocument{
		CanonicalURL: rawURL,
		Title:        title,
		HTML:         htmlValue,
		Links:        links,
	}, true
}

func buildLandingHTML(title string, abstract string, links []string) string {
	var builder strings.Builder
	if strings.TrimSpace(title) != "" {
		builder.WriteString("<h1>")
		builder.WriteString(html.EscapeString(strings.TrimSpace(title)))
		builder.WriteString("</h1>\n")
	}
	if strings.TrimSpace(abstract) != "" {
		builder.WriteString("<p>")
		builder.WriteString(html.EscapeString(strings.TrimSpace(abstract)))
		builder.WriteString("</p>\n")
	}

	seen := map[string]struct{}{}
	unique := make([]string, 0, len(links))
	for _, link := range links {
		if _, ok := seen[link]; ok {
			continue
		}
		seen[link] = struct{}{}
		unique = append(unique, link)
	}
	if len(unique) > 0 {
		builder.WriteString("<h2>Links</h2>\n<ul>\n")
		for _, href := range unique {
			builder.WriteString("<li><a href=\"")
			builder.WriteString(html.EscapeString(href))
			builder.WriteString("\">")
			builder.WriteString(html.EscapeString(href))
			builder.WriteString("</a></li>\n")
		}
		builder.WriteString("</ul>\n")
	}

	return strings.TrimSpace(builder.String())
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
		if err := xhtml.Render(&builder, node); err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}
