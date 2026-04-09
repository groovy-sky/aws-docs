// Mobile sidebar toggle
document.addEventListener('DOMContentLoaded', function() {
  const sidebar = document.querySelector('.sidebar');
  if (!sidebar) return;
  
  // Create toggle button
  const toggle = document.createElement('button');
  toggle.className = 'sidebar-toggle';
  toggle.innerHTML = '☰ Menu';
  toggle.setAttribute('aria-label', 'Toggle navigation');
  
  const header = document.querySelector('.site-header');
  if (header) {
    header.appendChild(toggle);
  }
  
  toggle.addEventListener('click', function(e) {
    e.stopPropagation();
    sidebar.classList.toggle('mobile-open');
  });
  
  // Close on outside click
  document.addEventListener('click', function(e) {
    if (!sidebar.contains(e.target) && e.target !== toggle) {
      sidebar.classList.remove('mobile-open');
    }
  });
  
  // Add copy buttons to code blocks
  document.querySelectorAll('pre').forEach(function(pre) {
    const button = document.createElement('button');
    button.className = 'copy-button';
    button.textContent = 'Copy';
    button.style.cssText = 'position: absolute; top: 0.5rem; right: 0.5rem; padding: 0.25rem 0.75rem; background: var(--bg-tertiary); border: 1px solid var(--border); color: var(--text-primary); border-radius: 4px; cursor: pointer; font-size: 0.875rem;';
    
    pre.style.position = 'relative';
    
    button.addEventListener('click', async function() {
      const code = pre.querySelector('code').textContent;
      try {
        await navigator.clipboard.writeText(code);
        button.textContent = 'Copied!';
        setTimeout(function() {
          button.textContent = 'Copy';
        }, 2000);
      } catch (err) {
        console.error('Failed to copy:', err);
        button.textContent = 'Failed';
        setTimeout(function() {
          button.textContent = 'Copy';
        }, 2000);
      }
    });
    
    pre.appendChild(button);
  });
});
