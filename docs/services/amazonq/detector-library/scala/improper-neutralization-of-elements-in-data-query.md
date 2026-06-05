---
title: "Improper Neutralization of Special Elements in Data Query High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Improper Neutralization of Special Elements in Data Query [High](severity/high.md)

The software generates queries aimed at accessing or altering data stored in a database. However, it overlooks the proper neutralization or incorrectly neutralizes special elements within these queries, opening up the possibility of unintended alterations to the query's logic.

**Detector ID**

scala/improper-neutralization-of-elements-in-data-query@v1.0

**Category**

[Security](categories/security.md)

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

All content copied from https://docs.aws.amazon.com/.
