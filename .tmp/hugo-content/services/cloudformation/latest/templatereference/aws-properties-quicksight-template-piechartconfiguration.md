This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PieChartConfiguration

The configuration of a pie chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "ContributionAnalysisDefaults" : [ ContributionAnalysisDefault, ... ],
  "DataLabels" : DataLabelOptions,
  "DonutOptions" : DonutOptions,
  "FieldWells" : PieChartFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "SmallMultiplesOptions" : SmallMultiplesOptions,
  "SortConfiguration" : PieChartSortConfiguration,
  "Tooltip" : TooltipOptions,
  "ValueLabelOptions" : ChartAxisLabelOptions,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  CategoryLabelOptions:
    ChartAxisLabelOptions
  ContributionAnalysisDefaults:
    - ContributionAnalysisDefault
  DataLabels:
    DataLabelOptions
  DonutOptions:
    DonutOptions
  FieldWells:
    PieChartFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  SmallMultiplesOptions:
    SmallMultiplesOptions
  SortConfiguration:
    PieChartSortConfiguration
  Tooltip:
    TooltipOptions
  ValueLabelOptions:
    ChartAxisLabelOptions
  VisualPalette:
    VisualPalette

```

## Properties

`CategoryLabelOptions`

The label options of the group/color that is displayed in a pie chart.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContributionAnalysisDefaults`

The contribution analysis (anomaly configuration) setup of the visual.

_Required_: No

_Type_: Array of [ContributionAnalysisDefault](aws-properties-quicksight-template-contributionanalysisdefault.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The options that determine if visual data labels are displayed.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-template-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DonutOptions`

The options that determine the shape of the chart. This option determines whether the chart is a pie chart or a donut chart.

_Required_: No

_Type_: [DonutOptions](aws-properties-quicksight-template-donutoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [PieChartFieldWells](aws-properties-quicksight-template-piechartfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-template-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend display setup of the visual.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-template-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiplesOptions`

The small multiples setup for the visual.

_Required_: No

_Type_: [SmallMultiplesOptions](aws-properties-quicksight-template-smallmultiplesoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a pie chart.

_Required_: No

_Type_: [PieChartSortConfiguration](aws-properties-quicksight-template-piechartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-template-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueLabelOptions`

The label options for the value that is displayed in a pie chart.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The palette (chart color) display setup of the visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-template-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PieChartAggregatedFieldWells

PieChartFieldWells

All content copied from https://docs.aws.amazon.com/.
