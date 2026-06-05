---
title: "Server-side request forgery High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Server-side request forgery [High](severity/high.md)

Server-side request forgery (SSRF) is a web application vulnerability where an attacker can cause the server to make requests to unintended resources or systems. This can lead to unauthorized access to data or systems that the server can access but the attacker cannot directly access. Proper input validation, whitelisting, and access controls are necessary to mitigate SSRF vulnerabilities.

**Detector ID**

kotlin/server-side-request-forgery@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-918](https://cwe.mitre.org/data/definitions/918.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

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

All content copied from https://docs.aws.amazon.com/.
