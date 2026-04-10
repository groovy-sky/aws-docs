This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleBooleanToEvaluate

The union type representing the allowed types of operands for a boolean
condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Analysis" : Analysis,
  "Attribute" : String,
  "IsInAddressList" : RuleIsInAddressList
}

```

### YAML

```yaml

  Analysis:
    Analysis
  Attribute: String
  IsInAddressList:
    RuleIsInAddressList

```

## Properties

`Analysis`

The Add On ARN and its returned value to evaluate in a boolean condition expression.

_Required_: No

_Type_: [Analysis](aws-properties-ses-mailmanagerruleset-analysis.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attribute`

The boolean type representing the allowed attribute types for an email.

_Required_: No

_Type_: String

_Allowed values_: `READ_RECEIPT_REQUESTED | TLS | TLS_WRAPPED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsInAddressList`

The structure representing the address lists and address list attribute that will be used in evaluation of boolean expression.

_Required_: No

_Type_: [RuleIsInAddressList](aws-properties-ses-mailmanagerruleset-ruleisinaddresslist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleBooleanExpression

RuleCondition

All content copied from https://docs.aws.amazon.com/.
