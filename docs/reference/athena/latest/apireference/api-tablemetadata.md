# TableMetadata

Contains metadata for a table.

## Contents

**Name**

The name of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**Columns**

A list of the columns in the table.

Type: Array of [Column](api-column.md) objects

Required: No

**CreateTime**

The time that the table was created.

Type: Timestamp

Required: No

**LastAccessTime**

The last time the table was accessed.

Type: Timestamp

Required: No

**Parameters**

A set of custom key/value pairs for table properties.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 255.

Key Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Value Length Constraints: Maximum length of 51200.

Required: No

**PartitionKeys**

A list of the partition keys in the table.

Type: Array of [Column](api-column.md) objects

Required: No

**TableType**

The type of table. In Athena, only `EXTERNAL_TABLE` is
supported.

Type: String

Length Constraints: Maximum length of 255.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/tablemetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/tablemetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/tablemetadata.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SessionSummary

Tag
