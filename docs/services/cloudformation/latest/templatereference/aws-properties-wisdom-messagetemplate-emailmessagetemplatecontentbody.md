This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate EmailMessageTemplateContentBody

The body to use in email messages.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Html" : MessageTemplateBodyContentProvider,
  "PlainText" : MessageTemplateBodyContentProvider
}

```

### YAML

```yaml

  Html:
    MessageTemplateBodyContentProvider
  PlainText:
    MessageTemplateBodyContentProvider

```

## Properties

`Html`

The message body, in HTML format, to use in email messages that are based on the message template. We recommend
using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an
HTML message.

_Required_: No

_Type_: [MessageTemplateBodyContentProvider](aws-properties-wisdom-messagetemplate-messagetemplatebodycontentprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlainText`

The message body, in plain text format, to use in email messages that are based on the message template. We
recommend using plain text format for email clients that don't render HTML content and clients that are connected to
high-latency networks, such as mobile devices.

_Required_: No

_Type_: [MessageTemplateBodyContentProvider](aws-properties-wisdom-messagetemplate-messagetemplatebodycontentprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmailMessageTemplateContent

EmailMessageTemplateHeader

All content copied from https://docs.aws.amazon.com/.
