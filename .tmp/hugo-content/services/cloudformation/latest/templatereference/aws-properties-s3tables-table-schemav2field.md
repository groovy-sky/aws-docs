This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table SchemaV2Field

Contains details about a schema field version 2 (V2) that supports [Apache Iceberg Nested Types](https://iceberg.apache.org/spec). Nested data type support includes struct, list and map. Primitive types are also supported.

###### Important

`IcebergSchemaV2` is mutually exclusive with `IcebergSchema`. Specify `IcebergSchema` for flat schemas with primitive types only, or `IcebergSchemaV2` for schemas that include nested types (struct, list, map). You cannot specify both.

###### Note

Top-level CloudFormation properties use PascalCase ( `Id`, `Name`, `Type`, `Required`, `Doc`). When specifying nested type objects inside `Type`, use lowercase keys ( `type`, `fields`, `id`, `name`, `required`, `element-id`, `element`, `element-required`, `key-id`, `key`, `value-id`, `value`, `value-required`) as these follow the Apache Iceberg specification format.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Doc" : String,
  "Id" : Integer,
  "Name" : String,
  "Required" : Boolean,
  "Type" : String
}

```

### YAML

```yaml

  Doc: String
  Id: Integer
  Name: String
  Required: Boolean
  Type: String

```

## Properties

`Doc`

An optional description of the field.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Id`

A unique identifier for the schema field V2. Field IDs are used by Apache Iceberg to track schema evolution and maintain compatibility across schema changes.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the field.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Required`

A Boolean value that specifies whether values are required for each row in this field. If this is `true` the field does not allow null values.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The field type. S3 Tables supports all Apache Iceberg primitive types and nested types. For more information, see the [Apache Iceberg documentation](https://iceberg.apache.org/spec).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaField

SnapshotManagement

All content copied from https://docs.aws.amazon.com/.
