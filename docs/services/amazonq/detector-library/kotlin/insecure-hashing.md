![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Insecure hashing [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

A hashing algorithm is weak if it is easy to determine the original input from the hash or to find another input that yields the same hash. Weak hashing algorithms can lead to security vulnerabilities.

**Detector ID**

kotlin/insecure-hashing@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-327](https://cwe.mitre.org/data/definitions/327.html) [CWE-328](https://cwe.mitre.org/data/definitions/328.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Used `NullCipher`, which will not use any encryption.
2fun noncompliant(plainText: String): Array<Byte> {
3    val doNothingCipher: Cipher = NullCipher()
4    val cipherText: Cipher = doNothingCihper.doFinal(plainText)
5    return cipherText
6}

```

#### Compliant example

```kotlin
1// Compliant: Avoided use of `NullCipher` while encrypting value
2fun compliant(plainText: String): Void {
3    val cipher: Cipher = Cipher.getInstance("AES/CBC/PKCS5Padding")
4    val cipherText: Array<Byte> = cipher.doFinal(plainText)
5    return cipherText
6}

```

#### Noncompliant example

```kotlin
1// Noncompliant: Using weak hashing algorithm which is insecure
2fun noncompliant(password: String): ByteArray {
3    val md5Digest: MessageDigest = MessageDigest.getInstance("MD5")
4    md5Digest.update(password.getBytes())
5    val hashValue: ByteArray = md5Digest.digest()
6    return hashValue
7}

```

#### Compliant example

```kotlin
1// Compliant: Using secure hashing algorithm
2fun compliant(password: String): ByteArray {
3    val shaDigest: MessageDigest = MessageDigest.getInstance("SHA-256")
4    shaDigest.update(password.getBytes())
5    val hashValue: ByteArray = shaDigest.digest()
6    return hashValue
7}

```
