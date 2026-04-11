This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template BoxPlotAggregatedFieldWells

The aggregated field well for a box plot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupBy" : [ DimensionField, ... ],
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  GroupBy:
    - DimensionField
  Values:
    - MeasureField

```

## Properties

`GroupBy`

The group by field well of a box plot chart. Values are grouped based on group by fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-template-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value field well of a box plot chart. Values are aggregated based on group by fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-template-measurefield.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BodySectionRepeatPageBreakConfiguration

BoxPlotChartConfiguration

All content copied from https://docs.aws.amazon.com/.
