This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel HlsSettings

The settings for an HLS output.

The parent of this entity is HlsOutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioOnlyHlsSettings" : AudioOnlyHlsSettings,
  "Fmp4HlsSettings" : Fmp4HlsSettings,
  "FrameCaptureHlsSettings" : Json,
  "StandardHlsSettings" : StandardHlsSettings
}

```

### YAML

```yaml

  AudioOnlyHlsSettings:
    AudioOnlyHlsSettings
  Fmp4HlsSettings:
    Fmp4HlsSettings
  FrameCaptureHlsSettings: Json
  StandardHlsSettings:
    StandardHlsSettings

```

## Properties

`AudioOnlyHlsSettings`

The settings for an audio-only output.

_Required_: No

_Type_: [AudioOnlyHlsSettings](aws-properties-medialive-channel-audioonlyhlssettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fmp4HlsSettings`

The settings for an fMP4 container.

_Required_: No

_Type_: [Fmp4HlsSettings](aws-properties-medialive-channel-fmp4hlssettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameCaptureHlsSettings`

Settings for a frame capture output in an HLS output group.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandardHlsSettings`

The settings for a standard output (an output that is not audio-only).

_Required_: No

_Type_: [StandardHlsSettings](aws-properties-medialive-channel-standardhlssettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsS3Settings

HlsWebdavSettings

All content copied from https://docs.aws.amazon.com/.
