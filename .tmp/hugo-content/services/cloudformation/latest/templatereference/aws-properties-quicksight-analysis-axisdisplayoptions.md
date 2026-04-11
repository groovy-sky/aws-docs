This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AxisDisplayOptions

The display options for the axis label.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisLineVisibility" : String,
  "AxisOffset" : String,
  "DataOptions" : AxisDataOptions,
  "GridLineVisibility" : String,
  "ScrollbarOptions" : ScrollBarOptions,
  "TickLabelOptions" : AxisTickLabelOptions
}

```

### YAML

```yaml

  AxisLineVisibility: String
  AxisOffset: String
  DataOptions:
    AxisDataOptions
  GridLineVisibility: String
  ScrollbarOptions:
    ScrollBarOptions
  TickLabelOptions:
    AxisTickLabelOptions

```

## Properties

`AxisLineVisibility`

Determines whether or not the axis line is visible.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AxisOffset`

The offset value that determines the starting placement of the axis within a visual's bounds.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataOptions`

The data options for an axis.

_Required_: No

_Type_: [AxisDataOptions](aws-properties-quicksight-analysis-axisdataoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GridLineVisibility`

Determines whether or not the grid line is visible.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScrollbarOptions`

The scroll bar options for an axis.

_Required_: No

_Type_: [ScrollBarOptions](aws-properties-quicksight-analysis-scrollbaroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TickLabelOptions`

The tick label options of an axis.

_Required_: No

_Type_: [AxisTickLabelOptions](aws-properties-quicksight-analysis-axisticklabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisDisplayMinMaxRange

AxisDisplayRange

All content copied from https://docs.aws.amazon.com/.
