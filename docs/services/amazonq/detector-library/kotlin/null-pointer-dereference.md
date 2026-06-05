---
title: "Null Pointer Dereference Medium"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Null Pointer Dereference [Medium](severity/medium.md)

Dereferencing a null pointer can lead to unexpected null pointer exceptions. Consider adding a null check before dereferencing the pointer.

**Detector ID**

kotlin/null-pointer-dereference@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-476](https://cwe.mitre.org/data/definitions/476.html)

**Tags**

[\# top10-cwes](tags/top10-cwes.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Dereferencing freed pointer
2fun noncompliant() {
3    val byteBuffer = ByteBuffer.allocateDirect(Int.SIZE_BYTES)
4    val ptr = byteBuffer.asIntBuffer()
5    byteBuffer.clear()
6    val value = ptr[0]
7}

```

#### Compliant example

```kotlin
1// Compliant: Added a null check before dereferencing the pointer.
2fun compliant() {
3    val byteBuffer = ByteBuffer.allocateDirect(Int.SIZE_BYTES)
4    val ptr = byteBuffer.asIntBuffer()
5    if(ptr != null) {
6        val value = ptr[0]
7        byteBuffer.clear()
8    }
9}

```

All content copied from https://docs.aws.amazon.com/.
