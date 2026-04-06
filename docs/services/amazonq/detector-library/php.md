![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### PHP detectors(34/34)

[Server Side Request Forgery](https://docs.aws.amazon.com/amazonq/detector-library/php/server-side-request-forgery) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sql-injection) [Activated Debug Feature](https://docs.aws.amazon.com/amazonq/detector-library/php/activated-debug-feature) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-information-leak) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/log-injection) [Origins-verified cross-origin communications](https://docs.aws.amazon.com/amazonq/detector-library/php/origins-verified-cross-origin-communications) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/php/cross-site-scripting) [Dangerous Function Usage](https://docs.aws.amazon.com/amazonq/detector-library/php/dangerous-function-usage) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/php/path-traversal) [Avoiding Exceptions in PHP](https://docs.aws.amazon.com/amazonq/detector-library/php/avoid-exit-die) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/php/os-command-injection) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/php/incorrect-comparison) [Ldap Bind Without Password](https://docs.aws.amazon.com/amazonq/detector-library/php/ldap-bind-without-password) [Sendfile Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sendfile-injection) [Assert Use](https://docs.aws.amazon.com/amazonq/detector-library/php/assert-use) [Loose file permissions](https://docs.aws.amazon.com/amazonq/detector-library/php/loose-file-permissions) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-authentication) [Insecure connection](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-connection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/php/weak-random-number-generation) [Open Redirect](https://docs.aws.amazon.com/amazonq/detector-library/php/open-redirect) [Allow Url Fopen Or Include](https://docs.aws.amazon.com/amazonq/detector-library/php/allow-url-fopen-or-include) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cryptography) [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/php/object-input-stream-insecure-deserialization) [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-cookie-without-http-only-flag) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection) [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection) [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling) [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization) [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector) [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie) [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control) [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

# PHP detectors

Showing all detectors for the PHP language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/php/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/php/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/php/categories)

* * *

### Browse all detectors

### [Server Side Request Forgery](https://docs.aws.amazon.com/amazonq/detector-library/php/server-side-request-forgery)

The web server lacks proper validation when processing a URL or comparable request from an upstream component, creating a potential security risk associated with the function and its payload.

### [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sql-injection)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Activated Debug Feature](https://docs.aws.amazon.com/amazonq/detector-library/php/activated-debug-feature)

Set `APP_DEBUG = false` in production to avoid exposing debug settings.

### [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-information-leak)

The `phpinfo` function may reveal sensitive information about your environment.

### [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/log-injection)

Improper Output Neutralization for Logs.

### [Origins-verified cross-origin communications](https://docs.aws.amazon.com/amazonq/detector-library/php/origins-verified-cross-origin-communications)

Unverified origins of messages and identities in cross-origin communications can lead to security vulnerabilities.

### [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/php/cross-site-scripting)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Dangerous Function Usage](https://docs.aws.amazon.com/amazonq/detector-library/php/dangerous-function-usage)

Identifies the use of deprecated `Mcrypt` functions in PHP, encouraging a switch to secure alternatives.

### [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/php/path-traversal)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Avoiding Exceptions in PHP](https://docs.aws.amazon.com/amazonq/detector-library/php/avoid-exit-die)

Maintaining application stability with graceful PHP exits and exceptional handling.

### [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/php/os-command-injection)

OS command injection from untrusted input.

### [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/php/incorrect-comparison)

Secure comparison type safety recommendation.

### [Ldap Bind Without Password](https://docs.aws.amazon.com/amazonq/detector-library/php/ldap-bind-without-password)

Uncovered an anonymous LDAP bind,posing a risk of unauthorized execution of LDAP statements. Strengthen LDAP security with authentication enforcement.

### [Sendfile Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/sendfile-injection)

Functions require input validation and sanitization to prevent security risks from untrusted user data.

### [Assert Use](https://docs.aws.amazon.com/amazonq/detector-library/php/assert-use)

Executing assert with user-provided input is comparable to invoking dynamic code evaluations.

### [Loose file permissions](https://docs.aws.amazon.com/amazonq/detector-library/php/loose-file-permissions)

Weak file permissions can lead to privilege escalation.

### [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-authentication)

The `wp_ajax_priv_upload` and `wp_ajax_nopriv_upload` hooks allow developers to define callbacks for authenticated and anonymous AJAX requests.

### [Insecure connection](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-connection)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/php/weak-random-number-generation)

Use secure random functions like `random_bytes()` instead of non-cryptographic PRNGs in security code.

### [Open Redirect](https://docs.aws.amazon.com/amazonq/detector-library/php/open-redirect)

Redirecting to the current request URL might direct to another domain if the current path begins with two slashes.

### [Allow Url Fopen Or Include](https://docs.aws.amazon.com/amazonq/detector-library/php/allow-url-fopen-or-include)

Using unvalidated URLs with allow\_url\_fopen enables server-side request forgery attacks.

### [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cryptography)

Use of a broken or risky cryptographic algorithm identified as weak.

### [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/php/object-input-stream-insecure-deserialization)

Sanitize user input to prevent PHP object injection vulnerabilities from direct injection without sanitization.

### [Cookie Without Http Only Flag](https://docs.aws.amazon.com/amazonq/detector-library/php/sensitive-cookie-without-http-only-flag)

Sensitive cookie without 'HttpOnly' flag

### [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/php/code-injection)

Avoid running dynamic commands to prevent command injection vulnerabilities.

### [Zip bomb attack](https://docs.aws.amazon.com/amazonq/detector-library/php/zip-bomb-attack)

Expanding unverified archive files without controlling the size of the expanded data can lead to zip bomb attacks.

### [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/php/unsafe-reflection)

Use of externally-controlled input in reflection.

### [Secure Signal Handling](https://docs.aws.amazon.com/amazonq/detector-library/php/secure-signal-handling)

Ensure secure coding by validating process signaling parameters to prevent application instability and potential denial of services.

### [Deserialization of untrusted data](https://docs.aws.amazon.com/amazonq/detector-library/php/untrusted-deserialization)

Deserialization of untrusted data can lead to security vulnerabilities, such as inadvertently running remote code.

### [Static Initialization Vector (IV)](https://docs.aws.amazon.com/amazonq/detector-library/php/static-initialization-vector)

Using a static initialization vector (IV) for a cryptographic cipher is security sensitive.

### [Coral Csrf Rule](https://docs.aws.amazon.com/amazonq/detector-library/php/coral-csrf-rule)

Modify input validation check to guarantee expected behavior for all inputs, not just an invalid assumption that causes errors.

### [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-cookie)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Improper access control](https://docs.aws.amazon.com/amazonq/detector-library/php/improper-access-control)

The software does not restrict or incorrectly restrict access to a resource from an unauthorized actor.

### [Insecure Object Attribute Modification](https://docs.aws.amazon.com/amazonq/detector-library/php/insecure-object-attribute-modification)

Overrides Eloquent's mass assignment protections by setting $guarded to empty array.
