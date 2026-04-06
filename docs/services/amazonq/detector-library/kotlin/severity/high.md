![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](../insecure-cookie.md) [Cookie Without Http Only Flag](../sensitive-cookie-without-http-only-flag.md) [Improper Authentication](../improper-authentication.md) [Cryptographic key generator](../cryptographic-key-generator.md) [Weak pseudorandom number generation](../weak-random-number-generation.md) [Path traversal](../path-traversal.md) [Cross-site scripting](../cross-site-scripting.md) [Reusing Nonce and key in encryption](../reusing-nonce-key-in-encryption.md) [Code Injection](../code-injection.md) [Server-side request forgery](../server-side-request-forgery.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Log injection](../log-injection.md) [Hardcoded credentials](../hardcoded-credentials.md) [Enabling and overriding debug feature](../detect-activated-debug-feature.md) [Null Pointer Dereference](../null-pointer-dereference.md) [Insecure hashing](../insecure-hashing.md) [Missing encryption of sensitive data](../missing-encryption-of-sensitive-data.md) [Improper verification of Intent](../improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [OS Command Injection](../os-command-injection.md) [Insecure Bean Validation](../insecure-bean-validation.md) [SQL injection](../sql-injection.md) [Incorrect Type Conversion](../incorrect-type-conversion.md)

# High

Showing all detectors for the Kotlin language withhigh severity.

### [Insecure cookie](../insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Cookie Without Http Only Flag](../sensitive-cookie-without-http-only-flag.md)

Sensitive cookie without 'HttpOnly' flag

### [Improper Authentication](../improper-authentication.md)

Improper authentication from insufficient identity verification.

### [Cryptographic key generator](../cryptographic-key-generator.md)

Insufficient key sizes can lead to brute force attacks.

### [Weak pseudorandom number generation](../weak-random-number-generation.md)

Insufficiently random generators (or hardcoded seeds) can make pseudorandom sequences predictable.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Reusing Nonce and key in encryption](../reusing-nonce-key-in-encryption.md)

GCM Cipher with reused initialization vector is detected.

### [Server-side request forgery](../server-side-request-forgery.md)

Server-side request forgery (SSRF) is a vulnerability that allows an attacker to manipulate a web application to make unintended requests from the server.

### [Cross-site request forgery](../cross-site-request-forgery.md)

Insecure configuration can lead to a cross-site request forgery (CSRF) vulnerability.

### [Log injection](../log-injection.md)

Using untrusted inputs in a log statement can enable attackers to break the log's format, forge log entries, and bypass log monitors.

### [Insecure hashing](../insecure-hashing.md)

Obsolete, broken, or weak hashing algorithms can lead to security vulnerabilities.

### [Missing encryption of sensitive data](../missing-encryption-of-sensitive-data.md)

The product does not encrypt sensitive or critical information before storage or transmission.

### [Improper verification of Intent](../improper-verification-of-intent.md)

Intent receiver method is registered without specifying any broadcast permission.

### [Insecure connection using unencrypted protocol](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [OS Command Injection](../os-command-injection.md)

Possible unintended system commands could be executed through user input.

### [Insecure Bean Validation](../insecure-bean-validation.md)

Passing user-controlled input directly to bean validation APIs can lead to code injection attacks.

### [SQL injection](../sql-injection.md)

Use of untrusted inputs in SQL database query can enable attackers to read, modify, or delete sensitive data in the database

### [Incorrect Type Conversion](../incorrect-type-conversion.md)

Failure to properly transform an object, resource, or structure from one type to a safer one.
