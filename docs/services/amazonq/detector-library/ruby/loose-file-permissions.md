![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Unsafe File Permissions [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Potentially dangerous settings are being assigned to a file. Either the group or world levels are being given read or write access to a file. This might allow unintended users to access private information.

**Detector ID**

ruby/loose-file-permissions@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-732](https://cwe.mitre.org/data/definitions/732.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1require "fileutils"
2
3def open_file_permission_noncompliant(filename)
4
5  # Noncompliant: sets file world writable.
6  FileUtils.chmod 0222, filename
7end
```

#### Compliant example

```ruby
1def open_file_permission_compliant(filename)
2
3  # Compliant: restricts group/world access.
4  FileUtils.chmod 0700, filename
5end
```
