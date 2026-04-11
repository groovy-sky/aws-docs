This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SubtotalOptions

The subtotal options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabel" : String,
  "FieldLevel" : String,
  "FieldLevelOptions" : [ PivotTableFieldSubtotalOptions, ... ],
  "MetricHeaderCellStyle" : TableCellStyle,
  "StyleTargets" : [ TableStyleTarget, ... ],
  "TotalCellStyle" : TableCellStyle,
  "TotalsVisibility" : String,
  "ValueCellStyle" : TableCellStyle
}

```

### YAML

```yaml

  CustomLabel: String
  FieldLevel: String
  FieldLevelOptions:
    - PivotTableFieldSubtotalOptions
  MetricHeaderCellStyle:
    TableCellStyle
  StyleTargets:
    - TableStyleTarget
  TotalCellStyle:
    TableCellStyle
  TotalsVisibility: String
  ValueCellStyle:
    TableCellStyle

```

## Properties

`CustomLabel`

The custom label string for the subtotal cells.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldLevel`

The field level (all, custom, last) for the subtotal cells.

_Required_: No

_Type_: String

_Allowed values_: `ALL | CUSTOM | LAST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldLevelOptions`

The optional configuration of subtotal cells.

_Required_: No

_Type_: Array of [PivotTableFieldSubtotalOptions](aws-properties-quicksight-dashboard-pivottablefieldsubtotaloptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricHeaderCellStyle`

The cell styling options for the subtotals of header cells.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StyleTargets`

The style targets options for subtotals.

_Required_: No

_Type_: Array of [TableStyleTarget](aws-properties-quicksight-dashboard-tablestyletarget.md)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalCellStyle`

The cell styling options for the subtotal cells.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalsVisibility`

The visibility configuration for the subtotal cells.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueCellStyle`

The cell styling options for the subtotals of value cells.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringValueWhenUnsetConfiguration

TableAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
