This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template NumericAxisOptions

The options for an axis with a numeric field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Range" : AxisDisplayRange,
  "Scale" : AxisScale
}

```

### YAML

```yaml

  Range:
    AxisDisplayRange
  Scale:
    AxisScale

```

## Properties

`Range`

The range setup of a numeric axis.

_Required_: No

_Type_: [AxisDisplayRange](aws-properties-quicksight-template-axisdisplayrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scale`

The scale setup of a numeric axis.

_Required_: No

_Type_: [AxisScale](aws-properties-quicksight-template-axisscale.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericalMeasureField

NumericEqualityDrillDownFilter

All content copied from https://docs.aws.amazon.com/.
