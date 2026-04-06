![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C\# detectors(44/44)

[Method Input Validation](method-input-validation.md) [Password Complexity](password-complexity.md) [Xml External Entity](xml-external-entity.md) [Memory Marshal CreateSpan](memory-marshal-create-span.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Module Injection](module-injection.md) [Improper Cryptographic Signature Verification](improper-cryptographic-signature-verification.md) [Obsolete Cryptography](obsolete-cryptography.md) [Inefficient Regular Expression](inefficient-regular-expression.md) [Double Epsilon Equality](double-epsilon-equality.md) [Unrestricted File Upload](unrestricted-file-upload.md) [Output Cache Conflicts](output-cache-conflicts.md) [Unsafe XSLT Setting Used](unsafe-xslt-setting-used.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Weak Cipher Algorithm](weak-cipher-algorithm.md) [Stack Trace Exposure](stack-trace-exposure.md) [XPath Injection](xpath-injection.md) [Thread Safety Violation](thread-safety-violation.md) [OS Command Injection](os-command-injection.md) [Unvalidated Redirect](unvalidated-redirect.md) [Integer Overflow](integer-overflow.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Untrusted Deserialization](untrusted-deserialization.md) [LDAP Injection](ldap-injection.md) [Weak Random Number Generation](weak-random-number-generation.md) [SQL Injection](sql-injection.md) [Path Traversal](path-traversal.md) [Debug Binary](debug-binary.md) [Sensitive Information Leak](sensitive-information-leak.md) [Webconfig Trace Enabled](webconfig-trace-enabled.md) [Inter Process Write of RegionInfo](region-info-inter-process-write.md) [Code Injection](code-injection.md) [Missing Authorization](missing-authorization.md) [JWT TokenValidationParameters No Expiry](jwt-no-expiry.md) [Razor Use of html string](razor-use-of-html-string.md) [Server-Side Request Forgery (SSRF)](server-side-request-forgery.md) [Origins Verified Cross Origin Communications](origins-verified-cross-origin-communications.md) [Prevent Excessive Authentication](prevent-excessive-authentication.md) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/csharp/improper-authentication) [Certificate Validation Disabled](https://docs.aws.amazon.com/amazonq/detector-library/csharp/certificate-validation-disabled) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/csharp/insecure-cryptography) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/csharp/log-injection) [Mass Assignment](https://docs.aws.amazon.com/amazonq/detector-library/csharp/mass-assignment) [Cookie Without SSL Flag](https://docs.aws.amazon.com/amazonq/detector-library/csharp/cookie-without-ssl-flag)

# Improper Authentication [High](https://docs.aws.amazon.com/amazonq/detector-library/csharp/severity/high)

Failure to verify a user's identity results in improper authentication. This can allow an attacker to acquire privileges to access sensitive data in your application.

**Detector ID**

csharp/improper-authentication@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/csharp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-287](https://cwe.mitre.org/data/definitions/287.html)

**Tags**

-

* * *

#### Noncompliant example

```csharp
1public void ImproperAuthenticationNoncompliant(int userId)
2{
3    var mySecret = "asdv234234^&%&^%&^hjsdfb2%%%";
4    var mySecurityKey = new SymmetricSecurityKey(Encoding.ASCII.GetBytes(mySecret));
5    var myIssuer = "http://mysite.com";
6    var myAudience = "http://myaudience.com";
7    var tokenHandler = new JwtSecurityTokenHandler();
8    var tokenDescriptor = new SecurityTokenDescriptor
9    {
10        Subject = new ClaimsIdentity(new Claim[]
11        {
12            new Claim(ClaimTypes.NameIdentifier, userId.ToString()),
13        }),
14        Expires = DateTime.UtcNow.AddDays(7),
15        Issuer = myIssuer,
16        Audience = myAudience,
17        SigningCredentials = new SigningCredentials(mySecurityKey, SecurityAlgorithms.HmacSha256Signature)
18    };
19    // Noncompliant: `JwtSecurityTokenHandler` is not validating before using it.
20    var token = tokenHandler.CreateToken(tokenDescriptor);
21}

```

#### Compliant example

```csharp
1public void ImproperAuthenticationCompliant(string token)
2{
3    var mySecret = "asdv234234^&%&^%&^hjsdfb2%%%";
4    var mySecurityKey = new SymmetricSecurityKey(Encoding.ASCII.GetBytes(mySecret));
5    var myIssuer = "http://mysite.com";
6    var myAudience = "http://myaudience.com";
7    var tokenHandler = new JwtSecurityTokenHandler();
8    // Compliant: `JwtSecurityTokenHandler` is validating before using it.
9    tokenHandler.ValidateToken(token, new TokenValidationParameters
10        {
11            ValidateIssuerSigningKey = true,
12            ValidateIssuer = true,
13            ValidateAudience = true,
14            ValidIssuer = myIssuer,
15            ValidAudience = myAudience,
16            IssuerSigningKey = mySecurityKey
17        }, out SecurityToken validatedToken);
18}

```
