This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PivotTableVisual

A pivot table.

For more information, see [Using pivot tables](../../../quicksight/latest/user/pivot-table.md) in the _Amazon Quick Suite User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ VisualCustomAction, ... ],
  "ChartConfiguration" : PivotTableConfiguration,
  "ConditionalFormatting" : PivotTableConditionalFormatting,
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
    PivotTableConfiguration
  ConditionalFormatting:
    PivotTableConditionalFormatting
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

_Type_: Array of [VisualCustomAction](aws-properties-quicksight-analysis-visualcustomaction.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChartConfiguration`

The configuration settings of the visual.

_Required_: No

_Type_: [PivotTableConfiguration](aws-properties-quicksight-analysis-pivottableconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConditionalFormatting`

The conditional formatting for a `PivotTableVisual`.

_Required_: No

_Type_: [PivotTableConditionalFormatting](aws-properties-quicksight-analysis-pivottableconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtitle`

The subtitle that is displayed on the visual.

_Required_: No

_Type_: [VisualSubtitleLabelOptions](aws-properties-quicksight-analysis-visualsubtitlelabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title that is displayed on the visual.

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

The unique identifier of a visual. This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers..

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableTotalOptions

PivotTotalOptions

All content copied from https://docs.aws.amazon.com/.
