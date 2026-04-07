This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisVideo::Stream

Specifies a new Kinesis video stream.

When you create a new stream, Kinesis Video Streams assigns it a version number. When you
change the stream's metadata, Kinesis Video Streams updates the version.

`CreateStream` is an asynchronous operation.

For information about how the service works, see [How it Works](https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/how-it-works.html).

You must have permissions for the `KinesisVideo:CreateStream` action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisVideo::Stream",
  "Properties" : {
      "DataRetentionInHours" : Integer,
      "DeviceName" : String,
      "KmsKeyId" : String,
      "MediaType" : String,
      "Name" : String,
      "StreamStorageConfiguration" : StreamStorageConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KinesisVideo::Stream
Properties:
  DataRetentionInHours: Integer
  DeviceName: String
  KmsKeyId: String
  MediaType: String
  Name: String
  StreamStorageConfiguration:
    StreamStorageConfiguration
  Tags:
    - Tag

```

## Properties

`DataRetentionInHours`

How long the stream retains data, in hours.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `87600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceName`

The name of the device that is associated with the stream.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the AWS Key Management Service (AWS KMS) key that Kinesis Video Streams
uses to encrypt data on the stream.

_Required_: No

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaType`

The `MediaType` of the stream.

_Required_: No

_Type_: String

_Pattern_: `[\w\-\.\+]+/[\w\-\.\+]+(,[\w\-\.\+]+/[\w\-\.\+]+)*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the stream.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StreamStorageConfiguration`

The configuration for stream storage, including the default storage tier for stream data. This configuration determines how stream data is stored and accessed, with different tiers offering varying levels of performance and cost optimization.

_Required_: No

_Type_: [StreamStorageConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisvideo-stream-streamstorageconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisvideo-stream-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the stream.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

StreamStorageConfiguration
