![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Java detectors(132/132)

[Reflected cross site scripting](../reflected-cross-site-scripting.md) [Mandatory method not called after object creation](../mandatory-methods.md) [Process empty record list in Amazon KCL](../kcl-with-call-process-records.md) [AWS object presence check](../object-presence.md) [Missing timeout check on CountDownLatch.await](../missing-timeout-check-on-latch-await.md) [Unspecified default value](../unspecified-default-value.md) [Device Permission Usage.](../device-permission-usage.md) [Deserialization of untrusted object](../untrusted-deserialization.md) [Preserve thread interruption status rule](../preserve-thread-interruption-status-rule.md) [Missing check on the value returned by moveToFirst API](../output-ignored-on-movetofirst-operation.md) [Missing timeout check on ExecutorService.awaitTermination](../missing-timeout-check-on-awaittermination.md) [Overflow when deserializing relational database objects](../javax-persistence-id.md) [Custom manual retries of AWS SDK calls](../aws-custom-retries.md) [Missing null check for cache response metadata](../null-check-cache-response-metadata.md) [Inefficient usage of Transaction library from AWS Labs](../dynamodb-transaction-library.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Inefficient additional authenticated data (AAD) authenticity](../cipher-update-aad.md) [Use of a deprecated method](../deprecated-method.md) [Error-prone AWS IAM policy creation](../aws-iam-error-prone-policy.md) [Use of externally-controlled input to build connection string](../connection-string-injection.md) [Inefficient Amazon S3 manual pagination](../amazon-s3-auto-paginated-with-prefix.md) [Mutually exclusive call](../mutually-exclusive-calls-found.md) [AWS Lambda client not reused](../aws-client-per-invocation.md) [Missing check on the result of createNewFile](../missing-check-on-createnewfile.md) [Sensitive data stored unencrypted due to partial encryption](../partial-encryption.md) [Missing statement to record cause of InvocationTargetException](../missing-getcause-on-invocationtargetexception.md) [Misconfigured Concurrency](../misconfigured-concurrency.md) [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md) [Improper Initialization](../improper-initialization.md) [Unexpected re-assignment of synchronized objects](../reassign-synchronized-object.md) [XPath injection](../xpath-injection.md) [AWS client not reused in a Lambda function](../lambda-client-reuse.md) [Long polling is not enabled in Amazon SQS](../amazon-sqs-enable-long-polling.md) [Insecure temporary file or directory](../insecure-temporary-file.md) [HTTP response splitting](../http-response-splitting.md) [Input and output values become out of sync](../out-of-sync-input-and-output.md) [Server-side request forgery](../server-side-request-forgery.md) [Missing Authorization for address id](../missing-authorization-for-address-id.md) [Do not catch and throw exception](../do-not-catch-and-throw-exception.md) [Concurrency deadlock](../concurrency-deadlock.md) [Not recommended aws credentials classes](../not-recommended-aws-credentials-classes.md) [Path traversal](../path-traversal.md) [Override of reserved variable names in a Lambda function](../lambda-override-reserved.md) [Missing byte array length of JSON parser](../json-parser-length.md) [Usage of an API that is not recommended](../not-recommended-apis.md) [Hardcoded credentials](../hardcoded-credentials.md) [Insecure JSON web token (JWT) parsing](../insecure-jwt-parsing.md) [Not calling finalize causes skipped cleanup steps](../finalize-on-super-class.md) [Unchecked S3 object metadata content length](../s3-object-metadata-content-length-check.md) [Untrusted data in security decision](../untrusted-data-in-decision.md) [Permissive cors configuration rule](../permissive-cors-configuration-rule.md) [Insecure cookie](../insecure-cookie.md) [Resource leak](../resource-leak.md) [XML External Entity](../xml-external-entity.md) [Bad parameters used with AWS API methods](../aws-bad-params.md) [Missing position check before getting substring](../missing-position-check-before-substring.md) [LDAP injection](../ldap-injection.md) [Avoid reset exception in Amazon S3](../avoid-reset-exception-rule.md) [Insecure hashing](../insecure-hashing.md) [Backward compatibility breaks with error message parsing](../aws-parse-error-message.md) [Inefficient map entry iteration](../iterate-on-map-entries.md) [Missing S3 bucket owner condition](../s3-verify-bucket-owner.md) [AWS DynamoDB getItem output is not null checked](../aws-dynamodb-getitem-null-check.md) [Invalid public method parameters](../method-input-validation.md) [Log injection](../log-injection.md) [Sensitive information leak](../sensitive-information-leak.md) [Usage of multiple date time pattern formatter](../date-time-pattern.md) [Synchronous publication of AWS Lambda metrics](../sync-metric-publish.md) [XML External Entity Document Builder Factory](../xml-external-entity-document-builder-factory.md) [Improper use of classes that aren't thread-safe](../thread-safety-class-violations.md) [Incorrect null check before setting a value](../incorrect-null-check-before-setting.md) [Insufficient use of name in Amazon SQS queue](../amazon-sqs-name-url.md) [Missing check on the value returned by ResultSet.next](../output-ignored-on-resultset-next.md) [Insecure TLS version](../insecure-tls-version.md) [Unsanitized input is run as code](../code-injection.md) [Use an enum to specify an AWS Region](../aws-region-enumeration.md) [Improperly formatted string arguments](../string-format-arguments.md) [Improper service shutdown](../sudden-service-shutdown.md) [Unrestricted upload of dangerous file type](../unrestricted-file-upload.md) [Untrusted AMI images](../untrusted-ami-images.md) [Insecure SAML parser configuration](../incorrect-authentication-exploitation.md) [Cross-site request forgery](../cross-site-request-forgery.md) [Case sensitive keys in S3 object user metadata](../s3-object-user-metadata-key-case-sensitivity.md) [Stack trace not included in re-thrown exception](../throw-exception-with-trace.md) [Region specification missing from AWS client initialization](../aws-service-client-initialization.md) [Insufficient number of PBEKeySpec iterations](../insufficient-pbekeyspec-iterations.md) [URL redirection to untrusted site](../open-redirect.md) [Use of externally-controlled input to select classes or code](../unsafe-reflection.md) [Missing encryption of sensitive data in storage](../missing-encryption-at-rest.md) [Ignored output of DynamoDBMapper operations](../aws-dynamodb-mapper-batch-output-ignored.md) [Null pointer dereference](../null-pointer-dereference.md) [Cross-site scripting](../cross-site-scripting.md) [Unauthenticated LDAP requests](../ldap-authentication.md) [Use of inefficient APIs](../inefficient-apis.md) [Low maintainability with old Android features](../old-android-features.md) [Atomicity violation](../concurrency-atomicity-violation.md) [Missing handling of specifically-thrown exceptions](../missing-specifically-thrown-exception-handling.md) [Weak obfuscation of web request](../weak-obfuscation-of-request.md) [Clear text credentials](../clear-text-credentials.md) [Session fixation](../session-fixation.md) [Catching and not re-throwing or logging exceptions](../unhandled-exceptions.md) [Missing check when launching an Android activity with an implicit intent](../missing-check-on-android-startactivity.md) [Client constructor deprecation](../client-constructor-deprecated-rule.md) [Inefficient use of stream sorting](../stream-min-max-vs-sort.md) [Arithmetic overflow or underflow](../arithmetic-overflow.md) [Simplifiable code](../simplifiable-code.md) [Loose file permissions](../loose-file-permissions.md) [Manual pagination](../manual-pagination.md) [Incorrect string equality operator](../string-equality-check.md) [Inefficient chain of AWS API calls](../aws-inefficient-chain.md) [OS command injection](../os-command-injection.md) [Internationalization](../internationalization.md) [Code clone](../code-clone.md) [SQL injection](../sql-injection.md) [Missing check on method output](../missing-check-on-method-output.md) [Missing pagination](../missing-pagination.md) [Resources used by an Amazon S3 TransferManager are not released](../amazon-s3-transfer-manager-shutdown.md) [Insecure cryptography](../insecure-cryptography.md) [Missing timezone of SimpleDateFormat](../simple-date-format-time-zone.md) [Low maintainability with low class cohesion](../code-quality-metrics-class-cohesion.md) [Oversynchronization](../concurrency-over-synchronization.md) [Infinite loop](../infinite-loop.md) [Batch operations preferred over looping](../batches-preferred-over-loops.md) [Object Input Stream Insecure Deserialization](../object-input-stream-insecure-deserialization.md) [Weak pseudorandom number generation](../weak-random-number-generation.md) [Insecure CORS policy](../insecure-cors-policy.md) [Missing handling of file deletion result](../missing-file-deletion-result-check.md) [Amazon SQS message visibility changed without a status check](../amazon-sqs-change-message-visibility-check-status.md) [State machine execution ARN is not logged](../step-function-arn-not-logged.md) [Client-side KMS reencryption](../aws-kms-reencryption.md) [Use Stream::anyMatch instead of Stream::findFirst or Stream::findAny](../stream-anymatch-vs-findfirst.md) [Batch request with unchecked failures](../aws-unchecked-batch-failures.md)

# Code Quality detectors

### [Mandatory method not called after object creation](../mandatory-methods.md)

Mandatory methods must be called after object creation.

### [Process empty record list in Amazon KCL](../kcl-with-call-process-records.md)

Setting `withCallProcessRecordsEvenForEmptyRecordList` to `TRUE` during Kinesis Client Library (KCL) initialization will treat empty records differently.

### [AWS object presence check](../object-presence.md)

Manually performing an object existence check is inefficient when a built-in operation is available.

### [Missing timeout check on CountDownLatch.await](../missing-timeout-check-on-latch-await.md)

Missing timeout check on CountDownLatch.await can cause synchronization errors.

### [Preserve thread interruption status rule](../preserve-thread-interruption-status-rule.md)

Preserve Thread Interruption Status Rule.

### [Missing timeout check on ExecutorService.awaitTermination](../missing-timeout-check-on-awaittermination.md)

Missing timeout checks on awaitTermination might make the code harder to debug.

### [Custom manual retries of AWS SDK calls](../aws-custom-retries.md)

Custom manual retries of calls to AWS SDK APIs are inefficient.

### [Missing null check for cache response metadata](../null-check-cache-response-metadata.md)

Accessing the cache response metadata without performing a `null` check might cause a null dereference error.

### [Inefficient usage of Transaction library from AWS Labs](../dynamodb-transaction-library.md)

The AWS Labs Transactions Library is a client-side solution and less efficient compared to DynamoDB native transactions.

### [Use of a deprecated method](../deprecated-method.md)

This code uses deprecated methods, which suggests that it has not been recently reviewed or maintained.

### [Inefficient Amazon S3 manual pagination](../amazon-s3-auto-paginated-with-prefix.md)

Use `S3Objects.withPrefix()` instead of manually paginating results.

### [Missing statement to record cause of InvocationTargetException](../missing-getcause-on-invocationtargetexception.md)

Missing statements to record the underlying cause of InvocationTargetException.

### [Misconfigured Concurrency](../misconfigured-concurrency.md)

Misconfigured CompletableFuture.get or Future.get.

### [Inefficient polling of AWS resource](../aws-polling-instead-of-waiter.md)

Custom polling can be inefficient and prone to error. Consider using AWS waiters instead.

### [Long polling is not enabled in Amazon SQS](../amazon-sqs-enable-long-polling.md)

Enable long polling for efficiency.

### [Do not catch and throw exception](../do-not-catch-and-throw-exception.md)

Do not catch and throw the same exception.

### [Not recommended aws credentials classes](../not-recommended-aws-credentials-classes.md)

Find usages of not recommended classes for AWS Credentials and suggest replacing them with something else.

### [Missing byte array length of JSON parser](../json-parser-length.md)

Specify the length of the input byte array when creating a `JSON` parser to avoid a deserialization problem.

### [Missing position check before getting substring](../missing-position-check-before-substring.md)

Getting a substring outside the range of a string can cause an index-out-of-bounds exception.

### [Backward compatibility breaks with error message parsing](../aws-parse-error-message.md)

Maintain your code's backward compatibility by checking the status code instead of parsing the error message.

### [Inefficient map entry iteration](../iterate-on-map-entries.md)

Iterating on `Map` entries is more efficient than iterating on the keys and asking for their respective values.

### [AWS DynamoDB getItem output is not null checked](../aws-dynamodb-getitem-null-check.md)

Outputs of AWS DynamoDB's `GetItem` method are not null checked.

### [Usage of multiple date time pattern formatter](../date-time-pattern.md)

Detects the usage of multiple date time pattern.

### [Synchronous publication of AWS Lambda metrics](../sync-metric-publish.md)

Synchronous publication of AWS Lambda metrics is inefficient.

### [Improper use of classes that aren't thread-safe](../thread-safety-class-violations.md)

Improper use of thread-unsafe classes in multi-threaded programs can cause the programs to be unstable.

### [Incorrect null check before setting a value](../incorrect-null-check-before-setting.md)

When a variable is assigned a value after checking if it's not `null` (for example, `x != null` insead of `x == null`), it might be inadvertently overwritten.

### [Insufficient use of name in Amazon SQS queue](../amazon-sqs-name-url.md)

Provide the full URL for the Amazon SQS queue.

### [Use an enum to specify an AWS Region](../aws-region-enumeration.md)

To minimize the risk of error, use an enum instead of a string to specify an AWS Region.

### [Improperly formatted string arguments](../string-format-arguments.md)

Format strings appropriately for their argument types. For example, use `%d`, not `%s`, for integers.

### [Improper service shutdown](../sudden-service-shutdown.md)

Sudden service shutdown might prevent a graceful termination of threads.

### [Region specification missing from AWS client initialization](../aws-service-client-initialization.md)

Set an explicit AWS Region to avoid cold start delays in AWS client initialization.

### [Use of inefficient APIs](../inefficient-apis.md)

Performance of this code can be enhanced by using alternative APIs.

### [Low maintainability with old Android features](../old-android-features.md)

Code uses older Android features.

### [Client constructor deprecation](../client-constructor-deprecated-rule.md)

Client constructors are now deprecated in favor of using builders to create the client.

### [Inefficient use of stream sorting](../stream-min-max-vs-sort.md)

Using `Stream::min` or `Stream::max` is more efficient than sorting and getting the first element in a stream.

### [Simplifiable code](../simplifiable-code.md)

Simplifiable code might be harder to read or maintain.

### [Manual pagination](../manual-pagination.md)

Suggest using auto-pagination instead of manual pagination.

### [Incorrect string equality operator](../string-equality-check.md)

Use `equals()`, not `==`, when checking if two strings are equal.

### [Inefficient chain of AWS API calls](../aws-inefficient-chain.md)

The chain of API calls can be replaced with a single, more efficient API call.

### [Internationalization](../internationalization.md)

Improper use of locals prevent internationalization.

### [Code clone](../code-clone.md)

Similar code fragments were detected in the same file.

### [Missing check on method output](../missing-check-on-method-output.md)

Missing checks might cause silent failures that are harder to debug.

### [Low maintainability with low class cohesion](../code-quality-metrics-class-cohesion.md)

Classes with low class cohesion contain unrelated operations which make them difficult to understand and less likely to be used.

### [Oversynchronization](../concurrency-over-synchronization.md)

Oversynchronization with `ConcurrentHashMap` or `ConcurrentLinkedQueue` can reduce program performance.

### [Infinite loop](../infinite-loop.md)

Use loop control flow to ensure that loops are exited, even if exceptional behaviors are encountered.

### [Batch operations preferred over looping](../batches-preferred-over-loops.md)

Batch operations are more efficient than looping to process several items at the same time.

### [Amazon SQS message visibility changed without a status check](../amazon-sqs-change-message-visibility-check-status.md)

When you change Amazon SQS message visibility, check for `MessageNotInFlight` exceptions.

### [State machine execution ARN is not logged](../step-function-arn-not-logged.md)

Log the ARN identifying the state machine execution for better debuggability.

### [Use Stream::anyMatch instead of Stream::findFirst or Stream::findAny](../stream-anymatch-vs-findfirst.md)

Using `Stream::anyMatch` is more readable and convenient than using a chain of `Stream::filter`, `Stream::findFirst` or `Stream::findAny` and `Optional::isPresent`.

### [Batch request with unchecked failures](../aws-unchecked-batch-failures.md)

Not checking which items have failed can lead to loss of data.
