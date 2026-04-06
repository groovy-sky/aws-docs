![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Improper Certificate Validation [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Lack of validation or insufficient validation of a security certificate can lead to host impersonation and sensitive data leaks.

**Detector ID**

ruby/improper-certificate-validation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-295](https://cwe.mitre.org/data/definitions/295.html)

**Tags**

[\# access-control](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/access-control) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/owasp-top10)

* * *

#### Noncompliant example

```ruby
1require "httparty"
2
3def certificate_validation_noncompliant
4
5  # Noncompliant: SSL certificate validation is disabled.
6  HTTParty.get("http://example.com/", verify: false)
7
8end
```

#### Compliant example

```ruby
1require "httparty"
2
3def certificate_validation_compliant
4
5  # Compliant: SSL certificate validation is enabled.
6  HTTParty.get("http://example.com/", verify: true)
7
8end
```
