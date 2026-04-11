This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GaugeChartOptions

The options that determine the presentation of the `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arc" : ArcConfiguration,
  "ArcAxis" : ArcAxisConfiguration,
  "Comparison" : ComparisonConfiguration,
  "PrimaryValueDisplayType" : String,
  "PrimaryValueFontConfiguration" : FontConfiguration
}

```

### YAML

```yaml

  Arc:
    ArcConfiguration
  ArcAxis:
    ArcAxisConfiguration
  Comparison:
    ComparisonConfiguration
  PrimaryValueDisplayType: String
  PrimaryValueFontConfiguration:
    FontConfiguration

```

## Properties

`Arc`

The arc configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [ArcConfiguration](aws-properties-quicksight-dashboard-arcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ArcAxis`

The arc axis configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [ArcAxisConfiguration](aws-properties-quicksight-dashboard-arcaxisconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Comparison`

The comparison configuration of a `GaugeChartVisual`.

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartFieldWells

GaugeChartPrimaryValueConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
