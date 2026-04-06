# PLAN: Mirroring AWS Technical Documentation to GitHub (Markdown)

## Introduction

This document describes a technical plan for automatically mirroring the public AWS technical documentation from `https://docs.aws.amazon.com` into a GitHub repository as Markdown files.

It is intended for engineers who will design, implement, and operate this mirroring system. The plan defines the overall architecture, tools, and processes needed to:

- Crawl and extract the main content from AWS documentation pages.
- Convert that content into clean, deterministic Markdown.
- Organize it in a logical, Git‑friendly directory structure.
- Keep the mirror up to date through incremental, automated updates (e.g., via CI).

The goal is to provide a concise yet complete blueprint that can be used to build a production‑quality documentation mirroring pipeline in Go.

---

## 1. High‑level System Architecture

**Objective:** A robust, incremental mirroring pipeline from `https://docs.aws.amazon.com` to a GitHub repo as Markdown.

**Components:**

1. **Crawler / Scheduler**
   - Discovers and schedules doc URLs.
   - Respects `robots.txt` and domain rules.
   - Persists frontier and visited state (resumable).

2. **Fetcher**
   - HTTP GET with timeouts, rate limiting, and retry.
   - Tracks `ETag`, `Last-Modified`, and status codes.

3. **Content Extractor**
   - Parses HTML.
   - Selects main documentation content (guide content, API reference text).
   - Strips navigation, header, footer, and non-doc UI.

4. **HTML → Markdown Converter**
   - Converts cleaned HTML to Markdown.
   - Preserves headings, code, tables, lists, links, images.
   - Normalizes output for deterministic Git diffs.

5. **Path Mapper**
   - Deterministic mapping from URL → repo path.
   - Logical grouping into `services/`, `general/`, `reference/`.
   - Supports internal link rewriting.

6. **Metadata Store / Crawl Index**
   - Embedded key-value DB (BoltDB/Badger).
   - Per‑URL metadata: `RepoPath`, `LastFetched`, `ContentHash`, `ETag`, `Last-Modified`, status, error.

7. **Git Integration**
   - Stages changed Markdown files.
   - Commits with descriptive messages.
   - Optional push to remote.

8. **CLI Orchestrator**
   - Modes: full crawl, incremental update, single‑URL refresh.
   - Configurable via file/flags.

9. **CI/CD Runner**
   - Scheduled (e.g., daily) GitHub Action or equivalent.
   - Runs incremental crawl and commits updates.

---

## 2. Crawling Strategy

**Discovery:**

- Start from a small set of **seed URLs**, e.g.:
  - `https://docs.aws.amazon.com/`
  - Known service guides (EC2, S3, IAM) and major references (CLI, SDKs).
- Optionally consume **sitemaps** if AWS exposes them under docs (check `robots.txt` for sitemap hints).
- Follow internal links:
  - Only within `docs.aws.amazon.com`.
  - Filter by path patterns associated with technical docs:
    - Service prefixes (e.g., `AWSEC2`, `AmazonS3`, `IAM`).
    - Guide and reference segments (`UserGuide`, `DeveloperGuide`, `APIReference`, `getting-started`, `ug`, `dg`).
  - Exclude obvious non-doc URLs (marketing, console, marketplace).

**Normalization & Deduplication:**

- Normalize URLs by:
  - Lowercasing host.
  - Removing fragments.
  - Normalizing trailing slashes consistently.
  - Dropping known tracking query parameters.
- Maintain a **persistent visited set** and frontier:
  - Avoid re‑enqueuing canonicalized URLs.
  - Persist in the embedded DB for resumability.

**Robots & Rate Limits:**

- Parse `https://docs.aws.amazon.com/robots.txt`.
- Respect disallow rules; restrict crawler to allowed paths.
- Apply rate limiting per host:
  - Conservative default (e.g., 1–3 requests/sec).
  - Configurable concurrency; tuned to avoid throttling.
- Add small random jitter to access patterns.

**Resumability & Fault Tolerance:**

- Frontier and metadata are persistent.
- On restart, resume from previous state.
- Track error counts and avoid hot‑looping on failing URLs.

---

## 3. Content Extraction Pipeline

**HTML Parsing:**

- Use a DOM parser (Go HTML + goquery) to work with AWS docs structure.

**Main Content Selection:**

- Most AWS docs pages embed the core text in a central container; common patterns:
  - IDs/classes like `#main-content`, `.awsdocs-content`, `.awsdocs-body`, `main`, `article`.
- Selection strategy:
  - Prefer known selectors from configuration (maintained and easily updated).
  - Fallback to “largest content block containing `<h1>` and substantive text”.

**Cleaning:**

- Remove non‑content elements from the selected root:
  - Global nav, sidebars, breadcrumbs.
  - Headers, footers, cookie banners, search forms.
  - TOCs if they are separate from the actual headings and not useful in Markdown.
  - `script`, `style`, `noscript`, iframes, tracking widgets.
- Keep:
  - Headings, paragraphs, lists.
  - Code examples and inline code.
  - Tables.
  - Links and images (with path rewriting).

**Normalization:**

- Decode HTML entities.
- Normalize whitespace while preserving code formatting.
- Ensure consistent spacing between sections and headings.

---

## 4. HTML → Markdown Conversion Approach

**Tooling:**

- Use a dedicated Go HTML→Markdown library that:
  - Is actively maintained.
  - Supports custom rules and hooks.
- A validated choice is `github.com/JohannesKaufmann/html-to-markdown`, which:
  - Works on Go’s standard HTML parser.
  - Allows customized handling of tags (e.g., tables, code).

**Conversion Rules:**

- Headings: map `h1`–`h6` to `#`–`######`.
- Code:
  - Convert block code to fenced blocks, including language when available from classes.
  - Convert inline `<code>` to backticks.
- Tables: generate Markdown tables with consistent column alignment.
- Lists: stable bullet style and numbering.
- Links:
  - Convert internal links based on the URL → path mapping.
  - Leave external links as absolute URLs.
- Images:
  - Convert to Markdown image syntax, with optional local mirroring or direct HTTPS URLs.

**Deterministic Output:**

- Ensure:
  - No trailing spaces.
  - Consistent line endings.
  - Stable heading and list formatting.
  - No automatic wrapping that changes between runs.
- Optional lightweight Markdown formatter step to enforce style.

---

## 5. Git Repository Structure

**Top-level layout:**

- `PLAN.md`, `README.md`: documentation of the mirror and process.
- `docs/`: mirrored Markdown content.
  - `services/`: per‑service manuals and guides.
  - `general/`: cross‑service concepts and getting-started docs.
  - `reference/`: CLI, API, SDK references.
  - `legacy/` or `misc/`: pages that don’t fit cleanly.
- `metadata/`:
  - Crawler configuration.
  - Embedded DB for crawl metadata (may or may not be committed).

Example logical structure inside `docs/`:

- `docs/services/ec2/...`
- `docs/services/s3/...`
- `docs/general/security/...`
- `docs/reference/cli/...`
- `docs/reference/api/...`

**URL → Path Mapping:**

- Strip protocol/host, then map by prefix:
  - Map known top-level product/path codes (e.g., `AWSEC2`, `AmazonS3`, `IAM`) to service directories.
  - Map generic guides (e.g., `general/latest/`, `aws-security`, `whitepapers`) to `general/`.
  - Map CLI/SDK reference paths to `reference/cli`, `reference/sdk` etc.
- Remove `.html` and use `.md`.
- Normalize `index` pages.
- Ensure mapping is stable and driven by a configurable mapping file to accommodate AWS naming patterns.

**Internal Link Rewriting:**

- Maintain a mapping from canonical URL to repo path.
- For each internal link:
  - Normalize the URL.
  - Look up its Markdown path.
  - Replace with a relative link from the current page.
- Preserve existing anchors (`#section`) when present.

---

## 6. Incremental Update Mechanism

**Per‑Page Metadata:**

- For each URL, store:
  - Canonical URL.
  - Repo path.
  - Last fetch time.
  - HTTP status.
  - `ETag` / `Last-Modified` if provided.
  - `ContentHash` of the cleaned HTML or final Markdown.

**Change Detection:**

- Fetch pages with conditional requests when possible (`If-None-Match`, `If-Modified-Since`).
- If server returns `304 Not Modified`, skip processing.
- If new content is fetched:
  - Compute hash of the cleaned content or final Markdown.
  - If hash unchanged vs stored, skip file write and Git changes.
  - If changed, overwrite Markdown file and update metadata.

**Scope Control:**

- Support:
  - Full recrawl (reset metadata or ignore historical timestamps).
  - Incremental (default): only act on new or changed pages.
  - Limits per run (max pages or time), suitable for CI.

---

## 7. Suggested Go Libraries

All of the following are widely used and suitable for production as of 2026:

- **HTTP & rate limiting**
  - `net/http` and `http.Transport` with timeouts, keep‑alive.
  - `golang.org/x/time/rate` for client‑side rate limiting.
- **Retry / backoff**
  - `github.com/cenkalti/backoff/v4` for exponential backoff on transient errors.
- **HTML parsing**
  - `github.com/PuerkitoBio/goquery` over `golang.org/x/net/html` for DOM traversal.
- **HTML → Markdown**
  - `github.com/JohannesKaufmann/html-to-markdown` for configurable conversion.
- **Embedded DB for metadata**
  - `go.etcd.io/bbolt` (BoltDB) or `github.com/dgraph-io/badger/v4`.
- **Sitemap parsing (if used)**
  - Small sitemap parsing library or custom XML parsing for standard sitemap XML.
- **Git integration**
  - `github.com/go-git/go-git/v5` for staging, committing, and pushing without external git binary.
- **Concurrency**
  - `golang.org/x/sync/errgroup` plus channels for worker pools.

---

## 8. Example Workflow (Pipeline Overview)

1. **Initialization**
   - Load configuration and mapping rules.
   - Open metadata DB and Git repository.
   - Load and interpret `robots.txt`.
   - Initialize frontier with seed URLs if needed.

2. **Crawling**
   - Run a worker pool consuming URLs from the frontier.
   - For each URL:
     - Enforce rate limiting.
     - Fetch with conditional headers and retry logic.
     - Update metadata status/error for failures.

3. **Extraction & Conversion**
   - Parse HTML and select main content region using configured selectors.
   - Strip non-doc elements.
   - Normalize HTML and pass to HTML→Markdown converter.
   - Normalize Markdown for deterministic output.
   - Compute content hash and compare with previous.

4. **Storage & Link Handling**
   - Map URL to repo path.
   - Write updated Markdown files only for changed content.
   - Record URL → path mapping in metadata for link rewriting.
   - Optionally run a second pass or in-place transform to rewrite internal links, once mappings are available.

5. **Git Operations**
   - Detect changed files.
   - Stage and commit under a consistent message format (e.g., `chore: sync aws docs (YYYY-MM-DD)`).
   - Optionally push to remote.

6. **CI/CD Execution**
   - Triggered on schedule or manual dispatch.
   - Runs incremental crawl with configured limits (e.g., time or page count).
   - Commits and pushes changes if any, else exits cleanly.

---

## 9. CI/CD Automation Design

**Runner (e.g., GitHub Actions):**

- Triggers:
  - Scheduled (e.g., daily at a fixed UTC time).
  - Manual dispatch for ad-hoc updates.
- Steps:
  - Checkout repository.
  - Set up Go environment.
  - Build the crawler CLI.
  - Run crawler in incremental mode with:
    - Config file checked into repo.
    - Rate limits tuned for CI runtime.
  - Configure Git user for automated commits.
  - If Markdown under `docs/` changed:
    - Commit and push.

**State Handling:**

- Prefer keeping the metadata DB in the repo or as a CI artifact:
  - Keeps incremental behavior effective across runs.
  - Without it, the crawler must rediscover content each run and will be slower.

**Operational Safeguards:**

- Max runtime safeguards (e.g., CI job timeout).
- Configurable limits on per‑run processed pages.
- Logging for:
  - URLs skipped due to robots.
  - Non‑200/304 responses.
  - Pages with ambiguous content extraction (for later tuning).
