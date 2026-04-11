This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AggregateColumn

Column in configured table that can be used in aggregate function in query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnNames" : [ String, ... ],
  "Function" : String
}

```

### YAML

```yaml

  ColumnNames:
    - String
  Function: String

```

## Properties

`ColumnNames`

Column names in configured table of aggregate columns.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Function`

Aggregation function that can be applied to aggregate column in query.

_Required_: Yes

_Type_: String

_Allowed values_: `SUM | SUM_DISTINCT | COUNT | COUNT_DISTINCT | AVG`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRooms::ConfiguredTable

AggregationConstraint

All content copied from https://docs.aws.amazon.com/.
