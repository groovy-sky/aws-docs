This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::EventBus DeadLetterConfig

Configuration details of the Amazon SQS queue for EventBridge to use as a
dead-letter queue (DLQ).

For more information, see [Using dead-letter queues to process undelivered events](../../../eventbridge/latest/userguide/eb-rule-event-delivery.md#eb-rule-dlq) in the _EventBridge User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String
}

```

### YAML

```yaml

  Arn: String

```

## Properties

`Arn`

The ARN of the SQS queue specified as the target for the dead-letter queue.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Events::EventBus

LogConfig

All content copied from https://docs.aws.amazon.com/.
