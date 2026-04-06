![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-ldap-configuration) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-connection) [Deserialization of Untrusted Data](https://docs.aws.amazon.com/amazonq/detector-library/scala/deserialization-of-untrusted-data) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-servlet-handling) [Use of Insufficiently Random Values](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-insufficiently-random-values) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cookie) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Insecure servlet handling [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Insecure LDAP configuration detected. Make sure anonymous/unauthenticated LDAP connections are not used in your code.

**Detector ID**

scala/insecure-ldap-configuration@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-20](https://cwe.mitre.org/data/definitions/20.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10)

* * *

#### Noncompliant example

```scala
1def nonCompliant(): Unit = {
2  // Noncompliant: `returnObjFlag` is set to true
3  new SearchControls(scope, countLimit, timeLimit, attributes, true, //!! It will flag line 14 ... the beginning of the call
4    deref)
5}

```

#### Compliant example

```scala
1def compliant(): Unit = {
2  // Compliant: `returnObjFlag` is set to false
3  new SearchControls(scope, countLimit, timeLimit, attributes, false,
4    deref)
5}

```
