This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TotalOptions

The total options for a table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabel" : String,
  "Placement" : String,
  "ScrollStatus" : String,
  "TotalAggregationOptions" : [ TotalAggregationOption, ... ],
  "TotalCellStyle" : TableCellStyle,
  "TotalsVisibility" : String
}

```

### YAML

```yaml

  CustomLabel: String
  Placement: String
  ScrollStatus: String
  TotalAggregationOptions:
    - TotalAggregationOption
  TotalCellStyle:
    TableCellStyle
  TotalsVisibility: String

```

## Properties

`CustomLabel`

The custom label string for the total cells.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Placement`

The placement (start, end) for the total cells.

_Required_: No

_Type_: String

_Allowed values_: `START | END | AUTO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScrollStatus`

The scroll status (pinned, scrolled) for the total cells.

_Required_: No

_Type_: String

_Allowed values_: `PINNED | SCROLLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalAggregationOptions`

The total aggregation settings for each value field.

_Required_: No

_Type_: Array of [TotalAggregationOption](aws-properties-quicksight-dashboard-totalaggregationoption.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalCellStyle`

Cell styling options for the total cells.

_Required_: No

_Type_: [TableCellStyle](aws-properties-quicksight-dashboard-tablecellstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalsVisibility`

The visibility configuration for the total cells.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TotalAggregationOption

TransposedTableOption

All content copied from https://docs.aws.amazon.com/.
