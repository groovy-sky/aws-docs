This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleStringToEvaluate

The string to evaluate in a string condition expression.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Analysis" : Analysis,
  "Attribute" : String,
  "MimeHeaderAttribute" : String
}

```

### YAML

```yaml

  Analysis:
    Analysis
  Attribute: String
  MimeHeaderAttribute: String

```

## Properties

`Analysis`

The Add On ARN and its returned value to evaluate in a string condition expression.

_Required_: No

_Type_: [Analysis](aws-properties-ses-mailmanagerruleset-analysis.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attribute`

The email attribute to evaluate in a string condition expression.

_Required_: No

_Type_: String

_Allowed values_: `MAIL_FROM | HELO | RECIPIENT | SENDER | FROM | SUBJECT | TO | CC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MimeHeaderAttribute`

The email MIME X-Header attribute to evaluate in a string condition expression.

_Required_: No

_Type_: String

_Pattern_: `^X-[a-zA-Z0-9-]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleStringExpression

RuleVerdictExpression

All content copied from https://docs.aws.amazon.com/.
