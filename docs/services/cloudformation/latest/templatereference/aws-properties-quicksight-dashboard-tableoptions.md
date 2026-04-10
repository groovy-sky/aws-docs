This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableOptions

The table options for a table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CellStyle" : TableCellStyle,
  "HeaderStyle" : TableCellStyle,
  "Orientation" : String,
  "RowAlternateColorOptions" : RowAlternateColorOptions
}

```

### YAML

```yaml

  CellStyle:
    TableCellStyle
  HeaderStyle:
    TableCellStyle
  Orientation: String
  RowAlternateColorOptions:
    RowAlternateColorOptions

```

## Properties

`CellStyle`

The table cell style of table cells.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderStyle`

The table cell style of a table header.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Orientation`

The orientation (vertical, horizontal) for a table.

_Required_: No

_Type_: String

_Allowed values_: `VERTICAL | HORIZONTAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowAlternateColorOptions`

The row alternate color options (widget status, row alternate colors) for a table.

_Required_: No

_Type_: [RowAlternateColorOptions](aws-properties-quicksight-dashboard-rowalternatecoloroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableInlineVisualization

TablePaginatedReportOptions

All content copied from https://docs.aws.amazon.com/.
