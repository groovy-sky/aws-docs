This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioSelectorSettings

Information about the audio to extract from the input.

The parent of this entity is AudioSelector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioHlsRenditionSelection" : AudioHlsRenditionSelection,
  "AudioLanguageSelection" : AudioLanguageSelection,
  "AudioPidSelection" : AudioPidSelection,
  "AudioTrackSelection" : AudioTrackSelection
}

```

### YAML

```yaml

  AudioHlsRenditionSelection:
    AudioHlsRenditionSelection
  AudioLanguageSelection:
    AudioLanguageSelection
  AudioPidSelection:
    AudioPidSelection
  AudioTrackSelection:
    AudioTrackSelection

```

## Properties

`AudioHlsRenditionSelection`

Selector for HLS audio rendition.

_Required_: No

_Type_: [AudioHlsRenditionSelection](aws-properties-medialive-channel-audiohlsrenditionselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioLanguageSelection`

The language code of the audio to select.

_Required_: No

_Type_: [AudioLanguageSelection](aws-properties-medialive-channel-audiolanguageselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioPidSelection`

The PID of the audio to select.

_Required_: No

_Type_: [AudioPidSelection](aws-properties-medialive-channel-audiopidselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioTrackSelection`

Information about the audio track to extract.

_Required_: No

_Type_: [AudioTrackSelection](aws-properties-medialive-channel-audiotrackselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioSelector

AudioSilenceFailoverSettings

All content copied from https://docs.aws.amazon.com/.
