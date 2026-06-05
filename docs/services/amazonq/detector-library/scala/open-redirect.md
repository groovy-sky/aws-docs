---
title: "URL redirection to untrusted site High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# URL redirection to untrusted site [High](severity/high.md)

An HTTP parameter could contain a URL value and cause the web application to redirect the request to the specified URL. By modifying the URL value to a malicious site, an attacker could successfully launch a phishing attack and steal user credentials.

**Detector ID**

scala/open-redirect@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-601](https://cwe.mitre.org/data/definitions/601.html)

**Tags**

[\# top25-cwes](tags/top25-cwes.md) [\# owasp-top10](tags/owasp-top10.md)

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

All content copied from https://docs.aws.amazon.com/.
