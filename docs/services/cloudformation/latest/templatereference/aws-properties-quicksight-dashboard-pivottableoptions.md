This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PivotTableOptions

The table options for a pivot table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CellStyle" : TableCellStyle,
  "CollapsedRowDimensionsVisibility" : String,
  "ColumnHeaderStyle" : TableCellStyle,
  "ColumnNamesVisibility" : String,
  "DefaultCellWidth" : String,
  "MetricPlacement" : String,
  "RowAlternateColorOptions" : RowAlternateColorOptions,
  "RowFieldNamesStyle" : TableCellStyle,
  "RowHeaderStyle" : TableCellStyle,
  "RowsLabelOptions" : PivotTableRowsLabelOptions,
  "RowsLayout" : String,
  "SingleMetricVisibility" : String,
  "ToggleButtonsVisibility" : String
}

```

### YAML

```yaml

  CellStyle:
    TableCellStyle
  CollapsedRowDimensionsVisibility: String
  ColumnHeaderStyle:
    TableCellStyle
  ColumnNamesVisibility: String
  DefaultCellWidth: String
  MetricPlacement: String
  RowAlternateColorOptions:
    RowAlternateColorOptions
  RowFieldNamesStyle:
    TableCellStyle
  RowHeaderStyle:
    TableCellStyle
  RowsLabelOptions:
    PivotTableRowsLabelOptions
  RowsLayout: String
  SingleMetricVisibility: String
  ToggleButtonsVisibility: String

```

## Properties

`CellStyle`

The table cell style of cells.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollapsedRowDimensionsVisibility`

The visibility setting of a pivot table's collapsed row dimension fields. If the value of this structure is `HIDDEN`, all collapsed columns in a pivot table are automatically hidden. The default value is `VISIBLE`.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnHeaderStyle`

The table cell style of the column header.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnNamesVisibility`

The visibility of the column names.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultCellWidth`

The default cell width of the pivot table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricPlacement`

The metric placement (row, column) options.

_Required_: No

_Type_: String

_Allowed values_: `ROW | COLUMN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowAlternateColorOptions`

The row alternate color options (widget status, row alternate colors).

_Required_: No

_Type_: [RowAlternateColorOptions](aws-properties-quicksight-dashboard-rowalternatecoloroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowFieldNamesStyle`

The table cell style of row field names.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowHeaderStyle`

The table cell style of the row headers.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowsLabelOptions`

The options for the label that is located above the row headers. This option is only applicable when `RowsLayout` is set to `HIERARCHY`.

_Required_: No

_Type_: [PivotTableRowsLabelOptions](aws-properties-quicksight-dashboard-pivottablerowslabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowsLayout`

The layout for the row dimension headers of a pivot table. Choose one of the following options.

- `TABULAR`: (Default) Each row field is displayed in a separate column.

- `HIERARCHY`: All row fields are displayed in a single column. Indentation is used to differentiate row headers of different fields.

_Required_: No

_Type_: String

_Allowed values_: `TABULAR | HIERARCHY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleMetricVisibility`

The visibility of the single metric options.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToggleButtonsVisibility`

Determines the visibility of the pivot table.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableFieldWells

PivotTablePaginatedReportOptions

All content copied from https://docs.aws.amazon.com/.
