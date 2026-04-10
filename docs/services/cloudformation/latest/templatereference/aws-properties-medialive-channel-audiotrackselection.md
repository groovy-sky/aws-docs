This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioTrackSelection

Information about the audio track to extract.

The parent of this entity is AudioSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DolbyEDecode" : AudioDolbyEDecode,
  "Tracks" : [ AudioTrack, ... ]
}

```

### YAML

```yaml

  DolbyEDecode:
    AudioDolbyEDecode
  Tracks:
    - AudioTrack

```

## Properties

`DolbyEDecode`

Property description not available.

_Required_: No

_Type_: [AudioDolbyEDecode](aws-properties-medialive-channel-audiodolbyedecode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tracks`

Selects one or more unique audio tracks from within a source.

_Required_: No

_Type_: Array of [AudioTrack](aws-properties-medialive-channel-audiotrack.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioTrack

AudioWatermarkSettings

All content copied from https://docs.aws.amazon.com/.
