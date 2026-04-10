This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VideoDescription

Encoding information for one output video.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodecSettings" : VideoCodecSettings,
  "Height" : Integer,
  "Name" : String,
  "RespondToAfd" : String,
  "ScalingBehavior" : String,
  "Sharpness" : Integer,
  "Width" : Integer
}

```

### YAML

```yaml

  CodecSettings:
    VideoCodecSettings
  Height: Integer
  Name: String
  RespondToAfd: String
  ScalingBehavior: String
  Sharpness: Integer
  Width: Integer

```

## Properties

`CodecSettings`

The video codec settings.

_Required_: No

_Type_: [VideoCodecSettings](aws-properties-medialive-channel-videocodecsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Height`

The output video height, in pixels. This must be an even number. For most codecs, you can keep this field and
width blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping
the field blank. For the Frame Capture codec, height and width are required.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of this VideoDescription. Outputs use this name to uniquely identify this description. Description
names should be unique within this channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RespondToAfd`

Indicates how to respond to the AFD values in the input stream. RESPOND causes input video to be clipped,
depending on the AFD value, input display aspect ratio, and output display aspect ratio, and (except for the
FRAMECAPTURE codec) includes the values in the output. PASSTHROUGH (does not apply to FRAMECAPTURE codec) ignores the
AFD values and includes the values in the output, so input video is not clipped. NONE ignores the AFD values and does
not include the values through to the output, so input video is not clipped.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingBehavior`

STRETCHTOOUTPUT configures the output position to stretch the video to the specified output resolution (height
and width). This option overrides any position value. DEFAULT might insert black boxes (pillar boxes or letter boxes)
around the video to provide the specified output resolution.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sharpness`

Changes the strength of the anti-alias filter used for scaling. 0 is the softest setting, and 100 is the
sharpest. We recommend a setting of 50 for most content.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Width`

The output video width, in pixels. It must be an even number. For most codecs, you can keep this field and
height blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping
the field blank. For the Frame Capture codec, height and width are required.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoCodecSettings

VideoSelector

All content copied from https://docs.aws.amazon.com/.
