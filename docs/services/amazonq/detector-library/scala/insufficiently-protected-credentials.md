---
title: "Insufficient Protected Credentials High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](incorrect-certificate-hostname-verification.md) [Improper privilege management](improper-privilege-management.md) [Cross-site scripting](cross-site-scripting.md) [Improper Certificate Validation](improper-certificate-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md)

# Insufficient Protected Credentials [High](severity/high.md)

The credentials being used do not have sufficient protection measures in place to prevent potential security breaches. Ensure that passwords and other sensitive information are stored in encrypted form.

**Detector ID**

scala/insufficiently-protected-credentials@v1.0

**Category**

[Security](categories/security.md)

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

All content copied from https://docs.aws.amazon.com/.
