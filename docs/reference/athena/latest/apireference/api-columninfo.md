# ColumnInfo

Information about the columns in a query execution result.

## Contents

**Name**

The name of the column.

Type: String

Required: Yes

**Type**

The data type of the column.

Type: String

Required: Yes

**CaseSensitive**

Indicates whether values in the column are case-sensitive.

Type: Boolean

Required: No

**CatalogName**

The catalog to which the query results belong.

Type: String

Required: No

**Label**

A column label.

Type: String

Required: No

**Nullable**

Unsupported constraint. This value always shows as `UNKNOWN`.

Type: String

Valid Values: `NOT_NULL | NULLABLE | UNKNOWN`

Required: No

**Precision**

For `DECIMAL` data types, specifies the total number of digits, up to 38.
For performance reasons, we recommend up to 18 digits.

Type: Integer

Required: No

**Scale**

For `DECIMAL` data types, specifies the total number of digits in the
fractional part of the value. Defaults to 0.

Type: Integer

Required: No

**SchemaName**

The schema name (database name) to which the query results belong.

Type: String

Required: No

**TableName**

The table name for the query results.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/columninfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/columninfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/columninfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Column

CustomerContentEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
