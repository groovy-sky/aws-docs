![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](https://docs.aws.amazon.com/amazonq/detector-library/scala/use-of-rsa-algorithm) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/scala/path-traversal) [URL redirection to untrusted site](https://docs.aws.amazon.com/amazonq/detector-library/scala/open-redirect) [Improper Validation Of Array Index](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-validation-of-array-index) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/scala/insufficiently-protected-credentials) [Insecure jax endpoint usage](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-jax-endpoint-usage) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/scala/xml-external-entity) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/scala/insecure-cors-policy) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# Use Of RSA Algorithm [Critical](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/critical)

Padding schemes are often used with cryptographic algorithms to make the plaintext less predictable and complicate attack efforts. The OAEP scheme is often used with RSA to nullify the impact of predictable common text.

**Detector ID**

scala/use-of-rsa-algorithm@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-780](https://cwe.mitre.org/data/definitions/780.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/scala/tags/owasp-top10)

* * *

#### Noncompliant example

```scala
1class UseOfRSAAlgorithmNoncompliant {
2    @throws[Exception]
3    def nonCompliant(): Unit = {
4        val cipher1 = null
5        Cipher.getInstance(cipher1)
6        val cipher2 = "RSA/NONE/NoPadding"
7        // Noncompliant: Use of RSA Algorithm without OAEP.
8        Cipher.getInstance(cipher2)
9    }
10}

```

#### Compliant example

```scala
1object UseOfRSAAlgorithmCompliant {
2  def compliant(args: Array[String]): Unit = {
3    // Compliant: Encrypt with RSA using OAEP padding.
4    val cipher = Cipher.getInstance("RSA/ECB/OAEPWithSHA-256AndMGF1Padding")
5    cipher.init(Cipher.ENCRYPT_MODE, publicKey)
6    val ciphertext = cipher.doFinal(plaintext.getBytes("UTF-8"))
7    println("Encrypted: " + javax.xml.bind.DatatypeConverter.printHexBinary(ciphertext))
8  }
9}

```
