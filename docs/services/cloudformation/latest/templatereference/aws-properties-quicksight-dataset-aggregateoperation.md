This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet AggregateOperation

A transform operation that groups rows by specified columns and applies aggregation functions
to calculate summary values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregations" : [ Aggregation, ... ],
  "Alias" : String,
  "GroupByColumnNames" : [ String, ... ],
  "Source" : TransformOperationSource
}

```

### YAML

```yaml

  Aggregations:
    - Aggregation
  Alias: String
  GroupByColumnNames:
    - String
  Source:
    TransformOperationSource

```

## Properties

`Aggregations`

The list of aggregation functions to apply to the grouped data, such as `SUM`,
`COUNT`, or `AVERAGE`.

_Required_: Yes

_Type_: Array of [Aggregation](aws-properties-quicksight-dataset-aggregation.md)

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Alias`

Alias for this operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupByColumnNames`

The list of column names to group by when performing the aggregation. Rows with the same values in
these columns will be grouped together.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `127 | 128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for the aggregation.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QuickSight::DataSet

Aggregation

All content copied from https://docs.aws.amazon.com/.
