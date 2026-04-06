![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-cookie) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sensitive-cookie-without-http-only-flag) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-authentication) [Cryptographic key generator](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cryptographic-key-generator) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Kotlin detectors

Showing all detectors for the Kotlin language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories)

* * *

### Browse all detectors

### [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-cookie)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sensitive-cookie-without-http-only-flag)

Sensitive cookie without 'HttpOnly' flag

### [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-authentication)

Improper authentication from insufficient identity verification.

### [Cryptographic key generator](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cryptographic-key-generator)

Insufficient key sizes can lead to brute force attacks.

### [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation)

Insufficiently random generators (or hardcoded seeds) can make pseudorandom sequences predictable.

### [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption)

GCM Cipher with reused initialization vector is detected.

### [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection)

Code injection occurs when an application executes untrusted code from an attacker.

### [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery)

Server-side request forgery (SSRF) is a vulnerability that allows an attacker to manipulate a web application to make unintended requests from the server.

### [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery)

Insecure configuration can lead to a cross-site request forgery (CSRF) vulnerability.

### [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection)

Using untrusted inputs in a log statement can enable attackers to break the log's format, forge log entries, and bypass log monitors.

### [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials)

Hardcoded credentials can be intercepted by malicious actors.

### [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature)

The Debug feature should not be enabled or overridden.

### [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference)

Dereferencing a null pointer can lead to unexpected null pointer exceptions.

### [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing)

Obsolete, broken, or weak hashing algorithms can lead to security vulnerabilities.

### [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data)

The product does not encrypt sensitive or critical information before storage or transmission.

### [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent)

Intent receiver method is registered without specifying any broadcast permission.

### [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection)

Possible unintended system commands could be executed through user input.

### [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation)

Passing user-controlled input directly to bean validation APIs can lead to code injection attacks.

### [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection)

Use of untrusted inputs in SQL database query can enable attackers to read, modify, or delete sensitive data in the database

### [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

Failure to properly transform an object, resource, or structure from one type to a safer one.
