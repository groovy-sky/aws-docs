![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JavaScript detectors(78/78)

[Improper access control](improper-access-control.md) [Sensitive data stored unencrypted due to partial encryption](partial-encryption.md) [Pseudorandom number generators](pseudorandom-number-generator.md) [OS command injection](os-command-injection.md) [URL redirection to untrusted site](open-redirect.md) [Integer overflow](integer-overflow.md) [Protection mechanism failure](protection-mechanism-failure.md) [Non-literal regular expression](non-literal-regular-expression.md) [Tainted input for Docker API](tainted-input-for-docker-api.md) [Usage of an API that is not recommended](not-recommended-apis.md) [XML external entity](xml-external-entity.md) [Server-side request forgery](server-side-request-forgery.md) [New function detected](new-function-detected.md) [Stack trace exposure](stack-trace-exposure.md) [Timing attack](timing-attack.md) [SNS don't bind subscribe and publish](sns-no-bind-subscribe-publish.md) [Invoke super appropriately](invoke-super-appropriately.md) [NoSQL injection](nosql-injection.md) [Hardcoded credentials](hardcoded-credentials.md) [Insecure cookie](insecure-cookie.md) [Cross-site scripting](cross-site-scripting.md) [Hardcoded IP address](hardcoded-ip-address.md) [AWS credentials logged](aws-logged-credentials.md) [XPath injection](xpath-injection.md) [Data loss in a batch request](unchecked-batch-failures.md) [Path traversal](path-traversal.md) [Least privilege violation](least-privilege-violation.md) [DNS prefetching](dns-prefetching.md) [Resource leak](resource-leak.md) [Insufficiently protected credentials](insufficiently-protected-credentials.md) [File extension validation](file-extension-validation.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Cross-site request forgery](cross-site-request-forgery.md) [Typeof expression](typeof-expression.md) [Set SNS Return Subscription ARN](sns-set-return-subscription-arn.md) [File and directory information exposure](file-and-directory-information-exposure.md) [Missing Amazon S3 bucket owner condition](s3-verify-bucket-owner.md) [Insecure hashing](insecure-hashing.md) [Numeric truncation error](numeric-truncation-error.md) [Client-side KMS reencryption](aws-kms-reencryption.md) [AWS client not reused in a Lambda function](lambda-client-reuse.md) [LDAP injection](ldap-injection.md) [Batch request with unchecked failures](aws-batch-write-output-ignored.md) [Cryptographic key generator](cryptographic-key-generator.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](sns-authenticate-on-unsubscribe.md) [Unverified hostname](unverified-hostname.md) [Origins-verified cross-origin communications](origins-verified-cross-origin-communications.md) [Loose file permissions](loose-file-permissions.md) [Unsanitized input is run as code](code-injection.md) [Missing pagination](missing-pagination.md) [Untrusted Amazon Machine Images](untrusted-amazon-machine-images.md) [Improper certificate validation](improper-certificate-validation.md) [Insecure CORS policy](insecure-cors-policy.md) [Deserialization of untrusted object](untrusted-deserialization.md) [Sensitive information leak](sensitive-information-leak.md) [Check failed records when using kinesis](kinesis-failed-record-check.md) [Weak obfuscation of web requests](weak-obfuscation-of-request.md) [Catch and swallow exception](swallow-exceptions.md) [Logging of sensitive information](logging-of-sensitive-information.md) [Limit request length](limit-on-request-content-length.md) [String passed to \`setInterval\` or \`setTimeout\`](do-not-pass-string-to-setinterval-or-settimeout.md) [Log injection](log-injection.md) [Override of reserved variable names in a Lambda function](lambda-override-reserved.md) [Improper restriction of rendered UI layers or frames](improper-restriction-of-frames.md) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-cryptography) [Insecure object attribute modification](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-object-attribute-modification) [Session fixation](https://docs.aws.amazon.com/amazonq/detector-library/javascript/session-fixation) [Avoid nan in comparison](https://docs.aws.amazon.com/amazonq/detector-library/javascript/avoid-nan-in-comparison) [Improper input validation](https://docs.aws.amazon.com/amazonq/detector-library/javascript/improper-input-validation) [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/javascript/do-not-disable-html-autoescape) [Use of a deprecated method](https://docs.aws.amazon.com/amazonq/detector-library/javascript/deprecated-method) [Unvalidated expansion of archive files](https://docs.aws.amazon.com/amazonq/detector-library/javascript/do-not-expand-archive-files-without-validating) [File injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/file-injection) [Sendfile injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sendfile-injection) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/sql-injection) [Header injection](https://docs.aws.amazon.com/amazonq/detector-library/javascript/header-injection) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/javascript/insecure-temp-file) [Inefficient polling of AWS resource](https://docs.aws.amazon.com/amazonq/detector-library/javascript/aws-polling-instead-of-waiter)

# Insecure cryptography [Critical](https://docs.aws.amazon.com/amazonq/detector-library/javascript/severity/critical)

Misuse of cryptography-related APIs can create security vulnerabilities. This includes algorithms with known weaknesses, certain padding modes, the lack of integrity checks, insufficiently large key sizes, and insecure combinations of the previous items.

**Detector ID**

javascript/insecure-cryptography@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/javascript/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-327](https://cwe.mitre.org/data/definitions/327.html) [CWE-347](https://cwe.mitre.org/data/definitions/347.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/cryptography) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/javascript/tags/owasp-top10)

* * *

#### Noncompliant example

```javascript
1function insecureCryptographyNoncompliant()
2{
3    var ciphers = [`TLS_DH_anon_WITH_AES_256_GCM_SHA384`,
4        `TLS_AES_128_GCM_SHA256`,
5        `ECDHE-ECDSA-AES128-GCM-SHA256`].join(':')
6    var options = {
7        hostname: 'www.example.com',
8        port: 443,
9        path: '/',
10        method: 'GET',
11        secureProtocol: 'TLSv1_2_method',
12        // Noncompliant: insecure TLS cipher suite is used.
13        ciphers:ciphers
14    }
15
16    var req = https.request(options, (res) => {
17        res.on('data', (d) => {
18            process.stdout.write(d)
19        })
20    })
21}

```

#### Compliant example

```javascript
1function insecureCryptographyCompliant()
2{
3    var ciphers = [`TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256` ,
4        `TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`,
5        `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`,
6        `TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`,
7        `TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256`,
8        `TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256`,
9        `TLS_AES_128_GCM_SHA256`,
10        `TLS_AES_256_GCM_SHA384`,
11        '!aNULL',
12        '!eNULL',
13        '!NULL',
14        '!DES',
15        '!RC4',
16        '!MD5'].join(':')
17    var options = {
18        hostname: 'www.example.com',
19        port: 443,
20        path: '/',
21        method: 'GET',
22        secureProtocol: 'TLSv1_2_method',
23        // Compliant: secure TLS cipher suite is used.
24        ciphers:ciphers
25    }
26
27    var req = https.request(options, (res) => {
28        res.on('data', (d) => {
29            process.stdout.write(d)
30        })
31    })
32}

```
