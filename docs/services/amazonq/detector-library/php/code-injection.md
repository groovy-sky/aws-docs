![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](path-traversal.md) [Avoiding Exceptions in PHP](avoid-exit-die.md) [OS command injection](os-command-injection.md) [Incorrect Comparison](incorrect-comparison.md) [Ldap Bind Without Password](ldap-bind-without-password.md) [Sendfile Injection](sendfile-injection.md) [Assert Use](assert-use.md) [Loose file permissions](loose-file-permissions.md) [Improper Authentication](improper-authentication.md) [Insecure connection](insecure-connection.md) [Weak Random Number Generation](weak-random-number-generation.md) [Open Redirect](open-redirect.md) [Allow Url Fopen Or Include](allow-url-fopen-or-include.md) [Insecure cryptography](insecure-cryptography.md) [Object Input Stream Insecure Deserialization](object-input-stream-insecure-deserialization.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection) [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Code Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

It is crucial to refrain from executing non-constant commands within a system, as doing so may expose it to the risk of command injection vulnerabilities. Command injection occurs when an application processes user input as part of a command, allowing an attacker to manipulate the input and inject malicious commands. By ensuring that only constant commands are executed, the system can mitigate the potential for unauthorized command manipulation and enhance overall security.

**Detector ID**

php/code-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-94](https://cwe.mitre.org/data/definitions/94.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/top25-cwes)

* * *
