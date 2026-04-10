This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ColumnSort

The sort configuration for a column that is not used in a field well.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationFunction" : AggregationFunction,
  "Direction" : String,
  "SortBy" : ColumnIdentifier
}

```

### YAML

```yaml

  AggregationFunction:
    AggregationFunction
  Direction: String
  SortBy:
    ColumnIdentifier

```

## Properties

`AggregationFunction`

The aggregation function that is defined in the column sort.

_Required_: No

_Type_: [AggregationFunction](aws-properties-quicksight-analysis-aggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Direction`

The sort direction.

_Required_: Yes

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortBy`

Property description not available.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnIdentifier

ColumnTooltipItem

All content copied from https://docs.aws.amazon.com/.
