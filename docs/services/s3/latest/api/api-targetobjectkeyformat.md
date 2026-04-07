# TargetObjectKeyFormat

Amazon S3 key format for log objects. Only one format, PartitionedPrefix or SimplePrefix, is
allowed.

## Contents

**PartitionedPrefix**

Partitioned S3 key for log objects.

Type: [PartitionedPrefix](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PartitionedPrefix.html) data type

Required: No

**SimplePrefix**

To use the simple format for S3 keys for log objects. To specify SimplePrefix format, set
SimplePrefix to {}.

Type: [SimplePrefix](https://docs.aws.amazon.com/AmazonS3/latest/API/API_SimplePrefix.html) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/TargetObjectKeyFormat)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/TargetObjectKeyFormat)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/TargetObjectKeyFormat)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetGrant

Tiering
