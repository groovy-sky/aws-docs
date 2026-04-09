# TableReplicationConfiguration

The replication configuration for an individual table. This configuration defines how the table is replicated to destination tables.

## Contents

**role**

The Amazon Resource Name (ARN) of the IAM role that S3 Tables assumes to replicate the table on your behalf.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:.+:iam::[0-9]{12}:role/.+`

Required: Yes

**rules**

An array of replication rules that define where this table should be replicated.

Type: Array of [TableReplicationRule](api-s3buckets-tablereplicationrule.md) objects

Array Members: Fixed number of 1 item.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/tablereplicationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/tablereplicationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/tablereplicationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableRecordExpirationSettings

TableReplicationRule

All content copied from https://docs.aws.amazon.com/.
