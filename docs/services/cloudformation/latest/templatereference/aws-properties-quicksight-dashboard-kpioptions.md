This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard KPIOptions

The options that determine the presentation of a KPI visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparison" : ComparisonConfiguration,
  "PrimaryValueDisplayType" : String,
  "PrimaryValueFontConfiguration" : FontConfiguration,
  "ProgressBar" : ProgressBarOptions,
  "SecondaryValue" : SecondaryValueOptions,
  "SecondaryValueFontConfiguration" : FontConfiguration,
  "Sparkline" : KPISparklineOptions,
  "TrendArrows" : TrendArrowOptions,
  "VisualLayoutOptions" : KPIVisualLayoutOptions
}

```

### YAML

```yaml

  Comparison:
    ComparisonConfiguration
  PrimaryValueDisplayType: String
  PrimaryValueFontConfiguration:
    FontConfiguration
  ProgressBar:
    ProgressBarOptions
  SecondaryValue:
    SecondaryValueOptions
  SecondaryValueFontConfiguration:
    FontConfiguration
  Sparkline:
    KPISparklineOptions
  TrendArrows:
    TrendArrowOptions
  VisualLayoutOptions:
    KPIVisualLayoutOptions

```

## Properties

`Comparison`

The comparison configuration of a KPI visual.

_Required_: No

_Type_: [ComparisonConfiguration](aws-properties-quicksight-dashboard-comparisonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryValueDisplayType`

The options that determine the primary value display type.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | COMPARISON | ACTUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryValueFontConfiguration`

The options that determine the primary value font configuration.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-dashboard-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgressBar`

The options that determine the presentation of the progress bar of a KPI visual.

_Required_: No

_Type_: [ProgressBarOptions](aws-properties-quicksight-dashboard-progressbaroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryValue`

The options that determine the presentation of the secondary value of a KPI visual.

_Required_: No

_Type_: [SecondaryValueOptions](aws-properties-quicksight-dashboard-secondaryvalueoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryValueFontConfiguration`

The options that determine the secondary value font configuration.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-dashboard-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sparkline`

The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.

_Required_: No

_Type_: [KPISparklineOptions](aws-properties-quicksight-dashboard-kpisparklineoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrendArrows`

The options that determine the presentation of trend arrows in a KPI visual.

_Required_: No

_Type_: [TrendArrowOptions](aws-properties-quicksight-dashboard-trendarrowoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualLayoutOptions`

The options that determine the layout a KPI visual.

_Required_: No

_Type_: [KPIVisualLayoutOptions](aws-properties-quicksight-dashboard-kpivisuallayoutoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KPIFieldWells

KPIPrimaryValueConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
