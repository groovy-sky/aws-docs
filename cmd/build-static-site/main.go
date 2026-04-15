package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

const indexHTML = `<!doctype html>
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
    .dir { color: #c4d4ea; font-weight: 600; }
    .file { color: var(--text); }
    .raw { color: var(--muted); margin-left:8px; font-size: 0.9em; }
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

    function renderNode(node, basePath = '') {
      const keys = Object.keys(node.children).sort((a, b) => a.localeCompare(b));
      const ul = document.createElement('ul');
      for (const key of keys) {
        const child = node.children[key];
        const li = document.createElement('li');
        const hasChildren = Object.keys(child.children).length > 0;
        if (hasChildren) {
          const label = document.createElement('span');
          label.className = 'dir';
          label.textContent = key + '/';
          li.appendChild(label);
          li.appendChild(renderNode(child, basePath ? basePath + '/' + key : key));
        } else {
          const view = document.createElement('a');
          view.href = 'viewer.html?path=' + encodeURIComponent(child.filePath);
          view.className = 'file';
          view.textContent = key;
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
      treeEl.appendChild(renderNode(root));
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
`

const viewerHTML = `<!doctype html>
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

    pathEl.textContent = 'docs/' + docPath;

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
`

func main() {
	rootFlag := flag.String("root", "", "project root directory (defaults to current working directory)")
	flag.Parse()

	rootDir, err := resolveRootDir(*rootFlag)
	if err != nil {
		fatalf("resolve root: %v", err)
	}

	publicDir := filepath.Join(rootDir, "public")
	staticDir := filepath.Join(rootDir, "static")
	docsDir := filepath.Join(rootDir, "docs")

	if err := os.MkdirAll(publicDir, 0o755); err != nil {
		fatalf("create public dir: %v", err)
	}

	if err := copyIfExists(filepath.Join(staticDir, "search-index.json"), filepath.Join(publicDir, "search-index.json")); err != nil {
		fatalf("copy search-index.json: %v", err)
	}

	if err := copyDirFilesIfExists(filepath.Join(staticDir, "css"), filepath.Join(publicDir, "css")); err != nil {
		fatalf("copy css: %v", err)
	}

	if err := copyDirFilesIfExists(filepath.Join(staticDir, "js"), filepath.Join(publicDir, "js")); err != nil {
		fatalf("copy js: %v", err)
	}

	paths, err := collectMarkdownPaths(docsDir)
	if err != nil {
		fatalf("walk docs: %v", err)
	}

	if err := writeJSON(filepath.Join(publicDir, "docs-index.json"), paths); err != nil {
		fatalf("write docs-index.json: %v", err)
	}

	if err := writeText(filepath.Join(publicDir, "index.html"), indexHTML); err != nil {
		fatalf("write index.html: %v", err)
	}

	if err := writeText(filepath.Join(publicDir, "viewer.html"), viewerHTML); err != nil {
		fatalf("write viewer.html: %v", err)
	}

	fmt.Printf("Static site built at %s\n", publicDir)
}

func resolveRootDir(rootFlag string) (string, error) {
	if rootFlag != "" {
		return filepath.Abs(rootFlag)
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func copyIfExists(src, dst string) error {
	if !pathExists(src) {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return copyFile(src, dst)
}

func copyDirFilesIfExists(srcDir, dstDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		return err
	}

	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())
		if err := copyFile(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode().Perm())
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}

func collectMarkdownPaths(docsDir string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(docsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		rel, err := filepath.Rel(docsDir, path)
		if err != nil {
			return err
		}
		paths = append(paths, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(paths)
	return paths, nil
}

func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func writeText(path, content string) error {
	return os.WriteFile(path, []byte(content), 0o644)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
