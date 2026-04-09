(function() {
  'use strict';
  
  const searchInput = document.getElementById('search-input');
  const searchResults = document.getElementById('search-results');
  
  if (!searchInput || !searchResults) return;
  
  let fuse = null;
  let searchIndex = [];
  
  // Load search index
  fetch('/aws-docs/index.json')
    .then(function(response) { return response.json(); })
    .then(function(data) {
      searchIndex = data;
      
      // Initialize Fuse.js
      const options = {
        keys: [
          { name: 'title', weight: 0.5 },
          { name: 'content', weight: 0.3 },
          { name: 'section', weight: 0.2 }
        ],
        threshold: 0.3,
        includeMatches: true,
        minMatchCharLength: 2,
        ignoreLocation: true
      };
      
      fuse = new Fuse(searchIndex, options);
      
      // Check URL params
      const urlParams = new URLSearchParams(window.location.search);
      const query = urlParams.get('q');
      if (query) {
        searchInput.value = query;
        performSearch(query);
      }
    })
    .catch(function(err) {
      console.error('Failed to load search index:', err);
      searchResults.innerHTML = '<p class="search-hint">Failed to load search index</p>';
    });
  
  // Search on input
  let debounceTimer;
  searchInput.addEventListener('input', function(e) {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(function() {
      const query = e.target.value.trim();
      
      if (query.length === 0) {
        searchResults.innerHTML = '<p class="search-hint">Enter a search term to see results</p>';
        updateURL('');
        return;
      }
      
      if (query.length < 2) {
        searchResults.innerHTML = '<p class="search-hint">Please enter at least 2 characters</p>';
        return;
      }
      
      performSearch(query);
      updateURL(query);
    }, 300);
  });
  
  function performSearch(query) {
    if (!fuse) {
      searchResults.innerHTML = '<p class="search-hint">Search index loading...</p>';
      return;
    }
    
    const results = fuse.search(query);
    
    if (results.length === 0) {
      searchResults.innerHTML = '<p class="search-hint">No results found for "<strong>' + escapeHtml(query) + '</strong>"</p>';
      return;
    }
    
    renderResults(results, query);
  }
  
  function renderResults(results, query) {
    const statsHTML = '<div class="search-stats">Found ' + results.length + ' result' + (results.length !== 1 ? 's' : '') + ' for "<strong>' + escapeHtml(query) + '</strong>"</div>';
    
    const resultsHTML = results.map(function(result) {
      const item = result.item;
      const excerpt = highlightMatches(item.content, result.matches);
      
      return '<div class="search-result">' +
        '<h3><a href="' + item.url + '">' + escapeHtml(item.title) + '</a></h3>' +
        '<div class="section-breadcrumb">' + escapeHtml(item.section) + '</div>' +
        '<div class="excerpt">' + excerpt + '</div>' +
        '</div>';
    }).join('');
    
    searchResults.innerHTML = statsHTML + resultsHTML;
  }
  
  function highlightMatches(content, matches) {
    if (!matches || matches.length === 0) {
      return escapeHtml(content.substring(0, 300)) + '...';
    }
    
    // Find matches in content
    const contentMatches = matches.filter(function(m) { return m.key === 'content'; });
    if (contentMatches.length === 0) {
      return escapeHtml(content.substring(0, 300)) + '...';
    }
    
    // Get first match
    const match = contentMatches[0];
    const indices = match.indices[0];
    
    if (!indices) {
      return escapeHtml(content.substring(0, 300)) + '...';
    }
    
    const start = indices[0];
    const end = indices[1];
    const contextStart = Math.max(0, start - 100);
    const contextEnd = Math.min(content.length, end + 200);
    
    let excerpt = content.substring(contextStart, contextEnd);
    const matchText = content.substring(start, end + 1);
    
    excerpt = escapeHtml(excerpt).replace(
      new RegExp(escapeRegex(escapeHtml(matchText)), 'gi'),
      function(match) { return '<mark>' + match + '</mark>'; }
    );
    
    return (contextStart > 0 ? '...' : '') + excerpt + (contextEnd < content.length ? '...' : '');
  }
  
  function updateURL(query) {
    const url = new URL(window.location);
    if (query) {
      url.searchParams.set('q', query);
    } else {
      url.searchParams.delete('q');
    }
    window.history.replaceState({}, '', url);
  }
  
  function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
  }
  
  function escapeRegex(text) {
    return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  }
})();
