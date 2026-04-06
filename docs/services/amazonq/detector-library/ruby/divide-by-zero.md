![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](https://docs.aws.amazon.com/amazonq/detector-library/ruby/divide-by-zero) [Sensitive HTTP Action](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-http-action) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/ruby/insufficiently-protected-credentials) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-information-leak) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-deserialization) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/log-injection) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity) [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Divide by Zero [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Divide by zero is detected, this can result in exceptions, or unintended values. Either wrap the expression in an if statement checking the denominator, or wrap in an exception handling begin block.

**Detector ID**

ruby/divide-by-zero@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-369](https://cwe.mitre.org/data/definitions/369.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1def divide_by_zero_noncompliant
2  zero = 0
3  # Noncompliant: divide by zero
4  bad = variable/zero
5end
```

#### Compliant example

```ruby
1def divide_by_zero_compliant
2  # Compliant: check before dividing
3  if zero != 0
4    variable / zero
5  end
6end
```
