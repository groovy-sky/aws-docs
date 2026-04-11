This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::RoutingProfile MediaConcurrency

Contains information about which channels are supported, and how many contacts an agent can have on a channel
simultaneously.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Channel" : String,
  "Concurrency" : Integer,
  "CrossChannelBehavior" : CrossChannelBehavior
}

```

### YAML

```yaml

  Channel: String
  Concurrency: Integer
  CrossChannelBehavior:
    CrossChannelBehavior

```

## Properties

`Channel`

The channels that agents can handle in the Contact Control Panel (CCP).

_Required_: Yes

_Type_: String

_Allowed values_: `VOICE | CHAT | TASK | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Concurrency`

The number of contacts an agent can have on a channel simultaneously.

Valid Range for `VOICE`: Minimum value of 1. Maximum value of 1.

Valid Range for `CHAT`: Minimum value of 1. Maximum value of 10.

Valid Range for `TASK`: Minimum value of 1. Maximum value of 10.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossChannelBehavior`

Defines the cross-channel routing behavior for each channel that is enabled for this Routing Profile. For
example, this allows you to offer an agent a different contact from another channel when they are currently working
with a contact from a Voice channel.

_Required_: No

_Type_: [CrossChannelBehavior](aws-properties-connect-routingprofile-crosschannelbehavior.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrossChannelBehavior

RoutingProfileManualAssignmentQueueConfig

All content copied from https://docs.aws.amazon.com/.
