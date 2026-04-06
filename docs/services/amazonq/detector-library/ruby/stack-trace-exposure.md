![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Stack Trace Exposure [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

Stack trace is exposed, which gives unnecessary architectural information. This information may give an attack information on how to trigger more attacks.

**Detector ID**

ruby/stack-trace-exposure@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-209](https://cwe.mitre.org/data/definitions/209.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1def exposure_of_stack_trace_noncompliant
2      # Process
3    rescue => e
4      # Noncompliant: Rendering the stack trace information
5      render body: e.backtrace, content_type: "text/plain"
6 end
```

#### Compliant example

```ruby
1def exposure_of_stack_trace_compliant
2  # Process
3  rescue => e
4    # Compliant: Rending a simple error message.
5    render body: "An error occurred", content_type: "text/plain"
6end
```
