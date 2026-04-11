This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleNumberExpression

A number expression to match numeric conditions with integers from the incoming
email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : RuleNumberToEvaluate,
  "Operator" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Evaluate:
    RuleNumberToEvaluate
  Operator: String
  Value: Number

```

## Properties

`Evaluate`

The number to evaluate in a numeric condition expression.

_Required_: Yes

_Type_: [RuleNumberToEvaluate](aws-properties-ses-mailmanagerruleset-rulenumbertoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The operator for a numeric condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | NOT_EQUALS | LESS_THAN | GREATER_THAN | LESS_THAN_OR_EQUAL | GREATER_THAN_OR_EQUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to evaluate in a numeric condition expression.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleIsInAddressList

RuleNumberToEvaluate

All content copied from https://docs.aws.amazon.com/.
