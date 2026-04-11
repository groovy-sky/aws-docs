This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic NamedEntityDefinitionMetric

A structure that represents a metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : String,
  "AggregationFunctionParameters" : {Key: Value, ...}
}

```

### YAML

```yaml

  Aggregation: String
  AggregationFunctionParameters:
    Key: Value

```

## Properties

`Aggregation`

The aggregation of a named entity. Valid values for this structure are `SUM`,
`MIN`, `MAX`, `COUNT`, `AVERAGE`,
`DISTINCT_COUNT`, `STDEV`, `STDEVP`, `VAR`,
`VARP`, `PERCENTILE`,
`MEDIAN`,
and `CUSTOM`.

_Required_: No

_Type_: String

_Allowed values_: `SUM | MIN | MAX | COUNT | AVERAGE | DISTINCT_COUNT | STDEV | STDEVP | VAR | VARP | PERCENTILE | MEDIAN | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AggregationFunctionParameters`

The additional parameters for an aggregation function.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NamedEntityDefinition

NegativeFormat

All content copied from https://docs.aws.amazon.com/.
