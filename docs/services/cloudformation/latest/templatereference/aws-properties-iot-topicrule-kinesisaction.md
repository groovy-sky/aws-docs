This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule KinesisAction

Describes an action to write data to an Amazon Kinesis stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PartitionKey" : String,
  "RoleArn" : String,
  "StreamName" : String
}

```

### YAML

```yaml

  PartitionKey: String
  RoleArn: String
  StreamName: String

```

## Properties

`PartitionKey`

The partition key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access to the Amazon Kinesis stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamName`

The name of the Amazon Kinesis stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KafkaActionHeader

LambdaAction

All content copied from https://docs.aws.amazon.com/.
