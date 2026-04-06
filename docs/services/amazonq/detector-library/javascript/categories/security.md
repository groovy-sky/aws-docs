![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### JavaScript detectors(78/78)

[Improper access control](../improper-access-control.md) [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md) [Pseudorandom number generators](../pseudorandom-number-generator.md) [OS command injection](../os-command-injection.md) [URL redirection to untrusted site](../open-redirect.md) [Integer overflow](../integer-overflow.md) [Protection mechanism failure](../protection-mechanism-failure.md) [Non-literal regular expression](../non-literal-regular-expression.md) [Tainted input for Docker API](../tainted-input-for-docker-api.md) [Usage of an API that is not recommended](../not-recommended-apis.md) [XML external entity](../xml-external-entity.md) [Server-side request forgery](../server-side-request-forgery.md) [New function detected](../new-function-detected.md) [Stack trace exposure](../stack-trace-exposure.md) [Timing attack](../timing-attack.md) [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md) [Invoke super appropriately](../invoke-super-appropriately.md) [NoSQL injection](../nosql-injection.md) [Hardcoded credentials](../hardcoded-credentials.md) [Insecure cookie](../insecure-cookie.md) [Cross-site scripting](../cross-site-scripting.md) [Hardcoded IP address](../hardcoded-ip-address.md) [AWS credentials logged](../aws-logged-credentials.md) [XPath injection](../xpath-injection.md) [Data loss in a batch request](../unchecked-batch-failures.md) [Path traversal](../path-traversal.md) [Least privilege violation](../least-privilege-violation.md) [DNS prefetching](../dns-prefetching.md) [Resource leak](../resource-leak.md) [Insufficiently protected credentials](../insufficiently-protected-credentials.md) [File extension validation](../file-extension-validation.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Typeof expression](../typeof-expression.md) [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md) [File and directory information exposure](../file-and-directory-information-exposure.md) [Missing Amazon S3 bucket owner condition](../s3-verify-bucket-owner.md) [Insecure hashing](../insecure-hashing.md) [Numeric truncation error](../numeric-truncation-error.md) [Client-side KMS reencryption](../aws-kms-reencryption.md) [AWS client not reused in a Lambda function](../lambda-client-reuse.md) [LDAP injection](../ldap-injection.md) [Batch request with unchecked failures](../aws-batch-write-output-ignored.md) [Cryptographic key generator](../cryptographic-key-generator.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md) [Unverified hostname](../unverified-hostname.md) [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md) [Loose file permissions](../loose-file-permissions.md) [Unsanitized input is run as code](../code-injection.md) [Missing pagination](../missing-pagination.md) [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md) [Improper certificate validation](../improper-certificate-validation.md) [Insecure CORS policy](../insecure-cors-policy.md) [Deserialization of untrusted object](../untrusted-deserialization.md) [Sensitive information leak](../sensitive-information-leak.md) [Check failed records when using kinesis](../kinesis-failed-record-check.md) [Weak obfuscation of web requests](../weak-obfuscation-of-request.md) [Catch and swallow exception](../swallow-exceptions.md) [Logging of sensitive information](../logging-of-sensitive-information.md) [Limit request length](../limit-on-request-content-length.md) [String passed to \`setInterval\` or \`setTimeout\`](../do-not-pass-string-to-setinterval-or-settimeout.md) [Log injection](../log-injection.md) [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md) [Improper restriction of rendered UI layers or frames](../improper-restriction-of-frames.md) [Insecure cryptography](../insecure-cryptography.md) [Insecure object attribute modification](../insecure-object-attribute-modification.md) [Session fixation](../session-fixation.md) [Avoid nan in comparison](../avoid-nan-in-comparison.md) [Improper input validation](../improper-input-validation.md) [Disabled HTML autoescape](../do-not-disable-html-autoescape.md) [Use of a deprecated method](../deprecated-method.md) [Unvalidated expansion of archive files](../do-not-expand-archive-files-without-validating.md) [File injection](../file-injection.md) [Sendfile injection](../sendfile-injection.md) [SQL injection](../sql-injection.md) [Header injection](../header-injection.md) [Insecure temporary file or directory](../insecure-temp-file.md) [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md)

# Security detectors

### [Improper access control](../improper-access-control.md)

The software does not restrict or incorrectly restrict access to a resource from an unauthorized actor.

### [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md)

Encryption that is dependent on conditional logic, such as an `if...then` clause, might cause unencrypted sensitive data to be stored.

### [Pseudorandom number generators](../pseudorandom-number-generator.md)

Using pseudorandom number generators (PRNGs) is security-sensitive.

### [OS command injection](../os-command-injection.md)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code.

### [URL redirection to untrusted site](../open-redirect.md)

User-controlled input that specifies a link to an external site could lead to phishing attacks and allow user credentials to be stolen.

### [Integer overflow](../integer-overflow.md)

An integer overflow might cause security issues when it is used for resource management or execution control.

### [Protection mechanism failure](../protection-mechanism-failure.md)

Disabled or incorrectly used protection mechanism can lead to security vulnerabilities.

### [Non-literal regular expression](../non-literal-regular-expression.md)

Non-literal input to a regular expression might lead to a denial of service attack.

### [Tainted input for Docker API](../tainted-input-for-docker-api.md)

Passing an unsanitized user argument to a function call makes your code insecure.

### [Usage of an API that is not recommended](../not-recommended-apis.md)

APIs that are not recommended were found.

### [XML external entity](../xml-external-entity.md)

Objects that parse or handle XML can lead to XML external entity (XXE) attacks when they are misconfigured.

### [Server-side request forgery](../server-side-request-forgery.md)

Insufficient sanitization of potentially untrusted URLs on the server side can allow server requests to unwanted destinations.

### [Stack trace exposure](../stack-trace-exposure.md)

Stack traces can be hard to use for debugging.

### [Timing attack](../timing-attack.md)

Insecure string comparison can lead to a timing-attack.

### [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md)

Do not bind the SNS Publish operation with the SNS Subscribe or Create Topic operation.

### [NoSQL injection](../nosql-injection.md)

User input can be vulnerable to injection attacks.

### [Hardcoded credentials](../hardcoded-credentials.md)

Hardcoded credentials can be intercepted by malicious actors.

### [Insecure cookie](../insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Hardcoded IP address](../hardcoded-ip-address.md)

Hardcoding an IP address can cause security problems.

### [AWS credentials logged](../aws-logged-credentials.md)

Logging unencrypted AWS credentials can expose them to an attacker.

### [XPath injection](../xpath-injection.md)

Potentially unsanitized user input in XPath queries can allow an attacker to control the query in unwanted or insecure ways.

### [Data loss in a batch request](../unchecked-batch-failures.md)

A batch request that doesn't check for failed items can lead to loss of data.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Least privilege violation](../least-privilege-violation.md)

The elevated privilege level should be dropped immediately after the operation is performed.

### [DNS prefetching](../dns-prefetching.md)

DNS prefetching can cause latency and privacy issues.

### [Resource leak](../resource-leak.md)

Allocated resources are not released properly.

### [Insufficiently protected credentials](../insufficiently-protected-credentials.md)

An object attribute constructed from a user-provided input should not be passed directly to a method.

### [File extension validation](../file-extension-validation.md)

Checks if the extension of a file uploaded by a user is validated before the file is saved.

### [Insecure connection using unencrypted protocol](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Cross-site request forgery](../cross-site-request-forgery.md)

Insecure configuration can lead to a cross-site request forgery (CRSF) vulnerability.

### [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md)

To always return the subscription ARN, set the `ReturnSubscriptionArn` argument to `True`.

### [File and directory information exposure](../file-and-directory-information-exposure.md)

Allowing hidden files while serving files from a given root directory can cause information leakage.

### [Missing Amazon S3 bucket owner condition](../s3-verify-bucket-owner.md)

Not setting the Amazon S3 bucket owner condition could lead to accidentally using the wrong bucket.

### [Insecure hashing](../insecure-hashing.md)

Obsolete, broken, or weak hashing algorithms can lead to security vulnerabilities.

### [Client-side KMS reencryption](../aws-kms-reencryption.md)

Client-side decryption followed by encryption is inefficient and can lead to sensitive data leaks.

### [AWS client not reused in a Lambda function](../lambda-client-reuse.md)

Recreating AWS clients in each Lambda function invocation is expensive.

### [LDAP injection](../ldap-injection.md)

LDAP queries that rely on potentially untrusted inputs can allow attackers to read or modify sensitive data, run code, and perform other unwanted actions.

### [Batch request with unchecked failures](../aws-batch-write-output-ignored.md)

Unchecked failures can lead to data loss.

### [Cryptographic key generator](../cryptographic-key-generator.md)

Insufficient key sizes can lead to brute force attacks.

### [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md)

Failing to set the `AuthenticateOnUnsubscribe` flag to `True` when confirming an SNS subscription can lead to unauthenticated cancellations.

### [Unverified hostname](../unverified-hostname.md)

Unverified hostnames lead to security vulnerabilities.

### [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md)

Unverified origins of messages and identities in cross-origin communications can lead to security vulnerabilities.

### [Loose file permissions](../loose-file-permissions.md)

Weak file permissions can lead to privilege escalation.

### [Unsanitized input is run as code](../code-injection.md)

Scripts generated from unsanitized inputs can lead to malicious behavior and inadvertently running code remotely.

### [Missing pagination](../missing-pagination.md)

Missing pagination on a paginated call can lead to inaccurate results.

### [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md)

Improper filtering of Amazon Machine Images (AMIs) can result in loading an untrusted image, which is a potential security vulnerability.

### [Improper certificate validation](../improper-certificate-validation.md)

Lack of validation of a security certificate can lead to host impersonation and sensitive data leaks.

### [Insecure CORS policy](../insecure-cors-policy.md)

Cross-origin resource sharing policies that are too permissive could lead to security vulnerabilities.

### [Deserialization of untrusted object](../untrusted-deserialization.md)

Deserialization of untrusted objects can lead to security vulnerabilities such as, inadvertently running remote code.

### [Sensitive information leak](../sensitive-information-leak.md)

Exposure of sensitive information can lead to an unauthorized actor having access to the information.

### [Weak obfuscation of web requests](../weak-obfuscation-of-request.md)

Weak obfuscation of web requests makes your application vulnerable.

### [Logging of sensitive information](../logging-of-sensitive-information.md)

The logging of sensitive information can expose the information to potential attackers.

### [Limit request length](../limit-on-request-content-length.md)

Significant content length can lead to denial of service.

### [Log injection](../log-injection.md)

Using untrusted inputs in a log statement can enable attackers to break the log's format, forge log entries, and bypass log monitors.

### [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md)

Overriding environment variables that are reserved by AWS Lambda might lead to unexpected behavior.

### [Improper restriction of rendered UI layers or frames](../improper-restriction-of-frames.md)

The application incorrectly restricts frame objects or UI layers that belong to another application or domain.

### [Insecure cryptography](../insecure-cryptography.md)

Weak, broken, or misconfigured cryptography can lead to security vulnerabilities.

### [Insecure object attribute modification](../insecure-object-attribute-modification.md)

Updating object attributes obtained from external sources is security sensitive.

### [Session fixation](../session-fixation.md)

Session fixation might allow an attacker to steal authenticated session IDs.

### [Improper input validation](../improper-input-validation.md)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Disabled HTML autoescape](../do-not-disable-html-autoescape.md)

Disabling the HTML autoescape mechanism exposes your web applications to attacks.

### [Use of a deprecated method](../deprecated-method.md)

This code uses deprecated methods, which suggests that it has not been recently reviewed or maintained.

### [Unvalidated expansion of archive files](../do-not-expand-archive-files-without-validating.md)

Expanding unverified archive files without controlling the size of the expanded data can lead to zip bomb attacks.

### [File injection](../file-injection.md)

Writing unsanitized user data to a file is unsafe.

### [Sendfile injection](../sendfile-injection.md)

The software allows user input to control or influence paths or file names that are used in file system operations.

### [SQL injection](../sql-injection.md)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Header injection](../header-injection.md)

Constructing HTTP response headers from user-controlled data is unsafe.

### [Insecure temporary file or directory](../insecure-temp-file.md)

Insecure ways of creating temporary files and directories can lead to race conditions, privilege escalation, and other security vulnerabilities.
