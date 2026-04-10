This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialGradientStepColor

The gradient step color for a single step.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "DataValue" : Number
}

```

### YAML

```yaml

  Color: String
  DataValue: Number

```

## Properties

`Color`

The color and opacity values for the gradient step color.

_Required_: Yes

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataValue`

The data value for the gradient step color.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialGradientColor

GeospatialHeatmapColorScale

All content copied from https://docs.aws.amazon.com/.
