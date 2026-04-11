This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate Content

The content of the message template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailMessageTemplateContent" : EmailMessageTemplateContent,
  "SmsMessageTemplateContent" : SmsMessageTemplateContent
}

```

### YAML

```yaml

  EmailMessageTemplateContent:
    EmailMessageTemplateContent
  SmsMessageTemplateContent:
    SmsMessageTemplateContent

```

## Properties

`EmailMessageTemplateContent`

The content of the message template that applies to the email channel subtype.

_Required_: No

_Type_: [EmailMessageTemplateContent](aws-properties-wisdom-messagetemplate-emailmessagetemplatecontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmsMessageTemplateContent`

The content of message template that applies to SMS channel subtype.

_Required_: No

_Type_: [SmsMessageTemplateContent](aws-properties-wisdom-messagetemplate-smsmessagetemplatecontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentAttributes

CustomerProfileAttributes

All content copied from https://docs.aws.amazon.com/.
