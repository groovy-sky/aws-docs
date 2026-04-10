This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule SqsParameters

The templated target type for the Amazon SQS [`SendMessage`](../../../../reference/awssimplequeueservice/latest/apireference/api-sendmessage.md) API operation.
Contains the message group ID to use when the target is a FIFO queue. If you specify an Amazon SQS FIFO queue as a target, the queue must have content-based deduplication enabled.
For more information, see [Using the Amazon SQS message deduplication ID](../../../awssimplequeueservice/latest/sqsdeveloperguide/using-messagededuplicationid-property.md) in the
_Amazon SQS Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MessageGroupId" : String
}

```

### YAML

```yaml

  MessageGroupId: String

```

## Properties

`MessageGroupId`

The FIFO message group ID to use as the target.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SageMakerPipelineParameters

Target

All content copied from https://docs.aws.amazon.com/.
