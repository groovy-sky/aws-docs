#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
BASE_URL="${1:-https://groovy-sky.github.io/aws-docs/}"

# Build bare minimum: Disable unnecessary page kinds (RSS, sitemap, taxonomies, terms)
exec hugo --baseURL "$BASE_URL" --contentDir "$ROOT_DIR/docs" --disableKinds "RSS,sitemap,taxonomy,term"