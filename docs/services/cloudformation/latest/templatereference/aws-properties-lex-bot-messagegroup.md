This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot MessageGroup

Provides one or more messages that Amazon Lex should send to the
user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Message" : Message,
  "Variations" : [ Message, ... ]
}

```

### YAML

```yaml

  Message:
    Message
  Variations:
    - Message

```

## Properties

`Message`

The primary message that Amazon Lex should send to the user.

_Required_: Yes

_Type_: [Message](aws-properties-lex-bot-message.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variations`

Message variations to send to the user. When variations are defined,
Amazon Lex chooses the primary message or one of the variations to send to
the user.

_Required_: No

_Type_: Array of [Message](aws-properties-lex-bot-message.md)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Message

MultipleValuesSetting

All content copied from https://docs.aws.amazon.com/.
