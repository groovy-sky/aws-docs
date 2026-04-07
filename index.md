---
title: AWS Docs Mirror
---

# AWS Documentation Mirror

A continuously updated mirror of public [AWS technical documentation](https://docs.aws.amazon.com) — crawled, converted to clean Markdown, and committed to Git.

**70+ services** · Written in **Go** · Incremental state via **BoltDB** · Automated via **GitHub Actions**

---

## Browse

See the full [Services index](SERVICES.md) for all mirrored services.

Documentation is organized under `docs/`:

| Folder | Contents |
|--------|----------|
| `docs/services/` | User guides and developer guides per service |
| `docs/reference/` | API references and SDK docs |
| `docs/general/` | Cross-service and general AWS docs |

---

## How it works

```
docs.aws.amazon.com
        │
        ▼
  HTTP Fetcher ── robots.txt / rate limiting / ETag caching
        │
        ▼
Content Extractor ── strips nav, header, footer
        │
        ▼
HTML → Markdown ─── deterministic, Git-friendly output
        │
        ▼
  Path Mapper ────── URL → docs/{section}/{service}/...
        │
        ▼
   BoltDB Store ──── per-URL metadata for incremental runs
        │
        ▼
   docs/ in Git ──── committed, diffable, searchable
```

---

## Run it yourself

```bash
# Clone
git clone https://github.com/groovy-sky/aws-docs.git && cd aws-docs

# Small partial crawl (~25 sections)
go run ./cmd/aws-docs -mode partial -max-sections 25

# Refresh a single page
go run ./cmd/aws-docs -mode refresh-url \
  -url https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html

# Full crawl
go run ./cmd/aws-docs -mode full
```

Crawl state is stored in `metadata/crawl.db`. On first run, seeds are derived from `robots.txt` and persisted for incremental updates.

---

*Source: [docs.aws.amazon.com](https://docs.aws.amazon.com) · Not affiliated with or endorsed by Amazon Web Services.*
