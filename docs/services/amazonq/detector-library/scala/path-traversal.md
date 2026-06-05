---
title: "Path Traversal High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Path Traversal [High](severity/high.md)

Input must be sanitized before use in path traversal. Unsanitized input enables unauthorized access to files or directories beyond the intended scope, potentially resulting in disclosure of sensitive information, unauthorized modification of data, or execution of arbitrary code.

**Detector ID**

scala/path-traversal@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html)

**Tags**

[\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

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

All content copied from https://docs.aws.amazon.com/.
