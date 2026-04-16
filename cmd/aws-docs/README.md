# aws-docs CLI

This directory contains the entrypoint for the AWS documentation mirroring pipeline. The CLI fetches public pages from `https://docs.aws.amazon.com`, extracts the main content, converts it to deterministic Markdown, and persists per-page metadata in BoltDB for incremental updates.

## Current scope

- CLI entrypoint with `partial`, `incremental`, `full`, `refresh-url`, `section`, and `scheduled` modes.
- HTTP fetching with rate limiting, retry, conditional requests, and `robots.txt` checks.
- Content extraction using configurable selectors.
- HTML to Markdown conversion and internal link rewriting.
- Deterministic URL to repository path mapping.
- Metadata persistence for page state, discovered links, and robots-derived seeds.

## Run

The crawler can run without any config file.

Seed flow:
- On first run, it uses `robots.txt` to derive crawl seeds and stores them in `metadata/crawl.db`.
- Later runs reuse the stored seeds.
- `full` mode additionally expands from sitemap page URLs for a deeper crawl.

1. Run a small partial crawl from stored robots-derived seeds:

```bash
go run ./cmd/aws-docs -mode partial -max-sections 25
```

2. Run a broader partial crawl:

```bash
go run ./cmd/aws-docs -mode partial -max-sections 500
```

3. Run a full crawl:

```bash
go run ./cmd/aws-docs -mode full
```

4. Refresh a single page:

```bash
go run ./cmd/aws-docs -mode refresh-url -url https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html
```

5. Optional: run with a custom config file:

```bash
go run ./cmd/aws-docs -config config.json -mode incremental
```

6. Crawl a single section:

```bash
go run ./cmd/aws-docs -mode section -name vpc
```

7. Crawl one scheduled section (for hourly cron):

```bash
go run ./cmd/aws-docs -mode scheduled
```

Notes:
- `partial` is the default mode and is currently an alias for incremental behavior from stored seeds.
- In `partial`/`incremental` mode, `-max-sections` limits how many discovered sections are processed in a run.
- In `section` mode, `-name` is required and must be a top-level section name (for example `vpc`, `iam`, `lambda`).
- In `scheduled` mode, crawler computes an hourly slot from UTC day-of-month and hour (`(day-1)*24+hour`) and selects one discovered section by `slot % section_count`.
- In `scheduled` mode, section ordering is deterministic (sorted by section key), so queue selection remains stable across runs while still adapting to added/removed services.
- CLI `-max-sections` overrides config only when value is greater than `0`.
- To run effectively unbounded in `partial`/`incremental` mode with a config file, set `"max_sections": 0` and do not pass `-max-sections`.
- Config file values `"include_path_patterns": []` and `"exclude_path_patterns": []` are respected as explicit empty filters (no path includes/excludes).

Mirrored output is written under `docs/`, and crawl state is stored in `metadata/crawl.db`.

On every `./cmd/aws-docs` invocation, a root `SERVICES.md` file is updated with run status metadata and links to currently parsed services.
