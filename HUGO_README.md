# Hugo Setup for AWS Documentation

This repository uses [Hugo](https://gohugo.io/) to build a static documentation site from Markdown files in the `docs/` directory.

## Quick Start

### Prerequisites

Install Hugo Extended:
```bash
# macOS
brew install hugo

# Linux (snap)
snap install hugo --channel=extended

# Or download from https://github.com/gohugoio/hugo/releases
```

### Local Development

1. **Start the development server:**
   ```bash
   hugo server -D
   ```

2. **View the site:**
   Open http://localhost:1313/aws-docs/ in your browser

3. **The server will auto-reload** when you edit files

### Build for Production

```bash
hugo --minify
```

Output will be in the `public/` directory.

## Project Structure

```
aws-docs/
├── hugo.toml              # Hugo configuration
├── docs/                  # Content directory (existing structure)
│   ├── reference/        # API references
│   ├── services/         # Service guides
│   └── general/          # General documentation
├── layouts/              # HTML templates
│   ├── _default/         # Default layouts
│   ├── partials/         # Reusable components
│   └── search/           # Search page layout
├── static/               # Static assets
│   ├── css/main.css      # Styling
│   └── js/               # JavaScript
└── content/              # Additional content (search page)
```

## Features

- ✅ **Dark minimalistic theme**
- ✅ **Client-side search** with Fuse.js
- ✅ **No template parsing issues** - code blocks are safe
- ✅ **Responsive design** with mobile navigation
- ✅ **Automatic section pages**
- ✅ **Table of contents** for long pages
- ✅ **Copy buttons** for code blocks
- ✅ **Fast builds** (Hugo builds in milliseconds)

## GitHub Pages Deployment

The site automatically deploys to GitHub Pages on push to the `main` branch via `.github/workflows/pages.yml`.

## Configuration

Edit `hugo.toml` to customize:
- Base URL
- Site title and description
- Menu items
- Permalink structure

## Adding Content

Simply add or edit Markdown files in any `docs/` subdirectory. Hugo will automatically:
- Generate pages
- Build navigation
- Index content for search

## Search

The search functionality:
- Indexes all pages automatically
- Runs entirely client-side (no server needed)
- Searches titles, content, and sections
- Highlights matches

## Troubleshooting

**Port already in use:**
```bash
hugo server -p 1314
```

**Clear cache:**
```bash
hugo --gc
```

**Verbose output:**
```bash
hugo server --verbose
```

## Customization

### Colors
Edit CSS variables in `static/css/main.css`:
```css
:root {
  --bg-primary: #1a1a1a;
  --accent: #58a6ff;
  /* ... */
}
```

### Layouts
Modify templates in `layouts/` directory.

### Navigation
Edit `hugo.toml` under `[menu]` section.

## Resources

- [Hugo Documentation](https://gohugo.io/documentation/)
- [Hugo Themes](https://themes.gohugo.io/)
- [Fuse.js Search](https://fusejs.io/)

## License

Same as repository license.
