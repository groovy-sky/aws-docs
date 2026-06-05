---
title: "Sensitive Information Leak High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](sql-injection.md) [Path Traversal](path-traversal.md) [Debug Binary](debug-binary.md) [Sensitive Information Leak](sensitive-information-leak.md) [Webconfig Trace Enabled](webconfig-trace-enabled.md) [Inter Process Write of RegionInfo](region-info-inter-process-write.md) [Code Injection](code-injection.md) [Missing Authorization](missing-authorization.md) [JWT TokenValidationParameters No Expiry](jwt-no-expiry.md) [Razor Use of html string](razor-use-of-html-string.md) [Server-Side Request Forgery (SSRF)](server-side-request-forgery.md) [Origins Verified Cross Origin Communications](origins-verified-cross-origin-communications.md) [Prevent Excessive Authentication](prevent-excessive-authentication.md) [Improper Authentication](improper-authentication.md) [Certificate Validation Disabled](certificate-validation-disabled.md) [Insecure Cryptography](insecure-cryptography.md) [Log Injection](log-injection.md) [Mass Assignment](mass-assignment.md) [Cookie Without SSL Flag](cookie-without-ssl-flag.md)

# Sensitive Information Leak [High](severity/high.md)

Sensitive information should not be exposed through log files or stack traces. Ensure that sensitive information is redacted and that logging is used only in debug mode with test data.

**Detector ID**

csharp/sensitive-information-leak@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-200](https://cwe.mitre.org/data/definitions/200.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1public void SensitiveInformationLeakNoncompliant()
2{
3    string url = "http://api.example.com";
4    string apiKey = "YOUR_API_KEY";
5    using (HttpClient client = new HttpClient())
6    {
7        // Noncompliant: Send sensitive data over HTTP connection
8        HttpResponseMessage response = client.GetAsync(url + "?api_key=" + apiKey).Result;
9    }
10}

```

#### Compliant example

```csharp
1public void SensitiveInformationLeakCompliant()
2{
3    string url = "http://api.example.com";
4    string apiKey = "YOUR_API_KEY";
5    using (HttpClient client = new HttpClient())
6    {
7        // Noncompliant: Do not send sensitive data over HTTP connection
8        HttpResponseMessage response = client.GetAsync(url).Result;
9    }
10}

```

All content copied from https://docs.aws.amazon.com/.
