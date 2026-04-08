# aws-docs

This repository contains a Go-based AWS documentation mirroring pipeline. It fetches public pages from `https://docs.aws.amazon.com`, extracts the main content, converts it to deterministic Markdown, and persists crawl metadata for incremental updates.

See [SERVICES.md](SERVICES.md) for the current mirrored service index and run status.

Detailed usage notes live in [cmd/aws-docs/README.md](cmd/aws-docs/README.md).