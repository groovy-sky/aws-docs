![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unrestricted-file-upload) [Output Cache Conflicts](https://docs.aws.amazon.com/amazonq/detector-library/csharp/output-cache-conflicts) [Unsafe XSLT Setting Used](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unsafe-xslt-setting-used) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cross-site-scripting) [Weak Cipher Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/csharp/weak-cipher-algorithm) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/csharp/stack-trace-exposure) [XPath Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/xpath-injection) [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/csharp/thread-safety-violation) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/os-command-injection) [Unvalidated Redirect](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unvalidated-redirect) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/csharp/integer-overflow) [Avoid Persistent Cookies](https://docs.aws.amazon.com/amazonq/detector-library/csharp/avoid-persistent-cookies) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/untrusted-deserialization) [LDAP Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/ldap-injection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/csharp/weak-random-number-generation) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sql-injection) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/csharp/path-traversal) [Debug Binary](https://docs.aws.amazon.com/amazonq/detector-library/csharp/debug-binary) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sensitive-information-leak) [Webconfig Trace Enabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/webconfig-trace-enabled) [Inter Process Write of RegionInfo](https://docs.aws.amazon.com/amazonq/detector-library/csharp/region-info-inter-process-write) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/code-injection) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/missing-authorization) [JWT TokenValidationParameters No Expiry](https://docs.aws.amazon.com/amazonq/detector-library/csharp/jwt-no-expiry) [Razor Use of html string](https://docs.aws.amazon.com/amazonq/detector-library/csharp/razor-use-of-html-string) [Server-Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/server-side-request-forgery) [Origins Verified Cross Origin Communications](https://docs.aws.amazon.com/amazonq/detector-library/csharp/origins-verified-cross-origin-communications) [Prevent Excessive Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/prevent-excessive-authentication) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# Unrestricted File Upload [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

The product allows attacker to upload or transfer files within product environment, even though file type is dangerous.

**Detector ID**

csharp/unrestricted-file-upload@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-434](https://cwe.mitre.org/data/definitions/434.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1public void UnrestrictedFileUploadNoncompliant()
2{
3    try {
4        // Noncompliant: the filename is user controlled.
5        string completePathNewFile= safeUploadFolder + System.IO.Path.PathSeparator + uploadedFile;
6        if (!File.Exists(completePathNewFile))
7        {
8            using (StreamWriter sw = File.CreateText(completePathNewFile))
9            {
10                sw.WriteLine(fileContent);
11            }
12        }
13        Console.WriteLine("SUCCESS");
14    } catch (System.Exception e) {
15        Console.WriteLine("ERROR");
16    }
17}

```

#### Compliant example

```csharp
1public void UnrestrictedFileUploadCompliant()
2{
3    try {
4        // Compliant: Restrict the upload path, and ensure it is outside of the webroot.
5        string fileNameSanitized = System.IO.Path.GetFileName(uploadedFile);
6        string completePathNewFile= safeUploadFolder + System.IO.Path.PathSeparator + fileNameSanitized;
7        if (!File.Exists(completePathNewFile))
8        {
9            using (StreamWriter sw = File.CreateText(completePathNewFile))
10            {
11                sw.WriteLine(fileContent);
12            }
13        }
14        Console.WriteLine("SUCCESS");
15    } catch (System.Exception e) {
16        Console.WriteLine("ERROR");
17    }
18}

```
