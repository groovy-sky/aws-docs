This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard BarChartAggregatedFieldWells

The aggregated field wells of a bar chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : [ DimensionField, ... ],
  "Colors" : [ DimensionField, ... ],
  "SmallMultiples" : [ DimensionField, ... ],
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Category:
    - DimensionField
  Colors:
    - DimensionField
  SmallMultiples:
    - DimensionField
  Values:
    - MeasureField

```

## Properties

`Category`

The category (y-axis) field well of a bar chart.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Colors`

The color (group/color) field well of a bar chart.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiples`

The small multiples field well of a bar chart.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value field wells of a bar chart. Values are aggregated by
category.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisTickLabelOptions

BarChartConfiguration

All content copied from https://docs.aws.amazon.com/.
