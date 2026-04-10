This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel H264Settings

The settings for the H.264 codec in the output.

The parent of this entity is VideoCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdaptiveQuantization" : String,
  "AfdSignaling" : String,
  "Bitrate" : Integer,
  "BufFillPct" : Integer,
  "BufSize" : Integer,
  "ColorMetadata" : String,
  "ColorSpaceSettings" : H264ColorSpaceSettings,
  "EntropyEncoding" : String,
  "FilterSettings" : H264FilterSettings,
  "FixedAfd" : String,
  "FlickerAq" : String,
  "ForceFieldPictures" : String,
  "FramerateControl" : String,
  "FramerateDenominator" : Integer,
  "FramerateNumerator" : Integer,
  "GopBReference" : String,
  "GopClosedCadence" : Integer,
  "GopNumBFrames" : Integer,
  "GopSize" : Number,
  "GopSizeUnits" : String,
  "Level" : String,
  "LookAheadRateControl" : String,
  "MaxBitrate" : Integer,
  "MinBitrate" : Integer,
  "MinIInterval" : Integer,
  "MinQp" : Integer,
  "NumRefFrames" : Integer,
  "ParControl" : String,
  "ParDenominator" : Integer,
  "ParNumerator" : Integer,
  "Profile" : String,
  "QualityLevel" : String,
  "QvbrQualityLevel" : Integer,
  "RateControlMode" : String,
  "ScanType" : String,
  "SceneChangeDetect" : String,
  "Slices" : Integer,
  "Softness" : Integer,
  "SpatialAq" : String,
  "SubgopLength" : String,
  "Syntax" : String,
  "TemporalAq" : String,
  "TimecodeBurninSettings" : TimecodeBurninSettings,
  "TimecodeInsertion" : String
}

```

### YAML

```yaml

  AdaptiveQuantization: String
  AfdSignaling: String
  Bitrate: Integer
  BufFillPct: Integer
  BufSize: Integer
  ColorMetadata: String
  ColorSpaceSettings:
    H264ColorSpaceSettings
  EntropyEncoding: String
  FilterSettings:
    H264FilterSettings
  FixedAfd: String
  FlickerAq: String
  ForceFieldPictures: String
  FramerateControl: String
  FramerateDenominator: Integer
  FramerateNumerator: Integer
  GopBReference: String
  GopClosedCadence: Integer
  GopNumBFrames: Integer
  GopSize: Number
  GopSizeUnits: String
  Level: String
  LookAheadRateControl: String
  MaxBitrate: Integer
  MinBitrate: Integer
  MinIInterval: Integer
  MinQp: Integer
  NumRefFrames: Integer
  ParControl: String
  ParDenominator: Integer
  ParNumerator: Integer
  Profile: String
  QualityLevel: String
  QvbrQualityLevel: Integer
  RateControlMode: String
  ScanType: String
  SceneChangeDetect: String
  Slices: Integer
  Softness: Integer
  SpatialAq: String
  SubgopLength: String
  Syntax: String
  TemporalAq: String
  TimecodeBurninSettings:
    TimecodeBurninSettings
  TimecodeInsertion: String

```

## Properties

`AdaptiveQuantization`

The adaptive quantization. This allows intra-frame quantizers to vary to improve visual quality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AfdSignaling`

Indicates that AFD values will be written into the output stream. If afdSignaling is auto, the system tries to
preserve the input AFD value (in cases where multiple AFD values are valid). If set to fixed, the AFD value is the
value configured in the fixedAfd parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bitrate`

The average bitrate in bits/second. This is required when the rate control mode is VBR or CBR. It isn't used for
QVBR. In a Microsoft Smooth output group, each output must have a unique value when its bitrate is rounded down to
the nearest multiple of 1000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufFillPct`

The percentage of the buffer that should initially be filled (HRD buffer model).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufSize`

The size of the buffer (HRD buffer model) in bits/second.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorMetadata`

Includes color space metadata in the output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSpaceSettings`

Settings to configure the color space handling for the video.

_Required_: No

_Type_: [H264ColorSpaceSettings](aws-properties-medialive-channel-h264colorspacesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntropyEncoding`

The entropy encoding mode. Use cabac (must be in Main or High profile) or cavlc.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterSettings`

Optional filters that you can apply to an encode.

_Required_: No

_Type_: [H264FilterSettings](aws-properties-medialive-channel-h264filtersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixedAfd`

A four-bit AFD value to write on all frames of video in the output stream. Valid only when afdSignaling is set
to Fixed.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlickerAq`

If set to enabled, adjusts the quantization within each frame to reduce flicker or pop on I-frames.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceFieldPictures`

This setting applies only when scan type is "interlaced." It controls whether coding is performed on a field basis or on a frame basis. (When the video is progressive, the coding is always performed on a frame basis.)
enabled: Force MediaLive to code on a field basis, so that odd and even sets of fields are coded separately.
disabled: Code the two sets of fields separately (on a field basis) or together (on a frame basis using PAFF), depending on what is most appropriate for the content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateControl`

Indicates how the output video frame rate is specified. If you select "specified," the output video frame rate
is determined by framerateNumerator and framerateDenominator. If you select "initializeFromSource," the output video
frame rate is set equal to the input video frame rate of the first input.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateDenominator`

The frame rate denominator.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateNumerator`

The frame rate numerator. The frame rate is a fraction, for example, 24000/1001 = 23.976 fps.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopBReference`

If enabled, uses reference B frames for GOP structures that have B frames > 1.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopClosedCadence`

The frequency of closed GOPs. In streaming applications, we recommend that you set this to 1 so that a decoder
joining mid-stream will receive an IDR frame as quickly as possible. Setting this value to 0 will break output
segmenting.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopNumBFrames`

The number of B-frames between reference frames.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopSize`

The GOP size (keyframe interval) in units of either frames or seconds per gopSizeUnits. The value must be
greater than zero.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopSizeUnits`

Indicates if the gopSize is specified in frames or seconds. If seconds, the system converts the gopSize into a
frame count at runtime.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Level`

The H.264 level.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LookAheadRateControl`

The amount of lookahead. A value of low can decrease latency and memory usage, while high can produce better
quality for certain content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxBitrate`

For QVBR: See the tooltip for Quality level. For VBR: Set the maximum bitrate in order to accommodate expected
spikes in the complexity of the video.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinBitrate`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinIInterval`

Meaningful only if sceneChangeDetect is set to enabled. This setting enforces separation between repeated
(cadence) I-frames and I-frames inserted by Scene Change Detection. If a scene change I-frame is within I-interval
frames of a cadence I-frame, the GOP is shrunk or stretched to the scene change I-frame. GOP stretch requires
enabling lookahead as well as setting the I-interval. The normal cadence resumes for the next GOP. Note that the
maximum GOP stretch = GOP size + Min-I-interval - 1.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinQp`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRefFrames`

The number of reference frames to use. The encoder might use more than requested if you use B-frames or
interlaced encoding.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParControl`

Indicates how the output pixel aspect ratio is specified. If "specified" is selected, the output video pixel
aspect ratio is determined by parNumerator and parDenominator. If "initializeFromSource" is selected, the output
pixels aspect ratio will be set equal to the input video pixel aspect ratio of the first input.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParDenominator`

The Pixel Aspect Ratio denominator.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParNumerator`

The Pixel Aspect Ratio numerator.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Profile`

An H.264 profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QualityLevel`

Leave as STANDARD\_QUALITY or choose a different value (which might result in additional costs to run the channel).
\- ENHANCED\_QUALITY: Produces a slightly better video quality without an increase in the bitrate. Has an effect only when the Rate control mode is QVBR or CBR. If this channel is in a MediaLive multiplex, the value must be ENHANCED\_QUALITY.
\- STANDARD\_QUALITY: Valid for any Rate control mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QvbrQualityLevel`

Controls the target quality for the video encode. This applies only when the rate control mode is QVBR. Set
values for the QVBR quality level field and Max bitrate field that suit your most important viewing devices.
Recommended values are: - Primary screen: Quality level: 8 to 10. Max bitrate: 4M - PC or tablet: Quality level: 7.
Max bitrate: 1.5M to 3M - Smartphone: Quality level: 6. Max bitrate: 1M to 1.5M.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RateControlMode`

The rate control mode. QVBR: The quality will match the specified quality level except when it is constrained by
the maximum bitrate. We recommend this if you or your viewers pay for bandwidth. VBR: The quality and bitrate vary,
depending on the video complexity. We recommend this instead of QVBR if you want to maintain a specific average
bitrate over the duration of the channel. CBR: The quality varies, depending on the video complexity. We recommend
this only if you distribute your assets to devices that can't handle variable bitrates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanType`

Sets the scan type of the output to progressive or top-field-first interlaced.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SceneChangeDetect`

The scene change detection. On: inserts I-frames when the scene change is detected. Off: does not force an
I-frame when the scene change is detected.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slices`

The number of slices per picture. The number must be less than or equal to the number of macroblock rows for
progressive pictures, and less than or equal to half the number of macroblock rows for interlaced pictures. This
field is optional. If you don't specify a value, MediaLive chooses the number of slices based on the encode
resolution.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Softness`

Softness. Selects a quantizer matrix. Larger values reduce high-frequency content in the encoded image.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpatialAq`

If set to enabled, adjusts quantization within each frame based on the spatial variation of content
complexity.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubgopLength`

If set to fixed, uses gopNumBFrames B-frames per sub-GOP. If set to dynamic, optimizes the number of B-frames
used for each sub-GOP to improve visual quality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Syntax`

Produces a bitstream that is compliant with SMPTE RP-2027.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemporalAq`

If set to enabled, adjusts quantization within each frame based on the temporal variation of content
complexity.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeBurninSettings`

Property description not available.

_Required_: No

_Type_: [TimecodeBurninSettings](aws-properties-medialive-channel-timecodeburninsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeInsertion`

Determines how timecodes should be inserted into the video elementary stream. disabled: don't include timecodes.
picTimingSei: pass through picture timing SEI messages from the source specified in Timecode Config.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

H264FilterSettings

H265ColorSpaceSettings

All content copied from https://docs.aws.amazon.com/.
