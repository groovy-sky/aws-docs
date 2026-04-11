This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TreeMapConfiguration

The configuration of a tree map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorLabelOptions" : ChartAxisLabelOptions,
  "ColorScale" : ColorScale,
  "DataLabels" : DataLabelOptions,
  "FieldWells" : TreeMapFieldWells,
  "GroupLabelOptions" : ChartAxisLabelOptions,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "SizeLabelOptions" : ChartAxisLabelOptions,
  "SortConfiguration" : TreeMapSortConfiguration,
  "Tooltip" : TooltipOptions
}

```

### YAML

```yaml

  ColorLabelOptions:
    ChartAxisLabelOptions
  ColorScale:
    ColorScale
  DataLabels:
    DataLabelOptions
  FieldWells:
    TreeMapFieldWells
  GroupLabelOptions:
    ChartAxisLabelOptions
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  SizeLabelOptions:
    ChartAxisLabelOptions
  SortConfiguration:
    TreeMapSortConfiguration
  Tooltip:
    TooltipOptions

```

## Properties

`ColorLabelOptions`

The label options (label text, label visibility) for the colors displayed in a tree map.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorScale`

The color options (gradient color, point of divergence) of a tree map.

_Required_: No

_Type_: [ColorScale](aws-properties-quicksight-template-colorscale.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabels`

The options that determine if visual data labels are displayed.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-template-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [TreeMapFieldWells](aws-properties-quicksight-template-treemapfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupLabelOptions`

The label options (label text, label visibility) of the groups that are displayed in a tree map.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

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

`SizeLabelOptions`

The label options (label text, label visibility) of the sizes that are displayed in a tree map.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a tree map.

_Required_: No

_Type_: [TreeMapSortConfiguration](aws-properties-quicksight-template-treemapsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-template-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TreeMapAggregatedFieldWells

TreeMapFieldWells

All content copied from https://docs.aws.amazon.com/.
