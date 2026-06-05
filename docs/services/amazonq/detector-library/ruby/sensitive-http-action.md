---
title: "Sensitive HTTP Action High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# Sensitive HTTP Action [High](severity/high.md)

An issue is discovered with the control flow block that utilizes `request.get?`. This issue can lead to unexpected behavior, as Rails treats HEAD requests as GET requests. To mitigate this, it is recommended to include an `elif` condition to handle HEAD requests separately and avoid any potential complications that may arise.

**Detector ID**

ruby/sensitive-http-action@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-650](https://cwe.mitre.org/data/definitions/650.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1class AccountsController < ApplicationController
2    def sensitive_http_get_noncompliant
3        # Noncompliant: GET request with a catch all 'else' block which might catch HEAD requests unknowingly
4        if request.get?
5            # Process request
6        else
7            # Process request
8        end
9    end
10end
```

#### Compliant example

```ruby
1class AccountsController < ApplicationController
2    def sensitive_http_get_compliant
3        # Compliant: GET request with 'elsif' which means exclusive blocks for other http methods
4        if request.get?
5            # Process request
6        elsif request.post?
7            # Process request
8        end
9    end
10end
```

All content copied from https://docs.aws.amazon.com/.
