This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template NumericalAggregationFunction

Aggregation for numerical values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PercentileAggregation" : PercentileAggregation,
  "SimpleNumericalAggregation" : String
}

```

### YAML

```yaml

  PercentileAggregation:
    PercentileAggregation
  SimpleNumericalAggregation: String

```

## Properties

`PercentileAggregation`

An aggregation based on the percentile of values in a dimension or measure.

_Required_: No

_Type_: [PercentileAggregation](aws-properties-quicksight-template-percentileaggregation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SimpleNumericalAggregation`

Built-in aggregation functions for numerical values.

- `SUM`: The sum of a dimension or measure.

- `AVERAGE`: The average of a dimension or measure.

- `MIN`: The minimum value of a dimension or measure.

- `MAX`: The maximum value of a dimension or measure.

- `COUNT`: The count of a dimension or measure.

- `DISTINCT_COUNT`: The count of distinct values in a dimension or measure.

- `VAR`: The variance of a dimension or measure.

- `VARP`: The partitioned variance of a dimension or measure.

- `STDEV`: The standard deviation of a dimension or measure.

- `STDEVP`: The partitioned standard deviation of a dimension or measure.

- `MEDIAN`: The median value of a dimension or measure.

_Required_: No

_Type_: String

_Allowed values_: `SUM | AVERAGE | MIN | MAX | COUNT | DISTINCT_COUNT | VAR | VARP | STDEV | STDEVP | MEDIAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumberFormatConfiguration

NumericalDimensionField

All content copied from https://docs.aws.amazon.com/.
