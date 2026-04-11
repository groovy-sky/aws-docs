This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GaugeChartConfiguration

The configuration of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorConfiguration" : GaugeChartColorConfiguration,
  "DataLabels" : DataLabelOptions,
  "FieldWells" : GaugeChartFieldWells,
  "GaugeChartOptions" : GaugeChartOptions,
  "Interactions" : VisualInteractionOptions,
  "TooltipOptions" : TooltipOptions,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  ColorConfiguration:
    GaugeChartColorConfiguration
  DataLabels:
    DataLabelOptions
  FieldWells:
    GaugeChartFieldWells
  GaugeChartOptions:
    GaugeChartOptions
  Interactions:
    VisualInteractionOptions
  TooltipOptions:
    TooltipOptions
  VisualPalette:
    VisualPalette

```

## Properties

`ColorConfiguration`

The color configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartColorConfiguration](aws-properties-quicksight-analysis-gaugechartcolorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The data label configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-analysis-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartFieldWells](aws-properties-quicksight-analysis-gaugechartfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GaugeChartOptions`

The options that determine the presentation of the `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartOptions](aws-properties-quicksight-analysis-gaugechartoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipOptions`

The tooltip configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-analysis-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The visual palette configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-analysis-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartConditionalFormattingOption

GaugeChartFieldWells

All content copied from https://docs.aws.amazon.com/.
