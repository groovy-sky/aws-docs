This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet ReplaceRecipientAction

This action replaces the email envelope recipients with the given list of recipients.
If the condition of this action applies only to a subset of recipients, only those
recipients are replaced with the recipients specified in the action. The message
contents and headers are unaffected by this action, only the envelope recipients are
updated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplaceWith" : [ String, ... ]
}

```

### YAML

```yaml

  ReplaceWith:
    - String

```

## Properties

`ReplaceWith`

This action specifies the replacement recipient email addresses to insert.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 1`

_Maximum_: `254 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RelayAction

Rule

All content copied from https://docs.aws.amazon.com/.
