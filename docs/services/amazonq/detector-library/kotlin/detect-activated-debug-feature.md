![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Enabling and overriding debug feature [Medium](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/medium)

Don't enable or override an application's debug feature. Instead, use OS environment variables to set up the debug feature.

**Detector ID**

kotlin/detect-activated-debug-feature@v1.0

**Category**

[Code Quality](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/code-quality)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-489](https://cwe.mitre.org/data/definitions/489.html) [CWE-215](https://cwe.mitre.org/data/definitions/215.html)

**Tags**

[\# efficiency](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/efficiency) [\# maintainability](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/maintainability)

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
