![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-authentication) [Argument Injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/argument-injection) [Insecure host name verifier](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-host-name-verifier) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cryptography) [Template Injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/template-injection) [Untrusted data in http session](https://docs.aws.amazon.com/amazonq/detector-library/scala/untrusted-data-in-http-session) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-ldap-configuration) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-connection) [Deserialization of Untrusted Data](https://docs.aws.amazon.com/amazonq/detector-library/scala/deserialization-of-untrusted-data) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-servlet-handling) [Use of Insufficiently Random Values](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-insufficiently-random-values) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cookie) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Improper Authentication [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Security issue where software mishandles XML data from unreliable sources, creating a way for attackers to sneak in harmful commands or access secret information.

**Detector ID**

scala/improper-authentication@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-287](https://cwe.mitre.org/data/definitions/287.html)

**Tags**

-

* * *

#### Noncompliant example

```scala
1def nonCompliant(): BasicParserPool = {
2    val parserPool = new BasicParserPool
3    // Noncompliant: Defining a method to create a parser pool with ignoreComments set to false.
4    parserPool.setIgnoreComments(false)
5    parserPool
6}

```

#### Compliant example

```scala
1def compliant(): BasicParserPool = {
2    val parserPool = new BasicParserPool
3    // Compliant: Defining a method to create a parser pool with ignoreComments set to true.
4    parserPool.setIgnoreComments(true)
5    parserPool
6}

```
