This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetSqsQueueParameters

The parameters for using a Amazon SQS stream as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MessageDeduplicationId" : String,
  "MessageGroupId" : String
}

```

### YAML

```yaml

  MessageDeduplicationId: String
  MessageGroupId: String

```

## Properties

`MessageDeduplicationId`

This parameter applies only to FIFO (first-in-first-out) queues.

The token used for deduplication of sent messages.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageGroupId`

The FIFO message group ID to use as the target.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetSageMakerPipelineParameters

PipeTargetStateMachineParameters

All content copied from https://docs.aws.amazon.com/.
