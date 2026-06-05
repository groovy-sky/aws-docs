---
title: "Insufficient Protected Credentials High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# Insufficient Protected Credentials [High](severity/high.md)

The credentials being used do not have sufficient protection measures in place to prevent potential security breaches. Ensure that passwords and other sensitive information are stored in encrypted form.

**Detector ID**

ruby/insufficiently-protected-credentials@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-522](https://cwe.mitre.org/data/definitions/522.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1require 'jwt'
2
3def insufficiently_protected_credentials_noncompliant(hmac_secret)
4  # Noncompliant: JWT password is hardcoded in payload.
5  payload = { data: 'data', password: 12345 }
6  token = JWT.encode payload, hmac_secret, 'HS256'
7  puts token
8end
```

#### Compliant example

```ruby
1def insufficiently_protected_credentials_compliant(hmac_secret)
2  # Compliant: JWT password is not hardcoded.
3  payload = { data: 'data', nbf: nbf }
4  token = JWT.encode payload, hmac_secret, 'HS256'
5  puts token
6end
```

All content copied from https://docs.aws.amazon.com/.
