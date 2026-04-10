This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AnalysisRuleAggregation

A type of analysis rule that enables query structure and specified queries that produce
aggregate statistics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalAnalyses" : String,
  "AggregateColumns" : [ AggregateColumn, ... ],
  "AllowedJoinOperators" : [ String, ... ],
  "DimensionColumns" : [ String, ... ],
  "JoinColumns" : [ String, ... ],
  "JoinRequired" : String,
  "OutputConstraints" : [ AggregationConstraint, ... ],
  "ScalarFunctions" : [ String, ... ]
}

```

### YAML

```yaml

  AdditionalAnalyses: String
  AggregateColumns:
    - AggregateColumn
  AllowedJoinOperators:
    - String
  DimensionColumns:
    - String
  JoinColumns:
    - String
  JoinRequired: String
  OutputConstraints:
    - AggregationConstraint
  ScalarFunctions:
    - String

```

## Properties

`AdditionalAnalyses`

An indicator as to whether additional analyses (such as AWS Clean Rooms ML) can be applied to the output of the direct query.

The `additionalAnalyses` parameter is currently supported for the list
analysis rule ( `AnalysisRuleList`) and the custom analysis rule
( `AnalysisRuleCustom`).

_Required_: No

_Type_: String

_Allowed values_: `ALLOWED | REQUIRED | NOT_ALLOWED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AggregateColumns`

The columns that query runners are allowed to use in aggregation queries.

_Required_: Yes

_Type_: Array of [AggregateColumn](aws-properties-cleanrooms-configuredtable-aggregatecolumn.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedJoinOperators`

Which logical operators (if any) are to be used in an INNER JOIN match condition.
Default is `AND`.

_Required_: No

_Type_: Array of String

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionColumns`

The columns that query runners are allowed to select, group by, or filter by.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinColumns`

Columns in configured table that can be used in join statements and/or as aggregate
columns. They can never be outputted directly.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinRequired`

Control that requires member who runs query to do a join with their configured table
and/or other configured table in query.

_Required_: No

_Type_: String

_Allowed values_: `QUERY_RUNNER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputConstraints`

Columns that must meet a specific threshold value (after an aggregation function is
applied to it) for each output row to be returned.

_Required_: Yes

_Type_: Array of [AggregationConstraint](aws-properties-cleanrooms-configuredtable-aggregationconstraint.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalarFunctions`

Set of scalar functions that are allowed to be used on dimension columns and the output
of aggregation of metrics.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisRule

AnalysisRuleCustom

All content copied from https://docs.aws.amazon.com/.
