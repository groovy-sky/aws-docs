This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataPrepAggregationFunction

Defines the type of aggregation function to apply to data during data preparation, supporting simple
and list aggregations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ListAggregation" : DataPrepListAggregationFunction,
  "PercentileAggregation" : DataPrepPercentileAggregationFunction,
  "SimpleAggregation" : DataPrepSimpleAggregationFunction
}

```

### YAML

```yaml

  ListAggregation:
    DataPrepListAggregationFunction
  PercentileAggregation:
    DataPrepPercentileAggregationFunction
  SimpleAggregation:
    DataPrepSimpleAggregationFunction

```

## Properties

`ListAggregation`

A list aggregation function that concatenates values from multiple rows into a single delimited string.

_Required_: No

_Type_: [DataPrepListAggregationFunction](aws-properties-quicksight-dataset-datapreplistaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PercentileAggregation`

Property description not available.

_Required_: No

_Type_: [DataPrepPercentileAggregationFunction](aws-properties-quicksight-dataset-datapreppercentileaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SimpleAggregation`

A simple aggregation function such as `SUM`, `COUNT`, `AVERAGE`,
`MIN`, `MAX`, `MEDIAN`, `VARIANCE`, or `STANDARD_DEVIATION`.

_Required_: No

_Type_: [DataPrepSimpleAggregationFunction](aws-properties-quicksight-dataset-dataprepsimpleaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomSql

DataPrepConfiguration

All content copied from https://docs.aws.amazon.com/.
