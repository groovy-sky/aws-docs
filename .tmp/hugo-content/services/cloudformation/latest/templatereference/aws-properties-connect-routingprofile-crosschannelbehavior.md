This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::RoutingProfile CrossChannelBehavior

Defines the cross-channel routing behavior that allows an agent working on a contact in one channel to be
offered a contact from a different channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BehaviorType" : String
}

```

### YAML

```yaml

  BehaviorType: String

```

## Properties

`BehaviorType`

Specifies the other channels that can be routed to an agent handling their current channel.

_Required_: Yes

_Type_: String

_Allowed values_: `ROUTE_CURRENT_CHANNEL_ONLY | ROUTE_ANY_CHANNEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::RoutingProfile

MediaConcurrency

All content copied from https://docs.aws.amazon.com/.
