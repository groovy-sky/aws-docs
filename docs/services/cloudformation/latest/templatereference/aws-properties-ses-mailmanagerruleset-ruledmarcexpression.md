This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleDmarcExpression

A DMARC policy expression. The condition matches if the given DMARC policy matches
that of the incoming email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Operator: String
  Values:
    - String

```

## Properties

`Operator`

The operator to apply to the DMARC policy of the incoming email.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | NOT_EQUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values to use for the given DMARC policy operator. For the operator EQUALS, if
multiple values are given, they are evaluated as an OR. That is, if any of the given
values match, the condition is deemed to match. For the operator NOT\_EQUALS, if multiple
values are given, they are evaluated as an AND. That is, only if the email's DMARC
policy is not equal to any of the given values, then the condition is deemed to
match.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleCondition

RuleIpExpression

All content copied from https://docs.aws.amazon.com/.
