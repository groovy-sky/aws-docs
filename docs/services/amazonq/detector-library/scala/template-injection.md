---
title: "Template Injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Template Injection [High](severity/high.md)

Potential Template Injection vulnerability. User input is directly used in rendering or evaluating templates without proper validation or sanitization.

**Detector ID**

scala/template-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-94](https://cwe.mitre.org/data/definitions/94.html)

**Tags**

[\# injection](tags/injection.md)

* * *

#### Noncompliant example

```scala
1class TemplateInjectionNoncompliant {
2  @throws[FileNotFoundException]
3  def nonCompliant(inputFile: String): Unit = {
4    Velocity.init
5    val context = new VelocityContext
6    context.put("author", "Elliot A.")
7    context.put("address", "217 E Broadway")
8    context.put("phone", "555-1337")
9    val file = new FileInputStream(inputFile)
10    val swOut = new StringWriter
11    // Noncompliant: User input is directly used in evaluating templates without proper validation or sanitization.
12    Velocity.evaluate(context, swOut, "test", file.toString)
13    val result = swOut.getBuffer.toString
14    System.out.println(result)
15  }
16}

```

#### Compliant example

```scala
1class TemplateInjectionCompliant {
2  @throws[IOException]
3  def compliant(inputFile: String): String = {
4    val engine = new PebbleEngine.Builder().build
5    var compiledTemplate: PebbleTemplate = null
6    val context = new HashMap[String, Object]
7    context.put("name", "Shivam")
8    val writer = new StringWriter
9    try {
10      // Compliant: User input is not directly used in any code.
11      compiledTemplate.evaluate(writer, context)
12    } catch {
13      case e: Exception =>
14        e.printStackTrace()
15        throw new IOException("Error while evaluating template", e)
16    }
17    writer.toString
18  }
19}

```

All content copied from https://docs.aws.amazon.com/.
