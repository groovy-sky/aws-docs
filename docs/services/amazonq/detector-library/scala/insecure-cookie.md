---
title: "Insecure cookie High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Insecure cookie [High](severity/high.md)

Insecure cookie settings can lead to unencrypted cookie transmission. Even if a cookie doesn't contain sensitive data now, sensitive data could be added later. It's good practice to transmit all cookies only through secure channels.

**Detector ID**

scala/insecure-cookie@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-614](https://cwe.mitre.org/data/definitions/614.html) [CWE-539](https://cwe.mitre.org/data/definitions/539.html)

**Tags**

[\# cookies](tags/cookies.md) [\# cryptography](tags/cryptography.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```scala
1def nonCompliant(res: HttpServletResponse): Unit = {
2    val cookie = new Cookie("key", "value")
3    cookie.setSecure(true)
4    cookie.setHttpOnly(true)
5    // Noncompliant: MaxAge set to one year.
6    cookie.setMaxAge(31536000)
7    res.addCookie(cookie)
8}

```

#### Compliant example

```scala
1def compliant(res: HttpServletResponse): Unit = {
2    val cookie = new Cookie("key", "value")
3    cookie.setSecure(true)
4    cookie.setHttpOnly(true)
5    // Compliant: MaxAge set to one week.
6    cookie.setMaxAge(604800)
7    res.addCookie(cookie)
8}

```

All content copied from https://docs.aws.amazon.com/.
