![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/csharp/thread-safety-violation) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/os-command-injection) [Unvalidated Redirect](https://docs.aws.amazon.com/amazonq/detector-library/csharp/unvalidated-redirect) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/csharp/integer-overflow) [Avoid Persistent Cookies](https://docs.aws.amazon.com/amazonq/detector-library/csharp/avoid-persistent-cookies) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/untrusted-deserialization) [LDAP Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/ldap-injection) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/csharp/weak-random-number-generation) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sql-injection) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/csharp/path-traversal) [Debug Binary](https://docs.aws.amazon.com/amazonq/detector-library/csharp/debug-binary) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sensitive-information-leak) [Webconfig Trace Enabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/webconfig-trace-enabled) [Inter Process Write of RegionInfo](https://docs.aws.amazon.com/amazonq/detector-library/csharp/region-info-inter-process-write) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/code-injection) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/missing-authorization) [JWT TokenValidationParameters No Expiry](https://docs.aws.amazon.com/amazonq/detector-library/csharp/jwt-no-expiry) [Razor Use of html string](https://docs.aws.amazon.com/amazonq/detector-library/csharp/razor-use-of-html-string) [Server-Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/server-side-request-forgery) [Origins Verified Cross Origin Communications](https://docs.aws.amazon.com/amazonq/detector-library/csharp/origins-verified-cross-origin-communications) [Prevent Excessive Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/prevent-excessive-authentication) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# Thread Safety Violation [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

Accessing global variables without a lock can lead to race condition.

**Detector ID**

csharp/thread-safety-violation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-362](https://cwe.mitre.org/data/definitions/362.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1public void ThreadSafetyViolationNoncompliant()
2{
3    for (int i = 0; i < 100000; i++)
4    {
5        // Noncompliant: This global variable is accessed without a lock.
6        result = result + 1;
7    }
8}

```

#### Compliant example

```csharp
1public void ThreadSafetyViolationCompliant()
2{
3    for (int i = 0; i < 100000; i++)
4    {
5        // Compliant: Provide lock to protect the integrity of the global variable.
6        Monitor.Enter(mainLock);
7        result = result + 1;
8        Monitor.Exit(mainLock);
9    }
10}

```
