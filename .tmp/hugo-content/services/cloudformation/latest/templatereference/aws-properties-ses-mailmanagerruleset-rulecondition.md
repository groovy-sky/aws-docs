This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleCondition

The conditional expression used to evaluate an email for determining if a rule action
should be taken.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanExpression" : RuleBooleanExpression,
  "DmarcExpression" : RuleDmarcExpression,
  "IpExpression" : RuleIpExpression,
  "NumberExpression" : RuleNumberExpression,
  "StringExpression" : RuleStringExpression,
  "VerdictExpression" : RuleVerdictExpression
}

```

### YAML

```yaml

  BooleanExpression:
    RuleBooleanExpression
  DmarcExpression:
    RuleDmarcExpression
  IpExpression:
    RuleIpExpression
  NumberExpression:
    RuleNumberExpression
  StringExpression:
    RuleStringExpression
  VerdictExpression:
    RuleVerdictExpression

```

## Properties

`BooleanExpression`

The condition applies to a boolean expression passed in this field.

_Required_: No

_Type_: [RuleBooleanExpression](aws-properties-ses-mailmanagerruleset-rulebooleanexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DmarcExpression`

The condition applies to a DMARC policy expression passed in this field.

_Required_: No

_Type_: [RuleDmarcExpression](aws-properties-ses-mailmanagerruleset-ruledmarcexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpExpression`

The condition applies to an IP address expression passed in this field.

_Required_: No

_Type_: [RuleIpExpression](aws-properties-ses-mailmanagerruleset-ruleipexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberExpression`

The condition applies to a number expression passed in this field.

_Required_: No

_Type_: [RuleNumberExpression](aws-properties-ses-mailmanagerruleset-rulenumberexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringExpression`

The condition applies to a string expression passed in this field.

_Required_: No

_Type_: [RuleStringExpression](aws-properties-ses-mailmanagerruleset-rulestringexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerdictExpression`

The condition applies to a verdict expression passed in this field.

_Required_: No

_Type_: [RuleVerdictExpression](aws-properties-ses-mailmanagerruleset-ruleverdictexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleBooleanToEvaluate

RuleDmarcExpression

All content copied from https://docs.aws.amazon.com/.
