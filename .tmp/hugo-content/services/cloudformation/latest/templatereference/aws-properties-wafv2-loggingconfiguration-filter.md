This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::LoggingConfiguration Filter

A single logging filter, used in `LoggingFilter`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Behavior" : String,
  "Conditions" : [ Condition, ... ],
  "Requirement" : String
}

```

### YAML

```yaml

  Behavior: String
  Conditions:
    - Condition
  Requirement: String

```

## Properties

`Behavior`

How to handle logs that satisfy the filter's conditions and requirement.

_Required_: Yes

_Type_: String

_Allowed values_: `KEEP | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

Match conditions for the filter.

_Required_: Yes

_Type_: Array of [Condition](aws-properties-wafv2-loggingconfiguration-condition.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Requirement`

Logic to apply to the filtering conditions. You can specify that, in order to satisfy
the filter, a log must match all conditions or must match at least one condition.

_Required_: Yes

_Type_: String

_Allowed values_: `MEETS_ALL | MEETS_ANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldToMatch

LabelNameCondition

All content copied from https://docs.aws.amazon.com/.
