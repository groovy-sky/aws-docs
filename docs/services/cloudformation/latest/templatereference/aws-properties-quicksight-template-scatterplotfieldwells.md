This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ScatterPlotFieldWells

The field well configuration of a scatter plot.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScatterPlotCategoricallyAggregatedFieldWells" : ScatterPlotCategoricallyAggregatedFieldWells,
  "ScatterPlotUnaggregatedFieldWells" : ScatterPlotUnaggregatedFieldWells
}

```

### YAML

```yaml

  ScatterPlotCategoricallyAggregatedFieldWells:
    ScatterPlotCategoricallyAggregatedFieldWells
  ScatterPlotUnaggregatedFieldWells:
    ScatterPlotUnaggregatedFieldWells

```

## Properties

`ScatterPlotCategoricallyAggregatedFieldWells`

The aggregated field wells of a scatter plot. The x and y-axes of scatter plots with aggregated field wells are aggregated by category, label, or both.

_Required_: No

_Type_: [ScatterPlotCategoricallyAggregatedFieldWells](aws-properties-quicksight-template-scatterplotcategoricallyaggregatedfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScatterPlotUnaggregatedFieldWells`

The unaggregated field wells of a scatter plot. The x and y-axes of these scatter plots are
unaggregated.

_Required_: No

_Type_: [ScatterPlotUnaggregatedFieldWells](aws-properties-quicksight-template-scatterplotunaggregatedfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScatterPlotConfiguration

ScatterPlotSortConfiguration

All content copied from https://docs.aws.amazon.com/.
