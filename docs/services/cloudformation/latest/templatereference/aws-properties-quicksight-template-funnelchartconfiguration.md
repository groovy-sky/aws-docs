This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FunnelChartConfiguration

The configuration of a `FunnelChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryLabelOptions" : ChartAxisLabelOptions,
  "DataLabelOptions" : FunnelChartDataLabelOptions,
  "FieldWells" : FunnelChartFieldWells,
  "Interactions" : VisualInteractionOptions,
  "SortConfiguration" : FunnelChartSortConfiguration,
  "Tooltip" : TooltipOptions,
  "ValueLabelOptions" : ChartAxisLabelOptions,
  "VisualPalette" : VisualPalette
}

```

### YAML

```yaml

  CategoryLabelOptions:
    ChartAxisLabelOptions
  DataLabelOptions:
    FunnelChartDataLabelOptions
  FieldWells:
    FunnelChartFieldWells
  Interactions:
    VisualInteractionOptions
  SortConfiguration:
    FunnelChartSortConfiguration
  Tooltip:
    TooltipOptions
  ValueLabelOptions:
    ChartAxisLabelOptions
  VisualPalette:
    VisualPalette

```

## Properties

`CategoryLabelOptions`

The label options of the categories that are displayed in a `FunnelChartVisual`.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabelOptions`

The options that determine the presentation of the data labels.

_Required_: No

_Type_: [FunnelChartDataLabelOptions](aws-properties-quicksight-template-funnelchartdatalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a `FunnelChartVisual`.

_Required_: No

_Type_: [FunnelChartFieldWells](aws-properties-quicksight-template-funnelchartfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-template-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a `FunnelChartVisual`.

_Required_: No

_Type_: [FunnelChartSortConfiguration](aws-properties-quicksight-template-funnelchartsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip configuration of a `FunnelChartVisual`.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-template-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueLabelOptions`

The label options for the values that are displayed in a `FunnelChartVisual`.

_Required_: No

_Type_: [ChartAxisLabelOptions](aws-properties-quicksight-template-chartaxislabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

The visual palette configuration of a `FunnelChartVisual`.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-template-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunnelChartAggregatedFieldWells

FunnelChartDataLabelOptions

All content copied from https://docs.aws.amazon.com/.
