This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PluginVisual

A flexible visualization type that allows engineers
to create new custom charts in Quick Sight.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChartConfiguration" : PluginVisualConfiguration,
  "PluginArn" : String,
  "Subtitle" : VisualSubtitleLabelOptions,
  "Title" : VisualTitleLabelOptions,
  "VisualContentAltText" : String,
  "VisualId" : String
}

```

### YAML

```yaml

  ChartConfiguration:
    PluginVisualConfiguration
  PluginArn: String
  Subtitle:
    VisualSubtitleLabelOptions
  Title:
    VisualTitleLabelOptions
  VisualContentAltText: String
  VisualId: String

```

## Properties

`ChartConfiguration`

A description of the plugin field wells and their persisted properties.

_Required_: No

_Type_: [PluginVisualConfiguration](aws-properties-quicksight-analysis-pluginvisualconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PluginArn`

The Amazon Resource Name (ARN) that reflects the plugin and version.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtitle`

Property description not available.

_Required_: No

_Type_: [VisualSubtitleLabelOptions](aws-properties-quicksight-analysis-visualsubtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

Property description not available.

_Required_: No

_Type_: [VisualTitleLabelOptions](aws-properties-quicksight-analysis-visualtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualContentAltText`

The alt text for the visual.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualId`

The ID of the visual that you want to use.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTotalOptions

PluginVisualConfiguration

All content copied from https://docs.aws.amazon.com/.
