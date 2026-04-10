This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialLineSymbolStyle

The symbol style for a line layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FillColor" : GeospatialColor,
  "LineWidth" : GeospatialLineWidth
}

```

### YAML

```yaml

  FillColor:
    GeospatialColor
  LineWidth:
    GeospatialLineWidth

```

## Properties

`FillColor`

The color and opacity values for the fill color.

_Required_: No

_Type_: [GeospatialColor](aws-properties-quicksight-dashboard-geospatialcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineWidth`

The width value for a line.

_Required_: No

_Type_: [GeospatialLineWidth](aws-properties-quicksight-dashboard-geospatiallinewidth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialLineStyle

GeospatialLineWidth

All content copied from https://docs.aws.amazon.com/.
