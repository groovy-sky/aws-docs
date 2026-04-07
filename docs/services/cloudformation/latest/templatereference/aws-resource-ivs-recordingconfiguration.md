This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::RecordingConfiguration

The `AWS::IVS::RecordingConfiguration` resource specifies an Amazon IVS recording configuration. A recording configuration
enables the recording of a channel’s live streams to a data store. Multiple channels can
reference the same recording configuration. For more information, see [RecordingConfiguration](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_RecordingConfiguration.html) in the _Amazon IVS Low-Latency Streaming API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::RecordingConfiguration",
  "Properties" : {
      "DestinationConfiguration" : DestinationConfiguration,
      "Name" : String,
      "RecordingReconnectWindowSeconds" : Integer,
      "RenditionConfiguration" : RenditionConfiguration,
      "Tags" : [ Tag, ... ],
      "ThumbnailConfiguration" : ThumbnailConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::IVS::RecordingConfiguration
Properties:
  DestinationConfiguration:
    DestinationConfiguration
  Name: String
  RecordingReconnectWindowSeconds: Integer
  RenditionConfiguration:
    RenditionConfiguration
  Tags:
    - Tag
  ThumbnailConfiguration:
    ThumbnailConfiguration

```

## Properties

`DestinationConfiguration`

A destination configuration describes an S3 bucket where recorded video will be
stored. See the [DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html) property type for more information.

_Required_: Yes

_Type_: [DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-destinationconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Recording-configuration name. The value does not need to be unique.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecordingReconnectWindowSeconds`

If a broadcast disconnects and then reconnects within the specified interval, the multiple
streams will be considered a single broadcast and merged together.

_Default_: `0`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RenditionConfiguration`

A rendition configuration describes which renditions should be recorded for a stream.
See the [RenditionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-renditionconfiguration.html)
property type for more information.

_Required_: No

_Type_: [RenditionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-renditionconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-tag.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThumbnailConfiguration`

A thumbnail configuration enables/disables the recording of thumbnails for a live session and controls the interval at which thumbnails are generated for the live session.
See the [ThumbnailConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html) property type for more information.

_Required_: No

_Type_: [ThumbnailConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the recording-configuration ARN. For example:

`{ "Ref": "myRecordingConfiguration" }`

For the Amazon IVS recording configuration
`myRecordingConfiguration`, `Ref` returns the
recording-configuration ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The recording configuration ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:recording-configuration/abcdABCDefgh`

`State`

Indicates the current state of the recording configuration. When the state is
`ACTIVE`, the configuration is ready to record a channel stream. Valid
values: `CREATING` \| `CREATE_FAILED` \| `ACTIVE`.

## Examples

### Recording Configuration Template Examples

The following examples specify an Amazon IVS Channel
that records live-channel streams to an S3 bucket.

#### JSON

```json

{
  "Resources": {
    "S3Bucket": {
      "Type": "AWS::S3::Bucket"
    },
    "RecordingConfiguration": {
      "Type": "AWS::IVS::RecordingConfiguration",
      "DependsOn": "S3Bucket",
      "Properties": {
        "Name": “MyRecordingConfiguration”,
        "DestinationConfiguration": {
          "S3": {
            "BucketName": {
              "Ref": "S3Bucket"
            }
          }
        },
        "ThumbnailConfiguration": {
          "RecordingMode": "INTERVAL",
            "TargetIntervalSeconds": 60,
            "Storage": ["SEQUENTIAL", "LATEST"],
          "Resolution": "HD"
        },
        "RenditionConfiguration": {
          "RenditionSelection": "CUSTOM",
          "Renditions": ["HD", "SD"]
        }
      }
    },
    "Channel": {
      "Type": "AWS::IVS::Channel",
      "DependsOn": "RecordingConfiguration",
      "Properties": {
        "Name": "MyRecordedChannel",
        "RecordingConfigurationArn": {
          "Ref": "RecordingConfiguration"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
 S3Bucket:
   Type: AWS::S3::Bucket

 RecordingConfiguration:
   Type: AWS::IVS::RecordingConfiguration
   DependsOn: S3Bucket
   Properties:
     Name: MyRecordingConfiguration
     DestinationConfiguration:
       S3:
         BucketName: !Ref S3Bucket
     ThumbnailConfiguration:
       RecordingMode: INTERVAL
       TargetIntervalSeconds: 60
       Resolution: HD
       Storage:
         - SEQUENTIAL
         - LATEST
     RenditionConfiguration:
       RenditionSelection: CUSTOM
       Renditions:
         - HD
         - SD

 Channel:
   Type: AWS::IVS::Channel
   DependsOn: RecordingConfiguration
   Properties:
     Name: MyRecordedChannel
     RecordingConfigurationArn: !Ref RecordingConfiguration

```

## See also

- [Getting\
Started with IVS Low-Latency Streaming](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/GSIVS.html)

- [Auto-Record to Amazon S3](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/record-to-s3.html)

- [RecordingConfiguration](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_RecordingConfiguration.html) API data type

- [DestinationConfiguration](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_DestinationConfiguration.html) API data type

- [S3DestinationConfiguration](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_S3DestinationConfiguration.html) API data type

- [CreateRecordingConfiguration](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_CreateRecordingConfiguration.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DestinationConfiguration
