This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot FulfillmentUpdateResponseSpecification

Provides settings for a message that is sent periodically to the
user while a fulfillment Lambda function is running.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowInterrupt" : Boolean,
  "FrequencyInSeconds" : Integer,
  "MessageGroups" : [ MessageGroup, ... ]
}

```

### YAML

```yaml

  AllowInterrupt: Boolean
  FrequencyInSeconds: Integer
  MessageGroups:
    - MessageGroup

```

## Properties

`AllowInterrupt`

Determines whether the user can interrupt an update message while it
is playing.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrequencyInSeconds`

The frequency that a message is sent to the user. When the period
ends, Amazon Lex chooses a message from the message groups and plays it to
the user. If the fulfillment Lambda returns before the first period
ends, an update message is not played to the user.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageGroups`

1 - 5 message groups that contain update messages. Amazon Lex chooses
one of the messages to play to the user.

_Required_: Yes

_Type_: Array of [MessageGroup](aws-properties-lex-bot-messagegroup.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FulfillmentStartResponseSpecification

FulfillmentUpdatesSpecification

All content copied from https://docs.aws.amazon.com/.
