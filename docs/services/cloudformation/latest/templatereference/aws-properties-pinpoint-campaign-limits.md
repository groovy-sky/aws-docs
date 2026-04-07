This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign Limits

Specifies the limits on the messages that a campaign can send.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Daily" : Integer,
  "MaximumDuration" : Integer,
  "MessagesPerSecond" : Integer,
  "Session" : Integer,
  "Total" : Integer
}

```

### YAML

```yaml

  Daily: Integer
  MaximumDuration: Integer
  MessagesPerSecond: Integer
  Session: Integer
  Total: Integer

```

## Properties

`Daily`

The maximum number of messages that a campaign can send to a single endpoint during a
24-hour period. The maximum value is 100.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumDuration`

The maximum amount of time, in seconds, that a campaign can attempt to deliver a
message after the scheduled start time for the campaign. The minimum value is 60
seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessagesPerSecond`

The maximum number of messages that a campaign can send each second. The minimum value
is 1. The maximum value is 20,000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Session`

The maximum number of messages that the campaign can send per user
session.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Total`

The maximum number of messages that a campaign can send to a single endpoint during
the course of the campaign. The maximum value is 100.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InAppMessageHeaderConfig

Message
