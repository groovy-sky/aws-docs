This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule DoubleCriteriaCondition

Defines a condition for filtering based on double/floating-point numeric values with comparison operators.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparison" : String,
  "Values" : [ Number, ... ]
}

```

### YAML

```yaml

  Comparison: String
  Values:
    - Number

```

## Properties

`Comparison`

The comparison operator to use, such as equals, greater than, less than, etc.

_Required_: No

_Type_: String

_Allowed values_: `StringEquals | StringNotEquals | StringEqualsIgnoreCase | StringNotEqualsIgnoreCase | StringLike | StringNotLike | NumericEquals | NumericNotEquals | NumericLessThan | NumericLessThanEquals | NumericGreaterThan | NumericGreaterThanEquals`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The list of double values to compare against using the specified comparison operator.

_Required_: No

_Type_: Array of Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Criteria

IntegerCriteriaCondition

All content copied from https://docs.aws.amazon.com/.
