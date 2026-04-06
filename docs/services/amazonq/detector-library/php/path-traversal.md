![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/php/path-traversal) [Avoiding Exceptions in PHP](https://docs.aws.amazon.com/amazonq/detector-library/php/avoid-exit-die) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/php/os-command-injection) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/php/incorrect-comparison) [Ldap Bind Without Password](https://docs.aws.amazon.com/amazonq/detector-library/php/ldap-bind-without-password) [Sendfile Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sendfile-injection) [Assert Use](https://docs.aws.amazon.com/amazonq/detector-library/php/assert-use) [Loose file permissions](https://docs.aws.amazon.com/amazonq/detector-library/php/loose-file-permissions) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-authentication) [Insecure connection](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-connection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/php/weak-random-number-generation) [Open Redirect](https://docs.aws.amazon.com/amazonq/detector-library/php/open-redirect) [Allow Url Fopen Or Include](https://docs.aws.amazon.com/amazonq/detector-library/php/allow-url-fopen-or-include) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cryptography) [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/php/object-input-stream-insecure-deserialization) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-cookie-without-http-only-flag) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection) [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Path Traversal [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

Creating file paths from untrusted input could allow a malicious actor to access arbitrary files on a disk by manipulating the file name in the path.

**Detector ID**

php/path-traversal@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html) [CWE-35](https://cwe.mitre.org/data/definitions/35.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/top25-cwes)

* * *

#### Noncompliant example

```php
1// Noncompliant: Direct utilization of path without adequate validation
2$path = '.../.../password';
3$localeFunctions = file_get_contents($path);
```

#### Compliant example

```php
1$user_input_compliant_4 = 'image.png';
2$path = BASE_PATH . "/" . $user_input_compliant_4;
3// Compliant: Validation of path before utilization
4if(realpath($path) !== BASE_PATH . $user_input_compliant_4) {
5    throw new InvalidPathException("Invalid path");
6}
7$json = file_get_contents($path);
```
