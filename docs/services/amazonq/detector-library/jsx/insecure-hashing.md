![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JSX detectors(78/78)

[Protection mechanism failure](protection-mechanism-failure.md) [Log injection](log-injection.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Use of a deprecated method](deprecated-method.md) [AWS credentials logged](aws-logged-credentials.md) [Improper input validation](improper-input-validation.md) [Insecure cryptography](insecure-cryptography.md) [Catch and swallow exception](swallow-exceptions.md) [File and directory information exposure](file-and-directory-information-exposure.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [SQL injection](sql-injection.md) [Non-literal regular expression](non-literal-regular-expression.md) [Typeof expression](typeof-expression.md) [Batch request with unchecked failures](aws-batch-write-output-ignored.md) [Pseudorandom number generators](pseudorandom-number-generator.md) [Cryptographic key generator](cryptographic-key-generator.md) [Server-side request forgery](server-side-request-forgery.md) [Sensitive information leak](sensitive-information-leak.md) [File injection](file-injection.md) [String passed to \`setInterval\` or \`setTimeout\`](do-not-pass-string-to-setinterval-or-settimeout.md) [Cross-site request forgery](cross-site-request-forgery.md) [Usage of an API that is not recommended](not-recommended-apis.md) [Tainted input for Docker API](tainted-input-for-docker-api.md) [Cross-site scripting](cross-site-scripting.md) [Weak obfuscation of web requests](weak-obfuscation-of-request.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](sns-authenticate-on-unsubscribe.md) [Set SNS Return Subscription ARN](sns-set-return-subscription-arn.md) [XML external entity](xml-external-entity.md) [Resource leak](resource-leak.md) [Improper access control](improper-access-control.md) [Loose file permissions](loose-file-permissions.md) [OS command injection](os-command-injection.md) [Client-side KMS reencryption](aws-kms-reencryption.md) [Insecure CORS policy](insecure-cors-policy.md) [Inefficient polling of AWS resource](aws-polling-instead-of-waiter.md) [New function detected](new-function-detected.md) [Missing pagination](missing-pagination.md) [Avoid nan in comparison](avoid-nan-in-comparison.md) [Header injection](header-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [File extension validation](file-extension-validation.md) [NoSQL injection](nosql-injection.md) [Missing Amazon S3 bucket owner condition](s3-verify-bucket-owner.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Least privilege violation](least-privilege-violation.md) [URL redirection to untrusted site](open-redirect.md) [Insufficiently protected credentials](insufficiently-protected-credentials.md) [Insecure hashing](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-hashing) [Unsanitized input is run as code](https://docs.aws.amazon.com/amazonq/detector-library/jsx/code-injection) [Check failed records when using kinesis](https://docs.aws.amazon.com/amazonq/detector-library/jsx/kinesis-failed-record-check) [Untrusted Amazon Machine Images](https://docs.aws.amazon.com/amazonq/detector-library/jsx/untrusted-amazon-machine-images) [Session fixation](https://docs.aws.amazon.com/amazonq/detector-library/jsx/session-fixation) [Data loss in a batch request](https://docs.aws.amazon.com/amazonq/detector-library/jsx/unchecked-batch-failures) [XPath injection](https://docs.aws.amazon.com/amazonq/detector-library/jsx/xpath-injection) [Deserialization of untrusted object](https://docs.aws.amazon.com/amazonq/detector-library/jsx/untrusted-deserialization) [Invoke super appropriately](https://docs.aws.amazon.com/amazonq/detector-library/jsx/invoke-super-appropriately) [Stack trace exposure](https://docs.aws.amazon.com/amazonq/detector-library/jsx/stack-trace-exposure) [Timing attack](https://docs.aws.amazon.com/amazonq/detector-library/jsx/timing-attack) [LDAP injection](https://docs.aws.amazon.com/amazonq/detector-library/jsx/ldap-injection) [Insecure cookie](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-cookie) [Sensitive data stored unencrypted due to partial encryption](https://docs.aws.amazon.com/amazonq/detector-library/jsx/partial-encryption) [Unvalidated expansion of archive files](https://docs.aws.amazon.com/amazonq/detector-library/jsx/do-not-expand-archive-files-without-validating) [Integer overflow](https://docs.aws.amazon.com/amazonq/detector-library/jsx/integer-overflow) [SNS don't bind subscribe and publish](https://docs.aws.amazon.com/amazonq/detector-library/jsx/sns-no-bind-subscribe-publish) [Unverified hostname](https://docs.aws.amazon.com/amazonq/detector-library/jsx/unverified-hostname) [Improper restriction of rendered UI layers or frames](https://docs.aws.amazon.com/amazonq/detector-library/jsx/improper-restriction-of-frames) [AWS client not reused in a Lambda function](https://docs.aws.amazon.com/amazonq/detector-library/jsx/lambda-client-reuse) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/jsx/path-traversal) [Override of reserved variable names in a Lambda function](https://docs.aws.amazon.com/amazonq/detector-library/jsx/lambda-override-reserved) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-temp-file) [Logging of sensitive information](https://docs.aws.amazon.com/amazonq/detector-library/jsx/logging-of-sensitive-information) [Hardcoded IP address](https://docs.aws.amazon.com/amazonq/detector-library/jsx/hardcoded-ip-address) [Insecure object attribute modification](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-object-attribute-modification) [Numeric truncation error](https://docs.aws.amazon.com/amazonq/detector-library/jsx/numeric-truncation-error) [DNS prefetching](https://docs.aws.amazon.com/amazonq/detector-library/jsx/dns-prefetching) [Limit request length](https://docs.aws.amazon.com/amazonq/detector-library/jsx/limit-on-request-content-length) [Sendfile injection](https://docs.aws.amazon.com/amazonq/detector-library/jsx/sendfile-injection) [Improper certificate validation](https://docs.aws.amazon.com/amazonq/detector-library/jsx/improper-certificate-validation)

# Insecure hashing [Medium](https://docs.aws.amazon.com/amazonq/detector-library/jsx/severity/medium)

A hashing algorithm is weak if it is easy to determine the original input from the hash or to find another input that yields the same hash. Weak hashing algorithms can lead to security vulnerabilities.

**Detector ID**

jsx/insecure-hashing@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/jsx/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-327](https://cwe.mitre.org/data/definitions/327.html) [CWE-328](https://cwe.mitre.org/data/definitions/328.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/jsx/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/jsx/tags/owasp-top10)

* * *

#### Noncompliant example

```javascript
1var crypto = require('crypto')
2function insecureHashingNoncompliant()
3{
4    // Noncompliant: 'md5' is weak hash algorithm.
5    var insecure_hash_algo = 'md5'
6    crypto.createHash(insecure_hash_algo)
7}

```

#### Compliant example

```javascript
1var crypto = require('crypto')
2function insecureHashingCompliant()
3{
4    // Compliant: 'SHA-256' is secure hash algorithm.
5    var secure_hash_algo = 'SHA-256'
6    crypto.createHash(secure_hash_algo)
7}

```
