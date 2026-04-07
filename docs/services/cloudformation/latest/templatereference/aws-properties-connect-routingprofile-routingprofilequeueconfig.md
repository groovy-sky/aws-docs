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
information, see [Queues: priority and delay](https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html) in the _Amazon Connect Administrator Guide_.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `9999`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The order in which contacts are to be handled for the queue. For more information, see [Queues: priority and\
delay](https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing-profiles-priority.html).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueReference`

Contains information about a queue resource.

_Required_: Yes

_Type_: [RoutingProfileQueueReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-routingprofile-routingprofilequeuereference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RoutingProfileManualAssignmentQueueConfig

RoutingProfileQueueReference
