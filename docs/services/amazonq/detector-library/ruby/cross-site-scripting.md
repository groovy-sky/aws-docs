![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Cross Site Scripting (XSS) [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Cross Site Scripting (XSS) is an attack which exploits a web application or system to treat user input as markup or script code. It is important to encode the data depending on the specific context it is used in.

**Detector ID**

ruby/cross-site-scripting@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-79](https://cwe.mitre.org/data/definitions/79.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1def crosssite_scripting_noncompliant
2  name = params[":name"]
3  # Noncompliant: The parameter is not escaped.
4  "<h2>#{name}</h2>".html_safe
5end
```

#### Compliant example

```ruby
1def crosssite_scripting_compliant
2  name = params[":name"]
3  # Compliant: Parameter is escaped.
4  "<h2>#{ERB::Util.html_escape(name)}</h2>".html_safe
5end
```
