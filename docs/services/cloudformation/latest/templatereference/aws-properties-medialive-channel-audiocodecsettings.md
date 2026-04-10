This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioCodecSettings

The configuration of the audio codec in the audio output.

The parent of this entity is AudioDescription.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AacSettings" : AacSettings,
  "Ac3Settings" : Ac3Settings,
  "Eac3AtmosSettings" : Eac3AtmosSettings,
  "Eac3Settings" : Eac3Settings,
  "Mp2Settings" : Mp2Settings,
  "PassThroughSettings" : Json,
  "WavSettings" : WavSettings
}

```

### YAML

```yaml

  AacSettings:
    AacSettings
  Ac3Settings:
    Ac3Settings
  Eac3AtmosSettings:
    Eac3AtmosSettings
  Eac3Settings:
    Eac3Settings
  Mp2Settings:
    Mp2Settings
  PassThroughSettings: Json
  WavSettings:
    WavSettings

```

## Properties

`AacSettings`

The setup of the AAC audio codec in the output.

_Required_: No

_Type_: [AacSettings](aws-properties-medialive-channel-aacsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ac3Settings`

The setup of an AC3 audio codec in the output.

_Required_: No

_Type_: [Ac3Settings](aws-properties-medialive-channel-ac3settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Eac3AtmosSettings`

Property description not available.

_Required_: No

_Type_: [Eac3AtmosSettings](aws-properties-medialive-channel-eac3atmossettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Eac3Settings`

The setup of an EAC3 audio codec in the output.

_Required_: No

_Type_: [Eac3Settings](aws-properties-medialive-channel-eac3settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mp2Settings`

The setup of an MP2 audio codec in the output.

_Required_: No

_Type_: [Mp2Settings](aws-properties-medialive-channel-mp2settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PassThroughSettings`

The setup to pass through the Dolby audio codec to the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WavSettings`

Settings for audio encoded with the WAV codec.

_Required_: No

_Type_: [WavSettings](aws-properties-medialive-channel-wavsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioChannelMapping

AudioDescription

All content copied from https://docs.aws.amazon.com/.
