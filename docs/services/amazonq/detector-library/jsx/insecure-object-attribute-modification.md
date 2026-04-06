![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JSX detectors(78/78)

[Protection mechanism failure](protection-mechanism-failure.md) [Log injection](log-injection.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Use of a deprecated method](deprecated-method.md) [AWS credentials logged](aws-logged-credentials.md) [Improper input validation](improper-input-validation.md) [Insecure cryptography](insecure-cryptography.md) [Catch and swallow exception](swallow-exceptions.md) [File and directory information exposure](file-and-directory-information-exposure.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [SQL injection](sql-injection.md) [Non-literal regular expression](non-literal-regular-expression.md) [Typeof expression](typeof-expression.md) [Batch request with unchecked failures](aws-batch-write-output-ignored.md) [Pseudorandom number generators](pseudorandom-number-generator.md) [Cryptographic key generator](cryptographic-key-generator.md) [Server-side request forgery](server-side-request-forgery.md) [Sensitive information leak](sensitive-information-leak.md) [File injection](file-injection.md) [String passed to \`setInterval\` or \`setTimeout\`](do-not-pass-string-to-setinterval-or-settimeout.md) [Cross-site request forgery](cross-site-request-forgery.md) [Usage of an API that is not recommended](not-recommended-apis.md) [Tainted input for Docker API](tainted-input-for-docker-api.md) [Cross-site scripting](cross-site-scripting.md) [Weak obfuscation of web requests](weak-obfuscation-of-request.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](sns-authenticate-on-unsubscribe.md) [Set SNS Return Subscription ARN](sns-set-return-subscription-arn.md) [XML external entity](xml-external-entity.md) [Resource leak](resource-leak.md) [Improper access control](improper-access-control.md) [Loose file permissions](loose-file-permissions.md) [OS command injection](os-command-injection.md) [Client-side KMS reencryption](aws-kms-reencryption.md) [Insecure CORS policy](insecure-cors-policy.md) [Inefficient polling of AWS resource](aws-polling-instead-of-waiter.md) [New function detected](new-function-detected.md) [Missing pagination](missing-pagination.md) [Avoid nan in comparison](avoid-nan-in-comparison.md) [Header injection](header-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [File extension validation](file-extension-validation.md) [NoSQL injection](nosql-injection.md) [Missing Amazon S3 bucket owner condition](s3-verify-bucket-owner.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Least privilege violation](least-privilege-violation.md) [URL redirection to untrusted site](open-redirect.md) [Insufficiently protected credentials](insufficiently-protected-credentials.md) [Insecure hashing](insecure-hashing.md) [Unsanitized input is run as code](code-injection.md) [Check failed records when using kinesis](kinesis-failed-record-check.md) [Untrusted Amazon Machine Images](untrusted-amazon-machine-images.md) [Session fixation](session-fixation.md) [Data loss in a batch request](unchecked-batch-failures.md) [XPath injection](xpath-injection.md) [Deserialization of untrusted object](untrusted-deserialization.md) [Invoke super appropriately](invoke-super-appropriately.md) [Stack trace exposure](stack-trace-exposure.md) [Timing attack](timing-attack.md) [LDAP injection](ldap-injection.md) [Insecure cookie](insecure-cookie.md) [Sensitive data stored unencrypted due to partial encryption](partial-encryption.md) [Unvalidated expansion of archive files](do-not-expand-archive-files-without-validating.md) [Integer overflow](integer-overflow.md) [SNS don't bind subscribe and publish](sns-no-bind-subscribe-publish.md) [Unverified hostname](unverified-hostname.md) [Improper restriction of rendered UI layers or frames](improper-restriction-of-frames.md) [AWS client not reused in a Lambda function](lambda-client-reuse.md) [Path traversal](path-traversal.md) [Override of reserved variable names in a Lambda function](lambda-override-reserved.md) [Insecure temporary file or directory](insecure-temp-file.md) [Logging of sensitive information](logging-of-sensitive-information.md) [Hardcoded IP address](hardcoded-ip-address.md) [Insecure object attribute modification](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-object-attribute-modification) [Numeric truncation error](https://docs.aws.amazon.com/amazonq/detector-library/jsx/numeric-truncation-error) [DNS prefetching](https://docs.aws.amazon.com/amazonq/detector-library/jsx/dns-prefetching) [Limit request length](https://docs.aws.amazon.com/amazonq/detector-library/jsx/limit-on-request-content-length) [Sendfile injection](https://docs.aws.amazon.com/amazonq/detector-library/jsx/sendfile-injection) [Improper certificate validation](https://docs.aws.amazon.com/amazonq/detector-library/jsx/improper-certificate-validation)

# Insecure object attribute modification [Medium](https://docs.aws.amazon.com/amazonq/detector-library/jsx/severity/medium)

An object attribute constructed from a user-provided input should be considered unsafe because this input can be used to make unexpected modifications to the object.

**Detector ID**

jsx/insecure-object-attribute-modification@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/jsx/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-915](https://cwe.mitre.org/data/definitions/915.html) [CWE-1321](https://cwe.mitre.org/data/definitions/1321.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/jsx/tags/owasp-top10)

* * *

#### Noncompliant example

```javascript
1var express = require('express')
2var app = express()
3function insecureObjectAttributeModificationNoncompliant() {
4    app.get('www.example.com', (req, res) => {
5        var userId = req.params.id
6        // Noncompliant: external input used as object property.
7        req.session.user[userId] = req.body['userDetails']
8    });
9}

```

#### Compliant example

```javascript
1var express = require('express')
2var app = express()
3function insecureObjectAttributeModificationCompliant() {
4    app.get('www.example.com', (req, res) => {
5        var userId = req.params.id
6        // Compliant: checks the type of userId as string.
7        if (typeof userId === 'string') {
8            req.session.user[userId] = req.body['userDetails']
9        }
10    });
11}

```
