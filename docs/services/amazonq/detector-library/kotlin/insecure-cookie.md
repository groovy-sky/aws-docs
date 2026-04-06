![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-cookie) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sensitive-cookie-without-http-only-flag) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-authentication) [Cryptographic key generator](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cryptographic-key-generator) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Insecure cookie [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

Insecure cookie settings can lead to unencrypted cookie transmission. Even if a cookie doesn't contain sensitive data now, sensitive data could be added later. It's good practice to transmit all cookies only through secure channels.

**Detector ID**

kotlin/insecure-cookie@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-614](https://cwe.mitre.org/data/definitions/614.html) [CWE-539](https://cwe.mitre.org/data/definitions/539.html)

**Tags**

[\# cookies](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/cookies) [\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant:  The `setSecure` attribute of a cookie is set to `false`
2fun noncompliant() {
3    var cookie: Cookie = Cookie("cookie", value)
4    cookie.setSecure(false)
5    cookie.setHttpOnly(false)
6    response.addCookie(cookie)
7}

```

#### Compliant example

```kotlin
1// Compliant: The `setSecure` attribute of a cookie is set to `true`
2fun compliant(@RequestParam value: String, response: HttpServletResponse) {
3    var cookie: Cookie = Cookie("cookie", value)
4    cookie.setSecure(true)
5    cookie.setHttpOnly(true)
6    response.addCookie(cookie)
7}

```
