![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](path-traversal.md) [Avoiding Exceptions in PHP](avoid-exit-die.md) [OS command injection](os-command-injection.md) [Incorrect Comparison](incorrect-comparison.md) [Ldap Bind Without Password](ldap-bind-without-password.md) [Sendfile Injection](sendfile-injection.md) [Assert Use](assert-use.md) [Loose file permissions](loose-file-permissions.md) [Improper Authentication](improper-authentication.md) [Insecure connection](insecure-connection.md) [Weak Random Number Generation](weak-random-number-generation.md) [Open Redirect](open-redirect.md) [Allow Url Fopen Or Include](allow-url-fopen-or-include.md) [Insecure cryptography](insecure-cryptography.md) [Object Input Stream Insecure Deserialization](object-input-stream-insecure-deserialization.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Code Injection](code-injection.md) [Zip bomb attack](zip-bomb-attack.md) [Unsafe Reflection](unsafe-reflection.md) [Secure Signal Handling](secure-signal-handling.md) [Deserialization of untrusted data](untrusted-deserialization.md) [Static Initialization Vector (IV)](static-initialization-vector.md) [Coral Csrf Rule](coral-csrf-rule.md) [Insecure cookie](insecure-cookie.md) [Improper access control](improper-access-control.md) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Insecure Object Attribute Modification [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

By default, Laravel Eloquent models protect attributes from mass assignment for security. Setting the `$guarded` property to an empty array in a model overrides this behavior, allowing all attributes to be mass assigned without restrictions.

**Detector ID**

php/insecure-object-attribute-modification@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-915](https://cwe.mitre.org/data/definitions/915.html)

**Tags**

-

* * *
