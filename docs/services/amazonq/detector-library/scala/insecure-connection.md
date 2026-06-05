---
title: "Insecure connection using unencrypted protocol High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Insecure connection using unencrypted protocol [High](severity/high.md)

Connections that use insecure protocols transmit data in cleartext. This introduces a risk of exposing sensitive data to third parties.

**Detector ID**

scala/insecure-connection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-319](https://cwe.mitre.org/data/definitions/319.html)

**Tags**

[\# cryptography](tags/cryptography.md) [\# information-leak](tags/information-leak.md) [\# networking](tags/networking.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```scala
1@throws[IOException]
2private[this] def nonCompliant(): Unit = {
3    // Noncompliant: The connection is not secure.
4    val soc = new Socket("www.google.com", 80)
5    doGetRequest(soc)
6}

```

#### Compliant example

```scala
1@throws[IOException]
2private[this] def compliant(): Unit = {
3    // Compliant: The connection is secure.
4    val soc = SSLSocketFactory.getDefault.createSocket("www.google.com", 443)
5    doGetRequest(soc)
6}

```

All content copied from https://docs.aws.amazon.com/.
