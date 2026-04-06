![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Java detectors(132/132)

[Reflected cross site scripting](../reflected-cross-site-scripting.md) [Mandatory method not called after object creation](../mandatory-methods.md) [Process empty record list in Amazon KCL](../kcl-with-call-process-records.md) [AWS object presence check](../object-presence.md) [Missing timeout check on CountDownLatch.await](../missing-timeout-check-on-latch-await.md) [Unspecified default value](../unspecified-default-value.md) [Device Permission Usage.](../device-permission-usage.md) [Deserialization of untrusted object](../untrusted-deserialization.md) [Preserve thread interruption status rule](../preserve-thread-interruption-status-rule.md) [Missing check on the value returned by moveToFirst API](../output-ignored-on-movetofirst-operation.md) [Missing timeout check on ExecutorService.awaitTermination](../missing-timeout-check-on-awaittermination.md) [Overflow when deserializing relational database objects](../javax-persistence-id.md) [Custom manual retries of AWS SDK calls](../aws-custom-retries.md) [Missing null check for cache response metadata](../null-check-cache-response-metadata.md) [Inefficient usage of Transaction library from AWS Labs](../dynamodb-transaction-library.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Inefficient additional authenticated data (AAD) authenticity](../cipher-update-aad.md) [Use of a deprecated method](../deprecated-method.md) [Error-prone AWS IAM policy creation](../aws-iam-error-prone-policy.md) [Use of externally-controlled input to build connection string](../connection-string-injection.md) [Inefficient Amazon S3 manual pagination](../amazon-s3-auto-paginated-with-prefix.md) [Mutually exclusive call](../mutually-exclusive-calls-found.md) [AWS Lambda client not reused](../aws-client-per-invocation.md) [Missing check on the result of createNewFile](../missing-check-on-createnewfile.md) [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md) [Missing statement to record cause of InvocationTargetException](../missing-getcause-on-invocationtargetexception.md) [Misconfigured Concurrency](../misconfigured-concurrency.md) [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md) [Improper Initialization](../improper-initialization.md) [Unexpected re-assignment of synchronized objects](../reassign-synchronized-object.md) [XPath injection](../xpath-injection.md) [AWS client not reused in a Lambda function](../lambda-client-reuse.md) [Long polling is not enabled in Amazon SQS](../amazon-sqs-enable-long-polling.md) [Insecure temporary file or directory](../insecure-temporary-file.md) [HTTP response splitting](../http-response-splitting.md) [Input and output values become out of sync](../out-of-sync-input-and-output.md) [Server-side request forgery](../server-side-request-forgery.md) [Missing Authorization for address id](../missing-authorization-for-address-id.md) [Do not catch and throw exception](../do-not-catch-and-throw-exception.md) [Concurrency deadlock](../concurrency-deadlock.md) [Not recommended aws credentials classes](../not-recommended-aws-credentials-classes.md) [Path traversal](../path-traversal.md) [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md) [Missing byte array length of JSON parser](../json-parser-length.md) [Usage of an API that is not recommended](../not-recommended-apis.md) [Hardcoded credentials](../hardcoded-credentials.md) [Insecure JSON web token (JWT) parsing](../insecure-jwt-parsing.md) [Not calling finalize causes skipped cleanup steps](../finalize-on-super-class.md) [Unchecked S3 object metadata content length](../s3-object-metadata-content-length-check.md) [Untrusted data in security decision](../untrusted-data-in-decision.md) [Permissive cors configuration rule](../permissive-cors-configuration-rule.md) [Insecure cookie](../insecure-cookie.md) [Resource leak](../resource-leak.md) [XML External Entity](../xml-external-entity.md) [Bad parameters used with AWS API methods](../aws-bad-params.md) [Missing position check before getting substring](../missing-position-check-before-substring.md) [LDAP injection](../ldap-injection.md) [Avoid reset exception in Amazon S3](../avoid-reset-exception-rule.md) [Insecure hashing](../insecure-hashing.md) [Backward compatibility breaks with error message parsing](../aws-parse-error-message.md) [Inefficient map entry iteration](../iterate-on-map-entries.md) [Missing S3 bucket owner condition](../s3-verify-bucket-owner.md) [AWS DynamoDB getItem output is not null checked](../aws-dynamodb-getitem-null-check.md) [Invalid public method parameters](../method-input-validation.md) [Log injection](../log-injection.md) [Sensitive information leak](../sensitive-information-leak.md) [Usage of multiple date time pattern formatter](../date-time-pattern.md) [Synchronous publication of AWS Lambda metrics](../sync-metric-publish.md) [XML External Entity Document Builder Factory](../xml-external-entity-document-builder-factory.md) [Improper use of classes that aren't thread-safe](../thread-safety-class-violations.md) [Incorrect null check before setting a value](../incorrect-null-check-before-setting.md) [Insufficient use of name in Amazon SQS queue](../amazon-sqs-name-url.md) [Missing check on the value returned by ResultSet.next](../output-ignored-on-resultset-next.md) [Insecure TLS version](../insecure-tls-version.md) [Unsanitized input is run as code](../code-injection.md) [Use an enum to specify an AWS Region](../aws-region-enumeration.md) [Improperly formatted string arguments](../string-format-arguments.md) [Improper service shutdown](../sudden-service-shutdown.md) [Unrestricted upload of dangerous file type](../unrestricted-file-upload.md) [Untrusted AMI images](../untrusted-ami-images.md) [Insecure SAML parser configuration](../incorrect-authentication-exploitation.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Case sensitive keys in S3 object user metadata](../s3-object-user-metadata-key-case-sensitivity.md) [Stack trace not included in re-thrown exception](../throw-exception-with-trace.md) [Region specification missing from AWS client initialization](../aws-service-client-initialization.md) [Insufficient number of PBEKeySpec iterations](../insufficient-pbekeyspec-iterations.md) [URL redirection to untrusted site](../open-redirect.md) [Use of externally-controlled input to select classes or code](../unsafe-reflection.md) [Missing encryption of sensitive data in storage](../missing-encryption-at-rest.md) [Ignored output of DynamoDBMapper operations](../aws-dynamodb-mapper-batch-output-ignored.md) [Null pointer dereference](../null-pointer-dereference.md) [Cross-site scripting](../cross-site-scripting.md) [Unauthenticated LDAP requests](../ldap-authentication.md) [Use of inefficient APIs](../inefficient-apis.md) [Low maintainability with old Android features](../old-android-features.md) [Atomicity violation](../concurrency-atomicity-violation.md) [Missing handling of specifically-thrown exceptions](../missing-specifically-thrown-exception-handling.md) [Weak obfuscation of web request](../weak-obfuscation-of-request.md) [Clear text credentials](../clear-text-credentials.md) [Session fixation](../session-fixation.md) [Catching and not re-throwing or logging exceptions](../unhandled-exceptions.md) [Missing check when launching an Android activity with an implicit intent](../missing-check-on-android-startactivity.md) [Client constructor deprecation](../client-constructor-deprecated-rule.md) [Inefficient use of stream sorting](../stream-min-max-vs-sort.md) [Arithmetic overflow or underflow](../arithmetic-overflow.md) [Simplifiable code](../simplifiable-code.md) [Loose file permissions](../loose-file-permissions.md) [Manual pagination](../manual-pagination.md) [Incorrect string equality operator](../string-equality-check.md) [Inefficient chain of AWS API calls](../aws-inefficient-chain.md) [OS command injection](../os-command-injection.md) [Internationalization](../internationalization.md) [Code clone](../code-clone.md) [SQL injection](../sql-injection.md) [Missing check on method output](../missing-check-on-method-output.md) [Missing pagination](../missing-pagination.md) [Resources used by an Amazon S3 TransferManager are not released](../amazon-s3-transfer-manager-shutdown.md) [Insecure cryptography](../insecure-cryptography.md) [Missing timezone of SimpleDateFormat](../simple-date-format-time-zone.md) [Low maintainability with low class cohesion](../code-quality-metrics-class-cohesion.md) [Oversynchronization](../concurrency-over-synchronization.md) [Infinite loop](../infinite-loop.md) [Batch operations preferred over looping](../batches-preferred-over-loops.md) [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md) [Weak pseudorandom number generation](../weak-random-number-generation.md) [Insecure CORS policy](../insecure-cors-policy.md) [Missing handling of file deletion result](../missing-file-deletion-result-check.md) [Amazon SQS message visibility changed without a status check](../amazon-sqs-change-message-visibility-check-status.md) [State machine execution ARN is not logged](../step-function-arn-not-logged.md) [Client-side KMS reencryption](../aws-kms-reencryption.md) [Use Stream::anyMatch instead of Stream::findFirst or Stream::findAny](../stream-anymatch-vs-findfirst.md) [Batch request with unchecked failures](../aws-unchecked-batch-failures.md)

# Tag: owasp-top10

### [Reflected cross site scripting](../reflected-cross-site-scripting.md)

Rule to detect reflected XSS.

### [Deserialization of untrusted object](../untrusted-deserialization.md)

Deserialization of untrusted objects can lead to security vulnerabilities such as, inadvertently running remote code.

### [Insecure connection using unencrypted protocol](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Error-prone AWS IAM policy creation](../aws-iam-error-prone-policy.md)

Manually creating text-based IAM policies is error-prone.

### [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md)

Encryption that is dependent on conditional logic, such as an `if...then` clause, might cause unencrypted sensitive data to be stored.

### [XPath injection](../xpath-injection.md)

Potentially unsanitized user input in XPath queries can allow an attacker to control the query in unwanted or insecure ways.

### [Insecure temporary file or directory](../insecure-temporary-file.md)

Insecure ways of creating temporary files and directories can lead to race conditions, privilege escalation, and other security vulnerabilities.

### [HTTP response splitting](../http-response-splitting.md)

Passing data from an untrusted source into a cookie or web response might expose the user to HTTP response splitting attacks.

### [Server-side request forgery](../server-side-request-forgery.md)

Insufficient sanitization of potentially untrusted URLs on the server side can allow server requests to unwanted destinations.

### [Missing Authorization for address id](../missing-authorization-for-address-id.md)

Rule to detect lack of authorization check when using address ID.

### [Not recommended aws credentials classes](../not-recommended-aws-credentials-classes.md)

Find usages of not recommended classes for AWS Credentials and suggest replacing them with something else.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Hardcoded credentials](../hardcoded-credentials.md)

Hardcoded credentials can be intercepted by malicious actors.

### [Insecure JSON web token (JWT) parsing](../insecure-jwt-parsing.md)

JWTs should not be parsed using the `parse` method.

### [Untrusted data in security decision](../untrusted-data-in-decision.md)

Security decisions should not depend on branching that can be influenced by untrusted or client-provided data.

### [Permissive cors configuration rule](../permissive-cors-configuration-rule.md)

Rule to enable detection for potential CORS vulnerabilities in services using the Coral or Spring frameworks.

### [Insecure cookie](../insecure-cookie.md)

Insecure cookies can lead to unencrypted transmission of sensitive data.

### [XML External Entity](../xml-external-entity.md)

Objects that parse or handle XML can lead to XML External Entity (XXE) attacks when misconfigured.

### [LDAP injection](../ldap-injection.md)

LDAP queries that rely on potentially untrusted inputs can allow attackers to read or modify sensitive data, run code, and perform other unwanted actions.

### [Insecure hashing](../insecure-hashing.md)

Obsolete, broken, or weak hashing algorithms can lead to security vulnerabilities.

### [Invalid public method parameters](../method-input-validation.md)

Public method parameters should be validated for nullness, unexpected values, and malicious values.

### [Log injection](../log-injection.md)

Using untrusted inputs in a log statement can enable attackers to break the log's format, forge log entries, and bypass log monitors.

### [Sensitive information leak](../sensitive-information-leak.md)

Sensitive information should not be exposed through log files or stack traces.

### [XML External Entity Document Builder Factory](../xml-external-entity-document-builder-factory.md)

Objects that parse or handle XML in XML document can lead to XML External Entity (XXE) attacks when misconfigured.

### [Insecure TLS version](../insecure-tls-version.md)

TLS versions older than TLS version 1.1 support weak, broken, or misconfigured cryptography.

### [Unsanitized input is run as code](../code-injection.md)

Scripts generated from unsanitized inputs can lead to malicious behavior and inadvertently running code remotely.

### [Unrestricted upload of dangerous file type](../unrestricted-file-upload.md)

Insufficiently restrictive file uploads can lead to inadvertently running malicious code.

### [Insecure SAML parser configuration](../incorrect-authentication-exploitation.md)

Comment parsing for OpenSAML2 might enable an attacker to bypass authentication.

### [Cross-site request forgery](../cross-site-request-forgery.md)

Insecure configuration can lead to a cross-site request forgery (CRSF) vulnerability.

### [URL redirection to untrusted site](../open-redirect.md)

User-controlled input that specifies a link to an external site could lead to phishing attacks and allow user credentials to be stolen.

### [Use of externally-controlled input to select classes or code](../unsafe-reflection.md)

Use of unsanitized external input in reflection can allow attackers to bypass security checks and run malicious code.

### [Cross-site scripting](../cross-site-scripting.md)

Relying on potentially untrusted user inputs when constructing web application outputs can lead to cross-site scripting vulnerabilities.

### [Unauthenticated LDAP requests](../ldap-authentication.md)

Unauthenticated LDAP requests can allow untrusted access to LDAP servers.

### [Weak obfuscation of web request](../weak-obfuscation-of-request.md)

Weak obfuscation while configuring a web request

### [Clear text credentials](../clear-text-credentials.md)

Credentials that are stored in clear text can be intercepted by a malicious actor.

### [Session fixation](../session-fixation.md)

Session fixation might allow an attacker to steal authenticated session IDs.

### [Loose file permissions](../loose-file-permissions.md)

Weak file permissions can lead to privilege escalation.

### [OS command injection](../os-command-injection.md)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code.

### [SQL injection](../sql-injection.md)

Use of untrusted inputs in SQL database query can enable attackers to read, modify, or delete sensitive data in the database

### [Insecure cryptography](../insecure-cryptography.md)

Weak, broken, or misconfigured cryptography can lead to security vulnerabilities.

### [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md)

Deserialization of untrusted data without sufficiently verifying that the resulting data will be valid.

### [Weak pseudorandom number generation](../weak-random-number-generation.md)

Insufficiently random generators (or hardcoded seeds) can make pseudorandom sequences predictable.

### [Insecure CORS policy](../insecure-cors-policy.md)

Cross-Origin Resource Sharing policies that are too permissive may lead to security vulnerabilities.

### [Client-side KMS reencryption](../aws-kms-reencryption.md)

Client-side decryption followed by reencryption is inefficient and can lead to sensitive data leaks.
