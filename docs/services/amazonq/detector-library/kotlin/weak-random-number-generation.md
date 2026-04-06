![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Weak pseudorandom number generation [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

Insufficiently random generators or hardcoded seeds can make pseudorandom sequences predictable, which may lead to security vulnerabilities.

**Detector ID**

kotlin/weak-random-number-generation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-338](https://cwe.mitre.org/data/definitions/338.html) [CWE-330](https://cwe.mitre.org/data/definitions/330.html) [CWE-326](https://cwe.mitre.org/data/definitions/326.html) [CWE-1241](https://cwe.mitre.org/data/definitions/1241.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/cryptography) [\# security-context](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/security-context) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10)

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
