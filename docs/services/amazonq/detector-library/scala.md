---
title: "Scala detectors"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](scala/improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](scala/avoid-persistent-cookies.md) [Improper Authentication](scala/improper-authentication.md) [Argument Injection](scala/argument-injection.md) [Insecure host name verifier](scala/insecure-host-name-verifier.md) [Insecure Cryptography](scala/insecure-cryptography.md) [Template Injection](scala/template-injection.md) [Untrusted data in http session](scala/untrusted-data-in-http-session.md) [Insecure servlet handling](scala/insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](scala/insecure-connection.md) [Deserialization of Untrusted Data](scala/deserialization-of-untrusted-data.md) [Insecure servlet handling](scala/insecure-servlet-handling.md) [Use of Insufficiently Random Values](scala/use-of-insufficiently-random-values.md) [Insecure cookie](scala/insecure-cookie.md) [Use Of RSA Algorithm](scala/use-of-rsa-algorithm.md) [Path Traversal](scala/path-traversal.md) [URL redirection to untrusted site](scala/open-redirect.md) [Improper Validation Of Array Index](scala/improper-validation-of-array-index.md) [Insufficient Protected Credentials](scala/insufficiently-protected-credentials.md) [Insecure jax endpoint usage](scala/insecure-jax-endpoint-usage.md) [XML External Entity](scala/xml-external-entity.md) [Insecure CORS policy](scala/insecure-cors-policy.md) [External Access to Files or Directories](scala/external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](scala/incorrect-certificate-hostname-verification.md) [Improper privilege management](scala/improper-privilege-management.md) [Cross-site scripting](scala/cross-site-scripting.md) [Improper Certificate Validation](scala/improper-certificate-validation.md) [Disabled HTML autoescape](scala/do-not-disable-html-autoescape.md)

# Scala detectors

Showing all detectors for the Scala language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](scala/tags.md)

##### Browse by severity

Browse all detectors by severity.

[Click here→](scala/severity.md)

##### Browse by category

Browse all detectors by category.

[Click here→](scala/categories.md)

* * *

### Browse all detectors

### [Improper Neutralization of Special Elements in Data Query](scala/improper-neutralization-of-elements-in-data-query.md)

The application constructs a query with inadequate neutralization of special elements, risking query logic manipulation.

### [Avoid Persistent Cookies](scala/avoid-persistent-cookies.md)

Persistent cookies are vulnerable to attacks.

### [Improper Authentication](scala/improper-authentication.md)

Security issue where software mishandles XML data from unreliable sources.

### [Argument Injection](scala/argument-injection.md)

Improper Neutralization of Argument Delimiters in a Command .

### [Insecure host name verifier](scala/insecure-host-name-verifier.md)

The software does not validate or improperly validate host name.

### [Insecure Cryptography](scala/insecure-cryptography.md)

Use of insecure cryptography

### [Template Injection](scala/template-injection.md)

User input is directly used in rendering or evaluating templates without proper validation or sanitization.

### [Untrusted data in http session](scala/untrusted-data-in-http-session.md)

User input in `setAttribute` could lead to trust boundary violation.

### [Insecure servlet handling](scala/insecure-ldap-configuration.md)

Insecure LDAP configuration detected.

### [Insecure connection using unencrypted protocol](scala/insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Deserialization of Untrusted Data](scala/deserialization-of-untrusted-data.md)

Deserializing of data from untrusted sources.

### [Insecure servlet handling](scala/insecure-servlet-handling.md)

The Servlet can read GET and POST parameters from various methods. The value obtained should be considered unsafe.

### [Use of Insufficiently Random Values](scala/use-of-insufficiently-random-values.md)

The product relies on random numbers or values that aren't random enough for security purposes, especially in situations where unpredictability is crucial.

### [Insecure cookie](scala/insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Use Of RSA Algorithm](scala/use-of-rsa-algorithm.md)

RSA algorithm does not incorporate Optimal Asymmetric Encryption Padding (OAEP), which might weaken the encryption.

### [Path Traversal](scala/path-traversal.md)

Improper input validation, sanitization, and access controls are can lead to path traversal vulnerabilities.

### [URL redirection to untrusted site](scala/open-redirect.md)

User-controlled input that specifies a link to an external site could lead to phishing attacks and allow user credentials to be stolen.

### [Improper Validation Of Array Index](scala/improper-validation-of-array-index.md)

Array Index Validation Failure enables attackers to execute code or cause a denial of service by manipulating array index values.

### [Insufficient Protected Credentials](scala/insufficiently-protected-credentials.md)

The credentials provided are not adequately protected against security threats.

### [Insecure jax endpoint usage](scala/insecure-jax-endpoint-usage.md)

Insecure usage of web service methods can enable attacks and lead to unwanted behavior.

### [XML External Entity](scala/xml-external-entity.md)

Objects that parse or handle XML can lead to XML External Entity (XXE) attacks when misconfigured.

### [Insecure CORS policy](scala/insecure-cors-policy.md)

Cross-origin resource sharing policies that are too permissive could lead to security vulnerabilities.

### [External Access to Files or Directories](scala/external-access-to-files-directories.md)

External parties gain unauthorized access to files or directories via the product.

### [Incorrect Certificate Hostname Verification](scala/incorrect-certificate-hostname-verification.md)

Improper Validation of Certificate with Host Mismatch.

### [Improper privilege management](scala/improper-privilege-management.md)

Granting unsafe permissions can lead to security vulnerabilities.

### [Cross-site scripting](scala/cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Improper Certificate Validation](scala/improper-certificate-validation.md)

Improper certificate validation might allow an attacker to spoof a trusted entity by interfering in the communication path between the host and client.

### [Disabled HTML autoescape](scala/do-not-disable-html-autoescape.md)

Disabling the HTML autoescape mechanism exposes your web applications to attacks.

All content copied from https://docs.aws.amazon.com/.
