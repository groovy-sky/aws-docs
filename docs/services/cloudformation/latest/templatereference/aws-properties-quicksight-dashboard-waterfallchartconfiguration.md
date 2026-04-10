This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard WaterfallChartConfiguration

The configuration for a waterfall visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryAxisDisplayOptions" : AxisDisplayOptions,
  "CategoryAxisLabelOptions" : ChartAxisLabelOptions,
  "ColorConfiguration" : WaterfallChartColorConfiguration,
  "DataLabels" : DataLabelOptions,
  "FieldWells" : WaterfallChartFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "PrimaryYAxisDisplayOptions" : AxisDisplayOptions,
  "PrimaryYAxisLabelOptions" : ChartAxisLabelOptions,
  "SortConfiguration" : WaterfallChartSortConfiguration,
  "VisualPalette" : VisualPalette,
  "WaterfallChartOptions" : WaterfallChartOptions
}

```

### YAML

```yaml

  CategoryAxisDisplayOptions:
    AxisDisplayOptions
  CategoryAxisLabelOptions:
    ChartAxisLabelOptions
  ColorConfiguration:
    WaterfallChartColorConfiguration
  DataLabels:
    DataLabelOptions
  FieldWells:
    WaterfallChartFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  PrimaryYAxisDisplayOptions:
    AxisDisplayOptions
  PrimaryYAxisLabelOptions:
    ChartAxisLabelOptions
  SortConfiguration:
    WaterfallChartSortConfiguration
  VisualPalette:
    VisualPalette
  WaterfallChartOptions:
    WaterfallChartOptions

```

## Properties

`CategoryAxisDisplayOptions`

The options that determine the presentation of the category axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryAxisLabelOptions`

The options that determine the presentation of the category axis label.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorConfiguration`

The color configuration of a waterfall visual.

_Required_: No

_Type_: [WaterfallChartColorConfiguration](aws-properties-quicksight-dashboard-waterfallchartcolorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The data label configuration of a waterfall visual.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-dashboard-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a waterfall visual.

_Required_: No

_Type_: [WaterfallChartFieldWells](aws-properties-quicksight-dashboard-waterfallchartfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-dashboard-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend configuration of a waterfall visual.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-dashboard-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisDisplayOptions`

The options that determine the presentation of the y-axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisLabelOptions`

The options that determine the presentation of the y-axis label.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a waterfall visual.

_Required_: No

_Type_: [WaterfallChartSortConfiguration](aws-properties-quicksight-dashboard-waterfallchartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The visual palette configuration of a waterfall visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-dashboard-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaterfallChartOptions`

The options that determine the presentation of a waterfall visual.

_Required_: No

_Type_: [WaterfallChartOptions](aws-properties-quicksight-dashboard-waterfallchartoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WaterfallChartColorConfiguration

WaterfallChartFieldWells

All content copied from https://docs.aws.amazon.com/.
