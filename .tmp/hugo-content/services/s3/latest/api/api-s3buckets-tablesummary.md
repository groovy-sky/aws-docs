# TableSummary

Contains details about a table.

## Contents

**createdAt**

The date and time the table was created at.

Type: Timestamp

Required: Yes

**modifiedAt**

The date and time the table was last modified at.

Type: Timestamp

Required: Yes

**name**

The name of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

Required: Yes

**namespace**

The name of the namespace.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

Required: Yes

**tableARN**

The Amazon Resource Name (ARN) of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

Required: Yes

**type**

The type of the table.

Type: String

Valid Values: `customer | aws`

Required: Yes

**managedByService**

The AWS service managing this table, if applicable. For example, a replicated table is managed by the S3 Tables replication service.

Type: String

Required: No

**namespaceId**

The unique identifier for the namespace that contains this table.

Type: String

Required: No

**tableBucketId**

The unique identifier for the table bucket that contains this table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/tablesummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/tablesummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/tablesummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableReplicationRule

Amazon S3 Vectors

All content copied from https://docs.aws.amazon.com/.
