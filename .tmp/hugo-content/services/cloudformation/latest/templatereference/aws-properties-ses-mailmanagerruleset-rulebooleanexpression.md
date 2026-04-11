This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleBooleanExpression

A boolean expression to be used in a rule condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : RuleBooleanToEvaluate,
  "Operator" : String
}

```

### YAML

```yaml

  Evaluate:
    RuleBooleanToEvaluate
  Operator: String

```

## Properties

`Evaluate`

The operand on which to perform a boolean condition operation.

_Required_: Yes

_Type_: [RuleBooleanToEvaluate](aws-properties-ses-mailmanagerruleset-rulebooleantoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The matching operator for a boolean condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `IS_TRUE | IS_FALSE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleAction

RuleBooleanToEvaluate

All content copied from https://docs.aws.amazon.com/.
