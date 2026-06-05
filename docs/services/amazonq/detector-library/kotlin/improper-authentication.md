---
title: "Improper Authentication High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Improper Authentication [High](severity/high.md)

Failing to properly verify user identities and authenticate against strong credentials enables attackers to bypass authentication controls. Weaknesses like hardcoded, empty, or missing credential checks allow unauthorized system and data access. User identities must be verified against secure credentials retrieved from env vars, vaults etc. before granting access. Proper authentication controls including credential strength verification are essential to prevent malicious login and account compromise.

**Detector ID**

kotlin/improper-authentication@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-287](https://cwe.mitre.org/data/definitions/287.html)

**Tags**

-

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Permiting anonymous users to execute LDAP statements
2fun noncompliant(env: Environment): Void {
3    env.put(Context.SECURITY_AUTHENTICATION, "none")
4    val ctx: DirContext = InitialDirContext(env)
5}

```

#### Compliant example

```kotlin
1// Compliant: Enforcing authentication for LDAP statements
2fun compliant(env: Environment): Void {
3    env.put(Context.SECURITY_AUTHENTICATION, "simple")
4    val ctx: DirContext = InitialDirContext(env)
5}

```

All content copied from https://docs.aws.amazon.com/.
