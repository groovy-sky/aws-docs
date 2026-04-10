This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template BarChartConfiguration

The configuration of a `BarChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BarsArrangement" : String,
  "CategoryAxis" : AxisDisplayOptions,
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "ColorLabelOptions" : ChartAxisLabelOptions,
  "ContributionAnalysisDefaults" : [ ContributionAnalysisDefault, ... ],
  "DataLabels" : DataLabelOptions,
  "FieldWells" : BarChartFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "Orientation" : String,
  "ReferenceLines" : [ ReferenceLine, ... ],
  "SmallMultiplesOptions" : SmallMultiplesOptions,
  "SortConfiguration" : BarChartSortConfiguration,
  "Tooltip" : TooltipOptions,
  "ValueAxis" : AxisDisplayOptions,
  "ValueLabelOptions" : ChartAxisLabelOptions,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  BarsArrangement: String
  CategoryAxis:
    AxisDisplayOptions
  CategoryLabelOptions:
    ChartAxisLabelOptions
  ColorLabelOptions:
    ChartAxisLabelOptions
  ContributionAnalysisDefaults:
    - ContributionAnalysisDefault
  DataLabels:
    DataLabelOptions
  FieldWells:
    BarChartFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  Orientation: String
  ReferenceLines:
    - ReferenceLine
  SmallMultiplesOptions:
    SmallMultiplesOptions
  SortConfiguration:
    BarChartSortConfiguration
  Tooltip:
    TooltipOptions
  ValueAxis:
    AxisDisplayOptions
  ValueLabelOptions:
    ChartAxisLabelOptions
  VisualPalette:
    VisualPalette

```

## Properties

`BarsArrangement`

Determines the arrangement of the bars. The orientation and arrangement of bars determine the type of bar that is used in the visual.

_Required_: No

_Type_: String

_Allowed values_: `CLUSTERED | STACKED | STACKED_PERCENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryAxis`

The label display options (grid line, range, scale, axis step) for bar chart category.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-template-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryLabelOptions`

The label options (label text, label visibility and sort icon visibility) for a bar chart.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorLabelOptions`

The label options (label text, label visibility and sort icon visibility) for a color that is used in a bar chart.

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

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [BarChartFieldWells](aws-properties-quicksight-template-barchartfieldwells.md)

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

`Orientation`

The orientation of the bars in a bar chart visual. There are two valid values in this structure:

- `HORIZONTAL`: Used for charts that have horizontal bars. Visuals that use this value are horizontal bar charts, horizontal stacked bar charts, and horizontal stacked 100% bar charts.

- `VERTICAL`: Used for charts that have vertical bars. Visuals that use this value are vertical bar charts, vertical stacked bar charts, and vertical stacked 100% bar charts.

_Required_: No

_Type_: String

_Allowed values_: `HORIZONTAL | VERTICAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferenceLines`

The reference line setup of the visual.

_Required_: No

_Type_: Array of [ReferenceLine](aws-properties-quicksight-template-referenceline.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiplesOptions`

The small multiples setup for the visual.

_Required_: No

_Type_: [SmallMultiplesOptions](aws-properties-quicksight-template-smallmultiplesoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a `BarChartVisual`.

_Required_: No

_Type_: [BarChartSortConfiguration](aws-properties-quicksight-template-barchartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-template-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueAxis`

The label display options (grid line, range, scale, axis step) for a bar chart value.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-template-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueLabelOptions`

The label options (label text, label visibility and sort icon visibility) for a bar chart value.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The palette (chart color) display setup of the visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-template-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BarChartAggregatedFieldWells

BarChartFieldWells

All content copied from https://docs.aws.amazon.com/.
