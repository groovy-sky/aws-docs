![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### TypeScript detectors(104/104)

[Integer overflow](../integer-overflow.md) [AWS insecure transmission CDK](../aws-insecure-transmission-cdk.md) [Insecure cookie](../insecure-cookie.md) [AWS credentials logged](../aws-logged-credentials.md) [SQL injection](../sql-injection.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [New function detected](../new-function-detected.md) [Batch request with unchecked failures](../aws-batch-write-output-ignored.md) [Unvalidated expansion of archive files](../do-not-expand-archive-files-without-validating.md) [Missing pagination](../missing-pagination.md) [XPath injection](../xpath-injection.md) [Logging of sensitive information](../logging-of-sensitive-information.md) [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md) [Insecure CORS policy](../insecure-cors-policy.md) [Improper input validation](../improper-input-validation.md) [Hardcoded credentials](../hardcoded-credentials.md) [Invoke super appropriately](../invoke-super-appropriately.md) [Numeric truncation error](../numeric-truncation-error.md) [Avoid nan in comparison](../avoid-nan-in-comparison.md) [Missing check on method output](../missing-check-on-method-output.md) [Insufficiently protected credentials](../insufficiently-protected-credentials.md) [Improper restriction of rendered UI layers or frames](../improper-restriction-of-frames.md) [Pseudorandom number generators](../pseudorandom-number-generator.md) [AWS missing encryption CDK](../aws-missing-encryption-cdk.md) [Catch and swallow exception](../swallow-exceptions.md) [Header injection](../header-injection.md) [Use {} instead of new Object()](../literal-instead-of-new-object.md) [Hardcoded IP address](../hardcoded-ip-address.md) [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md) [Weak obfuscation of web requests](../weak-obfuscation-of-request.md) [File Race Bad](../file-race-bad.md) [Improper certificate validation](../improper-certificate-validation.md) [Sendfile injection](../sendfile-injection.md) [Cryptographic key generator](../cryptographic-key-generator.md) [improper input validation cdk](../improper-input-validation-cdk.md) [XML external entity](../xml-external-entity.md) [Timing attack](../timing-attack.md) [Missing Amazon S3 bucket owner condition](../s3-verify-bucket-owner.md) [Insecure hashing](../insecure-hashing.md) [Session fixation](../session-fixation.md) [Use of Default Credentials CDK](../use-of-default-credentials-cdk.md) [File injection](../file-injection.md) [Improper Restriction of Operations within the Bounds of a Memory Buffer](../improper-restriction-of-operations.md) [Limit request length](../limit-on-request-content-length.md) [Log injection](../log-injection.md) [Type confusion](../type-confusion.md) [Server side request forgery](../server-side-request-forgery.md) [aws kmskey encryption cdk](../aws-kmskey-encryption-cdk.md) [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md) [Avoid Undefined As Variable Name](../avoid-undefined-as-variable-name.md) [Index of method comparison](../index-of-method-comparison.md) [String passed to \`setInterval\` or \`setTimeout\`](../do-not-pass-string-to-setinterval-or-settimeout.md) [Data loss in a batch request](../unchecked-batch-failures.md) [Tainted input for Docker API](../tainted-input-for-docker-api.md) [Insecure temporary file or directory](../insecure-temp-file.md) [Check failed records when using kinesis](../kinesis-failed-record-check.md) [AWS api logging disabled cdk](../api-logging-disabled-cdk.md) [Improper Access Control CDK](../improper-access-control-cdk.md) [Least privilege violation](../least-privilege-violation.md) [File extension validation](../file-extension-validation.md) [Resource leak](../resource-leak.md) [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md) [Insecure object attribute modification](../insecure-object-attribute-modification.md) [Missing Authentication for Critical Function CDK](../missing-authentication-for-critical-function-cdk.md) [Typeof expression](../typeof-expression.md) [AWS missing encryption of sensitive data cdk](../missing-encryption-of-sensitive-data-cdk.md) [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Client-side KMS reencryption](../aws-kms-reencryption.md) [Insecure cryptography](../insecure-cryptography.md) [Deserialization of untrusted object](../untrusted-deserialization.md) [Clear text credentials](../clear-text-credentials.md) [Improper handling of case sensitivity](../improper-handling-of-case-sensitivity.md) [Unsanitized input is run as code](../code-injection.md) [Protection mechanism failure](../protection-mechanism-failure.md) [Use of a deprecated method](../deprecated-method.md) [Sensitive information leak](../sensitive-information-leak.md) [DNS prefetching](../dns-prefetching.md) [Exposure of Sensitive Information CDK](../exposure-of-sensitive-information-cdk.md) [Insecure JWT parsing](../insecure-jwt-parsing.md) [Loose file permissions](../loose-file-permissions.md) [Missing Authorization CDK](../missing-authorization-cdk.md) [Sensitive query string](../sensitive-query-string.md) [Insufficient Logging CDK](../insufficient-logging-cdk.md) [OS command injection](../os-command-injection.md) [NoSQL injection](../nosql-injection.md) [Unverified hostname](../unverified-hostname.md) [Path traversal](../path-traversal.md) [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md) [LDAP injection](../ldap-injection.md) [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md) [S3 partial encrypt CDK](../s3-partial-encrypt-cdk.md) [Non-literal regular expression](../non-literal-regular-expression.md) [Stack trace exposure](../stack-trace-exposure.md) [File and directory information exposure](../file-and-directory-information-exposure.md) [Usage of an API that is not recommended](../not-recommended-apis.md) [Lazy Load Module](../lazy-load-module.md) [Disabled HTML autoescape](../do-not-disable-html-autoescape.md) [Improper access control](../improper-access-control.md) [Cross-site scripting](../cross-site-scripting.md) [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md) [AWS client not reused in a Lambda function](../lambda-client-reuse.md) [URL redirection to untrusted site](../open-redirect.md) [Untrusted data in security decision](../untrusted-data-in-decision.md)

# Security detectors

### [Integer overflow](../integer-overflow.md)

An integer overflow might cause security issues when it is used for resource management or execution control.

### [AWS insecure transmission CDK](../aws-insecure-transmission-cdk.md)

The product transmits sensitive or security-critical data in cleartext in a communication channel that can be sniffed by unauthorized actors.

### [Insecure cookie](../insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [AWS credentials logged](../aws-logged-credentials.md)

Logging unencrypted AWS credentials can expose them to an attacker.

### [SQL injection](../sql-injection.md)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Insecure connection using unencrypted protocol](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Batch request with unchecked failures](../aws-batch-write-output-ignored.md)

Unchecked failures can lead to data loss.

### [Unvalidated expansion of archive files](../do-not-expand-archive-files-without-validating.md)

Expanding unverified archive files without controlling the size of the expanded data can lead to zip bomb attacks.

### [Missing pagination](../missing-pagination.md)

Missing pagination on a paginated call can lead to inaccurate results.

### [XPath injection](../xpath-injection.md)

Potentially unsanitized user input in XPath queries can allow an attacker to control the query in unwanted or insecure ways.

### [Logging of sensitive information](../logging-of-sensitive-information.md)

The logging of sensitive information can expose the information to potential attackers.

### [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md)

Overriding environment variables that are reserved by AWS Lambda might lead to unexpected behavior.

### [Insecure CORS policy](../insecure-cors-policy.md)

Cross-origin resource sharing policies that are too permissive could lead to security vulnerabilities.

### [Improper input validation](../improper-input-validation.md)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Hardcoded credentials](../hardcoded-credentials.md)

Hardcoded credentials can be intercepted by malicious actors.

### [Insufficiently protected credentials](../insufficiently-protected-credentials.md)

An object attribute constructed from a user-provided input should not be passed directly to a method.

### [Improper restriction of rendered UI layers or frames](../improper-restriction-of-frames.md)

The application incorrectly restricts frame objects or UI layers that belong to another application or domain.

### [Pseudorandom number generators](../pseudorandom-number-generator.md)

Using pseudorandom number generators (PRNGs) is security-sensitive.

### [AWS missing encryption CDK](../aws-missing-encryption-cdk.md)

The AWS resource is missing appropriate encryption.

### [Header injection](../header-injection.md)

Constructing HTTP response headers from user-controlled data is unsafe.

### [Hardcoded IP address](../hardcoded-ip-address.md)

Hardcoding an IP address can cause security problems.

### [Untrusted Amazon Machine Images](../untrusted-amazon-machine-images.md)

Improper filtering of Amazon Machine Images (AMIs) can result in loading an untrusted image, which is a potential security vulnerability.

### [Weak obfuscation of web requests](../weak-obfuscation-of-request.md)

Weak obfuscation of web requests makes your application vulnerable.

### [File Race Bad](../file-race-bad.md)

File name should not be used to do the file operations.

### [Improper certificate validation](../improper-certificate-validation.md)

Lack of validation of a security certificate can lead to host impersonation and sensitive data leaks.

### [Sendfile injection](../sendfile-injection.md)

The software allows user input to control or influence paths or file names that are used in file system operations.

### [Cryptographic key generator](../cryptographic-key-generator.md)

Insufficient key sizes can lead to brute force attacks.

### [XML external entity](../xml-external-entity.md)

Objects that parse or handle XML can lead to XML external entity (XXE) attacks when they are misconfigured.

### [Timing attack](../timing-attack.md)

Insecure string comparison can lead to a timing-attack.

### [Missing Amazon S3 bucket owner condition](../s3-verify-bucket-owner.md)

Not setting the Amazon S3 bucket owner condition could lead to accidentally using the wrong bucket.

### [Insecure hashing](../insecure-hashing.md)

Obsolete, broken, or weak hashing algorithms can lead to security vulnerabilities.

### [Session fixation](../session-fixation.md)

Session fixation might allow an attacker to steal authenticated session IDs.

### [File injection](../file-injection.md)

Writing unsanitized user data to a file is unsafe.

### [Limit request length](../limit-on-request-content-length.md)

Significant content length can lead to denial of service.

### [Log injection](../log-injection.md)

Using untrusted inputs in a log statement can enable attackers to break the log's format, forge log entries, and bypass log monitors.

### [Type confusion](../type-confusion.md)

Type confusion occurs when an application accesses a resource using a type that is incompatible with its original type.

### [Server side request forgery](../server-side-request-forgery.md)

Insufficient sanitization of potentially untrusted URLs on the server side can allow server requests to unwanted destinations.

### [aws kmskey encryption cdk](../aws-kmskey-encryption-cdk.md)

Using an AWS KMS key helps follow the standard security advice of granting least privilege to objects generated by the project.

### [Data loss in a batch request](../unchecked-batch-failures.md)

A batch request that doesn't check for failed items can lead to loss of data.

### [Tainted input for Docker API](../tainted-input-for-docker-api.md)

Passing an unsanitized user argument to a function call makes your code insecure.

### [Insecure temporary file or directory](../insecure-temp-file.md)

Insecure ways of creating temporary files and directories can lead to race conditions, privilege escalation, and other security vulnerabilities.

### [AWS api logging disabled cdk](../api-logging-disabled-cdk.md)

Api Logging Disabled may lead to unable to access log and does not record the event.

### [Least privilege violation](../least-privilege-violation.md)

The elevated privilege level should be dropped immediately after the operation is performed.

### [File extension validation](../file-extension-validation.md)

Checks if the extension of a file uploaded by a user is validated before the file is saved.

### [Resource leak](../resource-leak.md)

Allocated resources are not released properly.

### [Set SNS Return Subscription ARN](../sns-set-return-subscription-arn.md)

To always return the subscription ARN, set the `ReturnSubscriptionArn` argument to `True`.

### [Insecure object attribute modification](../insecure-object-attribute-modification.md)

Updating object attributes obtained from external sources is security sensitive.

### [AWS missing encryption of sensitive data cdk](../missing-encryption-of-sensitive-data-cdk.md)

Sensitive or critical information is not encrypted before storage or transmission in the product.

### [Unauthenticated Amazon SNS unsubscribe requests might succeed](../sns-authenticate-on-unsubscribe.md)

Failing to set the `AuthenticateOnUnsubscribe` flag to `True` when confirming an SNS subscription can lead to unauthenticated cancellations.

### [Cross-site request forgery](../cross-site-request-forgery.md)

Insecure configuration can lead to a cross-site request forgery (CRSF) vulnerability.

### [Client-side KMS reencryption](../aws-kms-reencryption.md)

Client-side decryption followed by encryption is inefficient and can lead to sensitive data leaks.

### [Insecure cryptography](../insecure-cryptography.md)

Weak, broken, or misconfigured cryptography can lead to security vulnerabilities.

### [Deserialization of untrusted object](../untrusted-deserialization.md)

Deserialization of untrusted objects can lead to security vulnerabilities such as, inadvertently running remote code.

### [Clear text credentials](../clear-text-credentials.md)

Credentials that are stored in clear text can be intercepted by a malicious actor.

### [Improper handling of case sensitivity](../improper-handling-of-case-sensitivity.md)

Improper handling of case sensitivity when accessing a resource can lead to inconsistent results.

### [Unsanitized input is run as code](../code-injection.md)

Scripts generated from unsanitized inputs can lead to malicious behavior and inadvertently running code remotely.

### [Protection mechanism failure](../protection-mechanism-failure.md)

Disabled or incorrectly used protection mechanism can lead to security vulnerabilities.

### [Use of a deprecated method](../deprecated-method.md)

This code uses deprecated methods, which suggests that it has not been recently reviewed or maintained.

### [Sensitive information leak](../sensitive-information-leak.md)

Exposure of sensitive information can lead to an unauthorized actor having access to the information.

### [DNS prefetching](../dns-prefetching.md)

DNS prefetching can cause latency and privacy issues.

### [Insecure JWT parsing](../insecure-jwt-parsing.md)

Checks if the `none` algorithm is used in a `JWT token` parsing.

### [Loose file permissions](../loose-file-permissions.md)

Weak file permissions can lead to privilege escalation.

### [Missing Authorization CDK](../missing-authorization-cdk.md)

Improper Access Control.

### [Sensitive query string](../sensitive-query-string.md)

Do not fetch sensitive information from a GET request.

### [OS command injection](../os-command-injection.md)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code.

### [NoSQL injection](../nosql-injection.md)

User input can be vulnerable to injection attacks.

### [Unverified hostname](../unverified-hostname.md)

Unverified hostnames lead to security vulnerabilities.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Origins-verified cross-origin communications](../origins-verified-cross-origin-communications.md)

Unverified origins of messages and identities in cross-origin communications can lead to security vulnerabilities.

### [LDAP injection](../ldap-injection.md)

LDAP queries that rely on potentially untrusted inputs can allow attackers to read or modify sensitive data, run code, and perform other unwanted actions.

### [SNS don't bind subscribe and publish](../sns-no-bind-subscribe-publish.md)

Do not bind the SNS Publish operation with the SNS Subscribe or Create Topic operation.

### [S3 partial encrypt CDK](../s3-partial-encrypt-cdk.md)

An unencrypted bucket could lead to sensitive data exposure.

### [Non-literal regular expression](../non-literal-regular-expression.md)

Non-literal input to a regular expression might lead to a denial of service attack.

### [Stack trace exposure](../stack-trace-exposure.md)

Stack traces can be hard to use for debugging.

### [File and directory information exposure](../file-and-directory-information-exposure.md)

Allowing hidden files while serving files from a given root directory can cause information leakage.

### [Usage of an API that is not recommended](../not-recommended-apis.md)

APIs that are not recommended were found.

### [Disabled HTML autoescape](../do-not-disable-html-autoescape.md)

Disabling the HTML autoescape mechanism exposes your web applications to attacks.

### [Improper access control](../improper-access-control.md)

The software does not restrict or incorrectly restrict access to a resource from an unauthorized actor.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md)

Encryption that is dependent on conditional logic, such as an `if...then` clause, might cause unencrypted sensitive data to be stored.

### [AWS client not reused in a Lambda function](../lambda-client-reuse.md)

Recreating AWS clients in each Lambda function invocation is expensive.

### [URL redirection to untrusted site](../open-redirect.md)

User-controlled input that specifies a link to an external site could lead to phishing attacks and allow user credentials to be stolen.

### [Untrusted data in security decision](../untrusted-data-in-decision.md)

Security decisions should not depend on branching that can be influenced by untrusted or client-provided data.
