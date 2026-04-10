This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard AggregationFunction

An aggregation function aggregates values from a dimension or measure.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeAggregationFunction" : AttributeAggregationFunction,
  "CategoricalAggregationFunction" : String,
  "DateAggregationFunction" : String,
  "NumericalAggregationFunction" : NumericalAggregationFunction
}

```

### YAML

```yaml

  AttributeAggregationFunction:
    AttributeAggregationFunction
  CategoricalAggregationFunction: String
  DateAggregationFunction: String
  NumericalAggregationFunction:
    NumericalAggregationFunction

```

## Properties

`AttributeAggregationFunction`

Aggregation for attributes.

_Required_: No

_Type_: [AttributeAggregationFunction](aws-properties-quicksight-dashboard-attributeaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoricalAggregationFunction`

Aggregation for categorical values.

- `COUNT`: Aggregate by the total number of values, including duplicates.

- `DISTINCT_COUNT`: Aggregate by the total number of distinct values.

_Required_: No

_Type_: String

_Allowed values_: `COUNT | DISTINCT_COUNT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateAggregationFunction`

Aggregation for date values.

- `COUNT`: Aggregate by the total number of values, including duplicates.

- `DISTINCT_COUNT`: Aggregate by the total number of distinct values.

- `MIN`: Select the smallest date value.

- `MAX`: Select the largest date value.

_Required_: No

_Type_: String

_Allowed values_: `COUNT | DISTINCT_COUNT | MIN | MAX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericalAggregationFunction`

Aggregation for numerical values.

_Required_: No

_Type_: [NumericalAggregationFunction](aws-properties-quicksight-dashboard-numericalaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdHocFilteringOption

AggregationSortConfiguration

All content copied from https://docs.aws.amazon.com/.
