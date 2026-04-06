![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cross-site-request-forgery) [Module Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/module-injection) [Improper Cryptographic Signature Verification](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-cryptographic-signature-verification) [Obsolete Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/obsolete-cryptography) [Inefficient Regular Expression](https://docs.aws.amazon.com/amazonq/detector-library/csharp/inefficient-regular-expression) [Double Epsilon Equality](https://docs.aws.amazon.com/amazonq/detector-library/csharp/double-epsilon-equality) [Unrestricted File Upload](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unrestricted-file-upload) [Output Cache Conflicts](https://docs.aws.amazon.com/amazonq/detector-library/csharp/output-cache-conflicts) [Unsafe XSLT Setting Used](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unsafe-xslt-setting-used) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cross-site-scripting) [Weak Cipher Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/csharp/weak-cipher-algorithm) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/csharp/stack-trace-exposure) [XPath Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/xpath-injection) [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/csharp/thread-safety-violation) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/os-command-injection) [Unvalidated Redirect](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unvalidated-redirect) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/csharp/integer-overflow) [Avoid Persistent Cookies](https://docs.aws.amazon.com/amazonq/detector-library/csharp/avoid-persistent-cookies) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/untrusted-deserialization) [LDAP Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/ldap-injection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/csharp/weak-random-number-generation) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sql-injection) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/csharp/path-traversal) [Debug Binary](https://docs.aws.amazon.com/amazonq/detector-library/csharp/debug-binary) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sensitive-information-leak) [Webconfig Trace Enabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/webconfig-trace-enabled) [Inter Process Write of RegionInfo](https://docs.aws.amazon.com/amazonq/detector-library/csharp/region-info-inter-process-write) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/code-injection) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/missing-authorization) [JWT TokenValidationParameters No Expiry](https://docs.aws.amazon.com/amazonq/detector-library/csharp/jwt-no-expiry) [Razor Use of html string](https://docs.aws.amazon.com/amazonq/detector-library/csharp/razor-use-of-html-string) [Server-Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/server-side-request-forgery) [Origins Verified Cross Origin Communications](https://docs.aws.amazon.com/amazonq/detector-library/csharp/origins-verified-cross-origin-communications) [Prevent Excessive Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/prevent-excessive-authentication) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# Cross-Site Request Forgery (CSRF) [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

The application failed to protect against Cross-Site Request Forgery (CSRF) due to not including the 'ValidateAntiForgeryToken' attribute on an HTTP method handler that could change user state (usually in the form of POST or PUT methods).

**Detector ID**

csharp/cross-site-request-forgery@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-352](https://cwe.mitre.org/data/definitions/352.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1[HttpPost]
2// Noncompliant: Does not enforce anti-forgery token validation.
3public ActionResult CrossSiteRequestForgeryNoncompliant(User user) {
4  CreateUser(user);
5}

```

#### Compliant example

```csharp
1[HttpPost]
2// Compliant: Enforce anti-forgery token validation.
3[ValidateAntiForgeryToken]
4public IActionResult CrossSiteRequestForgeryNoncompliant(User user){
5  CreateUser(user);
6}

```
