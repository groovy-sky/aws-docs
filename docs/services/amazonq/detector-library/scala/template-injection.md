![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/template-injection) [Untrusted data in http session](https://docs.aws.amazon.com/amazonq/detector-library/scala/untrusted-data-in-http-session) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-ldap-configuration) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-connection) [Deserialization of Untrusted Data](https://docs.aws.amazon.com/amazonq/detector-library/scala/deserialization-of-untrusted-data) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-servlet-handling) [Use of Insufficiently Random Values](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-insufficiently-random-values) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cookie) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Template Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Potential Template Injection vulnerability. User input is directly used in rendering or evaluating templates without proper validation or sanitization.

**Detector ID**

scala/template-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-94](https://cwe.mitre.org/data/definitions/94.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/injection)

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
