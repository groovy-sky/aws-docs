This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard AxisScale

The scale setup
options for a numeric axis display.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Linear" : AxisLinearScale,
  "Logarithmic" : AxisLogarithmicScale
}

```

### YAML

```yaml

  Linear:
    AxisLinearScale
  Logarithmic:
    AxisLogarithmicScale

```

## Properties

`Linear`

The linear axis scale setup.

_Required_: No

_Type_: [AxisLinearScale](aws-properties-quicksight-dashboard-axislinearscale.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logarithmic`

The logarithmic axis scale setup.

_Required_: No

_Type_: [AxisLogarithmicScale](aws-properties-quicksight-dashboard-axislogarithmicscale.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisLogarithmicScale

AxisTickLabelOptions

All content copied from https://docs.aws.amazon.com/.
