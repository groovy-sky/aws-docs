![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](path-traversal.md) [Avoiding Exceptions in PHP](avoid-exit-die.md) [OS command injection](os-command-injection.md) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/php/incorrect-comparison) [Ldap Bind Without Password](https://docs.aws.amazon.com/amazonq/detector-library/php/ldap-bind-without-password) [Sendfile Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sendfile-injection) [Assert Use](https://docs.aws.amazon.com/amazonq/detector-library/php/assert-use) [Loose file permissions](https://docs.aws.amazon.com/amazonq/detector-library/php/loose-file-permissions) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-authentication) [Insecure connection](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-connection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/php/weak-random-number-generation) [Open Redirect](https://docs.aws.amazon.com/amazonq/detector-library/php/open-redirect) [Allow Url Fopen Or Include](https://docs.aws.amazon.com/amazonq/detector-library/php/allow-url-fopen-or-include) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cryptography) [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/php/object-input-stream-insecure-deserialization) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-cookie-without-http-only-flag) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection) [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Incorrect Comparison [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

Encouraging subject practices in comparing hash values or sensitive data, this recommendation underscores the use of type-safe comparison ( `===`) over loose equality ( `==`). Mitigating risk tied to type juggling, this enhances reliability and security in the codebase.

**Detector ID**

php/incorrect-comparison@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-697](https://cwe.mitre.org/data/definitions/697.html)

**Tags**

-

* * *

#### Noncompliant example

```php
1// Noncompliant: Used loose equality (`==`)
2md5("240610708") == "0";
```

#### Compliant example

```php
1// Compliant: used type-safe comparison (`===`)
2md5("240610708") === "0";
```
