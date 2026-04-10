This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::Stage ParticipantThumbnailConfiguration

Object specifying a configuration of thumbnails for recorded video from an individual participant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordingMode" : String,
  "Storage" : [ String, ... ],
  "TargetIntervalSeconds" : Integer
}

```

### YAML

```yaml

  RecordingMode: String
  Storage:
    - String
  TargetIntervalSeconds: Integer

```

## Properties

`RecordingMode`

Thumbnail recording mode. Default: `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `INTERVAL | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Storage`

Indicates the format in which thumbnails are recorded. `SEQUENTIAL` records all generated thumbnails in a serial manner,
to the media/thumbnails/high directory. `LATEST` saves the latest thumbnail in
media/latest\_thumbnail/high/thumb.jpg and overwrites it at the interval specified by `targetIntervalSeconds`.
You can enable both `SEQUENTIAL` and `LATEST`. Default: `SEQUENTIAL`.

_Required_: No

_Type_: Array of String

_Allowed values_: `SEQUENTIAL | LATEST`

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIntervalSeconds`

The targeted thumbnail-generation interval in seconds. This is configurable only if `recordingMode` is `INTERVAL`. Default: 60.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParticipantRecordingHlsConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
