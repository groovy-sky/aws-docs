This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule KinesisParameters

This object enables you to specify a JSON path to extract from the event and use as the
partition key for the Amazon Kinesis data stream, so that you can control the shard to which
the event goes. If you do not include this parameter, the default is to use the
`eventId` as the partition key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PartitionKeyPath" : String
}

```

### YAML

```yaml

  PartitionKeyPath: String

```

## Properties

`PartitionKeyPath`

The JSON path to be extracted from the event and used as the partition key. For more
information, see [Amazon Kinesis Streams Key\
Concepts](../../../streams/latest/dev/key-concepts.md#partition-key) in the _Amazon Kinesis Streams Developer Guide_.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputTransformer

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
