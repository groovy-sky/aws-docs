This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleVerdictToEvaluate

The verdict to evaluate in a verdict condition expression.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Analysis" : Analysis,
  "Attribute" : String
}

```

### YAML

```yaml

  Analysis:
    Analysis
  Attribute: String

```

## Properties

`Analysis`

The Add On ARN and its returned value to evaluate in a verdict condition
expression.

_Required_: No

_Type_: [Analysis](aws-properties-ses-mailmanagerruleset-analysis.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attribute`

The email verdict attribute to evaluate in a string verdict expression.

_Required_: No

_Type_: String

_Allowed values_: `SPF | DKIM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleVerdictExpression

S3Action

All content copied from https://docs.aws.amazon.com/.
