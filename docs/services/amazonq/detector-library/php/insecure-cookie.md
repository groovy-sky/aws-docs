![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](server-side-request-forgery.md) [SQL Injection](sql-injection.md) [Activated Debug Feature](activated-debug-feature.md) [Sensitive information leak](sensitive-information-leak.md) [Log Injection](log-injection.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Cross-site scripting](cross-site-scripting.md) [Dangerous Function Usage](dangerous-function-usage.md) [Path Traversal](path-traversal.md) [Avoiding Exceptions in PHP](avoid-exit-die.md) [OS command injection](os-command-injection.md) [Incorrect Comparison](incorrect-comparison.md) [Ldap Bind Without Password](ldap-bind-without-password.md) [Sendfile Injection](sendfile-injection.md) [Assert Use](assert-use.md) [Loose file permissions](loose-file-permissions.md) [Improper Authentication](improper-authentication.md) [Insecure connection](insecure-connection.md) [Weak Random Number Generation](weak-random-number-generation.md) [Open Redirect](open-redirect.md) [Allow Url Fopen Or Include](allow-url-fopen-or-include.md) [Insecure cryptography](insecure-cryptography.md) [Object Input Stream Insecure Deserialization](object-input-stream-insecure-deserialization.md) [Cookie Without Http Only Flag](sensitive-cookie-without-http-only-flag.md) [Code Injection](code-injection.md) [Zip bomb attack](zip-bomb-attack.md) [Unsafe Reflection](unsafe-reflection.md) [Secure Signal Handling](secure-signal-handling.md) [Deserialization of untrusted data](untrusted-deserialization.md) [Static Initialization Vector (IV)](static-initialization-vector.md) [Coral Csrf Rule](coral-csrf-rule.md) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# Insecure cookie [High](https://docs.aws.amazon.com/amazonq/detector-library/php/severity/high)

Insecure cookie settings can lead to unencrypted cookie transmission. Even if a cookie doesn't contain sensitive data now, sensitive data could be added later. It's good practice to transmit all cookies only through secure channels.

**Detector ID**

php/insecure-cookie@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/php/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-614](https://cwe.mitre.org/data/definitions/614.html)

**Tags**

[\# cookies](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/cookies) [\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/php/tags/owasp-top10)

* * *

#### Noncompliant example

```php
1// Noncompliant: Used insecure FTP functions that transmit credentials in plain text, such as ftp_login.
2$login_result = ftp_login($conn_id, $ftp_user_name, $ftp_user_pass);
```

#### Compliant example

```php
1// Compliant: Used secure file transfer functions like ssh2_scp_send
2ssh2_scp_send($connection, '/local/filename', '/remote/filename', 0644);
```
