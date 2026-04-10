This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::RoutingProfile RoutingProfileQueueReference

Contains the channel and queue identifier for a routing profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Channel" : String,
  "QueueArn" : String
}

```

### YAML

```yaml

  Channel: String
  QueueArn: String

```

## Properties

`Channel`

The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.

_Required_: Yes

_Type_: String

_Allowed values_: `VOICE | CHAT | TASK | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueArn`

The Amazon Resource Name (ARN) of the queue.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/queue/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RoutingProfileQueueConfig

Tag

All content copied from https://docs.aws.amazon.com/.
