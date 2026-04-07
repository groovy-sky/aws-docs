This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::RoutingProfile RoutingProfileManualAssignmentQueueConfig

Contains information about the queue and channel for manual assignment behaviour can be enabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueueReference" : RoutingProfileQueueReference
}

```

### YAML

```yaml

  QueueReference:
    RoutingProfileQueueReference

```

## Properties

`QueueReference`

Contains information about a queue resource.

_Required_: Yes

_Type_: [RoutingProfileQueueReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-routingprofile-routingprofilequeuereference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MediaConcurrency

RoutingProfileQueueConfig
