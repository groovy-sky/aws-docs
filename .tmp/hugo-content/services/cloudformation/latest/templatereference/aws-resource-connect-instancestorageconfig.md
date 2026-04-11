This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::InstanceStorageConfig

The storage configuration for the instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::InstanceStorageConfig",
  "Properties" : {
      "InstanceArn" : String,
      "KinesisFirehoseConfig" : KinesisFirehoseConfig,
      "KinesisStreamConfig" : KinesisStreamConfig,
      "KinesisVideoStreamConfig" : KinesisVideoStreamConfig,
      "ResourceType" : String,
      "S3Config" : S3Config,
      "StorageType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::InstanceStorageConfig
Properties:
  InstanceArn: String
  KinesisFirehoseConfig:
    KinesisFirehoseConfig
  KinesisStreamConfig:
    KinesisStreamConfig
  KinesisVideoStreamConfig:
    KinesisVideoStreamConfig
  ResourceType: String
  S3Config:
    S3Config
  StorageType: String

```

## Properties

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KinesisFirehoseConfig`

The configuration of the Kinesis Firehose delivery stream.

_Required_: No

_Type_: [KinesisFirehoseConfig](aws-properties-connect-instancestorageconfig-kinesisfirehoseconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamConfig`

The configuration of the Kinesis data stream.

_Required_: No

_Type_: [KinesisStreamConfig](aws-properties-connect-instancestorageconfig-kinesisstreamconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisVideoStreamConfig`

The configuration of the Kinesis video stream.

_Required_: No

_Type_: [KinesisVideoStreamConfig](aws-properties-connect-instancestorageconfig-kinesisvideostreamconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

A valid resource type. Following are the valid resource types:
`CHAT_TRANSCRIPTS` \| `CALL_RECORDINGS` \|
`SCHEDULED_REPORTS` \| `MEDIA_STREAMS` \|
`CONTACT_TRACE_RECORDS` \| `AGENT_EVENTS`

_Required_: Yes

_Type_: String

_Allowed values_: `CHAT_TRANSCRIPTS | CALL_RECORDINGS | SCHEDULED_REPORTS | MEDIA_STREAMS | CONTACT_TRACE_RECORDS | AGENT_EVENTS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Config`

The S3 bucket
configuration.

_Required_: No

_Type_: [S3Config](aws-properties-connect-instancestorageconfig-s3config.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageType`

A valid storage type.

_Required_: Yes

_Type_: String

_Allowed values_: `S3 | KINESIS_VIDEO_STREAM | KINESIS_STREAM | KINESIS_FIREHOSE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the instance storage configuration. For
example:

`{ "Ref": "myInstanceStorageConfig" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

The existing association identifier that uniquely identifies the resource type and
storage config for the given instance ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EncryptionConfig

All content copied from https://docs.aws.amazon.com/.
