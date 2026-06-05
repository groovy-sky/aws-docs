---
title: "Unrestricted File Upload High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](sql-injection.md) [Path Traversal](path-traversal.md) [Debug Binary](debug-binary.md) [Sensitive Information Leak](sensitive-information-leak.md) [Webconfig Trace Enabled](webconfig-trace-enabled.md) [Inter Process Write of RegionInfo](region-info-inter-process-write.md) [Code Injection](code-injection.md) [Missing Authorization](missing-authorization.md) [JWT TokenValidationParameters No Expiry](jwt-no-expiry.md) [Razor Use of html string](razor-use-of-html-string.md) [Server-Side Request Forgery (SSRF)](server-side-request-forgery.md) [Origins Verified Cross Origin Communications](origins-verified-cross-origin-communications.md) [Prevent Excessive Authentication](prevent-excessive-authentication.md) [Improper Authentication](improper-authentication.md) [Certificate Validation Disabled](certificate-validation-disabled.md) [Insecure Cryptography](insecure-cryptography.md) [Log Injection](log-injection.md) [Mass Assignment](mass-assignment.md) [Cookie Without SSL Flag](cookie-without-ssl-flag.md)

# Unrestricted File Upload [High](severity/high.md)

The product allows attacker to upload or transfer files within product environment, even though file type is dangerous.

**Detector ID**

csharp/unrestricted-file-upload@v1.0

**Category**

[Security](categories/security.md)

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

All content copied from https://docs.aws.amazon.com/.
