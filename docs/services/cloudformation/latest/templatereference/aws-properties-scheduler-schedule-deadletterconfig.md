This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule DeadLetterConfig

An object that contains information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule. If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue.

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

The Amazon Resource Name (ARN) of the SQS queue specified as the destination for the dead-letter queue.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:sqs:[a-z0-9\-]+:\d{12}:[a-zA-Z0-9\-_]+$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityProviderStrategyItem

EcsParameters

All content copied from https://docs.aws.amazon.com/.
