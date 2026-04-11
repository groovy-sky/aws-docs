This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot Message

The object that provides message text and its type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomPayload" : CustomPayload,
  "ImageResponseCard" : ImageResponseCard,
  "PlainTextMessage" : PlainTextMessage,
  "SSMLMessage" : SSMLMessage
}

```

### YAML

```yaml

  CustomPayload:
    CustomPayload
  ImageResponseCard:
    ImageResponseCard
  PlainTextMessage:
    PlainTextMessage
  SSMLMessage:
    SSMLMessage

```

## Properties

`CustomPayload`

A message in a custom format defined by the client
application.

_Required_: No

_Type_: [CustomPayload](aws-properties-lex-bot-custompayload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageResponseCard`

A message that defines a response card that the client application
can show to the user.

_Required_: No

_Type_: [ImageResponseCard](aws-properties-lex-bot-imageresponsecard.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlainTextMessage`

A message in plain text format.

_Required_: No

_Type_: [PlainTextMessage](aws-properties-lex-bot-plaintextmessage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSMLMessage`

A message in Speech Synthesis Markup Language (SSML).

_Required_: No

_Type_: [SSMLMessage](aws-properties-lex-bot-ssmlmessage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaCodeHook

MessageGroup

All content copied from https://docs.aws.amazon.com/.
