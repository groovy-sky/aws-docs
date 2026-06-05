---
title: "Ruby detectors"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](ruby/sql-injection.md) [Divide by Zero](ruby/divide-by-zero.md) [Sensitive HTTP Action](ruby/sensitive-http-action.md) [Insufficient Protected Credentials](ruby/insufficiently-protected-credentials.md) [Sensitive Information Leak](ruby/sensitive-information-leak.md) [Untrusted Deserialization](ruby/untrusted-deserialization.md) [Log Injection](ruby/log-injection.md) [XML External Entity](ruby/xml-external-entity.md) [Path Injection](ruby/path-traversal.md) [Http to File Access](ruby/http-to-file-access.md) [Code Injection](ruby/code-injection.md) [OS Command Injection](ruby/os-command-injection.md) [Resource leak](ruby/resource-leak.md) [Cross Site Scripting (XSS)](ruby/cross-site-scripting.md) [Untrusted Open](ruby/untrusted-file-open.md) [Improper Input Validation](ruby/improper-input-validation.md) [Stack Trace Exposure](ruby/stack-trace-exposure.md) [Improper Certificate Validation](ruby/improper-certificate-validation.md) [send\_file Injection](ruby/sendfile-injection.md) [Unsafe File Permissions](ruby/loose-file-permissions.md) [Tainted Format](ruby/tainted-format.md)

# Ruby detectors

Showing all detectors for the Ruby language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](ruby/tags.md)

##### Browse by severity

Browse all detectors by severity.

[Click here→](ruby/severity.md)

##### Browse by category

Browse all detectors by category.

[Click here→](ruby/categories.md)

* * *

### Browse all detectors

### [SQL Injection](ruby/sql-injection.md)

User input may run unintended SQL commands.

### [Divide by Zero](ruby/divide-by-zero.md)

Potentially dividing by zero without proper handling.

### [Sensitive HTTP Action](ruby/sensitive-http-action.md)

Issue found with `request.get?` block, potential unexpected behavior.

### [Insufficient Protected Credentials](ruby/insufficiently-protected-credentials.md)

The credentials provided are not adequately protected against security threats.

### [Sensitive Information Leak](ruby/sensitive-information-leak.md)

Neglecting sensitive information can lead to severe data leaks and breaches.

### [Untrusted Deserialization](ruby/untrusted-deserialization.md)

User input is deserialized.

### [Log Injection](ruby/log-injection.md)

Input from the user may be logged, giving false data.

### [XML External Entity](ruby/xml-external-entity.md)

Objects that parse or handle XML can lead to XML External Entity (XXE) attacks when misconfigured.

### [Path Injection](ruby/path-traversal.md)

User input may lead to opening unintended files.

### [Http to File Access](ruby/http-to-file-access.md)

Hardcoded download and writing of potentially harmful file.

### [Code Injection](ruby/code-injection.md)

User input is used in eval command.

### [OS Command Injection](ruby/os-command-injection.md)

Possible unintended system commands could be executed through user input.

### [Resource leak](ruby/resource-leak.md)

Allocated resources are not released properly.

### [Cross Site Scripting (XSS)](ruby/cross-site-scripting.md)

Improper neutralization of input during web page generation ('Cross-site Scripting')

### [Untrusted Open](ruby/untrusted-file-open.md)

Non-static variables used to open files.

### [Improper Input Validation](ruby/improper-input-validation.md)

Improper input validation can lead to security vulnerabilities and data breaches.

### [Stack Trace Exposure](ruby/stack-trace-exposure.md)

Stack trace shows software architecture.

### [Improper Certificate Validation](ruby/improper-certificate-validation.md)

Lack of validation of a security certificate can lead to host impersonation and sensitive data leaks.

### [send\_file Injection](ruby/sendfile-injection.md)

External Control of File Name or Path.

### [Unsafe File Permissions](ruby/loose-file-permissions.md)

Setting potentially harmful access rights

### [Tainted Format](ruby/tainted-format.md)

User input decides output information.

All content copied from https://docs.aws.amazon.com/.
