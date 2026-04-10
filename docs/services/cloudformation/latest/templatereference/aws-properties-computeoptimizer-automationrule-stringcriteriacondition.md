This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule StringCriteriaCondition

Criteria condition for filtering based on string values, including comparison operators and target values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparison" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Comparison: String
  Values:
    - String

```

## Properties

`Comparison`

The comparison operator used to evaluate the string criteria, such as equals, not equals, or contains.

_Required_: No

_Type_: String

_Allowed values_: `StringEquals | StringNotEquals | StringEqualsIgnoreCase | StringNotEqualsIgnoreCase | StringLike | StringNotLike | NumericEquals | NumericNotEquals | NumericLessThan | NumericLessThanEquals | NumericGreaterThan | NumericGreaterThanEquals`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

List of string values to compare against when applying the criteria condition.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Schedule

Tag

All content copied from https://docs.aws.amazon.com/.
