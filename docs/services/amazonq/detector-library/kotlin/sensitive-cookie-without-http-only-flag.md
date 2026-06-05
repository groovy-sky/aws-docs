---
title: "Cookie Without Http Only Flag High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Cookie Without Http Only Flag [High](severity/high.md)

The `HttpOnly` attribute when set to `true` protects the cookie value from being accessed by client side JavaScript such as reading the `document.cookie` values. By enabling this protection, a website that is vulnerable to Cross-Site Scripting (XSS) will be able to block malicious scripts from accessing the cookie value from JavaScript.

**Detector ID**

kotlin/sensitive-cookie-without-http-only-flag@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-1004](https://cwe.mitre.org/data/definitions/1004.html)

**Tags**

-

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: The `httponly` attribute of cookies is set to `false`
2fun noncompliant(value: String, response: HttpServletResponse) {
3    val cookie: Cookie = Cookie("cookie", value)
4    cookie.setHttpOnly(false)
5    response.addCookie(cookie)
6}

```

#### Compliant example

```kotlin
1// Compliant: The `httponly` attribute of cookies is set to `true`
2fun compliant(value: String, response: HttpServletResponse) {
3    val cookie: Cookie = Cookie("cookie", value)
4    cookie.setSecure(true)
5    cookie.setHttpOnly(true)
6    response.addCookie(cookie)
7}

```

All content copied from https://docs.aws.amazon.com/.
