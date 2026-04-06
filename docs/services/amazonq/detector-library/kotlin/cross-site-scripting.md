![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Cross-site scripting [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

User-controllable input must be sanitized before it's included in output used to dynamically generate a web page. Unsanitized user input can introduce cross-side scripting (XSS) vulnerabilities that can lead to inadvertedly running malicious code in a trusted context.

**Detector ID**

kotlin/cross-site-scripting@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-79](https://cwe.mitre.org/data/definitions/79.html) [CWE-80](https://cwe.mitre.org/data/definitions/80.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/top25-cwes)

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
