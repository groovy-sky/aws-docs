![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](../method-input-validation.md) [Password Complexity](../password-complexity.md) [Xml External Entity](../xml-external-entity.md) [Memory Marshal CreateSpan](../memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](../cross-site-request-forgery.md) [Module Injection](../module-injection.md) [Improper Cryptographic Signature Verification](../improper-cryptographic-signature-verification.md) [Obsolete Cryptography](../obsolete-cryptography.md) [Inefficient Regular Expression](../inefficient-regular-expression.md) [Double Epsilon Equality](../double-epsilon-equality.md) [Unrestricted File Upload](../unrestricted-file-upload.md) [Output Cache Conflicts](../output-cache-conflicts.md) [Unsafe XSLT Setting Used](../unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](../cross-site-scripting.md) [Weak Cipher Algorithm](../weak-cipher-algorithm.md) [Stack Trace Exposure](../stack-trace-exposure.md) [XPath Injection](../xpath-injection.md) [Thread Safety Violation](../thread-safety-violation.md) [OS Command Injection](../os-command-injection.md) [Unvalidated Redirect](../unvalidated-redirect.md) [Integer Overflow](../integer-overflow.md) [Avoid Persistent Cookies](../avoid-persistent-cookies.md) [Untrusted Deserialization](../untrusted-deserialization.md) [LDAP Injection](../ldap-injection.md) [Weak Random Number Generation](../weak-random-number-generation.md) [SQL Injection](../sql-injection.md) [Path Traversal](../path-traversal.md) [Debug Binary](../debug-binary.md) [Sensitive Information Leak](../sensitive-information-leak.md) [Webconfig Trace Enabled](../webconfig-trace-enabled.md) [Inter Process Write of RegionInfo](../region-info-inter-process-write.md) [Code Injection](../code-injection.md) [Missing Authorization](../missing-authorization.md) [JWT TokenValidationParameters No Expiry](../jwt-no-expiry.md) [Razor Use of html string](../razor-use-of-html-string.md) [Server-Side Request Forgery (SSRF)](../server-side-request-forgery.md) [Origins Verified Cross Origin Communications](../origins-verified-cross-origin-communications.md) [Prevent Excessive Authentication](../prevent-excessive-authentication.md) [Improper Authentication](../improper-authentication.md) [Certificate Validation Disabled](../certificate-validation-disabled.md) [Insecure Cryptography](../insecure-cryptography.md) [Log Injection](../log-injection.md) [Mass Assignment](../mass-assignment.md) [Cookie Without SSL Flag](../cookie-without-ssl-flag.md)

# Security detectors

### [Method Input Validation](../method-input-validation.md)

ASP.NET input validation disabled

### [Password Complexity](../password-complexity.md)

Weak password requirements

### [Xml External Entity](../xml-external-entity.md)

Improper restriction of XML external entity reference ('XXE')

### [Memory Marshal CreateSpan](../memory-marshal-create-span.md)

Out-of-bounds read due to improper length check

### [Cross-Site Request Forgery (CSRF)](../cross-site-request-forgery.md)

Potential Cross-Site Request Forgery (CSRF)

### [Module Injection](../module-injection.md)

Potential use of top-level wildcard bindings

### [Improper Cryptographic Signature Verification](../improper-cryptographic-signature-verification.md)

Incorrect verification of signature for data.

### [Obsolete Cryptography](../obsolete-cryptography.md)

Use of obsolete cryptographic algorithm

### [Inefficient Regular Expression](../inefficient-regular-expression.md)

Regular expression Denial of Service attack.

### [Unrestricted File Upload](../unrestricted-file-upload.md)

Unrestricted upload of file whose type is dangerous.

### [Output Cache Conflicts](../output-cache-conflicts.md)

Use of cache containing sensitive information

### [Unsafe XSLT Setting Used](../unsafe-xslt-setting-used.md)

Improper restriction of XML external entity reference

### [Cross Site Scripting (XSS)](../cross-site-scripting.md)

Improper neutralization of input during web page generation ('Cross-site Scripting')

### [Weak Cipher Algorithm](../weak-cipher-algorithm.md)

Use of a broken or risky cryptographic algorithm.

### [Stack Trace Exposure](../stack-trace-exposure.md)

Expose sensitive information through stack trace.

### [XPath Injection](../xpath-injection.md)

Improper neutralization of data within XPath expressions ('XPathInjection').

### [Thread Safety Violation](../thread-safety-violation.md)

Thread safety violation can lead to race condition.

### [OS Command Injection](../os-command-injection.md)

Improper neutralization of special elements used in an OS command ('OS Command Injection')

### [Unvalidated Redirect](../unvalidated-redirect.md)

URL redirection to untrusted site 'open redirect'

### [Integer Overflow](../integer-overflow.md)

Integer Overflow or Wraparound.

### [Avoid Persistent Cookies](../avoid-persistent-cookies.md)

Persistent cookies are vulnerable to attacks.

### [Untrusted Deserialization](../untrusted-deserialization.md)

Deserialization of potentially untrusted data

### [LDAP Injection](../ldap-injection.md)

Improper neutralization of special elements used in an LDAP query ('LDAP Injection')

### [Weak Random Number Generation](../weak-random-number-generation.md)

Use of cryptographically weak Pseudo-Random Number Generator (PRNG)

### [SQL Injection](../sql-injection.md)

Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

### [Path Traversal](../path-traversal.md)

Improper limitation of a pathname to a restricted directory ('Path Traversal')

### [Debug Binary](../debug-binary.md)

Debugging messages can help attacker to form some sort of attack on system.

### [Sensitive Information Leak](../sensitive-information-leak.md)

Sensitive information should not be exposed through log files or stack traces.

### [Webconfig Trace Enabled](../webconfig-trace-enabled.md)

Net Webconfig Trace Enabled.

### [Code Injection](../code-injection.md)

Generation of code using external input without validation.

### [Missing Authorization](../missing-authorization.md)

Improper Access Control.

### [JWT TokenValidationParameters No Expiry](../jwt-no-expiry.md)

Insufficient Session Expiration.

### [Razor Use of html string](../razor-use-of-html-string.md)

Improper encoding or escaping can allow attackers to change the commands that are sent to another component.

### [Server-Side Request Forgery (SSRF)](../server-side-request-forgery.md)

Potential Server-Side Request Forgery.

### [Origins Verified Cross Origin Communications](../origins-verified-cross-origin-communications.md)

Origin Validation Error.

### [Prevent Excessive Authentication](../prevent-excessive-authentication.md)

Improper Restriction of Excessive Authentication Attempts.

### [Improper Authentication](../improper-authentication.md)

Your code doesn't sufficiently authenticate identities provided by its users.

### [Certificate Validation Disabled](../certificate-validation-disabled.md)

Certificate validation disabled.

### [Insecure Cryptography](../insecure-cryptography.md)

Use of risky or broken cryptographic algorithm

### [Mass Assignment](../mass-assignment.md)

Improperly Controlled Modification of Dynamically-Determined Object Attributes

### [Cookie Without SSL Flag](../cookie-without-ssl-flag.md)

Sensitive cookie in HTTPS session without 'Secure' attribute
