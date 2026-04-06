![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Kotlin detectors(23/23)

[Insecure cookie](insecure-cookie.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Improper Authentication](improper-authentication.md) [Cryptographic key generator](cryptographic-key-generator.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Path traversal](path-traversal.md) [Cross-site scripting](cross-site-scripting.md) [Reusing Nonce and key in encryption](reusing-nonce-key-in-encryption.md) [Code Injection](code-injection.md) [Server-side request forgery](server-side-request-forgery.md) [Cross-site request forgery](cross-site-request-forgery.md) [Log injection](log-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Enabling and overriding debug feature](detect-activated-debug-feature.md) [Null Pointer Dereference](null-pointer-dereference.md) [Insecure hashing](insecure-hashing.md) [Missing encryption of sensitive data](missing-encryption-of-sensitive-data.md) [Improper verification of Intent](improper-verification-of-intent.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [OS Command Injection](os-command-injection.md) [Insecure Bean Validation](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/insecure-bean-validation) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/sql-injection) [Incorrect Type Conversion](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/incorrect-type-conversion)

# Insecure Bean Validation [High](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/severity/high)

It's not advisable to pass user-controlled inputs to bean validation APIs. To mitigate this issue, consider implementing input validation and sanitization mechanisms. This can be achieved by using appropriate libraries or frameworks that provide built-in sanitization functions or by implementing custom validation logic specific to your application's needs.

**Detector ID**

kotlin/insecure-bean-validation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-94](https://cwe.mitre.org/data/definitions/94.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/owasp-top10) [\# injection](https://docs.aws.amazon.com/amazonq/detector-library/kotlin/tags/injection)

* * *

#### Noncompliant example

```kotlin
1// Noncompliant: Controlling the content of the message template supplied to `ConstraintValidatorContext.buildConstraintViolationWithTemplate()` parameter can result in arbitrary Java code execution.
2fun noncompliant(request: HttpServletRequest, response: HttpServletResponse, constraintContext: ConstraintValidatorContext) {
3    val constraintViolation: String = request.getAttribute("constraintViolation").toString()
4    constraintContext.buildConstraintViolationWithTemplate(constraintViolation)
5    .addConstraintViolation()
6    .disableDefaultConstraintViolation()
7}

```

#### Compliant example

```kotlin
1// Compliant: Safe Bean properties are passed to `buildConstraintViolationWithTemplate`
2fun compliant(request: HttpServletRequest, response: HttpServletResponse, constraintContext: ConstraintValidatorContext) {
3    val context: HibernateConstraintValidatorContext = constraintContext.unwrap(HibernateConstraintValidatorContext::class.java)
4    context.addMessageParameter("prop", request.getParameter("prop"))
5    context.buildConstraintViolationWithTemplate("{prop} is invalid").addConstraintViolation()
6    return false
7}

```
