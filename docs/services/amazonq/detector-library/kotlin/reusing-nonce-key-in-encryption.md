---
title: "Reusing Nonce and key in encryption High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Reusing Nonce and key in encryption [High](severity/high.md)

GCM Cipher with reused initialization vector is detected. Using the same initialization vector and Key to encrypt numerous plaintext blocks allows an attacker to compare the ciphertexts and glean information about the encrypted data based on assumptions about its content.

**Detector ID**

kotlin/reusing-nonce-key-in-encryption@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-323](https://cwe.mitre.org/data/definitions/323.html)

**Tags**

-

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: GCM Cipher with reused initialization vector `theInnerIV2` is detected.
2fun noncompliant(clearText: String): String {
3    val cipher1: Cipher = Cipher.getInstance("AES/GCM/NoPadding")
4    val cipher2: Cipher = Cipher.getInstance("AES/GCM/NoPadding")
5    val keySpec: SecretKeySpec = SecretKeySpec(theKey.getEncoded(), "AES")
6
7    private val theInnerIV2: Array<Byte>
8
9    val gcmParameterSpec: GCMParameterSpec = GCMParameterSpec(GCM_TAG_LENGTH * 8, theInnerIV2)
10    cipher1.init(Cipher.DECRYPT_MODE, keySpec, gcmParameterSpec)
11    val gcmParameterSpec: GCMParameterSpec = GCMParameterSpec(GCM_TAG_LENGTH * 8, theInnerIV2)
12    cipher2.init(Cipher.DECRYPT_MODE, keySpec, gcmParameterSpec)
13
14    val decoded: Array<Byte> = Base64.getDecoder().decode(cipherText)
15    val decryptedText: Array<Byte> = cipher.doFinal(decoded)
16    return String(decryptedText)
17}

```

#### Compliant example

```kotlin
1// Compliant: GCM Cipher with new initialization vector is used
2fun compliant(clearText: String): String {
3    val cipher: Cipher = Cipher.getInstance("AES/CBC/PKCS5Padding")
4
5    val keySpec: SecretKeySpec= SecretKeySpec(theKey.getEncoded(), "AES")
6
7    private val theInnerIV: Array<Byte>
8    val gcmParameterSpec: GCMParameterSpec = GCMParameterSpec(GCM_TAG_LENGTH * 8, theInnerIV)
9    cipher.init(Cipher.ENCRYPT_MODE, keySpec, gcmParameterSpec)
10
11    val cipherText: Array<Byte> = cipher.doFinal(clearText.getBytes())
12
13    val encoded = base64.getEncoder().encodeToString(cipherText)
14    return encoded
15}

```

All content copied from https://docs.aws.amazon.com/.
