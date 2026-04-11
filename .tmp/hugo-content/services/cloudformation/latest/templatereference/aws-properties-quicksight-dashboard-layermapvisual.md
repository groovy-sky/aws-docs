This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard LayerMapVisual

A layer map visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChartConfiguration" : GeospatialLayerMapConfiguration,
  "DataSetIdentifier" : String,
  "Subtitle" : VisualSubtitleLabelOptions,
  "Title" : VisualTitleLabelOptions,
  "VisualContentAltText" : String,
  "VisualId" : String
}

```

### YAML

```yaml

  ChartConfiguration:
    GeospatialLayerMapConfiguration
  DataSetIdentifier: String
  Subtitle:
    VisualSubtitleLabelOptions
  Title:
    VisualTitleLabelOptions
  VisualContentAltText: String
  VisualId: String

```

## Properties

`ChartConfiguration`

The configuration settings of the visual.

_Required_: No

_Type_: [GeospatialLayerMapConfiguration](aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetIdentifier`

The dataset that is used to create the layer map visual. You can't create a visual without a dataset.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtitle`

Property description not available.

_Required_: No

_Type_: [VisualSubtitleLabelOptions](aws-properties-quicksight-dashboard-visualsubtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

Property description not available.

_Required_: No

_Type_: [VisualTitleLabelOptions](aws-properties-quicksight-dashboard-visualtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualContentAltText`

The alt text for the visual.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualId`

The ID of the visual.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LayerCustomActionOperation

Layout

All content copied from https://docs.aws.amazon.com/.
