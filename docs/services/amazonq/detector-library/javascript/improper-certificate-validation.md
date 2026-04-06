![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JavaScript detectors(78/78)

[Improper access control](improper-access-control.md) [Sensitive data stored unencrypted due to partial encryption](partial-encryption.md) [Pseudorandom number generators](pseudorandom-number-generator.md) [OS command injection](os-command-injection.md) [URL redirection to untrusted site](open-redirect.md) [Integer overflow](integer-overflow.md) [Protection mechanism failure](protection-mechanism-failure.md) [Non-literal regular expression](non-literal-regular-expression.md) [Tainted input for Docker API](tainted-input-for-docker-api.md) [Usage of an API that is not recommended](not-recommended-apis.md) [XML external entity](xml-external-entity.md) [Server-side request forgery](server-side-request-forgery.md) [New function detected](new-function-detected.md) [Stack trace exposure](stack-trace-exposure.md) [Timing attack](timing-attack.md) [SNS don't bind subscribe and publish](sns-no-bind-subscribe-publish.md) [Invoke super appropriately](invoke-super-appropriately.md) [NoSQL injection](nosql-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Insecure cookie](insecure-cookie.md) [Cross-site scripting](cross-site-scripting.md) [Hardcoded IP address](hardcoded-ip-address.md) [AWS credentials logged](aws-logged-credentials.md) [XPath injection](xpath-injection.md) [Data loss in a batch request](unchecked-batch-failures.md) [Path traversal](path-traversal.md) [Least privilege violation](least-privilege-violation.md) [DNS prefetching](dns-prefetching.md) [Resource leak](resource-leak.md) [Insufficiently protected credentials](insufficiently-protected-credentials.md) [File extension validation](file-extension-validation.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Cross-site request forgery](cross-site-request-forgery.md) [Typeof expression](typeof-expression.md) [Set SNS Return Subscription ARN](sns-set-return-subscription-arn.md) [File and directory information exposure](file-and-directory-information-exposure.md) [Missing Amazon S3 bucket owner condition](s3-verify-bucket-owner.md) [Insecure hashing](insecure-hashing.md) [Numeric truncation error](numeric-truncation-error.md) [Client-side KMS reencryption](aws-kms-reencryption.md) [AWS client not reused in a Lambda function](lambda-client-reuse.md) [LDAP injection](ldap-injection.md) [Batch request with unchecked failures](aws-batch-write-output-ignored.md) [Cryptographic key generator](cryptographic-key-generator.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](sns-authenticate-on-unsubscribe.md) [Unverified hostname](unverified-hostname.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Loose file permissions](loose-file-permissions.md) [Unsanitized input is run as code](code-injection.md) [Missing pagination](missing-pagination.md) [Untrusted Amazon Machine Images](untrusted-amazon-machine-images.md) [Improper certificate validation](https://docs.aws.amazon.com/amazonq/detector-library/javascript/improper-certificate-validation) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-cors-policy) [Deserialization of untrusted object](https://docs.aws.amazon.com/amazonq/detector-library/javascript/untrusted-deserialization) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sensitive-information-leak) [Check failed records when using kinesis](https://docs.aws.amazon.com/amazonq/detector-library/javascript/kinesis-failed-record-check) [Weak obfuscation of web requests](https://docs.aws.amazon.com/amazonq/detector-library/javascript/weak-obfuscation-of-request) [Catch and swallow exception](https://docs.aws.amazon.com/amazonq/detector-library/javascript/swallow-exceptions) [Logging of sensitive information](https://docs.aws.amazon.com/amazonq/detector-library/javascript/logging-of-sensitive-information) [Limit request length](https://docs.aws.amazon.com/amazonq/detector-library/javascript/limit-on-request-content-length) [String passed to \`setInterval\` or \`setTimeout\`](https://docs.aws.amazon.com/amazonq/detector-library/javascript/do-not-pass-string-to-setinterval-or-settimeout) [Log injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/log-injection) [Override of reserved variable names in a Lambda function](https://docs.aws.amazon.com/amazonq/detector-library/javascript/lambda-override-reserved) [Improper restriction of rendered UI layers or frames](https://docs.aws.amazon.com/amazonq/detector-library/javascript/improper-restriction-of-frames) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-cryptography) [Insecure object attribute modification](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-object-attribute-modification) [Session fixation](https://docs.aws.amazon.com/amazonq/detector-library/javascript/session-fixation) [Avoid nan in comparison](https://docs.aws.amazon.com/amazonq/detector-library/javascript/avoid-nan-in-comparison) [Improper input validation](https://docs.aws.amazon.com/amazonq/detector-library/javascript/improper-input-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/javascript/do-not-disable-html-autoescape) [Use of a deprecated method](https://docs.aws.amazon.com/amazonq/detector-library/javascript/deprecated-method) [Unvalidated expansion of archive files](https://docs.aws.amazon.com/amazonq/detector-library/javascript/do-not-expand-archive-files-without-validating) [File injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/file-injection) [Sendfile injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sendfile-injection) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sql-injection) [Header injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/header-injection) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-temp-file) [Inefficient polling of AWS resource](https://docs.aws.amazon.com/amazonq/detector-library/javascript/aws-polling-instead-of-waiter)

# Improper certificate validation [High](https://docs.aws.amazon.com/amazonq/detector-library/javascript/severity/high)

Lack of validation or insufficient validation of a security certificate can lead to host impersonation and sensitive data leaks.

**Detector ID**

javascript/improper-certificate-validation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/javascript/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-295](https://cwe.mitre.org/data/definitions/295.html)

**Tags**

[\# access-control](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/access-control) [\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/owasp-top10)

* * *

#### Noncompliant example

```javascript
1var tls = require("tls")
2var fs = require("fs")
3
4function improperCertificateValidationNoncompliant() {
5    var options = {
6        host: 'encrypted.example.com',
7        // Noncompliant: rejectUnauthorized is set to 'false'.
8        rejectUnauthorized: false
9    }
10
11    tls.createServer(options, (req, res) => {
12        res.writeHead(200)
13        res.end()
14    }).listen(8000)
15}

```

#### Compliant example

```javascript
1var tls = require("tls")
2var fs = require("fs")
3
4function improperCertificateValidationCompliant() {
5    var options = {
6        host: 'encrypted.example.com',
7        // Compliant: certificate is provided.
8        key: fs.readFileSync('keys/client-key.pem'),
9        cert: fs.readFileSync('keys/client-cert.pem')
10    }
11
12    tls.createServer(options, (req, res) => {
13        res.writeHead(200)
14        res.end()
15    }).listen(8000)
16}

```
