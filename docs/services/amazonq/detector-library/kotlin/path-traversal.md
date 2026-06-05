---
title: "Path traversal High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Path traversal [High](severity/high.md)

Creating file paths from untrusted input could allow a malicious actor to access arbitrary files on a disk by manipulating the file name in the path.

**Detector ID**

kotlin/path-traversal@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html) [CWE-23](https://cwe.mitre.org/data/definitions/23.html)

**Tags**

[\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Using untrusted inputs to access a file path
2fun noncompliant() {
3    print("Enter your filename:")
4    val filename = readLine()
5
6    val file = File(filename)
7    val lines = file.readLines()
8    for (line in lines) {
9        println(line)
10    }
11}

```

#### Compliant example

```kotlin
1// Compliant: Using safe input to access a file path
2fun compliant(filename: String) {
3    val file = File(filename)
4    val lines = file.readLines()
5    for (line in lines) {
6        println(line)
7    }
8}

```

All content copied from https://docs.aws.amazon.com/.
