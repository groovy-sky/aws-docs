![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sql-injection) [Divide by Zero](https://docs.aws.amazon.com/amazonq/detector-library/ruby/divide-by-zero) [Sensitive HTTP Action](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-http-action) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/ruby/insufficiently-protected-credentials) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-information-leak) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-deserialization) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/log-injection) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity) [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# Ruby detectors

Showing all detectors for the Ruby language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories)

* * *

### Browse all detectors

### [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sql-injection)

User input may run unintended SQL commands.

### [Divide by Zero](https://docs.aws.amazon.com/amazonq/detector-library/ruby/divide-by-zero)

Potentially dividing by zero without proper handling.

### [Sensitive HTTP Action](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-http-action)

Issue found with `request.get?` block, potential unexpected behavior.

### [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/ruby/insufficiently-protected-credentials)

The credentials provided are not adequately protected against security threats.

### [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-information-leak)

Neglecting sensitive information can lead to severe data leaks and breaches.

### [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-deserialization)

User input is deserialized.

### [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/log-injection)

Input from the user may be logged, giving false data.

### [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity)

Objects that parse or handle XML can lead to XML External Entity (XXE) attacks when misconfigured.

### [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal)

User input may lead to opening unintended files.

### [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access)

Hardcoded download and writing of potentially harmful file.

### [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection)

User input is used in eval command.

### [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection)

Possible unintended system commands could be executed through user input.

### [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak)

Allocated resources are not released properly.

### [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting)

Improper neutralization of input during web page generation ('Cross-site Scripting')

### [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open)

Non-static variables used to open files.

### [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation)

Improper input validation can lead to security vulnerabilities and data breaches.

### [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure)

Stack trace shows software architecture.

### [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation)

Lack of validation of a security certificate can lead to host impersonation and sensitive data leaks.

### [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection)

External Control of File Name or Path.

### [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions)

Setting potentially harmful access rights

### [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

User input decides output information.
