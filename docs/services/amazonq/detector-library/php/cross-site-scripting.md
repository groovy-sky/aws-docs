![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/php/cross-site-scripting) [Dangerous Function Usage](https://docs.aws.amazon.com/amazonq/detector-library/php/dangerous-function-usage) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/php/path-traversal) [Avoiding Exceptions in PHP](https://docs.aws.amazon.com/amazonq/detector-library/php/avoid-exit-die) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/php/os-command-injection) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/php/incorrect-comparison) [Ldap Bind Without Password](https://docs.aws.amazon.com/amazonq/detector-library/php/ldap-bind-without-password) [Sendfile Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sendfile-injection) [Assert Use](https://docs.aws.amazon.com/amazonq/detector-library/php/assert-use) [Loose file permissions](https://docs.aws.amazon.com/amazonq/detector-library/php/loose-file-permissions) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-authentication) [Insecure connection](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-connection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/php/weak-random-number-generation) [Open Redirect](https://docs.aws.amazon.com/amazonq/detector-library/php/open-redirect) [Allow Url Fopen Or Include](https://docs.aws.amazon.com/amazonq/detector-library/php/allow-url-fopen-or-include) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cryptography) [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/php/object-input-stream-insecure-deserialization) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-cookie-without-http-only-flag) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection) [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Cross-site scripting [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

User-controllable input must be sanitized before it's included in output used to dynamically generate a web page. Unsanitized user input can introduce cross-side scripting (XSS) vulnerabilities that can lead to inadvertedly running malicious code in a trusted context.

**Detector ID**

php/cross-site-scripting@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-79](https://cwe.mitre.org/data/definitions/79.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/top25-cwes)

* * *

#### Noncompliant example

```php
1function nonCompliant() {
2    $name = $_REQUEST['name'];
3    // Noncompliant: '$name' statement is non-compliant as it can leave the application vulnerable to cross-site scripting attacks.
4    echo "Hello :".$name;
5}

```

#### Compliant example

```php
1function compliant() {
2    $name = $_REQUEST['name'];
3    // Compliant: 'htmlentities' provides a compliant way to display user input safely.
4    print("Hello : " . htmlentities($name));
5}

```
