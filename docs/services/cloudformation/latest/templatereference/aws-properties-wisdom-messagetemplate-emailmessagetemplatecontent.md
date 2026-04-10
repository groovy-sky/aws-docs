This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate EmailMessageTemplateContent

The content of the message template that applies to the email channel subtype.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Body" : EmailMessageTemplateContentBody,
  "Headers" : [ EmailMessageTemplateHeader, ... ],
  "Subject" : String
}

```

### YAML

```yaml

  Body:
    EmailMessageTemplateContentBody
  Headers:
    - EmailMessageTemplateHeader
  Subject: String

```

## Properties

`Body`

The body to use in email messages.

_Required_: Yes

_Type_: [EmailMessageTemplateContentBody](aws-properties-wisdom-messagetemplate-emailmessagetemplatecontentbody.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

The email headers to include in email messages.

_Required_: Yes

_Type_: Array of [EmailMessageTemplateHeader](aws-properties-wisdom-messagetemplate-emailmessagetemplateheader.md)

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subject`

The subject line, or title, to use in email messages.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerProfileAttributes

EmailMessageTemplateContentBody

All content copied from https://docs.aws.amazon.com/.
