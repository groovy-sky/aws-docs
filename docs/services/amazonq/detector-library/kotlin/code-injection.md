---
title: "Code Injection Critical"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Code Injection [Critical](severity/critical.md)

Code injection occurs when an application executes untrusted code from an attacker. User input gets concatenated with code. The input is executed without validation or sanitization.

**Detector ID**

kotlin/code-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-94](https://cwe.mitre.org/data/definitions/94.html)

**Tags**

[\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md)

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

All content copied from https://docs.aws.amazon.com/.
