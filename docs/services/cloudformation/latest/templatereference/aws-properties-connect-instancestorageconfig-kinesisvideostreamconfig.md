This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::InstanceStorageConfig KinesisVideoStreamConfig

Configuration information of a Kinesis video stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionConfig" : EncryptionConfig,
  "Prefix" : String,
  "RetentionPeriodHours" : Number
}

```

### YAML

```yaml

  EncryptionConfig:
    EncryptionConfig
  Prefix: String
  RetentionPeriodHours: Number

```

## Properties

`EncryptionConfig`

The encryption configuration.

_Required_: Yes

_Type_: [EncryptionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-instancestorageconfig-encryptionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix of the video stream.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriodHours`

The number of hours data is retained in the stream. Kinesis Video Streams retains the data in a data store that
is associated with the stream.

The default value is 0, indicating that the stream does not persist data.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `87600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KinesisStreamConfig

S3Config
