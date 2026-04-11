This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot StillWaitingResponseSpecification

Defines the messages that Amazon Lex sends to a user to remind them that
the bot is waiting for a response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowInterrupt" : Boolean,
  "FrequencyInSeconds" : Integer,
  "MessageGroupsList" : [ MessageGroup, ... ],
  "TimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  AllowInterrupt: Boolean
  FrequencyInSeconds: Integer
  MessageGroupsList:
    - MessageGroup
  TimeoutInSeconds: Integer

```

## Properties

`AllowInterrupt`

Indicates that the user can interrupt the response by speaking while
the message is being played.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrequencyInSeconds`

How often a message should be sent to the user. Minimum of 1 second,
maximum of 5 minutes.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageGroupsList`

One or more message groups, each containing one or more messages,
that define the prompts that Amazon Lex sends to the user.

_Required_: Yes

_Type_: Array of [MessageGroup](aws-properties-lex-bot-messagegroup.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInSeconds`

If Amazon Lex waits longer than this length of time for a response, it
will stop sending messages.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSMLMessage

SubSlotSetting

All content copied from https://docs.aws.amazon.com/.
