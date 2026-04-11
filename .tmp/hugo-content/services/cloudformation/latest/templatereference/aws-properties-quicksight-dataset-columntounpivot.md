This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet ColumnToUnpivot

Specifies a column to be unpivoted, transforming it from a column into rows with associated values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "NewValue" : String
}

```

### YAML

```yaml

  ColumnName: String
  NewValue: String

```

## Properties

`ColumnName`

The name of the column to unpivot from the source data.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewValue`

The value to assign to this column in the unpivoted result, typically the column name or a descriptive label.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2047`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnTag

CreateColumnsOperation

All content copied from https://docs.aws.amazon.com/.
