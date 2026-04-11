This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis LineChartAggregatedFieldWells

The field well configuration of a line chart.

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

The category field wells of a line chart. Values are grouped by category fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Colors`

The color field wells of a line chart. Values are grouped by category fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiples`

The small multiples field well of a line chart.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value field wells of a line chart. Values are aggregated based on categories.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-analysis-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LegendOptions

LineChartConfiguration

All content copied from https://docs.aws.amazon.com/.
