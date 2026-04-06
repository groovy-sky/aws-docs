![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Cross-site scripting [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

User-controllable input must be sanitized before it's included in output used to dynamically generate a web page. Unsanitized user input can introduce cross-side scripting (XSS) vulnerabilities that can lead to inadvertedly running malicious code in a trusted context.

**Detector ID**

scala/cross-site-scripting@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-79](https://cwe.mitre.org/data/definitions/79.html) [CWE-80](https://cwe.mitre.org/data/definitions/80.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/top25-cwes)

* * *

#### Noncompliant example

```scala
1@throws[ServletException]
2@throws[IOException]
3protected def nonCompliant(req: HttpServletRequest, resp: HttpServletResponse): Unit = {
4    val input = req.getParameter("input")
5    val map = req.getParameterMap
6    val vals = req.getParameterValues("input2")
7    val names = req.getParameterNames
8    val contentType = req.getContentType
9    val serverName = req.getServerName
10    // Noncompliant: Server response uses potentially unsanitized data.
11    resp.getWriter.write(input)
12}

```

#### Compliant example

```scala
1@throws[ServletException]
2@throws[IOException]
3protected def compliant(req: HttpServletRequest, resp: HttpServletResponse): Unit = {
4    val input = req.getParameter("input")
5    val map = req.getParameterMap
6    val vals = req.getParameterValues("input2")
7    val names = req.getParameterNames
8    val contentType = req.getContentType
9    val serverName = req.getServerName
10    // Compliant: Unsanitized input is encoded.
11    resp.getWriter.write(Encode.forHtml(input))
12}

```
