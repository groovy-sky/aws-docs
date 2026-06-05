---
title: "Untrusted data in http session High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Untrusted data in http session [High](severity/high.md)

User input is going into a session command, `setAttribute`. User input into such a command could lead to an attacker inputting malicious code into your session parameters, blurring the line between what's trusted and untrusted, and therefore leading to a trust boundary violation.

**Detector ID**

scala/untrusted-data-in-http-session@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-501](https://cwe.mitre.org/data/definitions/501.html)

**Tags**

[\# configuration](tags/configuration.md) [\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```scala
1class UntrustedDataInHttpSessionNoncompliant {
2
3  def nonCompliant(req: HttpServletRequest): Unit = {
4    val input = req.getParameter("input")
5    // Noncompliant: Unsanitized user input is used inside `setAttribute` method.
6    req.getSession.setAttribute(input, "true")
7  }
8}

```

#### Compliant example

```scala
1class UntrustedDataInHttpSessionCompliant {
2
3    def compliant(req: HttpServletRequest, input: String): Unit = {
4        if ("enable".equals(input)) req.getSession.setAttribute("user", "true")
5        // Compliant: Unsanitized user input is not used inside `setAttribute` method.
6        else req.getSession.setAttribute("user", "false")
7  }
8}

```

All content copied from https://docs.aws.amazon.com/.
