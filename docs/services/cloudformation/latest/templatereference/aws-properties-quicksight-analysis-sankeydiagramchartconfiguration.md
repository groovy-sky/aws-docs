This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SankeyDiagramChartConfiguration

The configuration of a sankey diagram.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataLabels" : DataLabelOptions,
  "FieldWells" : SankeyDiagramFieldWells,
  "Interactions" : VisualInteractionOptions,
  "SortConfiguration" : SankeyDiagramSortConfiguration
}

```

### YAML

```yaml

  DataLabels:
    DataLabelOptions
  FieldWells:
    SankeyDiagramFieldWells
  Interactions:
    VisualInteractionOptions
  SortConfiguration:
    SankeyDiagramSortConfiguration

```

## Properties

`DataLabels`

The data label configuration of a sankey diagram.

_Required_: No

_Type_: [DataLabelOptions](aws-properties-quicksight-analysis-datalabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldWells`

The field well configuration of a sankey diagram.

_Required_: No

_Type_: [SankeyDiagramFieldWells](aws-properties-quicksight-analysis-sankeydiagramfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a sankey diagram.

_Required_: No

_Type_: [SankeyDiagramSortConfiguration](aws-properties-quicksight-analysis-sankeydiagramsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SankeyDiagramAggregatedFieldWells

SankeyDiagramFieldWells

All content copied from https://docs.aws.amazon.com/.
