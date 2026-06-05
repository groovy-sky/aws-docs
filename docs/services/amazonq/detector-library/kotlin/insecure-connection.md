---
title: "Insecure connection using unencrypted protocol High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Insecure connection using unencrypted protocol [High](severity/high.md)

Connections that use insecure protocols transmit data in cleartext. This introduces a risk of exposing sensitive data to third parties.

**Detector ID**

kotlin/insecure-connection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-319](https://cwe.mitre.org/data/definitions/319.html) [CWE-200](https://cwe.mitre.org/data/definitions/200.html)

**Tags**

[\# cryptography](tags/cryptography.md) [\# information-leak](tags/information-leak.md) [\# networking](tags/networking.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Using clear-text protocols such as `ftp` which is insecure
2fun noncompliant() {
3    val ftpClient = FTPClient(); // Sensitive
4}

```

#### Compliant example

```kotlin
1// Compliant: Using clear-text protocols such as `ftps` which is secure
2fun compliant() {
3    val ftpsClient = FTPSClient(true);
4}

```

All content copied from https://docs.aws.amazon.com/.
