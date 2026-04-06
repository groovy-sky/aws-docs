![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Resource leak [Medium](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/medium)

Allocated resources are not released properly. This can slow down or crash your system. They must be closed along all paths to prevent a resource leak.

**Detector ID**

ruby/resource-leak@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-400](https://cwe.mitre.org/data/definitions/400.html) [CWE-664](https://cwe.mitre.org/data/definitions/664.html)

**Tags**

[\# availability](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/availability) [\# resource-leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/resource-leak) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags/top25-cwes)

* * *

#### Noncompliant example

```ruby
1def file_reading_noncompliant(filename)
2  # Noncompliant: File hasn't been closed
3  file = File.open(filename, 'r')
4  contents = file.read
5  puts contents
6end
```

#### Compliant example

```ruby
1def file_reading_compliant(filename)
2  # Compliant: File has been closed after read
3  File.open(filename, 'r') do |file|
4    file.each_line do |line|
5      puts line
6    end
7  end
8end
```
