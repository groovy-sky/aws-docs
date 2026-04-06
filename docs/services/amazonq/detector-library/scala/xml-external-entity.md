![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# XML External Entity [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Objects that parse or handle XML data can lead to XML External Entity (XXE) attacks when not configured properly. Improper restriction of XML external entity processing can lead to server-side request forgery and information disclosure.

**Detector ID**

scala/xml-external-entity@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-611](https://cwe.mitre.org/data/definitions/611.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/injection) [\# xml](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/xml)

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
