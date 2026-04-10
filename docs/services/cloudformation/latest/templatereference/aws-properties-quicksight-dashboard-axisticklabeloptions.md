This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard AxisTickLabelOptions

The tick label options of an axis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelOptions" : LabelOptions,
  "RotationAngle" : Number
}

```

### YAML

```yaml

  LabelOptions:
    LabelOptions
  RotationAngle: Number

```

## Properties

`LabelOptions`

Determines whether or not the axis ticks are visible.

_Required_: No

_Type_: [LabelOptions](aws-properties-quicksight-dashboard-labeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationAngle`

The rotation angle of the axis tick labels.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisScale

BarChartAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
