This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table Order

Specifies the sort order of a sorted column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : String,
  "SortOrder" : Integer
}

```

### YAML

```yaml

  Column: String
  SortOrder: Integer

```

## Properties

`Column`

The name of the column.

_Required_: Yes

_Type_: [String](aws-properties-glue-table-column.md)

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortOrder`

Indicates that the column is sorted in ascending order
( `== 1`), or in descending order ( `==0`).

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Order Structure](../../../glue/latest/dg/aws-glue-api-catalog-tables.md#aws-glue-api-catalog-tables-Order) in the _AWS Glue Developer_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenTableFormatInput

SchemaId

All content copied from https://docs.aws.amazon.com/.
