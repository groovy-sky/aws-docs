# IcebergSchemaV2

Contains details about the schema for an Iceberg table using the V2 format. This schema format supports nested and complex data types such as `struct`, `list`, and `map`, in addition to primitive types.

## Contents

**fields**

The schema fields for the table. Each field defines a column in the table, including its name, type, and whether it is required.

Type: Array of [SchemaV2Field](api-s3buckets-schemav2field.md) objects

Required: Yes

**type**

The type of the top-level schema, which is always a `struct` type as defined in the [Apache Iceberg specification](https://iceberg.apache.org/spec). This value must be `struct`.

Type: String

Valid Values: `struct`

Required: Yes

**identifier-field-ids**

A list of field IDs that are used as the identifier fields for the table. Identifier fields uniquely identify a row in the table.

Type: Array of integers

Required: No

**schema-id**

An optional unique identifier for the schema. Schema IDs are used by Apache Iceberg to track schema evolution.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/icebergschemav2.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/icebergschemav2.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/icebergschemav2.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergSchema

IcebergSnapshotManagementSettings

All content copied from https://docs.aws.amazon.com/.
