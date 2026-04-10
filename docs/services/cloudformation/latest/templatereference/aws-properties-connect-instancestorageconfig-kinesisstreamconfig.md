This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::InstanceStorageConfig KinesisStreamConfig

Configuration information of a Kinesis data stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StreamArn" : String
}

```

### YAML

```yaml

  StreamArn: String

```

## Properties

`StreamArn`

The Amazon Resource Name (ARN) of the data stream.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:kinesis:[-a-z0-9]*:[0-9]{12}:stream/[-a-zA-Z0-9_.]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseConfig

KinesisVideoStreamConfig

All content copied from https://docs.aws.amazon.com/.
