![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# OS Command Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

User input influences a system command. This allows a malicious user to inject custom commands and take control of a system. This can be sanitized with shellescape to avoid injection.

**Detector ID**

kotlin/os-command-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-78](https://cwe.mitre.org/data/definitions/78.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/top25-cwes)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: User input is being passed to `exec`
2fun noncompliant() {
3    print("Enter your input:")
4    val input = readLine()
5
6    val command = "echo Hello, $input"
7    val process = Runtime.getRuntime().exec(String.format("The value is: %s", command))
8}

```

#### Compliant example

```kotlin
1// Compliant: Hardcoded value is being passed to `exec`
2fun compliant() {
3    val directory = "hardcoded_value"
4
5    val command = "ls -l " + directory
6    val r: Runtime = Runtime.getRuntime()
7    val process = r.exec(command)
8}

```
