---
title: "Log Injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# Log Injection [High](severity/high.md)

User input has access to log output, which allows manipulation of logged data. This may allow for malicious users to log false information.

**Detector ID**

ruby/log-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-117](https://cwe.mitre.org/data/definitions/117.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1  def log_params_noncompliant
2    init_logger
3
4    unsanitized = params[:foo]
5    # Noncompliant: Unsanitized user-input is used in logger
6    @logger.error "input: " + unsanitized
7  end
```

#### Compliant example

```ruby
1  def log_params_compliant
2    init_logger
3
4    unsanitized = params[:foo]
5
6    sanitized = unsanitized.gsub("\n", "")
7    # Compliant: Sanitized user-input is used in logger
8    @logger.warn "input: " + sanitized
9  end
```

All content copied from https://docs.aws.amazon.com/.
