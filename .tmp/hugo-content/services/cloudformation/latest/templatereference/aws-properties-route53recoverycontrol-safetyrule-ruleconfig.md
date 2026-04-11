This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::SafetyRule RuleConfig

The rule configuration for an assertion rule. That is, the criteria that you set for specific assertion
controls (routing controls) that specify how many controls must be enabled after a transaction completes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Inverted" : Boolean,
  "Threshold" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  Inverted: Boolean
  Threshold: Integer
  Type: String

```

## Properties

`Inverted`

Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Threshold`

The value of N, when you specify an `ATLEAST` rule type. That is, `Threshold` is the number
of controls that must be set when you specify an `ATLEAST` type.

_Required_: Yes

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Type`

A rule can be one of the following: `ATLEAST`, `AND`, or `OR`.

_Required_: Yes

_Type_: String

_Allowed values_: `AND | OR | ATLEAST`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatingRule

Tag

All content copied from https://docs.aws.amazon.com/.
