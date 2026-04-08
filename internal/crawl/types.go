package crawl

type FetchResult struct {
	FinalURL              string
	StatusCode            int
	ContentType           string
	Body                  []byte
	ETag                  string
	LastModified          string
	NotModified           bool
	PotentialBotChallenge bool
	BotChallengeReason    string
	RequestFailed         error
}

type ExtractedDocument struct {
	CanonicalURL string
	Title        string
	HTML         string
	Links        []string
	RedirectURL  string
}

type MarkdownDocument struct {
	Markdown    string
	ContentHash string
}

type RunOptions struct {
	Mode        string
	URL         string
	MaxSections int
}
