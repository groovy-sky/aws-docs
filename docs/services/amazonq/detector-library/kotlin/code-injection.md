![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Code Injection [Critical](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/critical)

Code injection occurs when an application executes untrusted code from an attacker. User input gets concatenated with code. The input is executed without validation or sanitization.

**Detector ID**

kotlin/code-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-94](https://cwe.mitre.org/data/definitions/94.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: User input gets executed as a code.
2fun noncompliant() {
3    val shell = GroovyShell()
4    val script: String = request.getParameter("script")
5    shell.evaluate(script)
6}

```

#### Compliant example

```kotlin
1// Compliant: Pre-defined string gets executed as a code.
2fun compliant() {
3    val shell = GroovyShell()
4    val script: String = request.getParameter("script")
5    shell.evaluate("script")
6}

```

#### Noncompliant example

```kotlin
1// Noncompliant: User input gets executed as a expression.
2fun noncompliant() {
3    val input = request.getParameter("expr")
4    val jexl: JexlEngine = JexlBuilder().create()
5    val expression: JexlExpression = jexl.createExpression(input)
6    val context: JexlContext = MapContext()
7    expression.evaluate(context)
8}

```

#### Compliant example

```kotlin
1// Compliant: Pre-defined string gets executed as a expression.
2fun compliant() {
3    val input = "hardcoded-value"
4    val jexl: JexlEngine = JexlBuilder().create()
5    val expression: JexlExpression = jexl.createExpression(input)
6    val context: JexlContext = MapContext()
7    expression.evaluate(context)
8}

```
