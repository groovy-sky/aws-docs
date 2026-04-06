![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](../sql-injection.md) [Divide by Zero](../divide-by-zero.md) [Sensitive HTTP Action](../sensitive-http-action.md) [Insufficient Protected Credentials](../insufficiently-protected-credentials.md) [Sensitive Information Leak](../sensitive-information-leak.md) [Untrusted Deserialization](../untrusted-deserialization.md) [Log Injection](../log-injection.md) [XML External Entity](../xml-external-entity.md) [Path Injection](../path-traversal.md) [Http to File Access](../http-to-file-access.md) [Code Injection](../code-injection.md) [OS Command Injection](../os-command-injection.md) [Resource leak](../resource-leak.md) [Cross Site Scripting (XSS)](../cross-site-scripting.md) [Untrusted Open](../untrusted-file-open.md) [Improper Input Validation](../improper-input-validation.md) [Stack Trace Exposure](../stack-trace-exposure.md) [Improper Certificate Validation](../improper-certificate-validation.md) [send\_file Injection](../sendfile-injection.md) [Unsafe File Permissions](../loose-file-permissions.md) [Tainted Format](../tainted-format.md)

# Security detectors

### [SQL Injection](../sql-injection.md)

User input may run unintended SQL commands.

### [Divide by Zero](../divide-by-zero.md)

Potentially dividing by zero without proper handling.

### [Sensitive HTTP Action](../sensitive-http-action.md)

Issue found with `request.get?` block, potential unexpected behavior.

### [Insufficient Protected Credentials](../insufficiently-protected-credentials.md)

The credentials provided are not adequately protected against security threats.

### [Sensitive Information Leak](../sensitive-information-leak.md)

Neglecting sensitive information can lead to severe data leaks and breaches.

### [Untrusted Deserialization](../untrusted-deserialization.md)

User input is deserialized.

### [Log Injection](../log-injection.md)

Input from the user may be logged, giving false data.

### [XML External Entity](../xml-external-entity.md)

Objects that parse or handle XML can lead to XML External Entity (XXE) attacks when misconfigured.

### [Path Injection](../path-traversal.md)

User input may lead to opening unintended files.

### [Http to File Access](../http-to-file-access.md)

Hardcoded download and writing of potentially harmful file.

### [Code Injection](../code-injection.md)

User input is used in eval command.

### [OS Command Injection](../os-command-injection.md)

Possible unintended system commands could be executed through user input.

### [Resource leak](../resource-leak.md)

Allocated resources are not released properly.

### [Cross Site Scripting (XSS)](../cross-site-scripting.md)

Improper neutralization of input during web page generation ('Cross-site Scripting')

### [Untrusted Open](../untrusted-file-open.md)

Non-static variables used to open files.

### [Improper Input Validation](../improper-input-validation.md)

Improper input validation can lead to security vulnerabilities and data breaches.

### [Stack Trace Exposure](../stack-trace-exposure.md)

Stack trace shows software architecture.

### [Improper Certificate Validation](../improper-certificate-validation.md)

Lack of validation of a security certificate can lead to host impersonation and sensitive data leaks.

### [send\_file Injection](../sendfile-injection.md)

External Control of File Name or Path.

### [Unsafe File Permissions](../loose-file-permissions.md)

Setting potentially harmful access rights

### [Tainted Format](../tainted-format.md)

User input decides output information.
