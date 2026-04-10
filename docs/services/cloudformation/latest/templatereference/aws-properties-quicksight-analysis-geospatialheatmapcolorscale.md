This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialHeatmapColorScale

The color scale specification for the heatmap point style.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Colors" : [ GeospatialHeatmapDataColor, ... ]
}

```

### YAML

```yaml

  Colors:
    - GeospatialHeatmapDataColor

```

## Properties

`Colors`

The list of colors to be used in heatmap point style.

_Required_: No

_Type_: Array of [GeospatialHeatmapDataColor](aws-properties-quicksight-analysis-geospatialheatmapdatacolor.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialGradientStepColor

GeospatialHeatmapConfiguration

All content copied from https://docs.aws.amazon.com/.
