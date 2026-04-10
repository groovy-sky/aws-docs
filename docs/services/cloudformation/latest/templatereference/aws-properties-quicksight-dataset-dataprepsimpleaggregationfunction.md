This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataPrepSimpleAggregationFunction

A simple aggregation function that performs standard statistical operations on a column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FunctionType" : String,
  "InputColumnName" : String
}

```

### YAML

```yaml

  FunctionType: String
  InputColumnName: String

```

## Properties

`FunctionType`

The type of aggregation function to perform, such as `COUNT`, `SUM`, `AVERAGE`,
`MIN`, `MAX`, `MEDIAN`, `VARIANCE`, or `STANDARD_DEVIATION`.

_Required_: Yes

_Type_: String

_Allowed values_: `COUNT | DISTINCT_COUNT | SUM | AVERAGE | MEDIAN | MAX | MIN | VARIANCE | STANDARD_DEVIATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputColumnName`

The name of the column on which to perform the aggregation function.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPrepPercentileAggregationFunction

DataSetColumnIdMapping

All content copied from https://docs.aws.amazon.com/.
