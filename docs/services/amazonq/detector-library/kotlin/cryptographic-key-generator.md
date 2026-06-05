---
title: "Cryptographic key generator High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Cryptographic key generator [High](severity/high.md)

Insufficient key sizes used for an HMAC are not robust against brute force attacks. Even strong encryption algorithms are vulnerable to brute force attacks when small key sizes are used.

**Detector ID**

kotlin/cryptographic-key-generator@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-326](https://cwe.mitre.org/data/definitions/326.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: `DefaultHttpClient` is used for setting up HTTP connection.
2fun noncompliant() {
3    val client: HttpClient = DefaultHttpClient()
4    val request: HttpGet = HttpGet("http://google.com")
5    val response: HttpResponse= client.execute(request)
6}

```

#### Compliant example

```kotlin
1// Compliant: `DefaultHttpClient` is not used for setting up HTTP connection.
2fun compliant() {
3    val client: HttpClient = SystemDefaultHttpClient()
4    val request: HttpGet = HttpGet("http://google.com")
5    val response: HttpResponse= client.execute(request)
6}

```

#### Noncompliant example

```kotlin
1// Noncompliant: The key 256 is not secure key length.
2fun noncompliant() {
3    val keyGen: KeyPairGenerator = KeyPairGenerator.getInstance("RSA")
4    keyGen.initialize(256)
5}

```

#### Compliant example

```kotlin
1// Compliant: The key 2048 is secure key length.
2fun compliant() {
3    val keyGen = KeyPairGenerator.getInstance("RSA")
4    keyGen.initialize(2048);
5}

```

All content copied from https://docs.aws.amazon.com/.
