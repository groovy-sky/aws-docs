This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis LineChartSeriesSettings

The options that determine the presentation of a line series in the visual

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LineStyleSettings" : LineChartLineStyleSettings,
  "MarkerStyleSettings" : LineChartMarkerStyleSettings
}

```

### YAML

```yaml

  LineStyleSettings:
    LineChartLineStyleSettings
  MarkerStyleSettings:
    LineChartMarkerStyleSettings

```

## Properties

`LineStyleSettings`

Line styles options for a line series in `LineChartVisual`.

_Required_: No

_Type_: [LineChartLineStyleSettings](aws-properties-quicksight-analysis-linechartlinestylesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MarkerStyleSettings`

Marker styles options for a line series in `LineChartVisual`.

_Required_: No

_Type_: [LineChartMarkerStyleSettings](aws-properties-quicksight-analysis-linechartmarkerstylesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LineChartMarkerStyleSettings

LineChartSortConfiguration

All content copied from https://docs.aws.amazon.com/.
