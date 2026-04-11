This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ChartAxisLabelOptions

The label options for an axis on a chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisLabelOptions" : [ AxisLabelOptions, ... ],
  "SortIconVisibility" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  AxisLabelOptions:
    - AxisLabelOptions
  SortIconVisibility: String
  Visibility: String

```

## Properties

`AxisLabelOptions`

The label options for a chart axis.

_Required_: No

_Type_: [Array](aws-properties-quicksight-analysis-axislabeloptions.md) of [AxisLabelOptions](aws-properties-quicksight-analysis-axislabeloptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortIconVisibility`

The visibility configuration of the sort icon on a chart's axis label.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of an axis label on a chart. Choose one of the following options:

- `VISIBLE`: Shows the axis.

- `HIDDEN`: Hides the axis.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CategoryInnerFilter

ClusterMarker

All content copied from https://docs.aws.amazon.com/.
