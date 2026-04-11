This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Mpeg2Settings

The settings for the MPEG-2 codec in the output.

The parent of this entity is VideoCodecSetting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdaptiveQuantization" : String,
  "AfdSignaling" : String,
  "ColorMetadata" : String,
  "ColorSpace" : String,
  "DisplayAspectRatio" : String,
  "FilterSettings" : Mpeg2FilterSettings,
  "FixedAfd" : String,
  "FramerateDenominator" : Integer,
  "FramerateNumerator" : Integer,
  "GopClosedCadence" : Integer,
  "GopNumBFrames" : Integer,
  "GopSize" : Number,
  "GopSizeUnits" : String,
  "ScanType" : String,
  "SubgopLength" : String,
  "TimecodeBurninSettings" : TimecodeBurninSettings,
  "TimecodeInsertion" : String
}

```

### YAML

```yaml

  AdaptiveQuantization: String
  AfdSignaling: String
  ColorMetadata: String
  ColorSpace: String
  DisplayAspectRatio: String
  FilterSettings:
    Mpeg2FilterSettings
  FixedAfd: String
  FramerateDenominator: Integer
  FramerateNumerator: Integer
  GopClosedCadence: Integer
  GopNumBFrames: Integer
  GopSize: Number
  GopSizeUnits: String
  ScanType: String
  SubgopLength: String
  TimecodeBurninSettings:
    TimecodeBurninSettings
  TimecodeInsertion: String

```

## Properties

`AdaptiveQuantization`

Choose Off to disable adaptive quantization. Or choose another value to enable the quantizer and set its strength. The strengths are: Auto, Off, Low, Medium, High. When you enable this field, MediaLive allows intra-frame quantizers to vary, which might improve visual quality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AfdSignaling`

Indicates the AFD values that MediaLive will write into the video encode. If you do not know what AFD signaling is, or if your downstream system has not given you guidance, choose AUTO.
AUTO: MediaLive will try to preserve the input AFD value (in cases where multiple AFD values are valid).
FIXED: MediaLive will use the value you specify in fixedAFD.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorMetadata`

Specifies whether to include the color space metadata. The metadata describes the color space that applies to the video (the colorSpace field). We recommend that you insert the metadata.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSpace`

Choose the type of color space conversion to apply to the output. For detailed information on setting up both the input and the output to obtain the desired color space in the output, see the section on \\"MediaLive Features - Video - color space\\" in the MediaLive User Guide.
PASSTHROUGH: Keep the color space of the input content - do not convert it.
AUTO:Convert all content that is SD to rec 601, and convert all content that is HD to rec 709.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayAspectRatio`

Sets the pixel aspect ratio for the encode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterSettings`

Optionally specify a noise reduction filter, which can improve quality of compressed content. If you do not choose a filter, no filter will be applied.
TEMPORAL: This filter is useful for both source content that is noisy (when it has excessive digital artifacts) and source content that is clean.
When the content is noisy, the filter cleans up the source content before the encoding phase, with these two effects: First, it improves the output video quality because the content has been cleaned up. Secondly, it decreases the bandwidth because MediaLive does not waste bits on encoding noise.
When the content is reasonably clean, the filter tends to decrease the bitrate.

_Required_: No

_Type_: [Mpeg2FilterSettings](aws-properties-medialive-channel-mpeg2filtersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixedAfd`

Complete this field only when afdSignaling is set to FIXED. Enter the AFD value (4 bits) to write on all frames of the video encode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateDenominator`

description": "The framerate denominator. For example, 1001. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateNumerator`

The framerate numerator. For example, 24000. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopClosedCadence`

MPEG2: default is open GOP.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopNumBFrames`

Relates to the GOP structure. The number of B-frames between reference frames. If you do not know what a B-frame is, use the default.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopSize`

Relates to the GOP structure. The GOP size (keyframe interval) in the units specified in gopSizeUnits. If you do not know what GOP is, use the default.
If gopSizeUnits is frames, then the gopSize must be an integer and must be greater than or equal to 1.
If gopSizeUnits is seconds, the gopSize must be greater than 0, but does not need to be an integer.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopSizeUnits`

Relates to the GOP structure. Specifies whether the gopSize is specified in frames or seconds. If you do not plan to change the default gopSize, leave the default. If you specify SECONDS, MediaLive will internally convert the gop size to a frame count.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanType`

Set the scan type of the output to PROGRESSIVE or INTERLACED (top field first).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubgopLength`

Relates to the GOP structure. If you do not know what GOP is, use the default.
FIXED: Set the number of B-frames in each sub-GOP to the value in gopNumBFrames.
DYNAMIC: Let MediaLive optimize the number of B-frames in each sub-GOP, to improve visual quality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeBurninSettings`

Property description not available.

_Required_: No

_Type_: [TimecodeBurninSettings](aws-properties-medialive-channel-timecodeburninsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeInsertion`

Determines how MediaLive inserts timecodes in the output video. For detailed information about setting up the input and the output for a timecode, see the section on \\"MediaLive Features - Timecode configuration\\" in the MediaLive User Guide.
DISABLED: do not include timecodes.
GOP\_TIMECODE: Include timecode metadata in the GOP header.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Mpeg2FilterSettings

MsSmoothGroupSettings

All content copied from https://docs.aws.amazon.com/.
