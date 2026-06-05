---
title: "XML External Entity High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# XML External Entity [High](severity/high.md)

Objects that parse or handle XML data can lead to XML External Entity (XXE) attacks when not configured properly. Improper restriction of XML external entity processing can lead to server-side request forgery and information disclosure.

**Detector ID**

scala/xml-external-entity@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-611](https://cwe.mitre.org/data/definitions/611.html)

**Tags**

[\# injection](tags/injection.md) [\# xml](tags/xml.md)

* * *

#### Noncompliant example

```scala
1class XmlExternalEntityNoncompliant {
2
3  def nonCompliant(file: File) = {
4    // Noncompliant: XML parsing is not performed with appropriate configurations to disable external entity resolution.
5    val docBuilderFactory = DocumentBuilderFactory.newInstance()
6    val docBuilder = docBuilderFactory.newDocumentBuilder()
7    val doc = docBuilder.parse(file)
8    doc.getDocumentElement().normalize()
9    val foobarList = doc.getElementsByTagName("Foobar")
10    foobarList
11  }
12}

```

#### Compliant example

```scala
1class XmlExternalEntityCompliant {
2
3    def compliant(file: File) = {
4        val docBuilderFactory = DocumentBuilderFactory.newInstance()
5        val docBuilder = docBuilderFactory.newDocumentBuilder()
6        docBuilder.setXIncludeAware(true)
7        docBuilder.setNamespaceAware(true)
8        // Compliant: XML parsing is performed with appropriate configurations to disable external entity resolution.
9        docBuilder.setFeature("http://apache.org/xml/features/disallow-doctype-decl", true)
10        docBuilder.setFeature("http://xml.org/sax/features/external-general-entities", false)
11        docBuilder.setFeature("http://xml.org/sax/features/external-parameter-entities", false)
12
13        val doc = docBuilder.parse(file)
14        doc.getDocumentElement().normalize()
15        val foobarList = doc.getElementsByTagName("Foobar")
16        foobarList
17    }
18}

```

All content copied from https://docs.aws.amazon.com/.
