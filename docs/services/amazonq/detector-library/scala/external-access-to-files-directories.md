![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Scala detectors(28/28)

[Improper Neutralization of Special Elements in Data Query](improper-neutralization-of-elements-in-data-query.md) [Avoid Persistent Cookies](avoid-persistent-cookies.md) [Improper Authentication](improper-authentication.md) [Argument Injection](argument-injection.md) [Insecure host name verifier](insecure-host-name-verifier.md) [Insecure Cryptography](insecure-cryptography.md) [Template Injection](template-injection.md) [Untrusted data in http session](untrusted-data-in-http-session.md) [Insecure servlet handling](insecure-ldap-configuration.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Deserialization of Untrusted Data](deserialization-of-untrusted-data.md) [Insecure servlet handling](insecure-servlet-handling.md) [Use of Insufficiently Random Values](use-of-insufficiently-random-values.md) [Insecure cookie](insecure-cookie.md) [Use Of RSA Algorithm](use-of-rsa-algorithm.md) [Path Traversal](path-traversal.md) [URL redirection to untrusted site](open-redirect.md) [Improper Validation Of Array Index](improper-validation-of-array-index.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Insecure jax endpoint usage](insecure-jax-endpoint-usage.md) [XML External Entity](xml-external-entity.md) [Insecure CORS policy](insecure-cors-policy.md) [External Access to Files or Directories](https://docs.aws.amazon.com/amazonq/detector-library/scala/external-access-to-files-directories) [Incorrect Certificate Hostname Verification](https://docs.aws.amazon.com/amazonq/detector-library/scala/incorrect-certificate-hostname-verification) [Improper privilege management](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-privilege-management) [Cross-site scripting](https://docs.aws.amazon.com/amazonq/detector-library/scala/cross-site-scripting) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/scala/improper-certificate-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/scala/do-not-disable-html-autoescape)

# External Access to Files or Directories [High](https://docs.aws.amazon.com/amazonq/detector-library/scala/severity/high)

Files or directories become accessible to unauthorized actors due to the product's operation, violating security protocols.

**Detector ID**

scala/external-access-to-files-directories@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/scala/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-552](https://cwe.mitre.org/data/definitions/552.html)

**Tags**

-

* * *

#### Noncompliant example

```scala
1@throws[IOException]
2def nonComplaint(request: HttpServletRequest, response: HttpServletResponse): Unit = {
3    try {
4    val jspFile = request.getParameter("jspFile")
5    var requestDispatcher = request.getRequestDispatcher(jspFile)
6    // Noncompliant: Granting access to file to unauthorized users.
7    requestDispatcher.include(request, response)
8    requestDispatcher = request.getSession.getServletContext.getRequestDispatcher(jspFile)
9
10    } catch {
11    case e: Exception =>
12        System.out.println(e)
13    }
14}

```

#### Compliant example

```scala
1@throws[IOException]
2def complaint(request: HttpServletRequest, response: HttpServletResponse): Unit = {
3    try {
4    val jspFile = request.getParameter("jspFile")
5    val sanitizedPath = sanitizePath(jspFile)
6    var requestDispatcher = request.getRequestDispatcher(sanitizedPath)
7    // Compliant: Sanitizing the file path before using it with the RequestDispatcher.
8    requestDispatcher.include(request, response)
9    requestDispatcher = request.getSession.getServletContext.getRequestDispatcher(sanitizedPath)
10
11    } catch {
12    case e: Exception =>
13        System.out.println(e)
14    }
15}

```
