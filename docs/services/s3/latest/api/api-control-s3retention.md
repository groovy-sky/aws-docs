# S3Retention

Contains the S3 Object Lock retention mode to be applied to all objects in the
S3 Batch Operations job. If you don't provide `Mode` and `RetainUntilDate`
data types in your operation, you will remove the retention from your objects. For more
information, see [Using S3 Object Lock retention\
with S3 Batch Operations](../dev/batch-ops-retention-date.md) in the _Amazon S3 User Guide_.

## Contents

**Mode**

The Object Lock retention mode to be applied to all objects in the Batch Operations
job.

Type: String

Valid Values: `COMPLIANCE | GOVERNANCE`

Required: No

**RetainUntilDate**

The date when the applied Object Lock retention will expire on all objects set by the
Batch Operations job.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/s3retention.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/s3retention.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/s3retention.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3ReplicateObjectOperation

S3SetObjectAclOperation

All content copied from https://docs.aws.amazon.com/.
