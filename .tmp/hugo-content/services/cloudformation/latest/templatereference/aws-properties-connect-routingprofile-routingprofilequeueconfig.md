This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::RoutingProfile RoutingProfileQueueConfig

Contains information about the queue and channel for which priority and delay can be set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delay" : Integer,
  "Priority" : Integer,
  "QueueReference" : RoutingProfileQueueReference
}

```

### YAML

```yaml

  Delay: Integer
  Priority: Integer
  QueueReference:
    RoutingProfileQueueReference

```

## Properties

`Delay`

The delay, in seconds, a contact should be in the queue before they are routed to an available agent. For more
information, see [Queues: priority and delay](../../../connect/latest/adminguide/concepts-routing-profiles-priority.md) in the _Amazon Connect Administrator Guide_.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `9999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The order in which contacts are to be handled for the queue. For more information, see [Queues: priority and\
delay](../../../connect/latest/adminguide/concepts-routing-profiles-priority.md).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueReference`

Contains information about a queue resource.

_Required_: Yes

_Type_: [RoutingProfileQueueReference](aws-properties-connect-routingprofile-routingprofilequeuereference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RoutingProfileManualAssignmentQueueConfig

RoutingProfileQueueReference

All content copied from https://docs.aws.amazon.com/.
