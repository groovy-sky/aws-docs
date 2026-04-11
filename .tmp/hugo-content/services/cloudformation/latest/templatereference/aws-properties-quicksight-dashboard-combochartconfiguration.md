This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ComboChartConfiguration

The configuration of a `ComboChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BarDataLabels" : DataLabelOptions,
  "BarsArrangement" : String,
  "CategoryAxis" : AxisDisplayOptions,
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "ColorLabelOptions" : ChartAxisLabelOptions,
  "FieldWells" : ComboChartFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "LineDataLabels" : DataLabelOptions,
  "PrimaryYAxisDisplayOptions" : AxisDisplayOptions,
  "PrimaryYAxisLabelOptions" : ChartAxisLabelOptions,
  "ReferenceLines" : [ ReferenceLine, ... ],
  "SecondaryYAxisDisplayOptions" : AxisDisplayOptions,
  "SecondaryYAxisLabelOptions" : ChartAxisLabelOptions,
  "SingleAxisOptions" : SingleAxisOptions,
  "SortConfiguration" : ComboChartSortConfiguration,
  "Tooltip" : TooltipOptions,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  BarDataLabels:
    DataLabelOptions
  BarsArrangement: String
  CategoryAxis:
    AxisDisplayOptions
  CategoryLabelOptions:
    ChartAxisLabelOptions
  ColorLabelOptions:
    ChartAxisLabelOptions
  FieldWells:
    ComboChartFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  LineDataLabels:
    DataLabelOptions
  PrimaryYAxisDisplayOptions:
    AxisDisplayOptions
  PrimaryYAxisLabelOptions:
    ChartAxisLabelOptions
  ReferenceLines:
    - ReferenceLine
  SecondaryYAxisDisplayOptions:
    AxisDisplayOptions
  SecondaryYAxisLabelOptions:
    ChartAxisLabelOptions
  SingleAxisOptions:
    SingleAxisOptions
  SortConfiguration:
    ComboChartSortConfiguration
  Tooltip:
    TooltipOptions
  VisualPalette:
    VisualPalette

```

## Properties

`BarDataLabels`

The options that determine if visual data labels are displayed.

The data label options for a bar in a combo chart.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-dashboard-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BarsArrangement`

Determines the bar arrangement in a combo chart. The following are valid values in this structure:

- `CLUSTERED`: For clustered bar combo charts.

- `STACKED`: For stacked bar combo charts.

- `STACKED_PERCENT`: Do not use. If you use this value, the operation returns a validation error.

_Required_: No

_Type_: String

_Allowed values_: `CLUSTERED | STACKED | STACKED_PERCENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryAxis`

The category axis of a combo chart.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryLabelOptions`

The label options (label text, label visibility, and sort icon visibility) of a combo chart category (group/color) field well.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorLabelOptions`

The label options (label text, label visibility, and sort icon visibility) of a combo chart's color field well.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [ComboChartFieldWells](aws-properties-quicksight-dashboard-combochartfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-dashboard-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend display setup of the visual.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-dashboard-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineDataLabels`

The options that determine if visual data labels are displayed.

The data label options for a line in a combo chart.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-dashboard-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisDisplayOptions`

The label display options (grid line, range, scale, and axis step) of a combo chart's primary y-axis (bar) field well.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisLabelOptions`

The label options (label text, label visibility, and sort icon visibility) of a combo chart's primary y-axis (bar) field well.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferenceLines`

The reference line setup of the visual.

_Required_: No

_Type_: Array of [ReferenceLine](aws-properties-quicksight-dashboard-referenceline.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryYAxisDisplayOptions`

The label display options (grid line, range, scale, axis step) of a combo chart's secondary y-axis (line) field well.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryYAxisLabelOptions`

The label options (label text, label visibility, and sort icon visibility) of a combo chart's secondary y-axis(line) field well.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleAxisOptions`

Property description not available.

_Required_: No

_Type_: [SingleAxisOptions](aws-properties-quicksight-dashboard-singleaxisoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a `ComboChartVisual`.

_Required_: No

_Type_: [ComboChartSortConfiguration](aws-properties-quicksight-dashboard-combochartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The legend display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-dashboard-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The palette (chart color) display setup of the visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-dashboard-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComboChartAggregatedFieldWells

ComboChartFieldWells

All content copied from https://docs.aws.amazon.com/.
