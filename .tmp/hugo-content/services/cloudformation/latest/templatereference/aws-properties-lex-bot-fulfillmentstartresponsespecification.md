This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot FulfillmentStartResponseSpecification

Provides settings for a message that is sent to the user when a
fulfillment Lambda function starts running.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowInterrupt" : Boolean,
  "DelayInSeconds" : Integer,
  "MessageGroups" : [ MessageGroup, ... ]
}

```

### YAML

```yaml

  AllowInterrupt: Boolean
  DelayInSeconds: Integer
  MessageGroups:
    - MessageGroup

```

## Properties

`AllowInterrupt`

Determines whether the user can interrupt the start message while it
is playing.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DelayInSeconds`

The delay between when the Lambda fulfillment function starts running
and the start message is played. If the Lambda function returns before
the delay is over, the start message isn't played.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageGroups`

1 - 5 message groups that contain start messages. Amazon Lex chooses
one of the messages to play to the user.

_Required_: Yes

_Type_: Array of [MessageGroup](aws-properties-lex-bot-messagegroup.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FulfillmentCodeHookSetting

FulfillmentUpdateResponseSpecification

All content copied from https://docs.aws.amazon.com/.
