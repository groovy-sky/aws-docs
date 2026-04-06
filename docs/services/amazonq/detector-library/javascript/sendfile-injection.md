![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JavaScript detectors(78/78)

[Improper access control](improper-access-control.md) [Sensitive data stored unencrypted due to partial encryption](partial-encryption.md) [Pseudorandom number generators](pseudorandom-number-generator.md) [OS command injection](os-command-injection.md) [URL redirection to untrusted site](open-redirect.md) [Integer overflow](integer-overflow.md) [Protection mechanism failure](protection-mechanism-failure.md) [Non-literal regular expression](non-literal-regular-expression.md) [Tainted input for Docker API](tainted-input-for-docker-api.md) [Usage of an API that is not recommended](not-recommended-apis.md) [XML external entity](xml-external-entity.md) [Server-side request forgery](server-side-request-forgery.md) [New function detected](new-function-detected.md) [Stack trace exposure](stack-trace-exposure.md) [Timing attack](timing-attack.md) [SNS don't bind subscribe and publish](sns-no-bind-subscribe-publish.md) [Invoke super appropriately](invoke-super-appropriately.md) [NoSQL injection](nosql-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Insecure cookie](insecure-cookie.md) [Cross-site scripting](cross-site-scripting.md) [Hardcoded IP address](hardcoded-ip-address.md) [AWS credentials logged](aws-logged-credentials.md) [XPath injection](xpath-injection.md) [Data loss in a batch request](unchecked-batch-failures.md) [Path traversal](path-traversal.md) [Least privilege violation](least-privilege-violation.md) [DNS prefetching](dns-prefetching.md) [Resource leak](resource-leak.md) [Insufficiently protected credentials](insufficiently-protected-credentials.md) [File extension validation](file-extension-validation.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Cross-site request forgery](cross-site-request-forgery.md) [Typeof expression](typeof-expression.md) [Set SNS Return Subscription ARN](sns-set-return-subscription-arn.md) [File and directory information exposure](file-and-directory-information-exposure.md) [Missing Amazon S3 bucket owner condition](s3-verify-bucket-owner.md) [Insecure hashing](insecure-hashing.md) [Numeric truncation error](numeric-truncation-error.md) [Client-side KMS reencryption](aws-kms-reencryption.md) [AWS client not reused in a Lambda function](lambda-client-reuse.md) [LDAP injection](ldap-injection.md) [Batch request with unchecked failures](aws-batch-write-output-ignored.md) [Cryptographic key generator](cryptographic-key-generator.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](sns-authenticate-on-unsubscribe.md) [Unverified hostname](unverified-hostname.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Loose file permissions](loose-file-permissions.md) [Unsanitized input is run as code](code-injection.md) [Missing pagination](missing-pagination.md) [Untrusted Amazon Machine Images](untrusted-amazon-machine-images.md) [Improper certificate validation](improper-certificate-validation.md) [Insecure CORS policy](insecure-cors-policy.md) [Deserialization of untrusted object](untrusted-deserialization.md) [Sensitive information leak](sensitive-information-leak.md) [Check failed records when using kinesis](kinesis-failed-record-check.md) [Weak obfuscation of web requests](weak-obfuscation-of-request.md) [Catch and swallow exception](swallow-exceptions.md) [Logging of sensitive information](logging-of-sensitive-information.md) [Limit request length](limit-on-request-content-length.md) [String passed to \`setInterval\` or \`setTimeout\`](do-not-pass-string-to-setinterval-or-settimeout.md) [Log injection](log-injection.md) [Override of reserved variable names in a Lambda function](lambda-override-reserved.md) [Improper restriction of rendered UI layers or frames](improper-restriction-of-frames.md) [Insecure cryptography](insecure-cryptography.md) [Insecure object attribute modification](insecure-object-attribute-modification.md) [Session fixation](session-fixation.md) [Avoid nan in comparison](avoid-nan-in-comparison.md) [Improper input validation](improper-input-validation.md) [Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Use of a deprecated method](deprecated-method.md) [Unvalidated expansion of archive files](do-not-expand-archive-files-without-validating.md) [File injection](file-injection.md) [Sendfile injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sendfile-injection) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sql-injection) [Header injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/header-injection) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-temp-file) [Inefficient polling of AWS resource](https://docs.aws.amazon.com/amazonq/detector-library/javascript/aws-polling-instead-of-waiter)

# Sendfile injection [High](https://docs.aws.amazon.com/amazonq/detector-library/javascript/severity/high)

User-provided inputs must be sanitized before being passed to `res.sendFile`. Otherwise an attacker could arbitrarily read files on the system through path traversal.

**Detector ID**

javascript/sendfile-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/javascript/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-73](https://cwe.mitre.org/data/definitions/73.html)

**Tags**

[\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/top25-cwes) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/owasp-top10)

* * *

#### Noncompliant example

```javascript
1var express = require("express")
2var path = require("path")
3var app = express()
4function sendfileInjectionNoncompliant() {
5    app.get('www.example.com', (req, res) => {
6        var fileName = req.params.file
7        // Noncompliant: tainted-data is passed into 'res.sendfile'.
8        res.sendFile(fileName)
9    })
10}

```

#### Compliant example

```javascript
1var express = require("express")
2var path = require("path")
3var app = express()
4function sendfileInjectionCompliant() {
5    app.get('www.example.com', (req, res) => {
6        var fileName = "file.txt"
7        if (fileName !== req.params.file) {
8            // Compliant: validated fileName before passing into 'res.sendFile'.
9            res.sendFile(fileName)
10            console.log("Valid file name.")
11        } else {
12            throw new Error("Invalid file name.")
13        }
14    })
15}

```
