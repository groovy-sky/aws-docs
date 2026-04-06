![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-neutralization-of-elements-in-data-query) [Avoid Persistent Cookies](https://docs.aws.amazon.com/amazonq/detector-library/scala/avoid-persistent-cookies) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-authentication) [Argument Injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/argument-injection) [Insecure host name verifier](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-host-name-verifier) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cryptography) [Template Injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/template-injection) [Untrusted data in http session](https://docs.aws.amazon.com/amazonq/detector-library/scala/untrusted-data-in-http-session) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-ldap-configuration) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-connection) [Deserialization of Untrusted Data](https://docs.aws.amazon.com/amazonq/detector-library/scala/deserialization-of-untrusted-data) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-servlet-handling) [Use of Insufficiently Random Values](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-insufficiently-random-values) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cookie) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Improper Neutralization of Special Elements in Data Query [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

The software generates queries aimed at accessing or altering data stored in a database. However, it overlooks the proper neutralization or incorrectly neutralizes special elements within these queries, opening up the possibility of unintended alterations to the query's logic.

**Detector ID**

scala/improper-neutralization-of-elements-in-data-query@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-943](https://cwe.mitre.org/data/definitions/943.html)

**Tags**

-

* * *

#### Noncompliant example

```scala
1@throws[IOException]
2override def nonCompliant(request: HttpServletRequest, response: HttpServletResponse): Unit = {
3    try {
4        val customerID = request.getParameter("customerID")
5        val awsCredentials = new BasicAWSCredentials("test", "test")
6        val sdbc = new AmazonSimpleDBClient(awsCredentials)
7        val query = "select * from invoices where customerID = " + customerID
8        // Noncompliant: Using untrusted HTTP request parameters into SQL queries.
9        val sdbResult = sdbc.select(new SelectRequest(query))
10    } catch {
11        case _: Throwable =>
12    }
13}

```

#### Compliant example

```scala
1@throws[IOException]
2override def compliant(request: HttpServletRequest, response: HttpServletResponse): Unit = {
3    try {
4        val customerID = request.getParameter("customerID")
5        val awsCredentials = new BasicAWSCredentials("test", "test")
6        val sdbc = new AmazonSimpleDBClient(awsCredentials)
7        val query = "select * from invoices where customerID = 123"
8        // Compliant: No untrusted input is used in the query.
9        val sdbResult = sdbc.select(new SelectRequest(query))
10    } catch {
11        case _: Throwable =>
12    }
13}

```
