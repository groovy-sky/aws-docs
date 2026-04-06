![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JavaScript detectors(78/78)

[Improper access control](../improper-access-control.md) [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md) [Pseudorandom number generators](../pseudorandom-number-generator.md) [OS command injection](../os-command-injection.md) [URL redirection to untrusted site](../open-redirect.md) [Integer overflow](../integer-overflow.md) [Protection mechanism failure](../protection-mechanism-failure.md) [Non-literal regular expression](../non-literal-regular-expression.md) [Tainted input for Docker API](../tainted-input-for-docker-api.md) [Usage of an API that is not recommended](../not-recommended-apis.md) [XML external entity](../xml-external-entity.md) [Server-side request forgery](../server-side-request-forgery.md) [New function detected](../new-function-detected.md) [Stack trace exposure](../stack-trace-exposure.md) [Timing attack](../timing-attack.md) [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md) [Invoke super appropriately](../invoke-super-appropriately.md) [NoSQL injection](../nosql-injection.md) [Hardcoded credentials](../hardcoded-credentials.md) [Insecure cookie](../insecure-cookie.md) [Cross-site scripting](../cross-site-scripting.md) [Hardcoded IP address](../hardcoded-ip-address.md) [AWS credentials logged](../aws-logged-credentials.md) [XPath injection](../xpath-injection.md) [Data loss in a batch request](../unchecked-batch-failures.md) [Path traversal](../path-traversal.md) [Least privilege violation](../least-privilege-violation.md) [DNS prefetching](../dns-prefetching.md) [Resource leak](../resource-leak.md) [Insufficiently protected credentials](../insufficiently-protected-credentials.md) [File extension validation](../file-extension-validation.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Typeof expression](../typeof-expression.md) [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md) [File and directory information exposure](../file-and-directory-information-exposure.md) [Missing Amazon S3 bucket owner condition](../s3-verify-bucket-owner.md) [Insecure hashing](../insecure-hashing.md) [Numeric truncation error](../numeric-truncation-error.md) [Client-side KMS reencryption](../aws-kms-reencryption.md) [AWS client not reused in a Lambda function](../lambda-client-reuse.md) [LDAP injection](../ldap-injection.md) [Batch request with unchecked failures](../aws-batch-write-output-ignored.md) [Cryptographic key generator](../cryptographic-key-generator.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md) [Unverified hostname](../unverified-hostname.md) [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md) [Loose file permissions](../loose-file-permissions.md) [Unsanitized input is run as code](../code-injection.md) [Missing pagination](../missing-pagination.md) [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md) [Improper certificate validation](../improper-certificate-validation.md) [Insecure CORS policy](../insecure-cors-policy.md) [Deserialization of untrusted object](../untrusted-deserialization.md) [Sensitive information leak](../sensitive-information-leak.md) [Check failed records when using kinesis](../kinesis-failed-record-check.md) [Weak obfuscation of web requests](../weak-obfuscation-of-request.md) [Catch and swallow exception](../swallow-exceptions.md) [Logging of sensitive information](../logging-of-sensitive-information.md) [Limit request length](../limit-on-request-content-length.md) [String passed to \`setInterval\` or \`setTimeout\`](../do-not-pass-string-to-setinterval-or-settimeout.md) [Log injection](../log-injection.md) [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md) [Improper restriction of rendered UI layers or frames](../improper-restriction-of-frames.md) [Insecure cryptography](../insecure-cryptography.md) [Insecure object attribute modification](../insecure-object-attribute-modification.md) [Session fixation](../session-fixation.md) [Avoid nan in comparison](../avoid-nan-in-comparison.md) [Improper input validation](../improper-input-validation.md) [Disabled HTML autoescape](../do-not-disable-html-autoescape.md) [Use of a deprecated method](../deprecated-method.md) [Unvalidated expansion of archive files](../do-not-expand-archive-files-without-validating.md) [File injection](../file-injection.md) [Sendfile injection](../sendfile-injection.md) [SQL injection](../sql-injection.md) [Header injection](../header-injection.md) [Insecure temporary file or directory](../insecure-temp-file.md) [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md)

# Tag: injection

### [OS command injection](../os-command-injection.md)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code.

### [XML external entity](../xml-external-entity.md)

Objects that parse or handle XML can lead to XML external entity (XXE) attacks when they are misconfigured.

### [Server-side request forgery](../server-side-request-forgery.md)

Insufficient sanitization of potentially untrusted URLs on the server side can allow server requests to unwanted destinations.

### [New function detected](../new-function-detected.md)

Use of `new Function()` can be dangerous if used to evaluate dynamic content.

### [NoSQL injection](../nosql-injection.md)

User input can be vulnerable to injection attacks.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [XPath injection](../xpath-injection.md)

Potentially unsanitized user input in XPath queries can allow an attacker to control the query in unwanted or insecure ways.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Cross-site request forgery](../cross-site-request-forgery.md)

Insecure configuration can lead to a cross-site request forgery (CRSF) vulnerability.

### [LDAP injection](../ldap-injection.md)

LDAP queries that rely on potentially untrusted inputs can allow attackers to read or modify sensitive data, run code, and perform other unwanted actions.

### [Unsanitized input is run as code](../code-injection.md)

Scripts generated from unsanitized inputs can lead to malicious behavior and inadvertently running code remotely.

### [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md)

Improper filtering of Amazon Machine Images (AMIs) can result in loading an untrusted image, which is a potential security vulnerability.

### [Deserialization of untrusted object](../untrusted-deserialization.md)

Deserialization of untrusted objects can lead to security vulnerabilities such as, inadvertently running remote code.

### [Log injection](../log-injection.md)

Using untrusted inputs in a log statement can enable attackers to break the log's format, forge log entries, and bypass log monitors.

### [Improper input validation](../improper-input-validation.md)

Improper input validation can enable attacks and lead to unwanted behavior.

### [File injection](../file-injection.md)

Writing unsanitized user data to a file is unsafe.

### [SQL injection](../sql-injection.md)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Header injection](../header-injection.md)

Constructing HTTP response headers from user-controlled data is unsafe.
