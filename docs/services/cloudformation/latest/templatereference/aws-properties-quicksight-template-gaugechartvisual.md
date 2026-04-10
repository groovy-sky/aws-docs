This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template GaugeChartVisual

A gauge chart.

For more information, see [Using gauge charts](../../../quicksight/latest/user/gauge-chart.md) in the _Amazon Quick Suite User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ VisualCustomAction, ... ],
  "ChartConfiguration" : GaugeChartConfiguration,
  "ConditionalFormatting" : GaugeChartConditionalFormatting,
  "Subtitle" : VisualSubtitleLabelOptions,
  "Title" : VisualTitleLabelOptions,
  "VisualContentAltText" : String,
  "VisualId" : String
}

```

### YAML

```yaml

  Actions:
    - VisualCustomAction
  ChartConfiguration:
    GaugeChartConfiguration
  ConditionalFormatting:
    GaugeChartConditionalFormatting
  Subtitle:
    VisualSubtitleLabelOptions
  Title:
    VisualTitleLabelOptions
  VisualContentAltText: String
  VisualId: String

```

## Properties

`Actions`

The list of custom actions that are configured for a visual.

_Required_: No

_Type_: Array of [VisualCustomAction](aws-properties-quicksight-template-visualcustomaction.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChartConfiguration`

The configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartConfiguration](aws-properties-quicksight-template-gaugechartconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConditionalFormatting`

The conditional formatting of a `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartConditionalFormatting](aws-properties-quicksight-template-gaugechartconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtitle`

The subtitle that is displayed on the visual.

_Required_: No

_Type_: [VisualSubtitleLabelOptions](aws-properties-quicksight-template-visualsubtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title that is displayed on the visual.

_Required_: No

_Type_: [VisualTitleLabelOptions](aws-properties-quicksight-template-visualtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualContentAltText`

The alt text for the visual.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualId`

The unique identifier of a visual. This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartPrimaryValueConditionalFormatting

GeospatialCoordinateBounds

All content copied from https://docs.aws.amazon.com/.
