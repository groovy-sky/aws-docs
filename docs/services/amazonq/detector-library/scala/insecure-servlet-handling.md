---
title: "Insecure servlet handling High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Insecure servlet handling [High](severity/high.md)

Usage of servlet methods that may expose the application to XSS and injection attacks by concatenating or using user input without proper validation or sanitization.

**Detector ID**

scala/insecure-servlet-handling@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-20](https://cwe.mitre.org/data/definitions/20.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md)

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

All content copied from https://docs.aws.amazon.com/.
