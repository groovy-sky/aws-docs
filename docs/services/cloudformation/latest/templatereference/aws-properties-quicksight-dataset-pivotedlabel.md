This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet PivotedLabel

Specifies a label value to be pivoted into a separate column, including the new column name and identifier.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelName" : String,
  "NewColumnId" : String,
  "NewColumnName" : String
}

```

### YAML

```yaml

  LabelName: String
  NewColumnId: String
  NewColumnName: String

```

## Properties

`LabelName`

The label value from the source data to be pivoted.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2047`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewColumnId`

A unique identifier for the new column created from this pivoted label.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewColumnName`

The name for the new column created from this pivoted label.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotConfiguration

PivotOperation

All content copied from https://docs.aws.amazon.com/.
