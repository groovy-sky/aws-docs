![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JSX detectors(78/78)

[Protection mechanism failure](../protection-mechanism-failure.md) [Log injection](../log-injection.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Use of a deprecated method](../deprecated-method.md) [AWS credentials logged](../aws-logged-credentials.md) [Improper input validation](../improper-input-validation.md) [Insecure cryptography](../insecure-cryptography.md) [Catch and swallow exception](../swallow-exceptions.md) [File and directory information exposure](../file-and-directory-information-exposure.md) [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md) [SQL injection](../sql-injection.md) [Non-literal regular expression](../non-literal-regular-expression.md) [Typeof expression](../typeof-expression.md) [Batch request with unchecked failures](../aws-batch-write-output-ignored.md) [Pseudorandom number generators](../pseudorandom-number-generator.md) [Cryptographic key generator](../cryptographic-key-generator.md) [Server-side request forgery](../server-side-request-forgery.md) [Sensitive information leak](../sensitive-information-leak.md) [File injection](../file-injection.md) [String passed to \`setInterval\` or \`setTimeout\`](../do-not-pass-string-to-setinterval-or-settimeout.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Usage of an API that is not recommended](../not-recommended-apis.md) [Tainted input for Docker API](../tainted-input-for-docker-api.md) [Cross-site scripting](../cross-site-scripting.md) [Weak obfuscation of web requests](../weak-obfuscation-of-request.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md) [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md) [XML external entity](../xml-external-entity.md) [Resource leak](../resource-leak.md) [Improper access control](../improper-access-control.md) [Loose file permissions](../loose-file-permissions.md) [OS command injection](../os-command-injection.md) [Client-side KMS reencryption](../aws-kms-reencryption.md) [Insecure CORS policy](../insecure-cors-policy.md) [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md) [New function detected](../new-function-detected.md) [Missing pagination](../missing-pagination.md) [Avoid nan in comparison](../avoid-nan-in-comparison.md) [Header injection](../header-injection.md) [Hardcoded credentials](../hardcoded-credentials.md) [File extension validation](../file-extension-validation.md) [NoSQL injection](../nosql-injection.md) [Missing Amazon S3 bucket owner condition](../s3-verify-bucket-owner.md) [Disabled HTML autoescape](../do-not-disable-html-autoescape.md) [Least privilege violation](../least-privilege-violation.md) [URL redirection to untrusted site](../open-redirect.md) [Insufficiently protected credentials](../insufficiently-protected-credentials.md) [Insecure hashing](../insecure-hashing.md) [Unsanitized input is run as code](../code-injection.md) [Check failed records when using kinesis](../kinesis-failed-record-check.md) [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md) [Session fixation](../session-fixation.md) [Data loss in a batch request](../unchecked-batch-failures.md) [XPath injection](../xpath-injection.md) [Deserialization of untrusted object](../untrusted-deserialization.md) [Invoke super appropriately](../invoke-super-appropriately.md) [Stack trace exposure](../stack-trace-exposure.md) [Timing attack](../timing-attack.md) [LDAP injection](../ldap-injection.md) [Insecure cookie](../insecure-cookie.md) [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md) [Unvalidated expansion of archive files](../do-not-expand-archive-files-without-validating.md) [Integer overflow](../integer-overflow.md) [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md) [Unverified hostname](../unverified-hostname.md) [Improper restriction of rendered UI layers or frames](../improper-restriction-of-frames.md) [AWS client not reused in a Lambda function](../lambda-client-reuse.md) [Path traversal](../path-traversal.md) [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md) [Insecure temporary file or directory](../insecure-temp-file.md) [Logging of sensitive information](../logging-of-sensitive-information.md) [Hardcoded IP address](../hardcoded-ip-address.md) [Insecure object attribute modification](../insecure-object-attribute-modification.md) [Numeric truncation error](../numeric-truncation-error.md) [DNS prefetching](../dns-prefetching.md) [Limit request length](../limit-on-request-content-length.md) [Sendfile injection](../sendfile-injection.md) [Improper certificate validation](../improper-certificate-validation.md)

# Medium

Showing all detectors for the JSX language withmedium severity.

### [Improper input validation](../improper-input-validation.md)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Catch and swallow exception](../swallow-exceptions.md)

Swallowing exceptions, without rethrowing or logging them, can make it hard to understand why your application is failing.

### [File and directory information exposure](../file-and-directory-information-exposure.md)

Allowing hidden files while serving files from a given root directory can cause information leakage.

### [Pseudorandom number generators](../pseudorandom-number-generator.md)

Using pseudorandom number generators (PRNGs) is security-sensitive.

### [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md)

Failing to set the `AuthenticateOnUnsubscribe` flag to `True` when confirming an SNS subscription can lead to unauthenticated cancellations.

### [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md)

To always return the subscription ARN, set the `ReturnSubscriptionArn` argument to `True`.

### [Resource leak](../resource-leak.md)

Allocated resources are not released properly.

### [Improper access control](../improper-access-control.md)

The software does not restrict or incorrectly restrict access to a resource from an unauthorized actor.

### [Insecure CORS policy](../insecure-cors-policy.md)

Cross-origin resource sharing policies that are too permissive could lead to security vulnerabilities.

### [New function detected](../new-function-detected.md)

Use of `new Function()` can be dangerous if used to evaluate dynamic content.

### [Missing pagination](../missing-pagination.md)

Missing pagination on a paginated call can lead to inaccurate results.

### [Avoid nan in comparison](../avoid-nan-in-comparison.md)

Checks if nan is used is comparison.

### [File extension validation](../file-extension-validation.md)

Checks if the extension of a file uploaded by a user is validated before the file is saved.

### [Least privilege violation](../least-privilege-violation.md)

The elevated privilege level should be dropped immediately after the operation is performed.

### [Insufficiently protected credentials](../insufficiently-protected-credentials.md)

An object attribute constructed from a user-provided input should not be passed directly to a method.

### [Insecure hashing](../insecure-hashing.md)

Obsolete, broken, or weak hashing algorithms can lead to security vulnerabilities.

### [Check failed records when using kinesis](../kinesis-failed-record-check.md)

A batch request that doesn't check for failed records can lead to loss of data.

### [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md)

Improper filtering of Amazon Machine Images (AMIs) can result in loading an untrusted image, which is a potential security vulnerability.

### [Data loss in a batch request](../unchecked-batch-failures.md)

A batch request that doesn't check for failed items can lead to loss of data.

### [Stack trace exposure](../stack-trace-exposure.md)

Stack traces can be hard to use for debugging.

### [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md)

Encryption that is dependent on conditional logic, such as an `if...then` clause, might cause unencrypted sensitive data to be stored.

### [Integer overflow](../integer-overflow.md)

An integer overflow might cause security issues when it is used for resource management or execution control.

### [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md)

Do not bind the SNS Publish operation with the SNS Subscribe or Create Topic operation.

### [AWS client not reused in a Lambda function](../lambda-client-reuse.md)

Recreating AWS clients in each Lambda function invocation is expensive.

### [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md)

Overriding environment variables that are reserved by AWS Lambda might lead to unexpected behavior.

### [Insecure temporary file or directory](../insecure-temp-file.md)

Insecure ways of creating temporary files and directories can lead to race conditions, privilege escalation, and other security vulnerabilities.

### [Hardcoded IP address](../hardcoded-ip-address.md)

Hardcoding an IP address can cause security problems.

### [Insecure object attribute modification](../insecure-object-attribute-modification.md)

Updating object attributes obtained from external sources is security sensitive.
