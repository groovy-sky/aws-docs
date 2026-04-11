This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel GlobalConfiguration

The configuration settings that apply to the entire channel.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InitialAudioGain" : Integer,
  "InputEndAction" : String,
  "InputLossBehavior" : InputLossBehavior,
  "OutputLockingMode" : String,
  "OutputLockingSettings" : OutputLockingSettings,
  "OutputTimingSource" : String,
  "SupportLowFramerateInputs" : String
}

```

### YAML

```yaml

  InitialAudioGain: Integer
  InputEndAction: String
  InputLossBehavior:
    InputLossBehavior
  OutputLockingMode: String
  OutputLockingSettings:
    OutputLockingSettings
  OutputTimingSource: String
  SupportLowFramerateInputs: String

```

## Properties

`InitialAudioGain`

The value to set the initial audio gain for the channel.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputEndAction`

Indicates the action to take when the current input completes (for example, end-of-file). When
switchAndLoopInputs is configured, MediaLive restarts at the beginning of the first input. When "none" is configured,
MediaLive transcodes either black, a solid color, or a user-specified slate images per the "Input Loss Behavior"
configuration until the next input switch occurs (which is controlled through the Channel Schedule API).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLossBehavior`

The settings for system actions when the input is lost.

_Required_: No

_Type_: [InputLossBehavior](aws-properties-medialive-channel-inputlossbehavior.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputLockingMode`

Indicates how MediaLive pipelines are synchronized. PIPELINELOCKING - MediaLive attempts to synchronize the
output of each pipeline to the other. EPOCHLOCKING - MediaLive attempts to synchronize the output of each pipeline to
the Unix epoch.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputLockingSettings`

Property description not available.

_Required_: No

_Type_: [OutputLockingSettings](aws-properties-medialive-channel-outputlockingsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputTimingSource`

Indicates whether the rate of frames emitted by the Live encoder should be paced by its system clock (which
optionally might be locked to another source through NTP) or should be locked to the clock of the source that is
providing the input stream.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportLowFramerateInputs`

Adjusts the video input buffer for streams with very low video frame rates. This is commonly set to enabled for
music channels with less than one video frame per second.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameCaptureSettings

H264ColorSpaceSettings

All content copied from https://docs.aws.amazon.com/.
