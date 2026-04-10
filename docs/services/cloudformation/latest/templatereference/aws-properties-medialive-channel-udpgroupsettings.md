This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel UdpGroupSettings

The configuration of a UDP output group.

The parent of this entity is OutputGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputLossAction" : String,
  "TimedMetadataId3Frame" : String,
  "TimedMetadataId3Period" : Integer
}

```

### YAML

```yaml

  InputLossAction: String
  TimedMetadataId3Frame: String
  TimedMetadataId3Period: Integer

```

## Properties

`InputLossAction`

Specifies the behavior of the last resort when the input video is lost, and no more backup inputs are available.
When dropTs is selected, the entire transport stream stops emitting. When dropProgram is selected, the program can be
dropped from the transport stream (and replaced with null packets to meet the TS bitrate requirement). Or when
emitProgram is selected, the transport stream continues to be produced normally with repeat frames, black frames, or
slate frames substituted for the absent input video.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataId3Frame`

Indicates the ID3 frame that has the timecode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataId3Period`

The timed metadata interval in seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UdpContainerSettings

UdpOutputSettings

All content copied from https://docs.aws.amazon.com/.
