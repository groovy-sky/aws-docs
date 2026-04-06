![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](insecure-bean-validation.md) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# SQL injection [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

User-provided inputs must be sanitized before being used to generate a SQL database query. An untrusted input can be intentionally built by an attacker in order to run unwanted query statements, possibly allowing the attacker to read, modify, or delete database content.

**Detector ID**

kotlin/sql-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/injection) [\# sql](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/sql) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/top25-cwes)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: User data is being used in SQL query
2fun noncompliant(connection: Connection): ResultSet {
3    print("Enter your userId:")
4    val userId = readLine()
5    val query = "SELECT * FROM users WHERE userId = '$userId'"
6    val statement = connection.createStatement()
7    return statement.executeQuery(query)
8}

```

#### Compliant example

```kotlin
1// Compliant: Hardcoded value is being passed to SQL query
2fun compliant(connection: Connection): ResultSet {
3    val userId = "hardcoded_value"
4    val query = "SELECT * FROM users WHERE userId = '$userId'"
5    val statement = connection.createStatement()
6    return statement.executeQuery(query)
7}

```
