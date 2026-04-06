![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Improper Certificate Validation [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

When a certificate is invalid or malicious, it might allow an attacker to spoof a trusted entity by interfering in the communication path between the host and client. The product might connect to a malicious host while believing it is a trusted host, or the product might be deceived into accepting spoofed data that appears to originate from a trusted host.

**Detector ID**

scala/improper-certificate-validation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-295](https://cwe.mitre.org/data/definitions/295.html)

**Tags**

-

* * *

#### Noncompliant example

```scala
1def nonCompliant(args: Array): Unit = {
2  try {
3    // Noncompliant: the default hostname verifier is not used.
4    val context1 = SSLContext.getInstance("SSL")
5  } catch {
6    case e: NoSuchAlgorithmException =>
7      e.printStackTrace
8  }
9}

```

#### Compliant example

```scala
1def compliant(args: Array): Unit = {
2  try {
3    // Compliant: the default hostname verifier is used.
4    val context2 = SSLContext.getInstance("TLS")
5  } catch {
6    case e: NoSuchAlgorithmException =>
7      e.printStackTrace
8  }
9}

```
