This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergSchemaV2

Contains details about the schema version 2 (V2) for an Iceberg table that supports [Apache Iceberg Nested Types](https://iceberg.apache.org/spec). Nested data type support includes struct, list and map. Primitive types are also supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdentifierFieldIds" : [ Integer, ... ],
  "SchemaId" : Integer,
  "SchemaV2FieldList" : [ SchemaV2Field, ... ],
  "SchemaV2FieldType" : String
}

```

### YAML

```yaml

  IdentifierFieldIds:
    - Integer
  SchemaId: Integer
  SchemaV2FieldList:
    - SchemaV2Field
  SchemaV2FieldType: String

```

## Properties

`IdentifierFieldIds`

A list of field IDs that are used as the identifier fields for the table. Identifier fields uniquely identify a row in the table.

_Required_: No

_Type_: Array of Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaId`

An optional unique identifier for the schema.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaV2FieldList`

The schema fields for the table.

_Required_: Yes

_Type_: Array of [SchemaV2Field](aws-properties-s3tables-table-schemav2field.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaV2FieldType`

The type of the top-level schema, which is always a struct type.

_Required_: Yes

_Type_: String

_Allowed values_: `struct`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergSchema

IcebergSortField

All content copied from https://docs.aws.amazon.com/.
