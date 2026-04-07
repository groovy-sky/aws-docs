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

_Type_: [EmailMessageTemplateContent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-messagetemplate-emailmessagetemplatecontent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmsMessageTemplateContent`

The content of message template that applies to SMS channel subtype.

_Required_: No

_Type_: [SmsMessageTemplateContent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-messagetemplate-smsmessagetemplatecontent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AgentAttributes

CustomerProfileAttributes
