This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template LineChartDefaultSeriesSettings

The options that determine the default presentation of all line series in `LineChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisBinding" : String,
  "LineStyleSettings" : LineChartLineStyleSettings,
  "MarkerStyleSettings" : LineChartMarkerStyleSettings
}

```

### YAML

```yaml

  AxisBinding: String
  LineStyleSettings:
    LineChartLineStyleSettings
  MarkerStyleSettings:
    LineChartMarkerStyleSettings

```

## Properties

`AxisBinding`

The axis to which you are binding all line series to.

_Required_: No

_Type_: String

_Allowed values_: `PRIMARY_YAXIS | SECONDARY_YAXIS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineStyleSettings`

Line styles options for all line series in the visual.

_Required_: No

_Type_: [LineChartLineStyleSettings](aws-properties-quicksight-template-linechartlinestylesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MarkerStyleSettings`

Marker styles options for all line series in the visual.

_Required_: No

_Type_: [LineChartMarkerStyleSettings](aws-properties-quicksight-template-linechartmarkerstylesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LineChartConfiguration

LineChartFieldWells

All content copied from https://docs.aws.amazon.com/.
