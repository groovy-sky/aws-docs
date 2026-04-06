![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Server-side request forgery [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

Server-side request forgery (SSRF) is a web application vulnerability where an attacker can cause the server to make requests to unintended resources or systems. This can lead to unauthorized access to data or systems that the server can access but the attacker cannot directly access. Proper input validation, whitelisting, and access controls are necessary to mitigate SSRF vulnerabilities.

**Detector ID**

kotlin/server-side-request-forgery@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-918](https://cwe.mitre.org/data/definitions/918.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/top25-cwes)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: User input data passed to respondText().
2fun noncompliant() {
3    embeddedServer(Netty, port = 8080) {
4        routing {
5            post("/proxy/{url}") {
6                val url = call.request.queryParameters["url"]
7                if (url != null) {
8                    val data = URL(url).readText()
9                    call.respondText(data)
10                }
11            }
12        }
13    }.start(wait = true)
14}

```

#### Compliant example

```kotlin
1// Compliant: Pre-defined data passed to respondText().
2fun compliant() {
3    embeddedServer(Netty, port = 8080) {
4        routing {
5            post("/proxy/{url}") {
6                val url = "<hardcoded_url>"
7                if (url != null) {
8                    val data = URL(url).readText()
9                    call.respondText(data)
10                }
11            }
12        }
13    }.start(wait = true)
14}

```
