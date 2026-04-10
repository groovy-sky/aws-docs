This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AggregationConstraint

Constraint on query output removing output rows that do not meet a minimum number of
distinct values of a specified column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "Minimum" : Number,
  "Type" : String
}

```

### YAML

```yaml

  ColumnName: String
  Minimum: Number
  Type: String

```

## Properties

`ColumnName`

Column in aggregation constraint for which there must be a minimum number of distinct
values in an output row for it to be in the query output.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minimum`

The minimum number of distinct values that an output row must be an aggregation of.
Minimum threshold of distinct values for a specified column that must exist in an output
row for it to be in the query output.

_Required_: Yes

_Type_: Number

_Minimum_: `2`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of aggregation the constraint allows. The only valid value is currently
\`COUNT\_DISTINCT\`.

_Required_: Yes

_Type_: String

_Allowed values_: `COUNT_DISTINCT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AggregateColumn

AnalysisRule

All content copied from https://docs.aws.amazon.com/.
