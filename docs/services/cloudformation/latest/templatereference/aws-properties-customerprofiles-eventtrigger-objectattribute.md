This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventTrigger ObjectAttribute

The criteria that a specific object attribute must meet to trigger the
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "FieldName" : String,
  "Source" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  ComparisonOperator: String
  FieldName: String
  Source: String
  Values:
    - String

```

## Properties

`ComparisonOperator`

The operator used to compare an attribute against a list of values.

_Required_: Yes

_Type_: String

_Allowed values_: `INCLUSIVE | EXCLUSIVE | CONTAINS | BEGINS_WITH | ENDS_WITH | GREATER_THAN | LESS_THAN | GREATER_THAN_OR_EQUAL | LESS_THAN_OR_EQUAL | EQUAL | BEFORE | AFTER | ON | BETWEEN | NOT_BETWEEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldName`

A field defined within an object type.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

An attribute contained within a source object.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The amount of time of the specified unit.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventTriggerLimits

Period

All content copied from https://docs.aws.amazon.com/.
