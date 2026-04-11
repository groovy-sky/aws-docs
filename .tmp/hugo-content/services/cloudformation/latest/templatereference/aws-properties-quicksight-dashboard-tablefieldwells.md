This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableFieldWells

The field wells for a table visual.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TableAggregatedFieldWells" : TableAggregatedFieldWells,
  "TableUnaggregatedFieldWells" : TableUnaggregatedFieldWells
}

```

### YAML

```yaml

  TableAggregatedFieldWells:
    TableAggregatedFieldWells
  TableUnaggregatedFieldWells:
    TableUnaggregatedFieldWells

```

## Properties

`TableAggregatedFieldWells`

The aggregated field well for the table.

_Required_: No

_Type_: [TableAggregatedFieldWells](aws-properties-quicksight-dashboard-tableaggregatedfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableUnaggregatedFieldWells`

The unaggregated field well for the table.

_Required_: No

_Type_: [TableUnaggregatedFieldWells](aws-properties-quicksight-dashboard-tableunaggregatedfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldURLConfiguration

TableInlineVisualization

All content copied from https://docs.aws.amazon.com/.
