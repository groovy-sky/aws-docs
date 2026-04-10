This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard BoxPlotChartConfiguration

The configuration of a `BoxPlotVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BoxPlotOptions" : BoxPlotOptions,
  "CategoryAxis" : AxisDisplayOptions,
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "FieldWells" : BoxPlotFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "PrimaryYAxisDisplayOptions" : AxisDisplayOptions,
  "PrimaryYAxisLabelOptions" : ChartAxisLabelOptions,
  "ReferenceLines" : [ ReferenceLine, ... ],
  "SortConfiguration" : BoxPlotSortConfiguration,
  "Tooltip" : TooltipOptions,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  BoxPlotOptions:
    BoxPlotOptions
  CategoryAxis:
    AxisDisplayOptions
  CategoryLabelOptions:
    ChartAxisLabelOptions
  FieldWells:
    BoxPlotFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  PrimaryYAxisDisplayOptions:
    AxisDisplayOptions
  PrimaryYAxisLabelOptions:
    ChartAxisLabelOptions
  ReferenceLines:
    - ReferenceLine
  SortConfiguration:
    BoxPlotSortConfiguration
  Tooltip:
    TooltipOptions
  VisualPalette:
    VisualPalette

```

## Properties

`BoxPlotOptions`

The box plot chart options for a box plot visual

_Required_: No

_Type_: [BoxPlotOptions](aws-properties-quicksight-dashboard-boxplotoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryAxis`

The label display options (grid line, range, scale, axis step) of a box plot category.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryLabelOptions`

The label options (label text, label visibility and sort Icon visibility) of a box plot category.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-dashboard-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [BoxPlotFieldWells](aws-properties-quicksight-dashboard-boxplotfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-dashboard-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

Property description not available.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-dashboard-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisDisplayOptions`

The label display options (grid line, range, scale, axis step) of a box plot category.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-dashboard-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryYAxisLabelOptions`

The label options (label text, label visibility and sort icon visibility) of a box plot value.

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

`SortConfiguration`

The sort configuration of a `BoxPlotVisual`.

_Required_: No

_Type_: [BoxPlotSortConfiguration](aws-properties-quicksight-dashboard-boxplotsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-dashboard-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The palette (chart color) display setup of the visual.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-dashboard-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BoxPlotAggregatedFieldWells

BoxPlotFieldWells

All content copied from https://docs.aws.amazon.com/.
