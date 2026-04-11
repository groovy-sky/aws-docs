This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet AppendedColumn

Represents a column that will be included in the result of an append operation, combining data from multiple sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "NewColumnId" : String
}

```

### YAML

```yaml

  ColumnName: String
  NewColumnId: String

```

## Properties

`ColumnName`

The name of the column to include in the appended result.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewColumnId`

A unique identifier for the column in the appended result.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aggregation

AppendOperation

All content copied from https://docs.aws.amazon.com/.
