This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ScatterPlotConfiguration

The configuration of a scatter plot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataLabels" : DataLabelOptions,
  "FieldWells" : ScatterPlotFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "SortConfiguration" : ScatterPlotSortConfiguration,
  "Tooltip" : TooltipOptions,
  "VisualPalette" : VisualPalette,
  "XAxisDisplayOptions" : AxisDisplayOptions,
  "XAxisLabelOptions" : ChartAxisLabelOptions,
  "YAxisDisplayOptions" : AxisDisplayOptions,
  "YAxisLabelOptions" : ChartAxisLabelOptions
}

```

### YAML

```yaml

  DataLabels:
    DataLabelOptions
  FieldWells:
    ScatterPlotFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  SortConfiguration:
    ScatterPlotSortConfiguration
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
  YAxisLabelOptions:
    ChartAxisLabelOptions

```

## Properties

`DataLabels`

The options that determine if visual data labels are displayed.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-analysis-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [ScatterPlotFieldWells](aws-properties-quicksight-analysis-scatterplotfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend display setup of the visual.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-analysis-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a scatter plot.

_Required_: No

_Type_: [ScatterPlotSortConfiguration](aws-properties-quicksight-analysis-scatterplotsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The legend display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-analysis-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The palette (chart color) display setup of the visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-analysis-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisDisplayOptions`

The label display options (grid line, range, scale, and axis step) of the scatter plot's x-axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-analysis-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisLabelOptions`

The label options (label text, label visibility, and sort icon visibility) of the scatter plot's x-axis.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxisDisplayOptions`

The label display options (grid line, range, scale, and axis step) of the scatter plot's y-axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-analysis-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxisLabelOptions`

The label options (label text, label visibility, and sort icon visibility) of the scatter plot's y-axis.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScatterPlotCategoricallyAggregatedFieldWells

ScatterPlotFieldWells

All content copied from https://docs.aws.amazon.com/.
