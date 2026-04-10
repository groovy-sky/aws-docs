This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputSettings

Information about extracting content from the input and about handling the content.

The parent of this entity is InputAttachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioSelectors" : [ AudioSelector, ... ],
  "CaptionSelectors" : [ CaptionSelector, ... ],
  "DeblockFilter" : String,
  "DenoiseFilter" : String,
  "FilterStrength" : Integer,
  "InputFilter" : String,
  "NetworkInputSettings" : NetworkInputSettings,
  "Scte35Pid" : Integer,
  "Smpte2038DataPreference" : String,
  "SourceEndBehavior" : String,
  "VideoSelector" : VideoSelector
}

```

### YAML

```yaml

  AudioSelectors:
    - AudioSelector
  CaptionSelectors:
    - CaptionSelector
  DeblockFilter: String
  DenoiseFilter: String
  FilterStrength: Integer
  InputFilter: String
  NetworkInputSettings:
    NetworkInputSettings
  Scte35Pid: Integer
  Smpte2038DataPreference: String
  SourceEndBehavior: String
  VideoSelector:
    VideoSelector

```

## Properties

`AudioSelectors`

Information about the specific audio to extract from the input.

The parent of this entity is InputSettings.

_Required_: No

_Type_: Array of [AudioSelector](aws-properties-medialive-channel-audioselector.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionSelectors`

Information about the specific captions to extract from the input.

_Required_: No

_Type_: Array of [CaptionSelector](aws-properties-medialive-channel-captionselector.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeblockFilter`

Enables or disables the deblock filter when filtering.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DenoiseFilter`

Enables or disables the denoise filter when filtering.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterStrength`

Adjusts the magnitude of filtering from 1 (minimal) to 5 (strongest).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputFilter`

Turns on the filter for this input. MPEG-2 inputs have the deblocking filter enabled by default. 1) auto -
filtering is applied depending on input type/quality 2) disabled - no filtering is applied to the input 3) forced -
filtering is applied regardless of the input type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInputSettings`

Information about how to connect to the upstream system.

_Required_: No

_Type_: [NetworkInputSettings](aws-properties-medialive-channel-networkinputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35Pid`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Smpte2038DataPreference`

Specifies whether to extract applicable ancillary data from a SMPTE-2038 source in this input. Applicable data types are captions, timecode, AFD, and SCTE-104 messages.
\- PREFER: Extract from SMPTE-2038 if present in this input, otherwise extract from another source (if any).
\- IGNORE: Never extract any ancillary data from SMPTE-2038.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceEndBehavior`

The loop input if it is a file.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoSelector`

Information about one video to extract from the input.

_Required_: No

_Type_: [VideoSelector](aws-properties-medialive-channel-videoselector.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputLossFailoverSettings

InputSpecification

All content copied from https://docs.aws.amazon.com/.
