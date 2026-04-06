![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](path-traversal.md) [Avoiding Exceptions in PHP](avoid-exit-die.md) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/php/os-command-injection) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/php/incorrect-comparison) [Ldap Bind Without Password](https://docs.aws.amazon.com/amazonq/detector-library/php/ldap-bind-without-password) [Sendfile Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sendfile-injection) [Assert Use](https://docs.aws.amazon.com/amazonq/detector-library/php/assert-use) [Loose file permissions](https://docs.aws.amazon.com/amazonq/detector-library/php/loose-file-permissions) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-authentication) [Insecure connection](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-connection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/php/weak-random-number-generation) [Open Redirect](https://docs.aws.amazon.com/amazonq/detector-library/php/open-redirect) [Allow Url Fopen Or Include](https://docs.aws.amazon.com/amazonq/detector-library/php/allow-url-fopen-or-include) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cryptography) [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/php/object-input-stream-insecure-deserialization) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-cookie-without-http-only-flag) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection) [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# OS command injection [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

OS command injection is a critical vulnerability that can lead to a full system compromise as it may allow an adversary to pass in arbitrary commands or arguments to be executed.

**Detector ID**

php/os-command-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-78](https://cwe.mitre.org/data/definitions/78.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/top25-cwes)

* * *

#### Noncompliant example

```php
1$username = $_COOKIE['username'];
2// Noncompliant: Incorporating variable into command strings
3exec("wto -n \"$username\" -g", $ret);
```

#### Compliant example

```php
1$fullpath = $_POST['fullpath'];
2// Compliant: escapeshellarg() is used to prevent command injection
3$filesize = trim(shell_exec('stat -c %s ' . escapeshellarg($fullpath)));
```
