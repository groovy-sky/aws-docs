This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Eac3Settings

The settings for an EAC3 audio encode in the output.

The parent of this entity is AudioCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttenuationControl" : String,
  "Bitrate" : Number,
  "BitstreamMode" : String,
  "CodingMode" : String,
  "DcFilter" : String,
  "Dialnorm" : Integer,
  "DrcLine" : String,
  "DrcRf" : String,
  "LfeControl" : String,
  "LfeFilter" : String,
  "LoRoCenterMixLevel" : Number,
  "LoRoSurroundMixLevel" : Number,
  "LtRtCenterMixLevel" : Number,
  "LtRtSurroundMixLevel" : Number,
  "MetadataControl" : String,
  "PassthroughControl" : String,
  "PhaseControl" : String,
  "StereoDownmix" : String,
  "SurroundExMode" : String,
  "SurroundMode" : String
}

```

### YAML

```yaml

  AttenuationControl: String
  Bitrate: Number
  BitstreamMode: String
  CodingMode: String
  DcFilter: String
  Dialnorm: Integer
  DrcLine: String
  DrcRf: String
  LfeControl: String
  LfeFilter: String
  LoRoCenterMixLevel: Number
  LoRoSurroundMixLevel: Number
  LtRtCenterMixLevel: Number
  LtRtSurroundMixLevel: Number
  MetadataControl: String
  PassthroughControl: String
  PhaseControl: String
  StereoDownmix: String
  SurroundExMode: String
  SurroundMode: String

```

## Properties

`AttenuationControl`

When set to attenuate3Db, applies a 3 dB attenuation to the surround channels. Used only for the 3/2 coding
mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bitrate`

The average bitrate in bits/second. Valid bitrates depend on the coding mode.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BitstreamMode`

Specifies the bitstream mode (bsmod) for the emitted E-AC-3 stream. For more information, see ATSC A/52-2012
(Annex E).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodingMode`

The Dolby Digital Plus coding mode. This mode determines the number of channels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DcFilter`

When set to enabled, activates a DC highpass filter for all input channels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dialnorm`

Sets the dialnorm for the output. If blank and the input audio is Dolby Digital Plus, dialnorm will be passed
through.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrcLine`

Sets the Dolby dynamic range compression profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrcRf`

Sets the profile for heavy Dolby dynamic range compression, ensuring that the instantaneous signal peaks do not
exceed specified levels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LfeControl`

When encoding 3/2 audio, setting to lfe enables the LFE channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LfeFilter`

When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding. Valid only with a
codingMode32 coding mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRoCenterMixLevel`

The Left only/Right only center mix level. Used only for the 3/2 coding mode.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRoSurroundMixLevel`

The Left only/Right only surround mix level. Used only for a 3/2 coding mode.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LtRtCenterMixLevel`

The Left total/Right total center mix level. Used only for a 3/2 coding mode.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LtRtSurroundMixLevel`

The Left total/Right total surround mix level. Used only for the 3/2 coding mode.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetadataControl`

When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this
audio data. If the audio is not supplied from one of these streams, then the static metadata settings are
used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PassthroughControl`

When set to whenPossible, input DD+ audio will be passed through if it is present on the input. This detection
is dynamic over the life of the transcode. Inputs that alternate between DD+ and non-DD+ content will have a
consistent DD+ output as the system alternates between passthrough and encoding.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhaseControl`

When set to shift90Degrees, applies a 90-degree phase shift to the surround channels. Used only for a 3/2 coding
mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StereoDownmix`

A stereo downmix preference. Used only for the 3/2 coding mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SurroundExMode`

When encoding 3/2 audio, sets whether an extra center back surround channel is matrix encoded into the left and
right surround channels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SurroundMode`

When encoding 2/0 audio, sets whether Dolby Surround is matrix-encoded into the two channels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Eac3AtmosSettings

EbuTtDDestinationSettings

All content copied from https://docs.aws.amazon.com/.
