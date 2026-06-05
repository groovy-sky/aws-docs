---
title: "Cross-site request forgery High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](sql-injection.md) [Incorrect Type Conversion](incorrect-type-conversion.md)

# Cross-site request forgery [High](severity/high.md)

Insecure configuration can lead to a cross-site request forgery (CSRF) vulnerability. This can enable an attacker to trick end users into performing unwanted actions while authenticated.

**Detector ID**

kotlin/cross-site-request-forgery@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-352](https://cwe.mitre.org/data/definitions/352.html)

**Tags**

[\# configuration](tags/configuration.md) [\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: CSRF protection disabled
2@Configuration
3@EnableWebSecurity
4class WebSecurityConfig : WebSecurityConfigurerAdapter() {
5    @Throws(Exception::class)
6    protected fun configure(http: HttpSecurity) {
7        http {
8            csrf().disable()
9            // Other security configurations...
10        }
11    }
12    }

```

#### Compliant example

```kotlin
1// Compliant: By default CSRF protection is enabled
2@Configuration
3@EnableWebSecurity
4class WebSecurityConfig : WebSecurityConfigurerAdapter() {
5
6    @Throws(Exception::class)
7    override fun configure(http: HttpSecurity) {
8        http {
9            csrf { }
10            // Other security configurations...
11        }
12    }
13}

```

All content copied from https://docs.aws.amazon.com/.
