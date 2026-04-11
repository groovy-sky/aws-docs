This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis HeatMapConfiguration

The configuration of a heat map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorScale" : ColorScale,
  "ColumnLabelOptions" : ChartAxisLabelOptions,
  "DataLabels" : DataLabelOptions,
  "FieldWells" : HeatMapFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "RowLabelOptions" : ChartAxisLabelOptions,
  "SortConfiguration" : HeatMapSortConfiguration,
  "Tooltip" : TooltipOptions
}

```

### YAML

```yaml

  ColorScale:
    ColorScale
  ColumnLabelOptions:
    ChartAxisLabelOptions
  DataLabels:
    DataLabelOptions
  FieldWells:
    HeatMapFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  RowLabelOptions:
    ChartAxisLabelOptions
  SortConfiguration:
    HeatMapSortConfiguration
  Tooltip:
    TooltipOptions

```

## Properties

`ColorScale`

The color options (gradient color, point of divergence) in a heat map.

_Required_: No

_Type_: [ColorScale](aws-properties-quicksight-analysis-colorscale.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnLabelOptions`

The label options of the column that is displayed in a heat map.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The options that determine if visual data labels are displayed.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-analysis-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [HeatMapFieldWells](aws-properties-quicksight-analysis-heatmapfieldwells.md)

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

`RowLabelOptions`

The label options of the row that is displayed in a `heat map`.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-analysis-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a heat map.

_Required_: No

_Type_: [HeatMapSortConfiguration](aws-properties-quicksight-analysis-heatmapsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-analysis-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeatMapAggregatedFieldWells

HeatMapFieldWells

All content copied from https://docs.aws.amazon.com/.
