---
title: "Cross-site scripting High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Cross-site scripting [High](severity/high.md)

User-controllable input must be sanitized before it's included in output used to dynamically generate a web page. Unsanitized user input can introduce cross-side scripting (XSS) vulnerabilities that can lead to inadvertedly running malicious code in a trusted context.

**Detector ID**

kotlin/cross-site-scripting@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-79](https://cwe.mitre.org/data/definitions/79.html) [CWE-80](https://cwe.mitre.org/data/definitions/80.html)

**Tags**

[\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Using unsanitized external inputs which leads to XSS
2fun noncompliant() {
3    print("Enter your name:")
4    val name = readLine()
5
6    val writer = PrintWriter(System.out)
7    writer.write("<p>Hello, $name!</p>")
8}

```

#### Compliant example

```kotlin
1// Compliant: Not using any unsanitized external inputs
2fun compliant() {
3    print("Enter your name:")
4    val name = readLine()
5
6    val writer = PrintWriter(System.out)
7    writer.write("<p>Hello, name!</p>")
8}

```

#### Noncompliant example

```kotlin
1// Noncompliant: Enabled JavaScript support for WebViews
2fun noncompliant() {
3    val webView: WebView = findViewById(R.id.webview)
4    webView.getSettings().setJavaScriptEnabled(true) // Sensitive
5}

```

#### Compliant example

```kotlin
1// Compliant: Disabled JavaScript support for WebViews
2fun compliant() {
3    val webView: WebView = findViewById(R.id.webview)
4    webView.getSettings().setJavaScriptEnabled(false)
5}

```

All content copied from https://docs.aws.amazon.com/.
