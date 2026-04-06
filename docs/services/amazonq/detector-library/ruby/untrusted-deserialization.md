![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-deserialization) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/log-injection) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity) [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Untrusted Deserialization [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Object deserialized from potentially dangerous source. This could overload a system with a denial of service attack, or execute remote code to control a system.

**Detector ID**

ruby/untrusted-deserialization@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-502](https://cwe.mitre.org/data/definitions/502.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1def handler_noncompliant(event:, context:)
2  foobar = event['smth']
3  # Noncompliant: tainted data passed inside `CSV.load()`.
4  obj3 = CSV.load("o:" + event['data'])
5end
```

#### Compliant example

```ruby
1def handler_compliant(event:, context:)
2  # Compliant: no tainted data passed inside `CSV.load()`.
3  obj3 = CSV.load(get_safe_data())
4end
```
