![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](../server-side-request-forgery.md) [SQL Injection](../sql-injection.md) [Activated Debug Feature](../activated-debug-feature.md) [Sensitive information leak](../sensitive-information-leak.md) [Log Injection](../log-injection.md) [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md) [Cross-site scripting](../cross-site-scripting.md) [Dangerous Function Usage](../dangerous-function-usage.md) [Path Traversal](../path-traversal.md) [Avoiding Exceptions in PHP](../avoid-exit-die.md) [OS command injection](../os-command-injection.md) [Incorrect Comparison](../incorrect-comparison.md) [Ldap Bind Without Password](../ldap-bind-without-password.md) [Sendfile Injection](../sendfile-injection.md) [Assert Use](../assert-use.md) [Loose file permissions](../loose-file-permissions.md) [Improper Authentication](../improper-authentication.md) [Insecure connection](../insecure-connection.md) [Weak Random Number Generation](../weak-random-number-generation.md) [Open Redirect](../open-redirect.md) [Allow Url Fopen Or Include](../allow-url-fopen-or-include.md) [Insecure cryptography](../insecure-cryptography.md) [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md) [Cookie Without Http Only Flag](../sensitive-cookie-without-http-only-flag.md) [Code Injection](../code-injection.md) [Zip bomb attack](../zip-bomb-attack.md) [Unsafe Reflection](../unsafe-reflection.md) [Secure Signal Handling](../secure-signal-handling.md) [Deserialization of untrusted data](../untrusted-deserialization.md) [Static Initialization Vector (IV)](../static-initialization-vector.md) [Coral Csrf Rule](../coral-csrf-rule.md) [Insecure cookie](../insecure-cookie.md) [Improper access control](../improper-access-control.md) [Insecure Object Attribute Modification](../insecure-object-attribute-modification.md)

# Tag: owasp-top10

### [SQL Injection](../sql-injection.md)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Sensitive information leak](../sensitive-information-leak.md)

The `phpinfo` function may reveal sensitive information about your environment.

### [Log Injection](../log-injection.md)

Improper Output Neutralization for Logs.

### [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md)

Unverified origins of messages and identities in cross-origin communications can lead to security vulnerabilities.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Path Traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [OS command injection](../os-command-injection.md)

OS command injection from untrusted input.

### [Ldap Bind Without Password](../ldap-bind-without-password.md)

Uncovered an anonymous LDAP bind,posing a risk of unauthorized execution of LDAP statements. Strengthen LDAP security with authentication enforcement.

### [Sendfile Injection](../sendfile-injection.md)

Functions require input validation and sanitization to prevent security risks from untrusted user data.

### [Assert Use](../assert-use.md)

Executing assert with user-provided input is comparable to invoking dynamic code evaluations.

### [Loose file permissions](../loose-file-permissions.md)

Weak file permissions can lead to privilege escalation.

### [Improper Authentication](../improper-authentication.md)

The `wp_ajax_priv_upload` and `wp_ajax_nopriv_upload` hooks allow developers to define callbacks for authenticated and anonymous AJAX requests.

### [Insecure connection](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Open Redirect](../open-redirect.md)

Redirecting to the current request URL might direct to another domain if the current path begins with two slashes.

### [Insecure cryptography](../insecure-cryptography.md)

Use of a broken or risky cryptographic algorithm identified as weak.

### [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md)

Sanitize user input to prevent PHP object injection vulnerabilities from direct injection without sanitization.

### [Code Injection](../code-injection.md)

Avoid running dynamic commands to prevent command injection vulnerabilities.

### [Zip bomb attack](../zip-bomb-attack.md)

Expanding unverified archive files without controlling the size of the expanded data can lead to zip bomb attacks.

### [Unsafe Reflection](../unsafe-reflection.md)

Use of externally-controlled input in reflection.

### [Deserialization of untrusted data](../untrusted-deserialization.md)

Deserialization of untrusted data can lead to security vulnerabilities, such as inadvertently running remote code.

### [Static Initialization Vector (IV)](../static-initialization-vector.md)

Using a static initialization vector (IV) for a cryptographic cipher is security sensitive.

### [Coral Csrf Rule](../coral-csrf-rule.md)

Modify input validation check to guarantee expected behavior for all inputs, not just an invalid assumption that causes errors.

### [Insecure cookie](../insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Improper access control](../improper-access-control.md)

The software does not restrict or incorrectly restrict access to a resource from an unauthorized actor.
