This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::Stage ParticipantRecordingHlsConfiguration

Object specifying a configuration of participant HLS recordings for individual participant recording.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetSegmentDurationSeconds" : Integer
}

```

### YAML

```yaml

  TargetSegmentDurationSeconds: Integer

```

## Properties

`TargetSegmentDurationSeconds`

Defines the target duration for recorded segments generated when recording a stage participant.
Segments may have durations longer than the specified value when needed to ensure each segment begins with a keyframe. Default: 6.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsConfiguration

ParticipantThumbnailConfiguration

All content copied from https://docs.aws.amazon.com/.
