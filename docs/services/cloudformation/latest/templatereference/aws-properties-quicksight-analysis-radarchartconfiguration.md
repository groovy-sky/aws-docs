This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis RadarChartConfiguration

The configuration of a `RadarChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlternateBandColorsVisibility" : String,
  "AlternateBandEvenColor" : String,
  "AlternateBandOddColor" : String,
  "AxesRangeScale" : String,
  "BaseSeriesSettings" : RadarChartSeriesSettings,
  "CategoryAxis" : AxisDisplayOptions,
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "ColorAxis" : AxisDisplayOptions,
  "ColorLabelOptions" : ChartAxisLabelOptions,
  "FieldWells" : RadarChartFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "Shape" : String,
  "SortConfiguration" : RadarChartSortConfiguration,
  "StartAngle" : Number,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  AlternateBandColorsVisibility: String
  AlternateBandEvenColor: String
  AlternateBandOddColor: String
  AxesRangeScale: String
  BaseSeriesSettings:
    RadarChartSeriesSettings
  CategoryAxis:
    AxisDisplayOptions
  CategoryLabelOptions:
    ChartAxisLabelOptions
  ColorAxis:
    AxisDisplayOptions
  ColorLabelOptions:
    ChartAxisLabelOptions
  FieldWells:
    RadarChartFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  Shape: String
  SortConfiguration:
    RadarChartSortConfiguration
  StartAngle: Number
  VisualPalette:
    VisualPalette

```

## Properties

`AlternateBandColorsVisibility`

Determines the visibility of the colors of alternatign bands in a radar chart.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlternateBandEvenColor`

The color of the even-numbered alternate bands of a radar chart.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlternateBandOddColor`

The color of the odd-numbered alternate bands of a radar chart.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AxesRangeScale`

The axis behavior options of a radar chart.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | INDEPENDENT | SHARED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseSeriesSettings`

The base sreies settings of a radar chart.

_Required_: No

_Type_: [RadarChartSeriesSettings](aws-properties-quicksight-analysis-radarchartseriessettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryAxis`

The category axis of a radar chart.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-analysis-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryLabelOptions`

The category label options of a radar chart.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorAxis`

The color axis of a radar chart.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-analysis-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorLabelOptions`

The color label options of a radar chart.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a `RadarChartVisual`.

_Required_: No

_Type_: [RadarChartFieldWells](aws-properties-quicksight-analysis-radarchartfieldwells.md)

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

`Shape`

The shape of the radar chart.

_Required_: No

_Type_: String

_Allowed values_: `CIRCLE | POLYGON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a `RadarChartVisual`.

_Required_: No

_Type_: [RadarChartSortConfiguration](aws-properties-quicksight-analysis-radarchartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartAngle`

The start angle of a radar chart's axis.

_Required_: No

_Type_: Number

_Minimum_: `-360`

_Maximum_: `360`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The palette (chart color) display setup of the visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-analysis-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RadarChartAreaStyleSettings

RadarChartFieldWells

All content copied from https://docs.aws.amazon.com/.
