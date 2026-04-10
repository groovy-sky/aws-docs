This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Ac3Settings

The settings for an AC3 audio encode in the output.

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
  "Dialnorm" : Integer,
  "DrcProfile" : String,
  "LfeFilter" : String,
  "MetadataControl" : String
}

```

### YAML

```yaml

  AttenuationControl: String
  Bitrate: Number
  BitstreamMode: String
  CodingMode: String
  Dialnorm: Integer
  DrcProfile: String
  LfeFilter: String
  MetadataControl: String

```

## Properties

`AttenuationControl`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bitrate`

The average bitrate in bits/second. Valid bitrates depend on the coding mode.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BitstreamMode`

Specifies the bitstream mode (bsmod) for the emitted AC-3 stream. For more information about these values, see
ATSC A/52-2012.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodingMode`

The Dolby Digital coding mode. This determines the number of channels.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dialnorm`

Sets the dialnorm for the output. If excluded and the input audio is Dolby Digital, dialnorm is passed
through.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DrcProfile`

If set to filmStandard, adds dynamic range compression signaling to the output bitstream as defined in the Dolby
Digital specification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LfeFilter`

When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding. This is valid only in
codingMode32Lfe mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetadataControl`

When set to followInput, encoder metadata is sourced from the DD, DD+, or DolbyE decoder that supplies this
audio data. If the audio is supplied from one of these streams, the static metadata settings are used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AacSettings

AdditionalDestinations

All content copied from https://docs.aws.amazon.com/.
