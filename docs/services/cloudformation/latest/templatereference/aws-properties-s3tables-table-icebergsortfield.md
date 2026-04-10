This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergSortField

Defines a single sort field in an Iceberg sort order specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Direction" : String,
  "NullOrder" : String,
  "SourceId" : Integer,
  "Transform" : String
}

```

### YAML

```yaml

  Direction: String
  NullOrder: String
  SourceId: Integer
  Transform: String

```

## Properties

`Direction`

The sort direction. Valid values are `asc` for ascending order or `desc` for descending order.

_Required_: Yes

_Type_: String

_Allowed values_: `asc | desc`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NullOrder`

Specifies how null values are ordered. Valid values are `nulls-first` to place nulls before non-null values, or `nulls-last` to place nulls after non-null values.

_Required_: Yes

_Type_: String

_Allowed values_: `nulls-first | nulls-last`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceId`

The ID of the source schema field to sort by. This must reference a valid field ID from the table schema.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Transform`

The transform to apply to the source field before sorting. Use `identity` to sort by the field value directly, or specify other transforms as needed.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergSchemaV2

IcebergSortOrder

All content copied from https://docs.aws.amazon.com/.
