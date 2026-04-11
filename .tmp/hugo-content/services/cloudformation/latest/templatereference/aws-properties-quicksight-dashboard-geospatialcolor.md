This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialColor

The visualization properties for solid, gradient, and categorical colors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Categorical" : GeospatialCategoricalColor,
  "Gradient" : GeospatialGradientColor,
  "Solid" : GeospatialSolidColor
}

```

### YAML

```yaml

  Categorical:
    GeospatialCategoricalColor
  Gradient:
    GeospatialGradientColor
  Solid:
    GeospatialSolidColor

```

## Properties

`Categorical`

The visualization properties for the categorical color.

_Required_: No

_Type_: [GeospatialCategoricalColor](aws-properties-quicksight-dashboard-geospatialcategoricalcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gradient`

The visualization properties for the gradient color.

_Required_: No

_Type_: [GeospatialGradientColor](aws-properties-quicksight-dashboard-geospatialgradientcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Solid`

The visualization properties for the solid color.

_Required_: No

_Type_: [GeospatialSolidColor](aws-properties-quicksight-dashboard-geospatialsolidcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialCircleSymbolStyle

GeospatialCoordinateBounds

All content copied from https://docs.aws.amazon.com/.
