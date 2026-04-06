![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sql-injection) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/csharp/path-traversal) [Debug Binary](https://docs.aws.amazon.com/amazonq/detector-library/csharp/debug-binary) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/csharp/sensitive-information-leak) [Webconfig Trace Enabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/webconfig-trace-enabled) [Inter Process Write of RegionInfo](https://docs.aws.amazon.com/amazonq/detector-library/csharp/region-info-inter-process-write) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/code-injection) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/csharp/missing-authorization) [JWT TokenValidationParameters No Expiry](https://docs.aws.amazon.com/amazonq/detector-library/csharp/jwt-no-expiry) [Razor Use of html string](https://docs.aws.amazon.com/amazonq/detector-library/csharp/razor-use-of-html-string) [Server-Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/csharp/server-side-request-forgery) [Origins Verified Cross Origin Communications](https://docs.aws.amazon.com/amazonq/detector-library/csharp/origins-verified-cross-origin-communications) [Prevent Excessive Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/prevent-excessive-authentication) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# SQL Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

SQL Injection is a critical vulnerability that can lead to data or system compromise. By dynamically generating SQL query strings, user input may be able to influence the logic of the SQL statement. This could lead to an adversary accessing information they should not have access to, or in some circumstances, being able to execute OS functionality or code.

**Detector ID**

csharp/sql-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1public void SqlInjectionNoncompliant(string sql)
2{
3    using (SqlConnection connection =
4           new SqlConnection("Data Source=(local);Initial Catalog=Northwind;Integrated Security=SSPI;")) {
5        connection.Open();
6        // Noncompliant: Detected a formatted string in a SQL statement.
7        SqlCommand command= new SqlCommand(sql);
8    }
9}

```

#### Compliant example

```csharp
1public void SqlInjectionCompliant()
2{
3    using (SqlConnection connection =
4           new SqlConnection("Data Source=(local);Initial Catalog=Northwind;Integrated Security=SSPI;"))
5    {
6        string sql = "SELECT * FROM Customers WHERE EmployeeId = @EmployeeId";
7        // Compliant: Obtain a PreparedStatement 'SqlParameter'.
8        SqlCommand command = new SqlCommand(sql);
9        command.Parameters.Add(new SqlParameter("@EmployeeId", System.Data.SqlDbType.Int));
10    }
11}

```
