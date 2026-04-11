This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet CalculatedColumn

A calculated column for a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnId" : String,
  "ColumnName" : String,
  "Expression" : String
}

```

### YAML

```yaml

  ColumnId: String
  ColumnName: String
  Expression: String

```

## Properties

`ColumnId`

A unique ID to identify a calculated column. During a dataset update, if the column ID
of a calculated column matches that of an existing calculated column, Quick Sight
preserves the existing calculated column.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnName`

Column name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

An expression that defines the calculated column.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `250000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppendOperation

CastColumnTypeOperation

All content copied from https://docs.aws.amazon.com/.
