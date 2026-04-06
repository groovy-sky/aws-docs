![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](../improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](../avoid-persistent-cookies.md) [Improper Authentication](../improper-authentication.md) [Argument Injection](../argument-injection.md) [Insecure host name verifier](../insecure-host-name-verifier.md) [Insecure Cryptography](../insecure-cryptography.md) [Template Injection](../template-injection.md) [Untrusted data in http session](../untrusted-data-in-http-session.md) [Insecure servlet handling](../insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Deserialization of Untrusted Data](../deserialization-of-untrusted-data.md) [Insecure servlet handling](../insecure-servlet-handling.md) [Use of Insufficiently Random Values](../use-of-insufficiently-random-values.md) [Insecure cookie](../insecure-cookie.md) [Use Of RSA Algorithm](../use-of-rsa-algorithm.md) [Path Traversal](../path-traversal.md) [URL redirection to untrusted site](../open-redirect.md) [Improper Validation Of Array Index](../improper-validation-of-array-index.md) [Insufficient Protected Credentials](../insufficiently-protected-credentials.md) [Insecure jax endpoint usage](../insecure-jax-endpoint-usage.md) [XML External Entity](../xml-external-entity.md) [Insecure CORS policy](../insecure-cors-policy.md) [External Access to Files or Directories](../external-access-to-files-directories.md) [Incorrect Certificate Hostname Verification](../incorrect-certificate-hostname-verification.md) [Improper privilege management](../improper-privilege-management.md) [Cross-site scripting](../cross-site-scripting.md) [Improper Certificate Validation](../improper-certificate-validation.md) [Disabled HTML autoescape](../do-not-disable-html-autoescape.md)

# Tag: owasp-top10

### [Insecure Cryptography](../insecure-cryptography.md)

Use of insecure cryptography

### [Untrusted data in http session](../untrusted-data-in-http-session.md)

User input in `setAttribute` could lead to trust boundary violation.

### [Insecure servlet handling](../insecure-ldap-configuration.md)

Insecure LDAP configuration detected.

### [Insecure connection using unencrypted protocol](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Insecure servlet handling](../insecure-servlet-handling.md)

The Servlet can read GET and POST parameters from various methods. The value obtained should be considered unsafe.

### [Insecure cookie](../insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Use Of RSA Algorithm](../use-of-rsa-algorithm.md)

RSA algorithm does not incorporate Optimal Asymmetric Encryption Padding (OAEP), which might weaken the encryption.

### [Path Traversal](../path-traversal.md)

Improper input validation, sanitization, and access controls are can lead to path traversal vulnerabilities.

### [URL redirection to untrusted site](../open-redirect.md)

User-controlled input that specifies a link to an external site could lead to phishing attacks and allow user credentials to be stolen.

### [Insecure CORS policy](../insecure-cors-policy.md)

Cross-origin resource sharing policies that are too permissive could lead to security vulnerabilities.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Disabled HTML autoescape](../do-not-disable-html-autoescape.md)

Disabling the HTML autoescape mechanism exposes your web applications to attacks.
