This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template WaterfallChartAggregatedFieldWells

The field well configuration of a waterfall visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Breakdowns" : [ DimensionField, ... ],
  "Categories" : [ DimensionField, ... ],
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Breakdowns:
    - DimensionField
  Categories:
    - DimensionField
  Values:
    - MeasureField

```

## Properties

`Breakdowns`

The breakdown field wells of a waterfall visual.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-template-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Categories`

The category field wells of a waterfall visual.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-template-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value field wells of a waterfall visual.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-template-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualTitleLabelOptions

WaterfallChartColorConfiguration

All content copied from https://docs.aws.amazon.com/.
