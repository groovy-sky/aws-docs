![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Path Traversal [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Input must be sanitized before use in path traversal. Unsanitized input enables unauthorized access to files or directories beyond the intended scope, potentially resulting in disclosure of sensitive information, unauthorized modification of data, or execution of arbitrary code.

**Detector ID**

scala/path-traversal@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/top25-cwes)

* * *

#### Noncompliant example

```scala
1class PathTraversalNoncompliant {
2
3    @throws[FileUploadException]
4    override protected def doGet_compliant(req: HttpServletRequest, resp: HttpServletResponse): Unit = {
5        val input = req.getParameter("input")
6
7        // Noncompliant: Utilizes an unsanitized HTTP request parameter to form a file path.
8        val file = new File(input, "abs/path")
9    }
10}

```

#### Compliant example

```scala
1import javax.servlet.http.HttpServletRequest
2import javax.servlet.http.HttpServletResponse
3
4
5class PathTraversalCompliant {
6
7  @throws[FileUploadException]
8  override protected def doGet_compliant(req: HttpServletRequest, resp: HttpServletResponse): Unit = {
9    val input = req.getParameter("input")
10    val baseDir = "/some/fixed/base/directory"
11    // Compliant: No HTTP request parameters are used to construct a file path.
12    val file = new File(baseDir, "abs/path")
13  }
14}

```
