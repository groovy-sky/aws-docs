---
title: "XML External Entity High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# XML External Entity [High](severity/high.md)

Objects that parse or handle XML data can lead to XML External Entity (XXE) attacks when not configured properly. Improper restriction of XML external entity processing can lead to server-side request forgery and information disclosure.

**Detector ID**

ruby/xml-external-entity@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-611](https://cwe.mitre.org/data/definitions/611.html)

**Tags**

[\# injection](tags/injection.md) [\# xml](tags/xml.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

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

All content copied from https://docs.aws.amazon.com/.
