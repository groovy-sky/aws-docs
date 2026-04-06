![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-servlet-handling) [Use of Insufficiently Random Values](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-insufficiently-random-values) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cookie) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Insecure servlet handling [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Usage of servlet methods that may expose the application to XSS and injection attacks by concatenating or using user input without proper validation or sanitization.

**Detector ID**

scala/insecure-servlet-handling@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-20](https://cwe.mitre.org/data/definitions/20.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10)

* * *

#### Noncompliant example

```scala
1override protected def nonCompliant(req: HttpServletRequest, resp: HttpServletResponse): Unit = {
2  useParameters(req)
3  // Noncompliant: Unsanitized user input is used
4  resp.getWriter.print("<!--" + req.getContentType + "-->")
5  resp.getWriter.print("<!--" + req.getQueryString + "-->")
6  val referrer = req.getHeader("Referer") //Should have a higher priority
7  if (referrer != null && referrer.startsWith("http://company.ca")) {
8    req.getHeader("Host")
9    req.getHeader("User-Agent")
10    req.getHeader("X-Requested-With")
11  }
12}

```

#### Compliant example

```scala
1override def compliant(request: HttpServletRequest, response: HttpServletResponse): Unit = {
2    val inputParam = request.getParameter("param")
3    if (inputParam != null && !inputParam.isEmpty && inputParam.matches("[a-zA-Z0-9]+")) {
4        // Sanitize the input using Encode.forHtml
5        val sanitizedParam = Encode.forHtml(inputParam)
6        // Use the sanitizedParam safely
7        // Compliant: User input is sanitized
8        response.getWriter.println(s"Sanitized input: $sanitizedParam")
9    } else {
10        response.getWriter.println("Invalid input")
11    }
12}

```
