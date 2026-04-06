# aws-docs

![](logo.png)

This repository contains the first implementation pass of an AWS documentation mirroring pipeline in Go. It can fetch public pages from `https://docs.aws.amazon.com`, extract the main content, convert it to deterministic Markdown, and persist per-page metadata in BoltDB for incremental updates.

## Current scope

- CLI entrypoint with `incremental`, `full`, and `refresh-url` modes.
- HTTP fetching with rate limiting, retry, conditional requests, and `robots.txt` checks.
- Content extraction using configurable selectors.
- HTML to Markdown conversion and internal link rewriting.
- Deterministic URL to repository path mapping.
- Metadata persistence for page state and discovered links.

## Run

```bash
go run ./cmd/aws-docs -config config.example.json -mode incremental
```

To refresh a single page:

```bash
go run ./cmd/aws-docs -config config.example.json -mode refresh-url -url https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html
```

Mirrored output is written under `docs/`, and crawl state is stored in `metadata/crawl.db`.