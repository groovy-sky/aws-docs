This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::InstanceStorageConfig KinesisFirehoseConfig

Configuration information of a Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FirehoseArn" : String
}

```

### YAML

```yaml

  FirehoseArn: String

```

## Properties

`FirehoseArn`

The Amazon Resource Name (ARN) of the delivery stream.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:firehose:[-a-z0-9]*:[0-9]{12}:deliverystream/[-a-zA-Z0-9_.]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EncryptionConfig

KinesisStreamConfig
