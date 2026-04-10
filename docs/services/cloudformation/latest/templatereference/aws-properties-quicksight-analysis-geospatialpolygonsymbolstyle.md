This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialPolygonSymbolStyle

The polygon symbol style for a polygon layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FillColor" : GeospatialColor,
  "StrokeColor" : GeospatialColor,
  "StrokeWidth" : GeospatialLineWidth
}

```

### YAML

```yaml

  FillColor:
    GeospatialColor
  StrokeColor:
    GeospatialColor
  StrokeWidth:
    GeospatialLineWidth

```

## Properties

`FillColor`

The color and opacity values for the fill color.

_Required_: No

_Type_: [GeospatialColor](aws-properties-quicksight-analysis-geospatialcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrokeColor`

The color and opacity values for the stroke color.

_Required_: No

_Type_: [GeospatialColor](aws-properties-quicksight-analysis-geospatialcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrokeWidth`

The width of the border stroke.

_Required_: No

_Type_: [GeospatialLineWidth](aws-properties-quicksight-analysis-geospatiallinewidth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialPolygonStyle

GeospatialSolidColor

All content copied from https://docs.aws.amazon.com/.
