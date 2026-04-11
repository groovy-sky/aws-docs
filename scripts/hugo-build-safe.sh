#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
TMP_ROOT="$(mktemp -d)"
TMP_CONTENT_DIR="$TMP_ROOT/hugo-content"
BASE_URL="${1:-https://groovy-sky.github.io/aws-docs/}"

cleanup() {
  rm -rf "$TMP_ROOT"
}
trap cleanup EXIT

mkdir -p "$TMP_CONTENT_DIR"

# Copy docs content to a temporary build directory outside the repo tree.
cp -a "$ROOT_DIR/docs/." "$TMP_CONTENT_DIR/"

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
    tmp_file="$(mktemp)"
    {
      printf '<!-- hugo-normalized-leading-dash -->\n'
      cat "$file"
    } > "$tmp_file"
    mv "$tmp_file" "$file"
    fixed=$((fixed + 1))
  fi
done < <(find "$TMP_CONTENT_DIR" -type f -name '*.md' -print0)

echo "Normalized $fixed markdown files in temporary content directory"

exec hugo --minify --baseURL "$BASE_URL" --contentDir "$TMP_CONTENT_DIR"
