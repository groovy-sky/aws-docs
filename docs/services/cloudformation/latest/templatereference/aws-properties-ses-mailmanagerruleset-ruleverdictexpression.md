This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleVerdictExpression

A verdict expression is evaluated against verdicts of the email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : RuleVerdictToEvaluate,
  "Operator" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Evaluate:
    RuleVerdictToEvaluate
  Operator: String
  Values:
    - String

```

## Properties

`Evaluate`

The verdict to evaluate in a verdict condition expression.

_Required_: Yes

_Type_: [RuleVerdictToEvaluate](aws-properties-ses-mailmanagerruleset-ruleverdicttoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The matching operator for a verdict condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | NOT_EQUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values to match with the email's verdict using the given operator. For the EQUALS
operator, if multiple values are given, the condition is deemed to match if any of the
given verdicts match that of the email. For the NOT\_EQUALS operator, if multiple values
are given, the condition is deemed to match of none of the given verdicts match the
verdict of the email.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleStringToEvaluate

RuleVerdictToEvaluate

All content copied from https://docs.aws.amazon.com/.
