This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SankeyDiagramAggregatedFieldWells

The field well configuration of a sankey diagram.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : [ DimensionField, ... ],
  "Source" : [ DimensionField, ... ],
  "Weight" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Destination:
    - DimensionField
  Source:
    - DimensionField
  Weight:
    - MeasureField

```

## Properties

`Destination`

The destination field wells of a sankey diagram.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source field wells of a sankey diagram.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The weight field wells of a sankey diagram.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SameSheetTargetVisualConfiguration

SankeyDiagramChartConfiguration

All content copied from https://docs.aws.amazon.com/.
