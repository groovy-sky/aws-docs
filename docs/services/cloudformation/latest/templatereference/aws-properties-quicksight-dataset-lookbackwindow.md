This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet LookbackWindow

The lookback window setup of an incremental refresh configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "Size" : Number,
  "SizeUnit" : String
}

```

### YAML

```yaml

  ColumnName: String
  Size: Number
  SizeUnit: String

```

## Properties

`ColumnName`

The name of the lookback window column.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The lookback window column size.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeUnit`

The size unit that is used for the lookback window column. Valid values for this structure are `HOUR`, `DAY`, and `WEEK`.

_Required_: Yes

_Type_: String

_Allowed values_: `HOUR | DAY | WEEK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogicalTableSource

NewDefaultValues

All content copied from https://docs.aws.amazon.com/.
