This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::RecordingConfiguration ThumbnailConfiguration

The ThumbnailConfiguration property type describes a configuration of thumbnails for recorded video.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordingMode" : String,
  "Resolution" : String,
  "Storage" : [ String, ... ],
  "TargetIntervalSeconds" : Integer
}

```

### YAML

```yaml

  RecordingMode: String
  Resolution: String
  Storage:
    - String
  TargetIntervalSeconds: Integer

```

## Properties

`RecordingMode`

Thumbnail recording mode. Valid values:

- `DISABLED`: Use DISABLED to disable the generation of thumbnails for recorded video.

- `INTERVAL`: Use INTERVAL to enable the generation of thumbnails for recorded video at a time interval controlled by the
[TargetIntervalSeconds](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-targetintervalseconds) property.

_Default_: `INTERVAL`

_Required_: No

_Type_: String

_Allowed values_: `INTERVAL | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resolution`

The desired resolution of recorded thumbnails for a stream. Thumbnails are recorded at the
selected resolution if the corresponding rendition is available during the stream; otherwise,
they are recorded at source resolution. For more information about resolution values and their
corresponding height and width dimensions, see [Auto-Record to Amazon S3](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/record-to-s3.html).

_Required_: No

_Type_: String

_Allowed values_: `FULL_HD | HD | SD | LOWEST_RESOLUTION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Storage`

The format in which thumbnails are recorded for a stream. `SEQUENTIAL` records all generated
thumbnails in a serial manner, to the media/thumbnails directory. `LATEST` saves the latest
thumbnail in media/thumbnails/latest/thumb.jpg and overwrites it at the interval specified by
`targetIntervalSeconds`. You can enable both `SEQUENTIAL` and `LATEST`. Default: `SEQUENTIAL`.

_Required_: No

_Type_: Array of String

_Allowed values_: `SEQUENTIAL | LATEST`

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetIntervalSeconds`

The targeted thumbnail-generation interval in seconds. This is configurable (and required) only if
[RecordingMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-thumbnailconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration-recordingmode) is `INTERVAL`.

###### Note

Setting a value for `TargetIntervalSeconds` does not guarantee that thumbnails are generated at the specified interval.
For thumbnails to be generated at the `TargetIntervalSeconds` interval,
the `IDR/Keyframe` value for the input video must be less than the `TargetIntervalSeconds` value.
See [Amazon IVS Streaming Configuration](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/streaming-config.html)
for information on setting `IDR/Keyframe` to the recommended value in video-encoder settings.

_Default_: 60

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IVS::Stage
