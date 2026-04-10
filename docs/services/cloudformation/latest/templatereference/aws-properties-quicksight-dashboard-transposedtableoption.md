This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TransposedTableOption

The column option of the transposed table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnIndex" : Number,
  "ColumnType" : String,
  "ColumnWidth" : String
}

```

### YAML

```yaml

  ColumnIndex: Number
  ColumnType: String
  ColumnWidth: String

```

## Properties

`ColumnIndex`

The index of a columns in a transposed table. The index range is 0-9999.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `9999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnType`

The column type of the column in a transposed table. Choose one of the following options:

- `ROW_HEADER_COLUMN`: Refers to the leftmost column of the row header in the transposed table.

- `VALUE_COLUMN`: Refers to all value columns in the transposed table.

_Required_: Yes

_Type_: String

_Allowed values_: `ROW_HEADER_COLUMN | VALUE_COLUMN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnWidth`

The width of a column in a transposed table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TotalOptions

TreeMapAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
