This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule ResourceTagsCriteriaCondition

Criteria condition for filtering resources based on their tags, including comparison operators and values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparison" : String,
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Comparison: String
  Key: String
  Values:
    - String

```

## Properties

`Comparison`

The comparison operator used to evaluate the tag criteria, such as equals, not equals, or contains.

_Required_: No

_Type_: String

_Allowed values_: `StringEquals | StringNotEquals | StringEqualsIgnoreCase | StringNotEqualsIgnoreCase | StringLike | StringNotLike | NumericEquals | NumericNotEquals | NumericLessThan | NumericLessThanEquals | NumericGreaterThan | NumericGreaterThanEquals`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The tag key to use for comparison when filtering resources.

_Required_: No

_Type_: String

_Pattern_: `^[\w\s\.\-\:\/\=\+\@\*\?]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

List of tag values to compare against when filtering resources.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrganizationConfiguration

Schedule

All content copied from https://docs.aws.amazon.com/.
