![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](sql-injection.md) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/csharp/path-traversal) [Debug Binary](https://docs.aws.amazon.com/amazonq/detector-library/csharp/debug-binary) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sensitive-information-leak) [Webconfig Trace Enabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/webconfig-trace-enabled) [Inter Process Write of RegionInfo](https://docs.aws.amazon.com/amazonq/detector-library/csharp/region-info-inter-process-write) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/code-injection) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/missing-authorization) [JWT TokenValidationParameters No Expiry](https://docs.aws.amazon.com/amazonq/detector-library/csharp/jwt-no-expiry) [Razor Use of html string](https://docs.aws.amazon.com/amazonq/detector-library/csharp/razor-use-of-html-string) [Server-Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/server-side-request-forgery) [Origins Verified Cross Origin Communications](https://docs.aws.amazon.com/amazonq/detector-library/csharp/origins-verified-cross-origin-communications) [Prevent Excessive Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/prevent-excessive-authentication) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# Path Traversal [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

The application dynamically constructs file or path information. If the path information comes from user input, it could be abused to read sensitive files, access other users data, or aid in exploitation to gain further system access.

**Detector ID**

csharp/path-traversal@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1public static byte[] PathTraversalNoncompliant(string filename)
2{
3    if (string.IsNullOrEmpty(filename))
4    {
5        throw new ArgumentNullException("error");
6    }
7    string filepath = Path.Combine("/pathToMyImage/images/", filename);
8    // Noncompliant: Reading data from a file without sanitization.
9    return File.ReadAllBytes(filepath);
10}

```

#### Compliant example

```csharp
1public static byte[] PathTraversalCompliant(string filename)
2{
3    if (string.IsNullOrEmpty(filename))
4    {
5        throw new ArgumentNullException("error");
6    }
7    filename = Path.GetFileName(filename);
8    // Compliant: `Path.GetFileName` used for sanitization.
9    string filepath = Path.Combine("/pathToMyImage/images/", filename);
10    return File.ReadAllBytes(filepath);
11}

```
