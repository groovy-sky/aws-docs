![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](path-traversal.md) [Avoiding Exceptions in PHP](avoid-exit-die.md) [OS command injection](os-command-injection.md) [Incorrect Comparison](incorrect-comparison.md) [Ldap Bind Without Password](ldap-bind-without-password.md) [Sendfile Injection](sendfile-injection.md) [Assert Use](assert-use.md) [Loose file permissions](loose-file-permissions.md) [Improper Authentication](improper-authentication.md) [Insecure connection](insecure-connection.md) [Weak Random Number Generation](weak-random-number-generation.md) [Open Redirect](open-redirect.md) [Allow Url Fopen Or Include](allow-url-fopen-or-include.md) [Insecure cryptography](insecure-cryptography.md) [Object Input Stream Insecure Deserialization](object-input-stream-insecure-deserialization.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Code Injection](code-injection.md) [Zip bomb attack](zip-bomb-attack.md) [Unsafe Reflection](unsafe-reflection.md) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Secure Signal Handling [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

Signaling processes or process groups without proper validation may lead to instability and potentialdenial of services. Validate parameters for secure coding.

**Detector ID**

php/secure-signal-handling@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-283](https://cwe.mitre.org/data/definitions/283.html)

**Tags**

-

* * *

#### Noncompliant example

```php
1function nonCompliant1($param)  {
2    $targetPid = (int)$_GET["pid"];
3    // Noncompliant: kills the process without validation
4    posix_kill($targetPid, 9);
5}

```

#### Compliant example

```php
1function compliant1($param)  {
2    $targetPid = (int)$_GET["pid"];
3    // Compliant: kills the process with validation
4    if (isValidPid($targetPid)) {
5        posix_kill($targetPid, 9);
6    }
7 }

```
