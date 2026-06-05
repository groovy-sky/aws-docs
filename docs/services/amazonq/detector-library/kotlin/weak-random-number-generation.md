---
title: "Weak pseudorandom number generation High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Weak pseudorandom number generation [High](severity/high.md)

Insufficiently random generators or hardcoded seeds can make pseudorandom sequences predictable, which may lead to security vulnerabilities.

**Detector ID**

kotlin/weak-random-number-generation@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-338](https://cwe.mitre.org/data/definitions/338.html) [CWE-330](https://cwe.mitre.org/data/definitions/330.html) [CWE-326](https://cwe.mitre.org/data/definitions/326.html) [CWE-1241](https://cwe.mitre.org/data/definitions/1241.html)

**Tags**

[\# cryptography](tags/cryptography.md) [\# security-context](tags/security-context.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: `Random()` is not a secure random number generator
2fun noncompliant() {
3    val random = Random()
4    val bytes = ByteArray(20)
5    random.nextBytes(bytes)
6}

```

#### Compliant example

```kotlin
1// Compliant: Using `SecureRandom()` to generate random numbers
2fun compliant() {
3    val random = SecureRandom()
4    val bytes = ByteArray(20)
5    random.nextBytes(bytes)
6}

```

All content copied from https://docs.aws.amazon.com/.
