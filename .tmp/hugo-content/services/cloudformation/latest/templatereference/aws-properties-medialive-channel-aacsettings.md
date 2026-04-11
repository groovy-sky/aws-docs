This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AacSettings

The settings for an AAC audio encode in the output.

The parent of this entity is AudioCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bitrate" : Number,
  "CodingMode" : String,
  "InputType" : String,
  "Profile" : String,
  "RateControlMode" : String,
  "RawFormat" : String,
  "SampleRate" : Number,
  "Spec" : String,
  "VbrQuality" : String
}

```

### YAML

```yaml

  Bitrate: Number
  CodingMode: String
  InputType: String
  Profile: String
  RateControlMode: String
  RawFormat: String
  SampleRate: Number
  Spec: String
  VbrQuality: String

```

## Properties

`Bitrate`

The average bitrate in bits/second. Valid values depend on the rate control mode and profile.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodingMode`

Mono, stereo, or 5.1 channel layout. Valid values depend on the rate control mode and profile. The adReceiverMix
setting receives a stereo description plus control track, and emits a mono AAC encode of the description track, with
control data emitted in the PES header as per ETSI TS 101 154 Annex E.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputType`

Set to broadcasterMixedAd when the input contains pre-mixed main audio + AD (narration) as a stereo pair. The
Audio Type field (audioType) will be set to 3, which signals to downstream systems that this stream contains
broadcaster mixed AD. Note that the input received by the encoder must contain pre-mixed audio; MediaLive does not
perform the mixing. The values in audioTypeControl and audioType (in AudioDescription) are ignored when set to
broadcasterMixedAd. Leave this set to normal when the input does not contain pre-mixed audio + AD.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Profile`

The AAC profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RateControlMode`

The rate control mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RawFormat`

Sets the LATM/LOAS AAC output for raw containers.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleRate`

The sample rate in Hz. Valid values depend on the rate control mode and profile.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Spec`

Uses MPEG-2 AAC audio instead of MPEG-4 AAC audio for raw or MPEG-2 Transport Stream containers.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VbrQuality`

The VBR quality level. This is used only if rateControlMode is VBR.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaLive::Channel

Ac3Settings

All content copied from https://docs.aws.amazon.com/.
