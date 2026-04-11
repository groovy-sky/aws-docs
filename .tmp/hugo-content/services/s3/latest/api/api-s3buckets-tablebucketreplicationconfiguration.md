# TableBucketReplicationConfiguration

The replication configuration for a table bucket. This configuration defines how tables in the source bucket are replicated to destination table buckets, including the IAM role used for replication.

## Contents

**role**

The Amazon Resource Name (ARN) of the IAM role that S3 Tables assumes to replicate tables on your behalf.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:.+:iam::[0-9]{12}:role/.+`

Required: Yes

**rules**

An array of replication rules that define which tables to replicate and where to replicate them.

Type: Array of [TableBucketReplicationRule](api-s3buckets-tablebucketreplicationrule.md) objects

Array Members: Fixed number of 1 item.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/tablebucketreplicationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/tablebucketreplicationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/tablebucketreplicationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableBucketMaintenanceSettings

TableBucketReplicationRule

All content copied from https://docs.aws.amazon.com/.
