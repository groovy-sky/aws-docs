This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VideoCodecSettings

The settings for the video codec in the output.

The parent of this entity is VideoDescription.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Av1Settings" : Av1Settings,
  "FrameCaptureSettings" : FrameCaptureSettings,
  "H264Settings" : H264Settings,
  "H265Settings" : H265Settings,
  "Mpeg2Settings" : Mpeg2Settings
}

```

### YAML

```yaml

  Av1Settings:
    Av1Settings
  FrameCaptureSettings:
    FrameCaptureSettings
  H264Settings:
    H264Settings
  H265Settings:
    H265Settings
  Mpeg2Settings:
    Mpeg2Settings

```

## Properties

`Av1Settings`

Property description not available.

_Required_: No

_Type_: [Av1Settings](aws-properties-medialive-channel-av1settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameCaptureSettings`

The settings for the video codec in a frame capture output.

_Required_: No

_Type_: [FrameCaptureSettings](aws-properties-medialive-channel-framecapturesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`H264Settings`

The settings for the H.264 codec in the output.

_Required_: No

_Type_: [H264Settings](aws-properties-medialive-channel-h264settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`H265Settings`

Settings for video encoded with the H265 codec.

_Required_: No

_Type_: [H265Settings](aws-properties-medialive-channel-h265settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mpeg2Settings`

Settings for video encoded with the MPEG-2 codec.

_Required_: No

_Type_: [Mpeg2Settings](aws-properties-medialive-channel-mpeg2settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoBlackFailoverSettings

VideoDescription

All content copied from https://docs.aws.amazon.com/.
