#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
PUBLIC_DIR="$ROOT_DIR/public"
STATIC_DIR="$ROOT_DIR/static"

mkdir -p "$PUBLIC_DIR"

# Copy search-index.json if it exists
[ -f "$STATIC_DIR/search-index.json" ] && cp "$STATIC_DIR/search-index.json" "$PUBLIC_DIR/"

# Copy CSS/JS assets
[ -d "$STATIC_DIR/css" ] && mkdir -p "$PUBLIC_DIR/css" && cp "$STATIC_DIR/css"/* "$PUBLIC_DIR/css/" 2>/dev/null || true
[ -d "$STATIC_DIR/js" ] && mkdir -p "$PUBLIC_DIR/js" && cp "$STATIC_DIR/js"/* "$PUBLIC_DIR/js/" 2>/dev/null || true

# Populate docs-index.json by walking docs/ and finding all .md files
# Build Go program (ensures it's executable, especially in CI environments)
go build -o cmd/gen-index/gen-index cmd/gen-index/main.go
./cmd/gen-index/gen-index

cat > "$PUBLIC_DIR/index.html" <<'HTML'
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>AWS Docs Index</title>
  <style>
    :root { --bg:#0f1115; --panel:#151922; --text:#dce3ef; --muted:#9aa8bc; --accent:#4db6ff; --border:#263246; }
    body { margin:0; font-family: ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif; background:var(--bg); color:var(--text); }
    .wrap { max-width: 1100px; margin: 0 auto; padding: 20px; }
    h1 { margin: 0 0 12px; font-size: 1.6rem; }
    .sub { color: var(--muted); margin-bottom: 16px; }
    .toolbar { display:flex; gap:10px; margin-bottom:14px; }
    input { flex:1; padding:10px 12px; border:1px solid var(--border); border-radius:8px; background:var(--panel); color:var(--text); }
    .card { background:var(--panel); border:1px solid var(--border); border-radius:10px; padding:14px; }
    ul { list-style:none; margin:0; padding-left:16px; }
    li { margin: 4px 0; }
    a { color: var(--accent); text-decoration: none; }
    a:hover { text-decoration: underline; }
    .dir { color: #c4d4ea; font-weight: 600; cursor: default; }
    .file { color: var(--text); }
    .raw { color: var(--muted); margin-left:8px; font-size: 0.9em; }
    details > summary { cursor: pointer; list-style: none; display: inline; }
    details > summary::-webkit-details-marker { display: none; }
    details > summary.dir::before { content: '▶ '; font-size: 0.75em; opacity: 0.7; }
    details[open] > summary.dir::before { content: '▼ '; font-size: 0.75em; opacity: 0.7; }
  </style>
</head>
<body>
  <div class="wrap">
    <h1>AWS Docs File Index</h1>
    <div class="sub">Static index generated from docs/. Click a file to open in the markdown viewer.</div>
    <div class="toolbar">
      <input id="q" type="search" placeholder="Filter files..." />
    </div>
    <div class="card">
      <div id="tree">Loading...</div>
    </div>
  </div>

  <script>
    const treeEl = document.getElementById('tree');
    const qEl = document.getElementById('q');

    function insertPath(root, path) {
      const parts = path.split('/');
      let node = root;
      for (let i = 0; i < parts.length; i++) {
        const part = parts[i];
        if (!node.children[part]) {
          node.children[part] = { name: part, children: {}, filePath: null };
        }
        node = node.children[part];
      }
      node.filePath = path;
    }

    function makeNode(name) {
      return { name, children: {}, filePath: null };
    }

    function renderNode(node, basePath, afterLatest) {
      afterLatest = !!afterLatest;
      const keys = Object.keys(node.children).sort((a, b) => a.localeCompare(b));
      const ul = document.createElement('ul');
      for (const key of keys) {
        const child = node.children[key];
        const li = document.createElement('li');
        const hasChildren = Object.keys(child.children).length > 0;
        if (hasChildren) {
          const childAfterLatest = afterLatest || key === 'latest';
          if (afterLatest) {
            const details = document.createElement('details');
            const summary = document.createElement('summary');
            summary.className = 'dir';
            summary.textContent = key + '/';
            details.appendChild(summary);
            details.appendChild(renderNode(child, basePath ? basePath + '/' + key : key, childAfterLatest));
            li.appendChild(details);
          } else {
            const label = document.createElement('span');
            label.className = 'dir';
            label.textContent = key + '/';
            li.appendChild(label);
            li.appendChild(renderNode(child, basePath ? basePath + '/' + key : key, childAfterLatest));
          }
        } else {
          const view = document.createElement('a');
          view.href = 'viewer.html?path=' + encodeURIComponent(child.filePath);
          view.className = 'file';
          view.textContent = key.endsWith('.md') ? key.slice(0, -3) : key;
          li.appendChild(view);

          const raw = document.createElement('a');
          raw.href = 'https://raw.githubusercontent.com/groovy-sky/aws-docs/main/docs/' + child.filePath.split('/').map(encodeURIComponent).join('/');
          raw.className = 'raw';
          raw.target = '_blank';
          raw.rel = 'noopener noreferrer';
          raw.textContent = 'raw';
          li.appendChild(raw);
        }
        ul.appendChild(li);
      }
      return ul;
    }

    function buildTree(paths, query) {
      const root = makeNode('root');
      const q = (query || '').trim().toLowerCase();
      for (const p of paths) {
        if (!q || p.toLowerCase().includes(q)) {
          insertPath(root, p);
        }
      }
      treeEl.innerHTML = '';
      treeEl.appendChild(renderNode(root, '', false));
      if (q) {
        treeEl.querySelectorAll('details').forEach(d => d.open = true);
      }
    }

    fetch('docs-index.json')
      .then(r => r.json())
      .then(paths => {
        buildTree(paths, '');
        qEl.addEventListener('input', e => buildTree(paths, e.target.value));
      })
      .catch(err => {
        console.error(err);
        treeEl.textContent = 'Failed to load docs index.';
      });
  </script>
</body>
</html>
HTML

cat > "$PUBLIC_DIR/viewer.html" <<'HTML'
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>Markdown Viewer</title>
  <style>
    :root { --bg:#0f1115; --panel:#151922; --text:#dce3ef; --muted:#9aa8bc; --accent:#4db6ff; --border:#263246; }
    body { margin:0; font-family: ui-sans-serif, system-ui, -apple-system, Segoe UI, sans-serif; background:var(--bg); color:var(--text); }
    .top { position: sticky; top:0; z-index:10; background:rgba(15,17,21,.9); backdrop-filter: blur(6px); border-bottom:1px solid var(--border); padding:10px 16px; display:flex; gap:10px; align-items:center; }
    .top a { color:var(--accent); text-decoration:none; }
    .path { color:var(--muted); font-size:.9rem; overflow:auto; white-space:nowrap; }
    main { max-width:1000px; margin:0 auto; padding:18px; }
    article { background:var(--panel); border:1px solid var(--border); border-radius:10px; padding:18px; }
    pre { overflow:auto; background:#0b0d12; padding:10px; border-radius:6px; }
    code { font-family: ui-monospace, SFMono-Regular, Menlo, monospace; }
    img { max-width:100%; }
    a { color: var(--accent); }
  </style>
</head>
<body>
  <div class="top">
    <a href="index.html">Back to index</a>
    <div class="path" id="path"></div>
  </div>
  <main>
    <article id="content">Loading...</article>
  </main>

  <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
  <script>
    const params = new URLSearchParams(window.location.search);
    const path = params.get('path');
    const pathEl = document.getElementById('path');
    const contentEl = document.getElementById('content');
    const RAW_BASE = 'https://raw.githubusercontent.com/groovy-sky/aws-docs/main/docs/';

    function encodePath(path) {
      return path.split('/').map(encodeURIComponent).join('/');
    }

    let docPath = path || '';
    if (docPath.startsWith('docs/')) {
      docPath = docPath.slice(5);
    }

    if (!docPath || docPath.startsWith('/') || docPath.includes('..')) {
      contentEl.textContent = 'Missing or invalid path.';
      throw new Error('invalid path');
    }

    const githubUrl = 'https://github.com/groovy-sky/aws-docs/blob/main/docs/' + encodePath(docPath);
    const pathLink = document.createElement('a');
    pathLink.href = githubUrl;
    pathLink.target = '_blank';
    pathLink.rel = 'noopener noreferrer';
    pathLink.textContent = githubUrl;
    pathEl.appendChild(pathLink);

    fetch(RAW_BASE + encodePath(docPath))
      .then(r => {
        if (!r.ok) throw new Error('HTTP ' + r.status);
        return r.text();
      })
      .then(md => {
        contentEl.innerHTML = marked.parse(md, { mangle: false, headerIds: true });
      })
      .catch(err => {
        console.error(err);
        contentEl.textContent = 'Failed to load markdown: ' + err.message;
      });
  </script>
</body>
</html>
HTML

echo "Static site built at $PUBLIC_DIR"
