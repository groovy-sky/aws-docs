This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleStringExpression

A string expression is evaluated against strings or substrings of the email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : RuleStringToEvaluate,
  "Operator" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Evaluate:
    RuleStringToEvaluate
  Operator: String
  Values:
    - String

```

## Properties

`Evaluate`

The string to evaluate in a string condition expression.

_Required_: Yes

_Type_: [RuleStringToEvaluate](aws-properties-ses-mailmanagerruleset-rulestringtoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The matching operator for a string condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | NOT_EQUALS | STARTS_WITH | ENDS_WITH | CONTAINS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The string(s) to be evaluated in a string condition expression. For all operators,
except for NOT\_EQUALS, if multiple values are given, the values are processed as an OR.
That is, if any of the values match the email's string using the given operator, the
condition is deemed to match. However, for NOT\_EQUALS, the condition is only deemed to
match if none of the given strings match the email's string.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `4096 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleNumberToEvaluate

RuleStringToEvaluate

All content copied from https://docs.aws.amazon.com/.
