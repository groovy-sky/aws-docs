This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule SqsParameters

The custom parameters for EventBridge to use for a target that is an Amazon SQS fair or FIFO queue.

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

The ID of the message group to use as the target.

_Required_: Yes

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SageMakerPipelineParameters

Tag

All content copied from https://docs.aws.amazon.com/.
