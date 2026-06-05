---
title: "Log injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Log injection [High](severity/high.md)

User-provided inputs must be sanitized before they are logged. An attacker can use unsanitized input to break a log's integrity, forge log entries, or bypass log monitors.

**Detector ID**

kotlin/log-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-117](https://cwe.mitre.org/data/definitions/117.html)

**Tags**

[\# data-integrity](tags/data-integrity.md) [\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Unsanitized user data is being written to the logs
2fun noncompliant(request: ServletRequest) {
3    val xValue = request.getParameter("x")
4    logger.info("Value is: {}", xValue)
5}

```

#### Compliant example

```kotlin
1// Compliant: There is no user input being written to the logs.
2fun compliant(input: String) {
3    logger.info("Value is: {}", input)
4}

```

All content copied from https://docs.aws.amazon.com/.
