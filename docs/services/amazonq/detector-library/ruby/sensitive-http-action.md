![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-http-action) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/ruby/insufficiently-protected-credentials) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-information-leak) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-deserialization) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/log-injection) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity) [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Sensitive HTTP Action [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

An issue is discovered with the control flow block that utilizes `request.get?`. This issue can lead to unexpected behavior, as Rails treats HEAD requests as GET requests. To mitigate this, it is recommended to include an `elif` condition to handle HEAD requests separately and avoid any potential complications that may arise.

**Detector ID**

ruby/sensitive-http-action@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

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
