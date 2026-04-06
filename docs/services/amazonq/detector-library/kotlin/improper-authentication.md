![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-authentication) [Cryptographic key generator](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cryptographic-key-generator) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/weak-random-number-generation) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/path-traversal) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-scripting) [Reusing Nonce and key in encryption](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/reusing-nonce-key-in-encryption) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/code-injection) [Server-side request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/server-side-request-forgery) [Cross-site request forgery](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/cross-site-request-forgery) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/log-injection) [Hardcoded credentials](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/hardcoded-credentials) [Enabling and overriding debug feature](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/detect-activated-debug-feature) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/null-pointer-dereference) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-hashing) [Missing encryption of sensitive data](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/missing-encryption-of-sensitive-data) [Improper verification of Intent](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/improper-verification-of-intent) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-connection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/os-command-injection) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Improper Authentication [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

Failing to properly verify user identities and authenticate against strong credentials enables attackers to bypass authentication controls. Weaknesses like hardcoded, empty, or missing credential checks allow unauthorized system and data access. User identities must be verified against secure credentials retrieved from env vars, vaults etc. before granting access. Proper authentication controls including credential strength verification are essential to prevent malicious login and account compromise.

**Detector ID**

kotlin/improper-authentication@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-287](https://cwe.mitre.org/data/definitions/287.html)

**Tags**

-

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Permiting anonymous users to execute LDAP statements
2fun noncompliant(env: Environment): Void {
3    env.put(Context.SECURITY_AUTHENTICATION, "none")
4    val ctx: DirContext = InitialDirContext(env)
5}

```

#### Compliant example

```kotlin
1// Compliant: Enforcing authentication for LDAP statements
2fun compliant(env: Environment): Void {
3    env.put(Context.SECURITY_AUTHENTICATION, "simple")
4    val ctx: DirContext = InitialDirContext(env)
5}

```
