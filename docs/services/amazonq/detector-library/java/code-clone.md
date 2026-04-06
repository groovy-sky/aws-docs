![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Java detectors(132/132)

[Reflected cross site scripting](reflected-cross-site-scripting.md) [Mandatory method not called after object creation](mandatory-methods.md) [Process empty record list in Amazon KCL](kcl-with-call-process-records.md) [AWS object presence check](object-presence.md) [Missing timeout check on CountDownLatch.await](missing-timeout-check-on-latch-await.md) [Unspecified default value](unspecified-default-value.md) [Device Permission Usage.](device-permission-usage.md) [Deserialization of untrusted object](untrusted-deserialization.md) [Preserve thread interruption status rule](preserve-thread-interruption-status-rule.md) [Missing check on the value returned by moveToFirst API](output-ignored-on-movetofirst-operation.md) [Missing timeout check on ExecutorService.awaitTermination](missing-timeout-check-on-awaittermination.md) [Overflow when deserializing relational database objects](javax-persistence-id.md) [Custom manual retries of AWS SDK calls](aws-custom-retries.md) [Missing null check for cache response metadata](null-check-cache-response-metadata.md) [Inefficient usage of Transaction library from AWS Labs](dynamodb-transaction-library.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Inefficient additional authenticated data (AAD) authenticity](cipher-update-aad.md) [Use of a deprecated method](deprecated-method.md) [Error-prone AWS IAM policy creation](aws-iam-error-prone-policy.md) [Use of externally-controlled input to build connection string](connection-string-injection.md) [Inefficient Amazon S3 manual pagination](amazon-s3-auto-paginated-with-prefix.md) [Mutually exclusive call](mutually-exclusive-calls-found.md) [AWS Lambda client not reused](aws-client-per-invocation.md) [Missing check on the result of createNewFile](missing-check-on-createnewfile.md) [Sensitive data stored unencrypted due to partial encryption](partial-encryption.md) [Missing statement to record cause of InvocationTargetException](missing-getcause-on-invocationtargetexception.md) [Misconfigured Concurrency](misconfigured-concurrency.md) [Inefficient polling of AWS resource](aws-polling-instead-of-waiter.md) [Improper Initialization](improper-initialization.md) [Unexpected re-assignment of synchronized objects](reassign-synchronized-object.md) [XPath injection](xpath-injection.md) [AWS client not reused in a Lambda function](lambda-client-reuse.md) [Long polling is not enabled in Amazon SQS](amazon-sqs-enable-long-polling.md) [Insecure temporary file or directory](insecure-temporary-file.md) [HTTP response splitting](http-response-splitting.md) [Input and output values become out of sync](out-of-sync-input-and-output.md) [Server-side request forgery](server-side-request-forgery.md) [Missing Authorization for address id](missing-authorization-for-address-id.md) [Do not catch and throw exception](do-not-catch-and-throw-exception.md) [Concurrency deadlock](concurrency-deadlock.md) [Not recommended aws credentials classes](not-recommended-aws-credentials-classes.md) [Path traversal](path-traversal.md) [Override of reserved variable names in a Lambda function](lambda-override-reserved.md) [Missing byte array length of JSON parser](json-parser-length.md) [Usage of an API that is not recommended](not-recommended-apis.md) [Hardcoded credentials](hardcoded-credentials.md) [Insecure JSON web token (JWT) parsing](insecure-jwt-parsing.md) [Not calling finalize causes skipped cleanup steps](finalize-on-super-class.md) [Unchecked S3 object metadata content length](s3-object-metadata-content-length-check.md) [Untrusted data in security decision](untrusted-data-in-decision.md) [Permissive cors configuration rule](permissive-cors-configuration-rule.md) [Insecure cookie](insecure-cookie.md) [Resource leak](resource-leak.md) [XML External Entity](xml-external-entity.md) [Bad parameters used with AWS API methods](aws-bad-params.md) [Missing position check before getting substring](missing-position-check-before-substring.md) [LDAP injection](ldap-injection.md) [Avoid reset exception in Amazon S3](avoid-reset-exception-rule.md) [Insecure hashing](insecure-hashing.md) [Backward compatibility breaks with error message parsing](aws-parse-error-message.md) [Inefficient map entry iteration](iterate-on-map-entries.md) [Missing S3 bucket owner condition](s3-verify-bucket-owner.md) [AWS DynamoDB getItem output is not null checked](aws-dynamodb-getitem-null-check.md) [Invalid public method parameters](method-input-validation.md) [Log injection](log-injection.md) [Sensitive information leak](sensitive-information-leak.md) [Usage of multiple date time pattern formatter](date-time-pattern.md) [Synchronous publication of AWS Lambda metrics](sync-metric-publish.md) [XML External Entity Document Builder Factory](xml-external-entity-document-builder-factory.md) [Improper use of classes that aren't thread-safe](thread-safety-class-violations.md) [Incorrect null check before setting a value](incorrect-null-check-before-setting.md) [Insufficient use of name in Amazon SQS queue](amazon-sqs-name-url.md) [Missing check on the value returned by ResultSet.next](output-ignored-on-resultset-next.md) [Insecure TLS version](insecure-tls-version.md) [Unsanitized input is run as code](code-injection.md) [Use an enum to specify an AWS Region](aws-region-enumeration.md) [Improperly formatted string arguments](string-format-arguments.md) [Improper service shutdown](sudden-service-shutdown.md) [Unrestricted upload of dangerous file type](unrestricted-file-upload.md) [Untrusted AMI images](untrusted-ami-images.md) [Insecure SAML parser configuration](incorrect-authentication-exploitation.md) [Cross-site request forgery](cross-site-request-forgery.md) [Case sensitive keys in S3 object user metadata](s3-object-user-metadata-key-case-sensitivity.md) [Stack trace not included in re-thrown exception](throw-exception-with-trace.md) [Region specification missing from AWS client initialization](aws-service-client-initialization.md) [Insufficient number of PBEKeySpec iterations](insufficient-pbekeyspec-iterations.md) [URL redirection to untrusted site](open-redirect.md) [Use of externally-controlled input to select classes or code](unsafe-reflection.md) [Missing encryption of sensitive data in storage](missing-encryption-at-rest.md) [Ignored output of DynamoDBMapper operations](aws-dynamodb-mapper-batch-output-ignored.md) [Null pointer dereference](null-pointer-dereference.md) [Cross-site scripting](cross-site-scripting.md) [Unauthenticated LDAP requests](ldap-authentication.md) [Use of inefficient APIs](inefficient-apis.md) [Low maintainability with old Android features](old-android-features.md) [Atomicity violation](concurrency-atomicity-violation.md) [Missing handling of specifically-thrown exceptions](missing-specifically-thrown-exception-handling.md) [Weak obfuscation of web request](weak-obfuscation-of-request.md) [Clear text credentials](clear-text-credentials.md) [Session fixation](session-fixation.md) [Catching and not re-throwing or logging exceptions](unhandled-exceptions.md) [Missing check when launching an Android activity with an implicit intent](missing-check-on-android-startactivity.md) [Client constructor deprecation](client-constructor-deprecated-rule.md) [Inefficient use of stream sorting](stream-min-max-vs-sort.md) [Arithmetic overflow or underflow](arithmetic-overflow.md) [Simplifiable code](simplifiable-code.md) [Loose file permissions](loose-file-permissions.md) [Manual pagination](manual-pagination.md) [Incorrect string equality operator](string-equality-check.md) [Inefficient chain of AWS API calls](aws-inefficient-chain.md) [OS command injection](os-command-injection.md) [Internationalization](internationalization.md) [Code clone](https://docs.aws.amazon.com/amazonq/detector-library/java/code-clone) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/java/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/java/missing-check-on-method-output) [Missing pagination](https://docs.aws.amazon.com/amazonq/detector-library/java/missing-pagination) [Resources used by an Amazon S3 TransferManager are not released](https://docs.aws.amazon.com/amazonq/detector-library/java/amazon-s3-transfer-manager-shutdown) [Insecure cryptography](https://docs.aws.amazon.com/amazonq/detector-library/java/insecure-cryptography) [Missing timezone of SimpleDateFormat](https://docs.aws.amazon.com/amazonq/detector-library/java/simple-date-format-time-zone) [Low maintainability with low class cohesion](https://docs.aws.amazon.com/amazonq/detector-library/java/code-quality-metrics-class-cohesion) [Oversynchronization](https://docs.aws.amazon.com/amazonq/detector-library/java/concurrency-over-synchronization) [Infinite loop](https://docs.aws.amazon.com/amazonq/detector-library/java/infinite-loop) [Batch operations preferred over looping](https://docs.aws.amazon.com/amazonq/detector-library/java/batches-preferred-over-loops) [Object Input Stream Insecure Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/java/object-input-stream-insecure-deserialization) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/java/weak-random-number-generation) [Insecure CORS policy](https://docs.aws.amazon.com/amazonq/detector-library/java/insecure-cors-policy) [Missing handling of file deletion result](https://docs.aws.amazon.com/amazonq/detector-library/java/missing-file-deletion-result-check) [Amazon SQS message visibility changed without a status check](https://docs.aws.amazon.com/amazonq/detector-library/java/amazon-sqs-change-message-visibility-check-status) [State machine execution ARN is not logged](https://docs.aws.amazon.com/amazonq/detector-library/java/step-function-arn-not-logged) [Client-side KMS reencryption](https://docs.aws.amazon.com/amazonq/detector-library/java/aws-kms-reencryption) [Use Stream::anyMatch instead of Stream::findFirst or Stream::findAny](https://docs.aws.amazon.com/amazonq/detector-library/java/stream-anymatch-vs-findfirst) [Batch request with unchecked failures](https://docs.aws.amazon.com/amazonq/detector-library/java/aws-unchecked-batch-failures)

# Code clone [Medium](https://docs.aws.amazon.com/amazonq/detector-library/java/severity/medium)

Similar code fragments were detected in the same file. This could indicate a deeper problem and reduce code maintainability.

**Detector ID**

java/code-clone@v1.0

**Category**

[Code Quality](https://docs.aws.amazon.com/amazonq/detector-library/java/categories/code-quality)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

-

**Tags**

[\# maintainability](https://docs.aws.amazon.com/amazonq/detector-library/java/tags/maintainability)

* * *

#### Noncompliant example

```java
1private boolean doesVideoStreamExistForJobNoncompliant(final String videoStreamArn,
2                                   final String viewAngle,
3                                   final String activityJobArn,
4                                   final ActivityType activityType) throws Exception {
5    try {
6        ActivityJobItem activityJob = null;
7        // Noncompliant: uses similar code fragments in the same file.
8        if (activityType == ActivityType.TRAINING) {
9            activityJob = trainingJobDao.loadTrainingJob(activityJobArn);
10        } else if (activityType == ActivityType.EVALUATION) {
11            activityJob = evaluationJobDao.loadEvaluationJob(activityJobArn);
12        } else if (activityType == ActivityType.FINETUNING) {
13            activityJob = finetuningJobDao.loadFinetuningJob(activityJobArn);
14        }
15        if (activityJob == null) {
16            throw new Exception("Unexpected workflow activity job.");
17        }
18        return containsVideoStreamWithGivenAngleAndArn(videoStreamArn, viewAngle, activityJob);
19    } catch (Exception ex) {
20        log.error("Unable to get video stream data from DynamoDB.", ex);
21        throw ex;
22    }
23}
24private void updateVideoInfoInDynamoDBNoncompliant(final Path videoFilePath,
25                                       final String s3BucketName,
26                                       final String activityJobArn,
27	   final ActivityType activityType) {
28    String videoFileLocation = null;
29    if ((videoFilePath != null) && (s3BucketName != null)) {
30        String videoFileName = videoFilePath.toFile().getName();
31        videoFileLocation = "s3://" + s3BucketName + "/" + S3_OBJECT_KEY_PREFIX + videoFileName;
32    }
33    ActivityJobItem activityJob = null;
34    // Noncompliant: uses similar code fragments in the same file.
35    if (activityType == ActivityType.TRAINING) {
36        activityJob = trainingJobDao.loadTrainingJob(activityJobArn);
37    } else if (activityType == ActivityType.EVALUATION) {
38        activityJob = evaluationJobDao.loadEvaluationJob(activityJobArn);
39    } else if (activityType == ActivityType.FINETUNING) {
40        activityJob = finetuningJobDao.loadFinetuningJob(activityJobArn);
41    }
42    if (activityJob == null) {
43        return;
44    }
45    updateActivityJobItem(activityJob, videoFileLocation, activityType);
46}

```

#### Compliant example

```java
1public ActivityJobItem getJobFromArnAndActivityType(String activityJobArn, ActivityType activityType){
2    ActivityJobItem activityJob = null;
3    if (activityType == ActivityType.TRAINING) {
4        activityJob = trainingJobDao.loadTrainingJob(activityJobArn);
5    } else if (activityType == ActivityType.EVALUATION) {
6        activityJob = evaluationJobDao.loadEvaluationJob(activityJobArn);
7    } else if (activityType == ActivityType.FINETUNING) {
8        activityJob = finetuningJobDao.loadFinetuningJob(activityJobArn);
9    }
10    return activityJob;
11}
12private boolean doesVideoStreamExistForJobCompliant(final String videoStreamArn,
13                                            final String viewAngle,
14                                            final String activityJobArn,
15                                            final ActivityType activityType) throws Exception {
16    try {
17        ActivityJobItem activityJob = null;
18        // Compliant: avoids using similar code fragments in the same file.
19        activityJob = getJobFromArnAndActivityType(activityJobArn, activityType);
20        if (activityJob == null) {
21            throw new Exception("Unexpected workflow activity job.");
22        }
23        return containsVideoStreamWithGivenAngleAndArn(videoStreamArn, viewAngle, activityJob);
24    } catch (Exception ex) {
25        log.error("Unable to get video stream data from DynamoDB.", ex);
26        throw ex;
27    }
28}
29private void UpdateVideoInfoInDynamoDBCompliant(final Path videoFilePath,
30                                                final String s3BucketName,
31                                                final String activityJobArn,
32                                                final ActivityType activityType) {
33    String videoFileLocation = null;
34    if ((videoFilePath != null) && (s3BucketName != null)) {
35        String videoFileName = videoFilePath.toFile().getName();
36        videoFileLocation = "s3://" + s3BucketName + "/" + S3_OBJECT_KEY_PREFIX + videoFileName;
37    }
38    ActivityJobItem activityJob = null;
39    // Compliant: avoids using similar code fragments in the same file.
40    activityJob = getJobFromArnAndActivityType(activityJobArn, activityType);
41    if (activityJob == null) {
42        return;
43    }
44    updateActivityJobItem(activityJob, videoFileLocation, activityType);
45}

```
