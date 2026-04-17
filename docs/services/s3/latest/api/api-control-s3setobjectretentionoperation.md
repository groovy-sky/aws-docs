---
title: "S3SetObjectRetentionOperation"
---

# S3SetObjectRetentionOperation

Contains the configuration parameters for the Object Lock retention action for an
S3 Batch Operations job. Batch Operations passes every object to the underlying
`PutObjectRetention`
API
operation. For more information, see [Using S3 Object Lock retention\
with S3 Batch Operations](../dev/batch-ops-retention-date.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported by directory buckets.

## Contents

**Retention**

Contains the Object Lock retention mode to be applied to all objects in the Batch Operations
job. For more information, see [Using S3 Object Lock retention\
with S3 Batch Operations](../dev/batch-ops-retention-date.md) in the _Amazon S3 User Guide_.

Type: [S3Retention](api-control-s3retention.md) data type

Required: Yes

**BypassGovernanceRetention**

Indicates if the action should be applied to objects in the Batch Operations job even if they
have Object Lock ` GOVERNANCE` type in place.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/S3SetObjectRetentionOperation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/S3SetObjectRetentionOperation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/S3SetObjectRetentionOperation)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3SetObjectLegalHoldOperation

S3SetObjectTaggingOperation

All content copied from https://docs.aws.amazon.com/.
