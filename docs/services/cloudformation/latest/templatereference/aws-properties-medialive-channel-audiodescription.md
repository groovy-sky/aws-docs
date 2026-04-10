This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioDescription

The encoding information for one output audio.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioDashRoles" : [ String, ... ],
  "AudioNormalizationSettings" : AudioNormalizationSettings,
  "AudioSelectorName" : String,
  "AudioType" : String,
  "AudioTypeControl" : String,
  "AudioWatermarkingSettings" : AudioWatermarkSettings,
  "CodecSettings" : AudioCodecSettings,
  "DvbDashAccessibility" : String,
  "LanguageCode" : String,
  "LanguageCodeControl" : String,
  "Name" : String,
  "RemixSettings" : RemixSettings,
  "StreamName" : String
}

```

### YAML

```yaml

  AudioDashRoles:
    - String
  AudioNormalizationSettings:
    AudioNormalizationSettings
  AudioSelectorName: String
  AudioType: String
  AudioTypeControl: String
  AudioWatermarkingSettings:
    AudioWatermarkSettings
  CodecSettings:
    AudioCodecSettings
  DvbDashAccessibility: String
  LanguageCode: String
  LanguageCodeControl: String
  Name: String
  RemixSettings:
    RemixSettings
  StreamName: String

```

## Properties

`AudioDashRoles`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioNormalizationSettings`

The advanced audio normalization settings.

_Required_: No

_Type_: [AudioNormalizationSettings](aws-properties-medialive-channel-audionormalizationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioSelectorName`

The name of the AudioSelector that is used as the source for this AudioDescription.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioType`

Applies only if audioTypeControl is useConfigured. The values for audioType are defined in ISO-IEC
13818-1.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioTypeControl`

Determines how audio type is determined. followInput: If the input contains an ISO 639 audioType, then that
value is passed through to the output. If the input contains no ISO 639 audioType, the value in Audio Type is
included in the output. useConfigured: The value in Audio Type is included in the output. Note that this field and
audioType are both ignored if inputType is broadcasterMixedAd.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioWatermarkingSettings`

Settings to configure one or more solutions that insert audio watermarks in the audio encode

_Required_: No

_Type_: [AudioWatermarkSettings](aws-properties-medialive-channel-audiowatermarksettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodecSettings`

The audio codec settings.

_Required_: No

_Type_: [AudioCodecSettings](aws-properties-medialive-channel-audiocodecsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DvbDashAccessibility`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageCode`

Indicates the language of the audio output track. Used only if languageControlMode is useConfigured, or there is
no ISO 639 language code specified in the input.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageCodeControl`

Choosing followInput causes the ISO 639 language code of the output to follow the ISO 639 language code of the
input. The languageCode setting is used when useConfigured is set, or when followInput is selected but there is no
ISO 639 language code specified by the input.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of this AudioDescription. Outputs use this name to uniquely identify this AudioDescription. Description
names should be unique within this channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemixSettings`

The settings that control how input audio channels are remixed into the output audio channels.

_Required_: No

_Type_: [RemixSettings](aws-properties-medialive-channel-remixsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamName`

Used for Microsoft Smooth and Apple HLS outputs. Indicates the name displayed by the player (for example,
English or Director Commentary).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioCodecSettings

AudioDolbyEDecode

All content copied from https://docs.aws.amazon.com/.
