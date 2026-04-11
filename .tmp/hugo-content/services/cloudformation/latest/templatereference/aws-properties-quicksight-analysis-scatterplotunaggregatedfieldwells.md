This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ScatterPlotUnaggregatedFieldWells

The unaggregated field wells of a scatter plot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : [ DimensionField, ... ],
  "Label" : [ DimensionField, ... ],
  "Size" : [ MeasureField, ... ],
  "XAxis" : [ DimensionField, ... ],
  "YAxis" : [ DimensionField, ... ]
}

```

### YAML

```yaml

  Category:
    - DimensionField
  Label:
    - DimensionField
  Size:
    - MeasureField
  XAxis:
    - DimensionField
  YAxis:
    - DimensionField

```

## Properties

`Category`

The category field well of a scatter plot.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label field well of a scatter plot.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The size field well of a scatter plot.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-analysis-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxis`

The x-axis field well of a scatter plot.

The x-axis is a dimension field and cannot be aggregated.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxis`

The y-axis field well of a scatter plot.

The y-axis is a dimension field and cannot be aggregated.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScatterPlotSortConfiguration

ScatterPlotVisual

All content copied from https://docs.aws.amazon.com/.
