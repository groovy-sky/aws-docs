This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis LineChartConfiguration

The configuration of a line chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContributionAnalysisDefaults" : [ ContributionAnalysisDefault, ... ],
  "DataLabels" : DataLabelOptions,
  "DefaultSeriesSettings" : LineChartDefaultSeriesSettings,
  "FieldWells" : LineChartFieldWells,
  "ForecastConfigurations" : [ ForecastConfiguration, ... ],
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "PrimaryYAxisDisplayOptions" : LineSeriesAxisDisplayOptions,
  "PrimaryYAxisLabelOptions" : ChartAxisLabelOptions,
  "ReferenceLines" : [ ReferenceLine, ... ],
  "SecondaryYAxisDisplayOptions" : LineSeriesAxisDisplayOptions,
  "SecondaryYAxisLabelOptions" : ChartAxisLabelOptions,
  "Series" : [ SeriesItem, ... ],
  "SingleAxisOptions" : SingleAxisOptions,
  "SmallMultiplesOptions" : SmallMultiplesOptions,
  "SortConfiguration" : LineChartSortConfiguration,
  "Tooltip" : TooltipOptions,
  "Type" : String,
  "VisualPalette" : VisualPalette,
  "XAxisDisplayOptions" : AxisDisplayOptions,
  "XAxisLabelOptions" : ChartAxisLabelOptions
}

```

### YAML

```yaml

  ContributionAnalysisDefaults:
    - ContributionAnalysisDefault
  DataLabels:
    DataLabelOptions
  DefaultSeriesSettings:
    LineChartDefaultSeriesSettings
  FieldWells:
    LineChartFieldWells
  ForecastConfigurations:
    - ForecastConfiguration
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  PrimaryYAxisDisplayOptions:
    LineSeriesAxisDisplayOptions
  PrimaryYAxisLabelOptions:
    ChartAxisLabelOptions
  ReferenceLines:
    - ReferenceLine
  SecondaryYAxisDisplayOptions:
    LineSeriesAxisDisplayOptions
  SecondaryYAxisLabelOptions:
    ChartAxisLabelOptions
  Series:
    - SeriesItem
  SingleAxisOptions:
    SingleAxisOptions
  SmallMultiplesOptions:
    SmallMultiplesOptions
  SortConfiguration:
    LineChartSortConfiguration
  Tooltip:
    TooltipOptions
  Type: String
  VisualPalette:
    VisualPalette
  XAxisDisplayOptions:
    AxisDisplayOptions
  XAxisLabelOptions:
    ChartAxisLabelOptions

```

## Properties

`ContributionAnalysisDefaults`

The default configuration of a line chart's contribution analysis.

_Required_: No

_Type_: Array of [ContributionAnalysisDefault](aws-properties-quicksight-analysis-contributionanalysisdefault.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The data label configuration of a line chart.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-analysis-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultSeriesSettings`

The options that determine the default presentation of all line series in `LineChartVisual`.

_Required_: No

_Type_: [LineChartDefaultSeriesSettings](aws-properties-quicksight-analysis-linechartdefaultseriessettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a line chart.

_Required_: No

_Type_: [LineChartFieldWells](aws-properties-quicksight-analysis-linechartfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForecastConfigurations`

The forecast configuration of a line chart.

_Required_: No

_Type_: Array of [ForecastConfiguration](aws-properties-quicksight-analysis-forecastconfiguration.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend configuration of a line chart.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-analysis-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisDisplayOptions`

The series axis configuration of a line chart.

_Required_: No

_Type_: [LineSeriesAxisDisplayOptions](aws-properties-quicksight-analysis-lineseriesaxisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisLabelOptions`

The options that determine the presentation of the y-axis label.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferenceLines`

The reference lines configuration of a line chart.

_Required_: No

_Type_: Array of [ReferenceLine](aws-properties-quicksight-analysis-referenceline.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryYAxisDisplayOptions`

The series axis configuration of a line chart.

_Required_: No

_Type_: [LineSeriesAxisDisplayOptions](aws-properties-quicksight-analysis-lineseriesaxisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryYAxisLabelOptions`

The options that determine the presentation of the secondary y-axis label.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Series`

The series item configuration of a line chart.

_Required_: No

_Type_: Array of [SeriesItem](aws-properties-quicksight-analysis-seriesitem.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleAxisOptions`

Property description not available.

_Required_: No

_Type_: [SingleAxisOptions](aws-properties-quicksight-analysis-singleaxisoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiplesOptions`

The small multiples setup for the visual.

_Required_: No

_Type_: [SmallMultiplesOptions](aws-properties-quicksight-analysis-smallmultiplesoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a line chart.

_Required_: No

_Type_: [LineChartSortConfiguration](aws-properties-quicksight-analysis-linechartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip configuration of a line chart.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-analysis-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Determines the type of the line chart.

_Required_: No

_Type_: String

_Allowed values_: `LINE | AREA | STACKED_AREA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The visual palette configuration of a line chart.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-analysis-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisDisplayOptions`

The options that determine the presentation of the x-axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-analysis-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisLabelOptions`

The options that determine the presentation of the x-axis label.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LineChartAggregatedFieldWells

LineChartDefaultSeriesSettings

All content copied from https://docs.aws.amazon.com/.
