![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity) [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# XML External Entity [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Objects that parse or handle XML data can lead to XML External Entity (XXE) attacks when not configured properly. Improper restriction of XML external entity processing can lead to server-side request forgery and information disclosure.

**Detector ID**

ruby/xml-external-entity@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-611](https://cwe.mitre.org/data/definitions/611.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/injection) [\# xml](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/xml) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/top25-cwes)

* * *

#### Noncompliant example

```ruby
1def xml_external_entity_noncompliant()
2  # Noncompliant: Disabling encryption of sensitive data
3  config.force_ssl = false
4end
```

#### Compliant example

```ruby
1def xml_external_entity_compliant()
2  # Compliant: Sensitive data is encrypted
3  config.force_ssl = true
4end
```
