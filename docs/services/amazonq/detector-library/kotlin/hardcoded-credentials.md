---
title: "Hardcoded credentials Critical"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Hardcoded credentials [Critical](severity/critical.md)

Hardcoded credentials can be intercepted by malicious actors. Even after removing them from the code they may still pose a risk because an attacker might have recorded them to use them at a later point in time.

**Detector ID**

kotlin/hardcoded-credentials@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-798](https://cwe.mitre.org/data/definitions/798.html)

**Tags**

[\# secrets](tags/secrets.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Password is hardcoded
2fun noncompliant() {
3    val username = "admin"
4    val password = "password123"
5    val ssh = SSHClient()
6    ssh.authPassword(username, password)
7}

```

#### Compliant example

```kotlin
1// Compliant: Password is retrieved from environment variables.
2fun compliant() {
3    val username = System.getenv("SSH_USERNAME")
4    val password = System.getenv("SSH_PASSWORD")
5    val ssh = SSHClient()
6    ssh.authPassword(username, password)
7}

```

All content copied from https://docs.aws.amazon.com/.
