This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template AggregationSortConfiguration

The configuration options to sort aggregated values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationFunction" : AggregationFunction,
  "Column" : ColumnIdentifier,
  "SortDirection" : String
}

```

### YAML

```yaml

  AggregationFunction:
    AggregationFunction
  Column:
    ColumnIdentifier
  SortDirection: String

```

## Properties

`AggregationFunction`

The function that aggregates the values in `Column`.

_Required_: No

_Type_: [AggregationFunction](aws-properties-quicksight-template-aggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column that determines the sort order of aggregated values.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-template-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortDirection`

The sort direction of values.

- `ASC`: Sort in ascending order.

- `DESC`: Sort in descending order.

_Required_: Yes

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AggregationFunction

AnalysisDefaults

All content copied from https://docs.aws.amazon.com/.
