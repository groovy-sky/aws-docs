This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard HistogramConfiguration

The configuration for a `HistogramVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BinOptions" : HistogramBinOptions,
  "DataLabels" : DataLabelOptions,
  "FieldWells" : HistogramFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Tooltip" : TooltipOptions,
  "VisualPalette" : VisualPalette,
  "XAxisDisplayOptions" : AxisDisplayOptions,
  "XAxisLabelOptions" : ChartAxisLabelOptions,
  "YAxisDisplayOptions" : AxisDisplayOptions
}

```

### YAML

```yaml

  BinOptions:
    HistogramBinOptions
  DataLabels:
    DataLabelOptions
  FieldWells:
    HistogramFieldWells
  Interactions:
    VisualInteractionOptions
  Tooltip:
    TooltipOptions
  VisualPalette:
    VisualPalette
  XAxisDisplayOptions:
    AxisDisplayOptions
  XAxisLabelOptions:
    ChartAxisLabelOptions
  YAxisDisplayOptions:
    AxisDisplayOptions

```

## Properties

`BinOptions`

The options that determine the presentation of histogram bins.

_Required_: No

_Type_: [HistogramBinOptions](aws-properties-quicksight-dashboard-histogrambinoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The data label configuration of a histogram.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-dashboard-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a histogram.

_Required_: No

_Type_: [HistogramFieldWells](aws-properties-quicksight-dashboard-histogramfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-dashboard-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip configuration of a histogram.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-dashboard-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The visual palette configuration of a histogram.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-dashboard-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisDisplayOptions`

The options that determine the presentation of the x-axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisLabelOptions`

The options that determine the presentation of the x-axis label.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxisDisplayOptions`

The options that determine the presentation of the y-axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HistogramBinOptions

HistogramFieldWells

All content copied from https://docs.aws.amazon.com/.
