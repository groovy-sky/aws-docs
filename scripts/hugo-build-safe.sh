#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
BASE_URL="${1:-https://groovy-sky.github.io/aws-docs/}"

# We'll modify the docs directly in the workspace instead of duplicating them
CONTENT_DIR="$ROOT_DIR/docs"

fixed=0
while IFS= read -r -d '' file; do
  first_line=""
  IFS= read -r first_line < "$file" || true

  if [ "$first_line" = "---" ]; then
    # Keep valid front matter files unchanged.
    if awk 'NR > 1 && $0 == "---" { found=1; exit } END { exit(found ? 0 : 1) }' "$file"; then
      continue
    fi
  fi

  if [ "${first_line#- }" != "$first_line" ] || [ "$first_line" = "---" ]; then
    # In-place editing: Much faster than mktemp + cat + mv
    sed -i '1s/^/<!-- hugo-normalized-leading-dash -->\n/' "$file"
    fixed=$((fixed + 1))
  fi
done < <(find "$CONTENT_DIR" -type f -name '*.md' -print0)

echo "Normalized $fixed markdown files in $CONTENT_DIR"

# Build bare minimum: Disable unnecessary page kinds (RSS, sitemap, taxonomies, terms)
exec hugo --minify --baseURL "$BASE_URL" --contentDir "$CONTENT_DIR" --disableKinds "RSS,sitemap,taxonomy,term"