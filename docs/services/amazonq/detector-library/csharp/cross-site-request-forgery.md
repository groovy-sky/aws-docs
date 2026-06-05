---
title: "Cross-Site Request Forgery (CSRF) High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](sql-injection.md) [Path Traversal](path-traversal.md) [Debug Binary](debug-binary.md) [Sensitive Information Leak](sensitive-information-leak.md) [Webconfig Trace Enabled](webconfig-trace-enabled.md) [Inter Process Write of RegionInfo](region-info-inter-process-write.md) [Code Injection](code-injection.md) [Missing Authorization](missing-authorization.md) [JWT TokenValidationParameters No Expiry](jwt-no-expiry.md) [Razor Use of html string](razor-use-of-html-string.md) [Server-Side Request Forgery (SSRF)](server-side-request-forgery.md) [Origins Verified Cross Origin Communications](origins-verified-cross-origin-communications.md) [Prevent Excessive Authentication](prevent-excessive-authentication.md) [Improper Authentication](improper-authentication.md) [Certificate Validation Disabled](certificate-validation-disabled.md) [Insecure Cryptography](insecure-cryptography.md) [Log Injection](log-injection.md) [Mass Assignment](mass-assignment.md) [Cookie Without SSL Flag](cookie-without-ssl-flag.md)

# Cross-Site Request Forgery (CSRF) [High](severity/high.md)

The application failed to protect against Cross-Site Request Forgery (CSRF) due to not including the 'ValidateAntiForgeryToken' attribute on an HTTP method handler that could change user state (usually in the form of POST or PUT methods).

**Detector ID**

csharp/cross-site-request-forgery@v1.0

**Category**

[Security](categories/security.md)

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

All content copied from https://docs.aws.amazon.com/.
