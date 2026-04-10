This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule KinesisParameters

The templated target type for the Amazon Kinesis [`PutRecord`](../../../../reference/kinesis/latest/apireference/api-putrecord.md) API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PartitionKey" : String
}

```

### YAML

```yaml

  PartitionKey: String

```

## Properties

`PartitionKey`

Specifies the shard to which EventBridge Scheduler sends the event. For more information, see [Amazon Kinesis Data Streams terminology and concepts](../../../streams/latest/dev/key-concepts.md) in the
_Amazon Kinesis Streams Developer Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlexibleTimeWindow

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
