This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::Stage

The `AWS::IVS::Stage` resource specifies an Amazon IVS stage. A stage is a virtual space where participants can exchange video in real time. For more information, see [CreateStage](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_CreateStage.html) in the
_Amazon IVS Real-Time Streaming API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::Stage",
  "Properties" : {
      "AutoParticipantRecordingConfiguration" : AutoParticipantRecordingConfiguration,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVS::Stage
Properties:
  AutoParticipantRecordingConfiguration:
    AutoParticipantRecordingConfiguration
  Name: String
  Tags:
    - Tag

```

## Properties

`AutoParticipantRecordingConfiguration`

Configuration object for individual participant recording.

_Required_: No

_Type_: [AutoParticipantRecordingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Stage name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-tag.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-stage-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the stage ARN. For example:

`{ "Ref": "myStage" }`

For the Amazon IVS stage
`myStage`, `Ref` returns the
stage ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ActiveSessionId`

ID of the active session within the stage. For example:
`st-a1b2c3d4e5f6g`

`Arn`

The stage ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:stage/abcdABCDefgh`

## Examples

### Stage Template Examples

The following examples specify an Amazon IVS stage.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "StorageConfiguration": {
            "Type": "AWS::IVS::StorageConfiguration",
            "Properties": {
                "Name": "myStorageConfiguration",
                "S3": { "BucketName": "my-bucket" },
                "Tags": [
                    {
                        "Key": "MyKey",
                        "Value": "MyValue"
                    }
                ]
            }
        },
        "Stage": {
            "Type": "AWS::IVS::Stage",
            "Properties": {
                "Name": "myStage",
                "AutoParticipantRecordingConfiguration": {
                    "StorageConfigurationArn": {
                        "Ref": "StorageConfiguration"
                    },
                    "MediaTypes": [
                        "AUDIO_VIDEO"
                    ],
                    "HlsConfiguration": {
                        "ParticipantRecordingHlsConfiguration": {
                            "TargetSegmentDurationSeconds": 5
                        }
                    },
                    "RecordingReconnectWindowSeconds": 30,
                    "ThumbnailConfiguration": {
                        "ParticipantThumbnailConfiguration": {
                            "RecordingMode": "INTERVAL",
                            "TargetIntervalSeconds": 2,
                            "Storage": [
                                "SEQUENTIAL",
                                "LATEST"
                            ]
                        }
                    }
                },
                "Tags": [
                    {
                        "Key": "MyKey",
                        "Value": "MyValue"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  StorageConfiguration:
    Type: AWS::IVS::StorageConfiguration
    Properties:
      Name: myStorageConfiguration
      S3:
        BucketName: my-bucket
      Tags:
        - Key: myKey
          Value: myValue
  Stage:
    Type: AWS::IVS::Stage
    Properties:
      Name: myStage
      AutoParticipantRecordingConfiguration:
        HlsConfiguration:
          ParticipantRecordingHlsConfiguration:
            TargetSegmentDurationSeconds: 5
        MediaTypes:
          - AUDIO_VIDEO
        RecordingReconnectWindowSeconds: 30
        StorageConfigurationArn: !Ref StorageConfiguration
        ThumbnailConfiguration:
          ParticipantThumbnailConfiguration:
            RecordingMode: INTERVAL
            TargetIntervalSeconds: 2
            Storage:
              - SEQUENTIAL
              - LATEST
      Tags:
        - Key: myKey
          Value: myValue

```

## See also

- [Getting\
Started with IVS Real-Time Streaming](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/getting-started.html)

- [Stage](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_Stage.html) data
type

- [CreateStage](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_CreateStage.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ThumbnailConfiguration

AutoParticipantRecordingConfiguration
