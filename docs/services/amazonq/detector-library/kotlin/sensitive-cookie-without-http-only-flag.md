![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sensitive-cookie-without-http-only-flag) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-authentication) [Cryptographic key generator](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cryptographic-key-generator) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Cookie Without Http Only Flag [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

The `HttpOnly` attribute when set to `true` protects the cookie value from being accessed by client side JavaScript such as reading the `document.cookie` values. By enabling this protection, a website that is vulnerable to Cross-Site Scripting (XSS) will be able to block malicious scripts from accessing the cookie value from JavaScript.

**Detector ID**

kotlin/sensitive-cookie-without-http-only-flag@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

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
