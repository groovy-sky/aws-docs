![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# URL redirection to untrusted site [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

An HTTP parameter could contain a URL value and cause the web application to redirect the request to the specified URL. By modifying the URL value to a malicious site, an attacker could successfully launch a phishing attack and steal user credentials.

**Detector ID**

scala/open-redirect@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-601](https://cwe.mitre.org/data/definitions/601.html)

**Tags**

[\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/top25-cwes) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10)

* * *

#### Noncompliant example

```scala
1import javax.servlet.http.HttpServletRequest
2import javax.servlet.http.HttpServletResponse
3
4class OpenRedirectNoncompliant extends HttpServlet {
5    def nonCompliant(req: HttpServletRequest, res: HttpServletResponse): Unit = {
6      val forwardedUrl = req.getHeader("Forwarded")
7      if (forwardedUrl != null && !forwardedUrl.isEmpty) {
8        // Noncompliant: Using user-controlled input in the Forwarded header for redirection
9        res.sendRedirect(forwardedUrl)
10      }
11    }
12}

```

#### Compliant example

```scala
1import javax.servlet.http.HttpServletRequest
2import javax.servlet.http.HttpServletResponse
3
4class OpenRedirectCompliant extends HttpServlet {
5    def compliant(req: HttpServletRequest, res: HttpServletResponse): Unit = {
6      val forwardedUrl = req.getHeader("Forwarded")
7      if (forwardedUrl.getHost.contains("trusteddomain.com")) {
8        // Compliant: The forwarded URL is validated before use.
9        Redirect(validatedUrl)
10      } else {
11        BadRequest("Invalid URL")
12      }
13  }
14}

```
