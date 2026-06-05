---
title: "Path Injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# Path Injection [High](severity/high.md)

User input leads to file opening. This allows users to take control and open any readable file on the system which may leak sensitive information. This can be sanitized with the basename method.

**Detector ID**

ruby/path-traversal@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1def render_modern_param_noncompliant
2    page = params[:page]
3    # Noncompliant: Unsanitized user-input is used in render file.
4    render file: "/some/path/#{page}"
5end
```

#### Compliant example

```ruby
1def render_modern_param_compliant
2    page = params[:page]
3    # Compliant: User-input is sanitized before using it in render file.
4    render file: File.basename("/some/path/#{page}")
5end
```

All content copied from https://docs.aws.amazon.com/.
