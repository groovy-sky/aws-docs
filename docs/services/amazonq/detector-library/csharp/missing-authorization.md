![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](sql-injection.md) [Path Traversal](path-traversal.md) [Debug Binary](debug-binary.md) [Sensitive Information Leak](sensitive-information-leak.md) [Webconfig Trace Enabled](webconfig-trace-enabled.md) [Inter Process Write of RegionInfo](region-info-inter-process-write.md) [Code Injection](code-injection.md) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/missing-authorization) [JWT TokenValidationParameters No Expiry](https://docs.aws.amazon.com/amazonq/detector-library/csharp/jwt-no-expiry) [Razor Use of html string](https://docs.aws.amazon.com/amazonq/detector-library/csharp/razor-use-of-html-string) [Server-Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/server-side-request-forgery) [Origins Verified Cross Origin Communications](https://docs.aws.amazon.com/amazonq/detector-library/csharp/origins-verified-cross-origin-communications) [Prevent Excessive Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/prevent-excessive-authentication) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# Missing Authorization [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

The endpoint is potentially accessible to not authorized users. If it contains sensitive information, like log files for example, it may lead to privilege escalation.

**Detector ID**

csharp/missing-authorization@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-862](https://cwe.mitre.org/data/definitions/862.html) [CWE-285](https://cwe.mitre.org/data/definitions/285.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1
2// Noncompliant: Access control checks are missing.
3[AllowAnonymous]
4public class AtLeast21Controller : Controller
5{
6    public IActionResult Index() => View();
7    private IActionResult View()
8    {
9        throw new NotImplementedException();
10    }
11    public interface IActionResult
12    {
13    }
14}

```

#### Compliant example

```csharp
1
2// Compliant: Access only permitted for specific role.
3[Authorize(Roles = "LegalAdultGroup")]
4public class AtLeast21Controller2 : Controller
5{
6    public IActionResult Index() => View();
7    private IActionResult View()
8    {
9        throw new NotImplementedException();
10    }
11}

```
