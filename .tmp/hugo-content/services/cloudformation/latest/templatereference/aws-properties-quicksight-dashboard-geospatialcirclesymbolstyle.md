This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialCircleSymbolStyle

The properties for a circle symbol style.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CircleRadius" : GeospatialCircleRadius,
  "FillColor" : GeospatialColor,
  "StrokeColor" : GeospatialColor,
  "StrokeWidth" : GeospatialLineWidth
}

```

### YAML

```yaml

  CircleRadius:
    GeospatialCircleRadius
  FillColor:
    GeospatialColor
  StrokeColor:
    GeospatialColor
  StrokeWidth:
    GeospatialLineWidth

```

## Properties

`CircleRadius`

The radius of the circle.

_Required_: No

_Type_: [GeospatialCircleRadius](aws-properties-quicksight-dashboard-geospatialcircleradius.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FillColor`

The color and opacity values for the fill color.

_Required_: No

_Type_: [GeospatialColor](aws-properties-quicksight-dashboard-geospatialcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrokeColor`

The color and opacity values for the stroke color.

_Required_: No

_Type_: [GeospatialColor](aws-properties-quicksight-dashboard-geospatialcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrokeWidth`

The width of the stroke (border).

_Required_: No

_Type_: [GeospatialLineWidth](aws-properties-quicksight-dashboard-geospatiallinewidth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialCircleRadius

GeospatialColor

All content copied from https://docs.aws.amazon.com/.
