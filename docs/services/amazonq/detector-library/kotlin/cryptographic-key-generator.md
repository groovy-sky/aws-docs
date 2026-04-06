![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cryptographic-key-generator) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Cryptographic key generator [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

Insufficient key sizes used for an HMAC are not robust against brute force attacks. Even strong encryption algorithms are vulnerable to brute force attacks when small key sizes are used.

**Detector ID**

kotlin/cryptographic-key-generator@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-326](https://cwe.mitre.org/data/definitions/326.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10)

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
