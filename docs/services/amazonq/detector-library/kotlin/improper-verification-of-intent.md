---
title: "Improper verification of Intent High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Improper verification of Intent [High](severity/high.md)

Intent receiver method is registered without specifying any broadcast permission. Other applications can send potentially malicious broadcasts, so it is important to limit the applications that can send broadcasts to the receiver.

**Detector ID**

kotlin/improper-verification-of-intent@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-926](https://cwe.mitre.org/data/definitions/926.html) [CWE-925](https://cwe.mitre.org/data/definitions/925.html)

**Tags**

-

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Intent receiver method is registered without specifying any broadcast permission
2fun noncompliant(
3    context: Context, receiver: BroadcastReceiver?,
4    filter: IntentFilter?,
5    scheduler: Handler?,
6    flags: Int
7) {
8    context.registerReceiver(receiver, filter) // Sensitive
9
10    context.registerReceiver(receiver, filter, flags) // Sensitive
11
12}

```

#### Compliant example

```kotlin
1// Compliant: Intent receiver method is registered with a limiting broadcasting permission.
2fun compliant(
3    context: Context, receiver: BroadcastReceiver?,
4    filter: IntentFilter?,
5    broadcastPermission: String?,
6    scheduler: Handler?,
7    flags: Int
8) {
9    context.registerReceiver(receiver, filter, broadcastPermission, scheduler)
10}

```

All content copied from https://docs.aws.amazon.com/.
