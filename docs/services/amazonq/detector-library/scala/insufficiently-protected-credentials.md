![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Insufficient Protected Credentials [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

The credentials being used do not have sufficient protection measures in place to prevent potential security breaches. Ensure that passwords and other sensitive information are stored in encrypted form.

**Detector ID**

scala/insufficiently-protected-credentials@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-522](https://cwe.mitre.org/data/definitions/522.html)

**Tags**

-

* * *

#### Noncompliant example

```scala
1def nonCompliant(token: String) = {
2  val algo = JwtAlgorithm.HS256
3  // Noncompliant: hardcoded JWT secret is used.
4  JwtArgonaut.decodeJson(token, secretKey, algo)
5}

```

#### Compliant example

```scala
1class InsufficientlyProtectedCredentialsCompliant {
2    def compliant(token: String) = {
3    val algo = JwtAlgorithm.HS256
4    // Compliant: `getSecretFromEnv` method used to get JWT secret.
5    JwtArgonaut.decodeJson(token, getSecretFromEnv(), algo)
6  }
7}

```
