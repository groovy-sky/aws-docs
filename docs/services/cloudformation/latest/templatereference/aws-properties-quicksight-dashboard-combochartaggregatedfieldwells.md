This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ComboChartAggregatedFieldWells

The aggregated field wells of a combo chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BarValues" : [ MeasureField, ... ],
  "Category" : [ DimensionField, ... ],
  "Colors" : [ DimensionField, ... ],
  "LineValues" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  BarValues:
    - MeasureField
  Category:
    - DimensionField
  Colors:
    - DimensionField
  LineValues:
    - MeasureField

```

## Properties

`BarValues`

The aggregated `BarValues` field well of a combo chart.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Category`

The aggregated category field wells of a combo chart.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Colors`

The aggregated colors field well of a combo chart.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineValues`

The aggregated `LineValues` field well of a combo chart.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnTooltipItem

ComboChartConfiguration

All content copied from https://docs.aws.amazon.com/.
