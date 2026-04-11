# JobOperation

The operation that you want this job to perform on every object listed in the manifest.
For more information about the available operations, see [Operations](../dev/batch-ops-operations.md) in the
_Amazon S3 User Guide_.

## Contents

**LambdaInvoke**

Directs the specified job to invoke an AWS Lambda function on every object in the
manifest.

Type: [LambdaInvokeOperation](api-control-lambdainvokeoperation.md) data type

Required: No

**S3ComputeObjectChecksum**

Directs the specified job to compute checksum values for every object in the manifest.

Type: [S3ComputeObjectChecksumOperation](api-control-s3computeobjectchecksumoperation.md) data type

Required: No

**S3DeleteObjectTagging**

Directs the specified job to execute a DELETE Object tagging call on every object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3DeleteObjectTaggingOperation](api-control-s3deleteobjecttaggingoperation.md) data type

Required: No

**S3InitiateRestoreObject**

Directs the specified job to initiate restore requests for every archived object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3InitiateRestoreObjectOperation](api-control-s3initiaterestoreobjectoperation.md) data type

Required: No

**S3PutObjectAcl**

Directs the specified job to run a `PutObjectAcl` call on every object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectAclOperation](api-control-s3setobjectacloperation.md) data type

Required: No

**S3PutObjectCopy**

Directs the specified job to run a PUT Copy object call on every object in the
manifest.

Type: [S3CopyObjectOperation](api-control-s3copyobjectoperation.md) data type

Required: No

**S3PutObjectLegalHold**

Contains the configuration for an S3 Object Lock legal hold operation that an
S3 Batch Operations job passes
to
every object to the underlying
`PutObjectLegalHold`
API
operation. For more information, see [Using S3 Object Lock legal hold\
with S3 Batch Operations](../dev/batch-ops-legal-hold.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectLegalHoldOperation](api-control-s3setobjectlegalholdoperation.md) data type

Required: No

**S3PutObjectRetention**

Contains the configuration parameters for the Object Lock retention action for an
S3 Batch Operations job. Batch Operations passes every object to the underlying
`PutObjectRetention`
API
operation. For more information, see [Using S3 Object Lock retention\
with S3 Batch Operations](../dev/batch-ops-retention-date.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectRetentionOperation](api-control-s3setobjectretentionoperation.md) data type

Required: No

**S3PutObjectTagging**

Directs the specified job to run a PUT Object tagging call on every object in the
manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3SetObjectTaggingOperation](api-control-s3setobjecttaggingoperation.md) data type

Required: No

**S3ReplicateObject**

Directs the specified job to invoke `ReplicateObject` on every object in the
job's manifest.

###### Note

This functionality is not supported by directory buckets.

Type: [S3ReplicateObjectOperation](api-control-s3replicateobjectoperation.md) data type

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

Type: [S3UpdateObjectEncryptionOperation](api-control-s3updateobjectencryptionoperation.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/joboperation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/joboperation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/joboperation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobManifestSpec

JobProgressSummary

All content copied from https://docs.aws.amazon.com/.
