# 1. DOM Reduction Before Conversion

Goal: shrink the DOM **before Markdown conversion** so the converter processes far fewer nodes.

Large pages contain scripts, navigation, UI elements, ads, and hidden elements that are irrelevant to Markdown output.

### Research
Content extraction tools and readability libraries remove clutter like navigation, ads, scripts, and background elements to isolate readable content. ([github.com](https://github.com/go-shiori/go-readability))

These elements typically account for **50–80% of nodes** in modern pages.

### Plan

Step 1 — Parse HTML once
Use the standard parser:

- `golang.org/x/net/html`

Avoid parsing multiple times.

```
doc, _ := html.Parse(reader)
```

---

Step 2 — Remove heavy nodes

Delete entire subtrees for tags:

Remove always:
- script
- style
- noscript
- iframe
- svg
- canvas
- form

Often remove:
- nav
- footer
- header
- aside

Pseudo‑code:

```
func prune(n *html.Node) {
    if isRemovable(n) {
        removeNode(n)
        return
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        prune(c)
    }
}
```

---

Step 3 — Remove invisible elements

Check attributes:

Remove nodes with:

```
hidden
aria-hidden="true"
style="display:none"
style="visibility:hidden"
```

---

Step 4 — Remove tracking / UI classes

Drop nodes with class names like:

```
cookie
banner
advert
promo
subscribe
sidebar
modal
popup
```

Regex filtering works well.

---

Step 5 — Drop empty containers

After pruning, remove elements like:

```
<div></div>
<span></span>
<p></p>
```

This can reduce nodes another **10–20%**.

---

Expected improvement:

- DOM size: ↓ 40–70%
- Markdown generation CPU: ↓ ~30–50%

---

# 2. Extract Only Meaningful Content

Instead of converting the entire DOM, extract the **main article/content block**.

This is the **single biggest optimization**.

### Research

Readability-style algorithms extract the main readable content while removing navigation, ads, and other boilerplate. ([github.com](https://github.com/go-shiori/go-readability))

Libraries:

- `go-readability`
- `go-boilerpipe`

Boilerpipe focuses on text extraction and removing boilerplate sections. ([github.com](https://github.com/jlubawy/go-boilerpipe))

---

### Plan

Pipeline:

```
HTML
 ↓
Readability extraction
 ↓
clean article DOM
 ↓
Markdown conversion
```

---

Step 1 — Run readability extraction

Example:

```
article, err := readability.FromReader(reader, url)
```

Result:

```
article.Title
article.Content
article.TextContent
```

Use `article.Content` (clean HTML).

---

Step 2 — Reparse cleaned HTML

```
doc, _ := html.Parse(strings.NewReader(article.Content))
```

Now your DOM is **tiny compared to the original**.

---

Step 3 — Fallback strategy

Not all pages work with readability.

Fallback order:

1. readability
2. boilerpipe
3. manual heuristic

Manual heuristic:

Find node with largest:

```
text length / number of links
```

This is similar to Readability scoring.

---

Step 4 — Detect low‑quality extraction

Reject results when:

```
len(text) < 300
```

Then fallback to heuristic extraction.

---

Expected improvement:

Typical news/blog pages:

```
HTML size: 500KB
article HTML: 20–60KB
```

Reduction: **80–95%**

Markdown output shrinks accordingly.

---

# 3. Normalize HTML Before Markdown

HTML is extremely verbose.

Example:

```
<div>
   <span>
      <span>
         Text
      </span>
   </span>
</div>
```

Markdown converter generates messy output unless normalized.

---

### Plan

Step 1 — Flatten useless containers

Replace:

```
div
span
section
article
```

with their children if they:

- contain only text
- contain only one child
- have no semantic attributes

Example:

```
<div><p>text</p></div>
```

→

```
<p>text</p>
```

---

Step 2 — Remove attributes

Drop:

```
class
id
style
data-*
aria-*
onclick
```

Keep only:

```
href
src
alt
title
```

---

Step 3 — Normalize inline tags

Convert:

```
<b> → <strong>
<i> → <em>
```

Remove:

```
font
center
```

---

Step 4 — Normalize lists

Fix messy lists like:

```
<ul>
 <div><li>item</li></div>
</ul>
```

→

```
<ul>
 <li>item</li>
</ul>
```

---

Step 5 — Collapse whitespace

Before Markdown conversion:

```
multiple spaces → single
multiple newlines → max 2
```

---

Expected improvements

Markdown generators perform best on **semantic HTML**:

```
h1-h6
p
ul/ol/li
blockquote
pre/code
img
a
table
```

Everything else increases Markdown noise.

---

# Recommended Optimized Pipeline

```
Crawler
   ↓
Fetch HTML
   ↓
Initial DOM prune (scripts/nav/etc)
   ↓
Readability extraction
   ↓
Reparse extracted HTML
   ↓
HTML normalization
   ↓
Markdown conversion
   ↓
Whitespace cleanup
   ↓
Storage
```

---

# Realistic Performance Gains

Typical page:

```
Original HTML: 600KB
```

After pipeline:

```
After prune: 300KB
After readability: 30KB
After normalization: 20KB
Markdown output: 10–15KB
```

Result:

```
Size reduction: ~97%
CPU reduction: ~60–80%
```