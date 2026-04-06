![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Incorrect Type Conversion [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

Failure to properly transform an object, resource, or structure from one type to a safer one.

**Detector ID**

kotlin/incorrect-type-conversion@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-704](https://cwe.mitre.org/data/definitions/704.html)

**Tags**

-

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Using `Integer.toHexString()` which creates a weak hash
2fun noncompliant(password: String): String {
3    val md: MessageDigest = MessageDigest.getInstance("SHA-1")
4    val resultBytes: Array<Byte> = md.digest(password.getBytes("UTF-8"))
5
6    var stringBuilder: StringBuilder = StringBuilder()
7    for (b in resultBytes) {
8        stringBuilder.append(Integer.toHexString(b and 0xFF))
9    }
10
11    return stringBuilder.toString()
12}

```

#### Compliant example

```kotlin
1// Compliant: Using `String.format(\"%02X\",...)` which does not creates a weak hash
2fun compliant(password: String): String {
3    val md: MessageDigest = MessageDigest.getInstance("SHA-1")
4    val resultBytes: Array<Byte> = md.digest(password.getBytes("UTF-8"))
5
6    var stringBuilder: StringBuilder = StringBuilder()
7    for (b in resultBytes) {
8        stringBuilder.append(String.format("%02X", b))
9    }
10
11    return stringBuilder.toString()
12}

```
