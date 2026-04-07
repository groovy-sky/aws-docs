# Application Logic Reference

This document stores the current runtime logic of the crawler so future work can reuse the same mental model quickly.

## Sync Contract

- This file is the source of truth for crawler behavior.
- Any behavior change in crawler-related code must update this file in the same change.
- Crawler-related code includes these areas:
	- `internal/crawl/`
	- `internal/app/run.go`
	- `internal/config/config.go`
	- `internal/store/store.go`
	- `internal/write/services_index.go` when behavior depends on crawl output
- No fixed changelog section format is required; update the most relevant section directly.

## Runtime Entry and Orchestration

- CLI entrypoint is `cmd/aws-docs/main.go`.
- Flags are mapped into `app.Options` and passed to `app.Run`.
- Mode behavior:
	- `partial` is normalized to `incremental`.
	- `refresh-url` requires a single URL.
- `app.Run` constructs and wires these components in order:
	- `config.Load`
	- `store.Open`
	- `crawl.NewFetcher`
	- `crawl.LoadRobots`
	- `crawl.NewMapper`
	- `crawl.NewExtractor`
	- `crawl.NewConverter`
	- `crawl.NewCrawler`
- At the end of every run (success or failure), `write.WriteServicesIndex` updates root `SERVICES.md` with run metadata and discovered service links.

## Crawl Pipeline

High-level flow in `Crawler.Run`:

1. Build queue from `seedQueue`.
2. Apply section cap when mode is incremental-style and `max_sections > 0`.
3. Process queue breadth-first while deduplicating with an in-memory `seen` set.
4. Skip discovered URLs not allowed by include/exclude filters or `robots.txt`.
5. Emit run-summary logs with selected section count and processed URL count.

URL processing (`processURL`) sequence:

1. Validate robots rule.
2. Fetch page via `Fetcher.Fetch` (rate-limited with retry).
3. Extract content via `Extractor.Extract`.
4. If extractor returns `RedirectURL` (meta refresh wrapper pages), resolve target and enqueue it instead of writing markdown.
5. Normalize canonical URL.
6. Convert cleaned HTML to markdown using `Converter.Convert`.
7. Map canonical URL to repository path via `Mapper.RepoPath`.
8. Write markdown file.
9. Resolve extracted links and return discovered allowed URLs.

## Extraction Rules

`Extractor.Extract` chooses the first matching strategy:

1. Raw landing-page XML extraction from embedded `landing-page-xml` value.
2. DOM parse and meta refresh redirect detection (`extractDocumentRedirectURL`).
3. Structured landing-page extraction from parsed DOM.
4. Generic main-content extraction with configured selectors and fallback scoring.

Current redirect handling:

- Supports `<meta http-equiv="refresh" content="...;URL=target">` patterns.
- Returns `ExtractedDocument{RedirectURL: target}`.
- Avoids writing empty wrapper pages that only forward to real docs.

## Conversion and Linking

- `Converter` rewrites in-domain documentation links to local relative markdown links.
- Conversion uses `html-to-markdown` then normalizes whitespace for deterministic output.
- Content hash is SHA-256 of normalized markdown.

## URL to Path Mapping

- Mapping classifies URLs into `docs/services`, `docs/reference`, `docs/reference/cli`, or `docs/general`.
- Service bucket uses first path segment with specific overrides (for example `AWSEC2 -> ec2`).
- Output filenames are slugified and normalized to `.md`.
- If URL points to a section root ending at `latest/`, mapper writes an `index.md` path under that section.

## Seed and Metadata Behavior

- Metadata DB currently stores only seed URLs in BoltDB bucket `seeds`.
- On open, legacy buckets `pages` and `links` are removed.
- `seedQueue` prioritizes stored seeds, canonicalizes, deduplicates, and prunes unreachable seeds.
- In `incremental` and `partial` mode, when requested `max_sections` exceeds the section count represented by stored seeds, queue building backfills from configured seeds to reach the requested section window when possible.
- If no stored seeds exist, crawler discovers from robots-derived structure and sitemaps, then persists.
- In `full` mode, sitemap URLs are added to queue for broader coverage.
- In `incremental` and `partial` mode with `max_sections > 0`, section selection rotates across runs using a persisted cursor in metadata state (`section_cursor`) so each run advances to the next section window instead of always reusing the first sections.

## Practical Maintenance Notes

- If service landing pages list links but child markdown files are missing, inspect redirect-wrapper handling first.
- For crawl coverage issues, inspect `IsAllowedURL`, include/exclude path patterns, and section capping behavior.
- For wrong output location, inspect `Mapper.RepoPath` and `mapServiceCode` mappings.
- For root service listing issues, inspect `internal/write/services_index.go` link selection behavior.
