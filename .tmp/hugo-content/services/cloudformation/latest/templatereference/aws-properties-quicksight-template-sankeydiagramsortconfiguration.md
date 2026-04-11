This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SankeyDiagramSortConfiguration

The sort configuration of a sankey diagram.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationItemsLimit" : ItemsLimitConfiguration,
  "SourceItemsLimit" : ItemsLimitConfiguration,
  "WeightSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  DestinationItemsLimit:
    ItemsLimitConfiguration
  SourceItemsLimit:
    ItemsLimitConfiguration
  WeightSort:
    - FieldSortOptions

```

## Properties

`DestinationItemsLimit`

The limit on the number of destination nodes that are displayed in a sankey diagram.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-template-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceItemsLimit`

The limit on the number of source nodes that are displayed in a sankey diagram.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-template-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeightSort`

The sort configuration of the weight fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-template-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SankeyDiagramFieldWells

SankeyDiagramVisual

All content copied from https://docs.aws.amazon.com/.
