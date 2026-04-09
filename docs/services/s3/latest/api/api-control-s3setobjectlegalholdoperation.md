# S3SetObjectLegalHoldOperation

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

## Contents

**LegalHold**

Contains the Object Lock legal hold status to be applied to all objects in the
Batch Operations job.

Type: [S3ObjectLockLegalHold](api-control-s3objectlocklegalhold.md) data type

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/s3setobjectlegalholdoperation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/s3setobjectlegalholdoperation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/s3setobjectlegalholdoperation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3SetObjectAclOperation

S3SetObjectRetentionOperation

All content copied from https://docs.aws.amazon.com/.
