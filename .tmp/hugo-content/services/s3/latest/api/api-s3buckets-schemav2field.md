# SchemaV2Field

Contains details about a schema field in the V2 format. This field format supports nested and complex data types such as `struct`, `list`, and `map`, in addition to primitive types.

## Contents

**id**

The unique identifier for the schema field. Field IDs are used by Apache Iceberg to track schema evolution and maintain compatibility across schema changes.

Type: Integer

Required: Yes

**name**

The name of the field.

Type: String

Required: Yes

**required**

A Boolean value that specifies whether values are required for each row in this field. If this is `true`, the field does not allow null values.

Type: Boolean

Required: Yes

**type**

The data type of the field. This can be a primitive type string such as `boolean`, `int`, `long`, `float`, `double`, `string`, `binary`, `date`, `timestamp`, or `timestamptz`, or a complex type represented as a JSON object for nested types such as `struct`, `list`, or `map`. For more information, see the [Apache Iceberg schemas and data types documentation](https://iceberg.apache.org/spec).

Type: JSON value

Required: Yes

**doc**

An optional description of the field.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/schemav2field.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/schemav2field.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/schemav2field.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaField

StorageClassConfiguration

All content copied from https://docs.aws.amazon.com/.
