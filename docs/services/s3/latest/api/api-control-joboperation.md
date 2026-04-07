# JobOperation

The operation that you want this job to perform on every object listed in the manifest.
For more information about the available operations, see [Operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-operations.html) in the
_Amazon S3 User Guide_.

## Contents

**LambdaInvoke**

Directs the specified job to invoke an AWS Lambda function on every object in the
manifest.

Type: [LambdaInvokeOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_LambdaInvokeOperation.html) data type

Required: No

**S3ComputeObjectChecksum**

Directs the specified job to compute checksum values for every object in the manifest.

Type: [S3ComputeObjectChecksumOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3ComputeObjectChecksumOperation.html) data type

Required: No

**S3DeleteObjectTagging**

Directs the specified job to execute a DELETE Object tagging call on every object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3DeleteObjectTaggingOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3DeleteObjectTaggingOperation.html) data type

Required: No

**S3InitiateRestoreObject**

Directs the specified job to initiate restore requests for every archived object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3InitiateRestoreObjectOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3InitiateRestoreObjectOperation.html) data type

Required: No

**S3PutObjectAcl**

Directs the specified job to run a `PutObjectAcl` call on every object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectAclOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3SetObjectAclOperation.html) data type

Required: No

**S3PutObjectCopy**

Directs the specified job to run a PUT Copy object call on every object in the
manifest.

Type: [S3CopyObjectOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3CopyObjectOperation.html) data type

Required: No

**S3PutObjectLegalHold**

Contains the configuration for an S3 Object Lock legal hold operation that an
S3 Batch Operations job passes
to
every object to the underlying
`PutObjectLegalHold`
API
operation. For more information, see [Using S3 Object Lock legal hold\
with S3 Batch Operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-legal-hold.html) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectLegalHoldOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3SetObjectLegalHoldOperation.html) data type

Required: No

**S3PutObjectRetention**

Contains the configuration parameters for the Object Lock retention action for an
S3 Batch Operations job. Batch Operations passes every object to the underlying
`PutObjectRetention`
API
operation. For more information, see [Using S3 Object Lock retention\
with S3 Batch Operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-retention-date.html) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectRetentionOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3SetObjectRetentionOperation.html) data type

Required: No

**S3PutObjectTagging**

Directs the specified job to run a PUT Object tagging call on every object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectTaggingOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3SetObjectTaggingOperation.html) data type

Required: No

**S3ReplicateObject**

Directs the specified job to invoke `ReplicateObject` on every object in the
job's manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3ReplicateObjectOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3ReplicateObjectOperation.html) data type

Required: No

**S3UpdateObjectEncryption**

Updates the server-side encryption type of an existing encrypted
object in a general purpose bucket. You can use the `UpdateObjectEncryption`
operation to change encrypted objects from server-side encryption with
Amazon S3 managed keys (SSE-S3) to server-side encryption with AWS Key Management Service (AWS KMS)
keys (SSE-KMS), or to apply S3 Bucket Keys. You can also use the
`UpdateObjectEncryption` operation to change the customer-managed
KMS key used to encrypt your data so that you can comply with custom
key-rotation standards.

Type: [S3UpdateObjectEncryptionOperation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_S3UpdateObjectEncryptionOperation.html) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/JobOperation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/JobOperation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/JobOperation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobManifestSpec

JobProgressSummary
