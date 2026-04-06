![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cryptography) [Template Injection](https://docs.aws.amazon.com/amazonq/detector-library/scala/template-injection) [Untrusted data in http session](https://docs.aws.amazon.com/amazonq/detector-library/scala/untrusted-data-in-http-session) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-ldap-configuration) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-connection) [Deserialization of Untrusted Data](https://docs.aws.amazon.com/amazonq/detector-library/scala/deserialization-of-untrusted-data) [Insecure servlet handling](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-servlet-handling) [Use of Insufficiently Random Values](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-insufficiently-random-values) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cookie) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Insecure Cryptography [Critical](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/critical)

Using insecure cryptographic algorithms or configurations introduces vulnerabilities in applications. This includes weak ciphers like RC4 or DES, ECB mode, no integrity checking, insufficient key sizes, and other known cryptographic weaknesses. Modern secure ciphers like AES-GCM and recommended key sizes should be used instead. Following cryptography best practices is essential to prevent confidentiality and integrity loss.

**Detector ID**

scala/insecure-cryptography@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-310](https://cwe.mitre.org/data/definitions/310.html) [CWE-326](https://cwe.mitre.org/data/definitions/326.html) [CWE-327](https://cwe.mitre.org/data/definitions/327.html) [CWE-328](https://cwe.mitre.org/data/definitions/328.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10)

* * *

#### Noncompliant example

```scala
1@throws[NoSuchAlgorithmException]
2@throws[NoSuchProviderException]
3def weakKeySizeWithProviderString = {
4    val keyGen = KeyPairGenerator.getInstance("RSA", "BC")
5    // Noncompliant: A small key size makes the ciphertext vulnerable to brute force attacks.
6    keyGen.initialize(1024)
7    keyGen.generateKeyPair
8}

```

#### Compliant example

```scala
1@throws[NoSuchAlgorithmException]
2@throws[NoSuchProviderException]
3def strongKeySizeWithProviderString = {
4    val keyGen = KeyPairGenerator.getInstance("RSA", "BC")
5    // Compliant:Key size is 2048 bits.
6    keyGen.initialize(2048)
7    keyGen.generateKeyPair
8}

```
