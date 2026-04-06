![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](../server-side-request-forgery.md) [SQL Injection](../sql-injection.md) [Activated Debug Feature](../activated-debug-feature.md) [Sensitive information leak](../sensitive-information-leak.md) [Log Injection](../log-injection.md) [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md) [Cross-site scripting](../cross-site-scripting.md) [Dangerous Function Usage](../dangerous-function-usage.md) [Path Traversal](../path-traversal.md) [Avoiding Exceptions in PHP](../avoid-exit-die.md) [OS command injection](../os-command-injection.md) [Incorrect Comparison](../incorrect-comparison.md) [Ldap Bind Without Password](../ldap-bind-without-password.md) [Sendfile Injection](../sendfile-injection.md) [Assert Use](../assert-use.md) [Loose file permissions](../loose-file-permissions.md) [Improper Authentication](../improper-authentication.md) [Insecure connection](../insecure-connection.md) [Weak Random Number Generation](../weak-random-number-generation.md) [Open Redirect](../open-redirect.md) [Allow Url Fopen Or Include](../allow-url-fopen-or-include.md) [Insecure cryptography](../insecure-cryptography.md) [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md) [Cookie Without Http Only Flag](../sensitive-cookie-without-http-only-flag.md) [Code Injection](../code-injection.md) [Zip bomb attack](../zip-bomb-attack.md) [Unsafe Reflection](../unsafe-reflection.md) [Secure Signal Handling](../secure-signal-handling.md) [Deserialization of untrusted data](../untrusted-deserialization.md) [Static Initialization Vector (IV)](../static-initialization-vector.md) [Coral Csrf Rule](../coral-csrf-rule.md) [Insecure cookie](../insecure-cookie.md) [Improper access control](../improper-access-control.md) [Insecure Object Attribute Modification](../insecure-object-attribute-modification.md)

# High

Showing all detectors for the PHP language withhigh severity.

### [Server Side Request Forgery](../server-side-request-forgery.md)

The web server lacks proper validation when processing a URL or comparable request from an upstream component, creating a potential security risk associated with the function and its payload.

### [SQL Injection](../sql-injection.md)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Activated Debug Feature](../activated-debug-feature.md)

Set `APP_DEBUG = false` in production to avoid exposing debug settings.

### [Sensitive information leak](../sensitive-information-leak.md)

The `phpinfo` function may reveal sensitive information about your environment.

### [Log Injection](../log-injection.md)

Improper Output Neutralization for Logs.

### [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md)

Unverified origins of messages and identities in cross-origin communications can lead to security vulnerabilities.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Dangerous Function Usage](../dangerous-function-usage.md)

Identifies the use of deprecated `Mcrypt` functions in PHP, encouraging a switch to secure alternatives.

### [Path Traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Avoiding Exceptions in PHP](../avoid-exit-die.md)

Maintaining application stability with graceful PHP exits and exceptional handling.

### [OS command injection](../os-command-injection.md)

OS command injection from untrusted input.

### [Incorrect Comparison](../incorrect-comparison.md)

Secure comparison type safety recommendation.

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

### [Weak Random Number Generation](../weak-random-number-generation.md)

Use secure random functions like `random_bytes()` instead of non-cryptographic PRNGs in security code.

### [Open Redirect](../open-redirect.md)

Redirecting to the current request URL might direct to another domain if the current path begins with two slashes.

### [Allow Url Fopen Or Include](../allow-url-fopen-or-include.md)

Using unvalidated URLs with allow\_url\_fopen enables server-side request forgery attacks.

### [Insecure cryptography](../insecure-cryptography.md)

Use of a broken or risky cryptographic algorithm identified as weak.

### [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md)

Sanitize user input to prevent PHP object injection vulnerabilities from direct injection without sanitization.

### [Cookie Without Http Only Flag](../sensitive-cookie-without-http-only-flag.md)

Sensitive cookie without 'HttpOnly' flag

### [Code Injection](../code-injection.md)

Avoid running dynamic commands to prevent command injection vulnerabilities.

### [Zip bomb attack](../zip-bomb-attack.md)

Expanding unverified archive files without controlling the size of the expanded data can lead to zip bomb attacks.

### [Unsafe Reflection](../unsafe-reflection.md)

Use of externally-controlled input in reflection.

### [Secure Signal Handling](../secure-signal-handling.md)

Ensure secure coding by validating process signaling parameters to prevent application instability and potential denial of services.

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

### [Insecure Object Attribute Modification](../insecure-object-attribute-modification.md)

Overrides Eloquent's mass assignment protections by setting $guarded to empty array.
