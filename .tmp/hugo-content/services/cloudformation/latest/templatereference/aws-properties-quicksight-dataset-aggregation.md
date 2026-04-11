This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet Aggregation

Defines an aggregation function to be applied to grouped data, creating a new column with
the calculated result.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationFunction" : DataPrepAggregationFunction,
  "NewColumnId" : String,
  "NewColumnName" : String
}

```

### YAML

```yaml

  AggregationFunction:
    DataPrepAggregationFunction
  NewColumnId: String
  NewColumnName: String

```

## Properties

`AggregationFunction`

The aggregation function to apply, such as `SUM`, `COUNT`, `AVERAGE`,
`MIN`, `MAX`

_Required_: Yes

_Type_: [DataPrepAggregationFunction](aws-properties-quicksight-dataset-dataprepaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewColumnId`

A unique identifier for the new column that will contain the aggregated values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewColumnName`

The name for the new column that will contain the aggregated values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AggregateOperation

AppendedColumn

All content copied from https://docs.aws.amazon.com/.
