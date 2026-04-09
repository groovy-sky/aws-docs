# SchemaField

Contains details about a schema field.

## Contents

**name**

The name of the field.

Type: String

Required: Yes

**type**

The field type. S3 Tables supports all Apache Iceberg primitive types. For more information, see the [Apache Iceberg documentation](https://iceberg.apache.org/spec).

Type: String

Required: Yes

**id**

An optional unique identifier for the schema field. Field IDs are used by Apache Iceberg to track schema evolution and maintain compatibility across schema changes. If not specified, S3 Tables automatically assigns field IDs.

Type: Integer

Required: No

**required**

A Boolean value that specifies whether values are required for each row in this field. By default, this is `false` and null values are allowed in the field. If this is `true` the field does not allow null values.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/schemafield.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/schemafield.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/schemafield.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationInformation

SchemaV2Field

All content copied from https://docs.aws.amazon.com/.
