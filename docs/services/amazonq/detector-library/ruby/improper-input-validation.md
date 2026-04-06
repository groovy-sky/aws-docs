![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Improper Input Validation [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Improper input validation can result in potential exploits, leaving systems vulnerable to malicious attacks, compromising data integrity, and undermining overall system security, which should always be a top priority.

**Detector ID**

ruby/improper-input-validation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-185](https://cwe.mitre.org/data/definitions/185.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1class InputValidation
2  # Noncompliant: Improperly bounded regex passed to validates.
3  validates :username, :length => 6..20, :format => /([a-z][0-9])+/i
4
5  accepts_nested_attributes_for :author, :pages
6end
```

#### Compliant example

```ruby
1class InputValidation
2  # Compliant: Properly bounded regex passed to validates.
3  validates_format_of :good_valid, :with => /\A[a-zA-Z]\z/
4end
```
