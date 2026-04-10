This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioSilenceFailoverSettings

MediaLive will perform a failover if audio is not detected in this input for the specified period.

The parent of this entity is FailoverConditionSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioSelectorName" : String,
  "AudioSilenceThresholdMsec" : Integer
}

```

### YAML

```yaml

  AudioSelectorName: String
  AudioSilenceThresholdMsec: Integer

```

## Properties

`AudioSelectorName`

The name of the audio selector in the input that MediaLive should monitor to detect silence. Select your most important rendition. If you didn't create an audio selector in this input, leave blank.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioSilenceThresholdMsec`

The amount of time (in milliseconds) that the active input must be silent before automatic input failover occurs. Silence is defined as audio loss or audio quieter than -50 dBFS.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioSelectorSettings

AudioTrack

All content copied from https://docs.aws.amazon.com/.
