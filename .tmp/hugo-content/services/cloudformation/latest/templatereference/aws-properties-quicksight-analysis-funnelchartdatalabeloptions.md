This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FunnelChartDataLabelOptions

The options that determine the presentation of the data labels.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryLabelVisibility" : String,
  "LabelColor" : String,
  "LabelFontConfiguration" : FontConfiguration,
  "MeasureDataLabelStyle" : String,
  "MeasureLabelVisibility" : String,
  "Position" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  CategoryLabelVisibility: String
  LabelColor: String
  LabelFontConfiguration:
    FontConfiguration
  MeasureDataLabelStyle: String
  MeasureLabelVisibility: String
  Position: String
  Visibility: String

```

## Properties

`CategoryLabelVisibility`

The visibility of the category labels within the data labels.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelColor`

The color of the data label text.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelFontConfiguration`

The font configuration for the data labels.

Only the `FontSize` attribute of the font configuration is used for data labels.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-analysis-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureDataLabelStyle`

Determines the style of the metric labels.

_Required_: No

_Type_: String

_Allowed values_: `VALUE_ONLY | PERCENTAGE_BY_FIRST_STAGE | PERCENTAGE_BY_PREVIOUS_STAGE | VALUE_AND_PERCENTAGE_BY_FIRST_STAGE | VALUE_AND_PERCENTAGE_BY_PREVIOUS_STAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureLabelVisibility`

The visibility of the measure labels within the data labels.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

Determines the positioning of the data label relative to a section of the funnel.

_Required_: No

_Type_: String

_Allowed values_: `INSIDE | OUTSIDE | LEFT | TOP | BOTTOM | RIGHT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility option that determines if data labels are displayed.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunnelChartConfiguration

FunnelChartFieldWells

All content copied from https://docs.aws.amazon.com/.
