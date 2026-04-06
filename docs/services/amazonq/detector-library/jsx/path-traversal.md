![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JSX detectors(78/78)

[Protection mechanism failure](protection-mechanism-failure.md) [Log injection](log-injection.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Use of a deprecated method](deprecated-method.md) [AWS credentials logged](aws-logged-credentials.md) [Improper input validation](improper-input-validation.md) [Insecure cryptography](insecure-cryptography.md) [Catch and swallow exception](swallow-exceptions.md) [File and directory information exposure](file-and-directory-information-exposure.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [SQL injection](sql-injection.md) [Non-literal regular expression](non-literal-regular-expression.md) [Typeof expression](typeof-expression.md) [Batch request with unchecked failures](aws-batch-write-output-ignored.md) [Pseudorandom number generators](pseudorandom-number-generator.md) [Cryptographic key generator](cryptographic-key-generator.md) [Server-side request forgery](server-side-request-forgery.md) [Sensitive information leak](sensitive-information-leak.md) [File injection](file-injection.md) [String passed to \`setInterval\` or \`setTimeout\`](do-not-pass-string-to-setinterval-or-settimeout.md) [Cross-site request forgery](cross-site-request-forgery.md) [Usage of an API that is not recommended](not-recommended-apis.md) [Tainted input for Docker API](tainted-input-for-docker-api.md) [Cross-site scripting](cross-site-scripting.md) [Weak obfuscation of web requests](weak-obfuscation-of-request.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](sns-authenticate-on-unsubscribe.md) [Set SNS Return Subscription ARN](sns-set-return-subscription-arn.md) [XML external entity](xml-external-entity.md) [Resource leak](resource-leak.md) [Improper access control](improper-access-control.md) [Loose file permissions](loose-file-permissions.md) [OS command injection](os-command-injection.md) [Client-side KMS reencryption](aws-kms-reencryption.md) [Insecure CORS policy](insecure-cors-policy.md) [Inefficient polling of AWS resource](aws-polling-instead-of-waiter.md) [New function detected](new-function-detected.md) [Missing pagination](missing-pagination.md) [Avoid nan in comparison](avoid-nan-in-comparison.md) [Header injection](header-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [File extension validation](file-extension-validation.md) [NoSQL injection](nosql-injection.md) [Missing Amazon S3 bucket owner condition](s3-verify-bucket-owner.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Least privilege violation](least-privilege-violation.md) [URL redirection to untrusted site](open-redirect.md) [Insufficiently protected credentials](insufficiently-protected-credentials.md) [Insecure hashing](insecure-hashing.md) [Unsanitized input is run as code](code-injection.md) [Check failed records when using kinesis](kinesis-failed-record-check.md) [Untrusted Amazon Machine Images](untrusted-amazon-machine-images.md) [Session fixation](session-fixation.md) [Data loss in a batch request](unchecked-batch-failures.md) [XPath injection](xpath-injection.md) [Deserialization of untrusted object](untrusted-deserialization.md) [Invoke super appropriately](invoke-super-appropriately.md) [Stack trace exposure](stack-trace-exposure.md) [Timing attack](timing-attack.md) [LDAP injection](ldap-injection.md) [Insecure cookie](insecure-cookie.md) [Sensitive data stored unencrypted due to partial encryption](partial-encryption.md) [Unvalidated expansion of archive files](do-not-expand-archive-files-without-validating.md) [Integer overflow](integer-overflow.md) [SNS don't bind subscribe and publish](sns-no-bind-subscribe-publish.md) [Unverified hostname](unverified-hostname.md) [Improper restriction of rendered UI layers or frames](improper-restriction-of-frames.md) [AWS client not reused in a Lambda function](lambda-client-reuse.md) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/jsx/path-traversal) [Override of reserved variable names in a Lambda function](https://docs.aws.amazon.com/amazonq/detector-library/jsx/lambda-override-reserved) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-temp-file) [Logging of sensitive information](https://docs.aws.amazon.com/amazonq/detector-library/jsx/logging-of-sensitive-information) [Hardcoded IP address](https://docs.aws.amazon.com/amazonq/detector-library/jsx/hardcoded-ip-address) [Insecure object attribute modification](https://docs.aws.amazon.com/amazonq/detector-library/jsx/insecure-object-attribute-modification) [Numeric truncation error](https://docs.aws.amazon.com/amazonq/detector-library/jsx/numeric-truncation-error) [DNS prefetching](https://docs.aws.amazon.com/amazonq/detector-library/jsx/dns-prefetching) [Limit request length](https://docs.aws.amazon.com/amazonq/detector-library/jsx/limit-on-request-content-length) [Sendfile injection](https://docs.aws.amazon.com/amazonq/detector-library/jsx/sendfile-injection) [Improper certificate validation](https://docs.aws.amazon.com/amazonq/detector-library/jsx/improper-certificate-validation)

# Path traversal [High](https://docs.aws.amazon.com/amazonq/detector-library/jsx/severity/high)

Creating file paths from untrusted input could allow a malicious actor to access arbitrary files on a disk by manipulating the file name in the path.

**Detector ID**

jsx/path-traversal@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/jsx/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html) [CWE-23](https://cwe.mitre.org/data/definitions/23.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/jsx/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/jsx/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/jsx/tags/top25-cwes)

* * *

#### Noncompliant example

```javascript
1var express = require('express')
2var app = express()
3var path = require('path')
4function pathTraversalNoncompliant() {
5    app.get('/products', (req, res) => {
6        const basePath = '/data/product/images/'
7        // Noncompliant: user-supplied path is not sanitized and could contain malicious characters.
8        var targetPath = path.join(basePath, req.query.path)
9        retrieveProduct(targetPath)
10        res.send('Here is your requested product!')
11    })
12}

```

#### Compliant example

```javascript
1var express = require('express')
2var app = express()
3var path = require('path')
4function pathTraversalCompliant() {
5    app.get('/products', (req, res) => {
6        const basePath = '/data/product/images/'
7        // Compliant: user-supplied relative-path is sanitized.
8        const queryPath = sanitizer(req.query.path)
9        if(queryPath) {
10            const targetPath = path.join(basePath, queryPath)
11            retrieveProduct(targetPath)
12            res.send('Here is your requested product!')
13        }
14        else
15            res.send('Invalid product!')
16    })
17}
18function sanitizer(path) {
19    return path.match(/^[a-z]+$/) ? path : null
20}

```
