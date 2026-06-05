---
title: "Enabling and overriding debug feature Medium"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Enabling and overriding debug feature [Medium](severity/medium.md)

Don't enable or override an application's debug feature. Instead, use OS environment variables to set up the debug feature.

**Detector ID**

kotlin/detect-activated-debug-feature@v1.0

**Category**

[Code Quality](categories/code-quality.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-489](https://cwe.mitre.org/data/definitions/489.html) [CWE-215](https://cwe.mitre.org/data/definitions/215.html)

**Tags**

[\# efficiency](tags/efficiency.md) [\# maintainability](tags/maintainability.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: The Debug feature should not be enabled or overridden.
2fun noncompliant() {
3    WebView.setWebContentsDebuggingEnabled(true)
4}

```

#### Compliant example

```kotlin
1// Compliant: The Debug feature should not be enabled or overridden.
2fun compliant() {
3   WebView.setWebContentsDebuggingEnabled(false)
4}

```

All content copied from https://docs.aws.amazon.com/.
