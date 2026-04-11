This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel H265Settings

H265 Settings

The parent of this entity is VideoCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdaptiveQuantization" : String,
  "AfdSignaling" : String,
  "AlternativeTransferFunction" : String,
  "Bitrate" : Integer,
  "BufSize" : Integer,
  "ColorMetadata" : String,
  "ColorSpaceSettings" : H265ColorSpaceSettings,
  "Deblocking" : String,
  "FilterSettings" : H265FilterSettings,
  "FixedAfd" : String,
  "FlickerAq" : String,
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
  "MvOverPictureBoundaries" : String,
  "MvTemporalPredictor" : String,
  "ParDenominator" : Integer,
  "ParNumerator" : Integer,
  "Profile" : String,
  "QvbrQualityLevel" : Integer,
  "RateControlMode" : String,
  "ScanType" : String,
  "SceneChangeDetect" : String,
  "Slices" : Integer,
  "SubgopLength" : String,
  "Tier" : String,
  "TileHeight" : Integer,
  "TilePadding" : String,
  "TileWidth" : Integer,
  "TimecodeBurninSettings" : TimecodeBurninSettings,
  "TimecodeInsertion" : String,
  "TreeblockSize" : String
}

```

### YAML

```yaml

  AdaptiveQuantization: String
  AfdSignaling: String
  AlternativeTransferFunction: String
  Bitrate: Integer
  BufSize: Integer
  ColorMetadata: String
  ColorSpaceSettings:
    H265ColorSpaceSettings
  Deblocking: String
  FilterSettings:
    H265FilterSettings
  FixedAfd: String
  FlickerAq: String
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
  MvOverPictureBoundaries: String
  MvTemporalPredictor: String
  ParDenominator: Integer
  ParNumerator: Integer
  Profile: String
  QvbrQualityLevel: Integer
  RateControlMode: String
  ScanType: String
  SceneChangeDetect: String
  Slices: Integer
  SubgopLength: String
  Tier: String
  TileHeight: Integer
  TilePadding: String
  TileWidth: Integer
  TimecodeBurninSettings:
    TimecodeBurninSettings
  TimecodeInsertion: String
  TreeblockSize: String

```

## Properties

`AdaptiveQuantization`

Adaptive quantization. Allows intra-frame quantizers to vary to improve visual quality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AfdSignaling`

Indicates that AFD values will be written into the output stream. If afdSignaling is "auto", the system will try to preserve the input AFD value (in cases where multiple AFD values are valid). If set to "fixed", the AFD value will be the value configured in the fixedAfd parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlternativeTransferFunction`

Whether or not EML should insert an Alternative Transfer Function SEI message to support backwards compatibility with non-HDR decoders and displays.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bitrate`

Average bitrate in bits/second. Required when the rate control mode is VBR or CBR. Not used for QVBR. In an MS Smooth output group, each output must have a unique value when its bitrate is rounded down to the nearest multiple of 1000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufSize`

Size of buffer (HRD buffer model) in bits.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorMetadata`

Includes colorspace metadata in the output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSpaceSettings`

Color Space settings

_Required_: No

_Type_: [H265ColorSpaceSettings](aws-properties-medialive-channel-h265colorspacesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Deblocking`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterSettings`

Optional filters that you can apply to an encode.

_Required_: No

_Type_: [H265FilterSettings](aws-properties-medialive-channel-h265filtersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixedAfd`

Four bit AFD value to write on all frames of video in the output stream. Only valid when afdSignaling is set to 'Fixed'.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlickerAq`

If set to enabled, adjust quantization within each frame to reduce flicker or 'pop' on I-frames.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateDenominator`

Framerate denominator.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FramerateNumerator`

Framerate numerator - framerate is a fraction, e.g. 24000 / 1001 = 23.976 fps.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopBReference`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopClosedCadence`

Frequency of closed GOPs. In streaming applications, it is recommended that this be set to 1 so a decoder joining mid-stream will receive an IDR frame as quickly as possible. Setting this value to 0 will break output segmenting.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopNumBFrames`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopSize`

GOP size (keyframe interval) in units of either frames or seconds per gopSizeUnits.
If gopSizeUnits is frames, gopSize must be an integer and must be greater than or equal to 1.
If gopSizeUnits is seconds, gopSize must be greater than 0, but need not be an integer.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GopSizeUnits`

Indicates if the gopSize is specified in frames or seconds. If seconds the system will convert the gopSize into a frame count at run time.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Level`

H.265 Level.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LookAheadRateControl`

Amount of lookahead. A value of low can decrease latency and memory usage, while high can produce better quality for certain content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxBitrate`

For QVBR: See the tooltip for Quality level

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinBitrate`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinIInterval`

Only meaningful if sceneChangeDetect is set to enabled. Defaults to 5 if multiplex rate control is used. Enforces separation between repeated (cadence) I-frames and I-frames inserted by Scene Change Detection. If a scene change I-frame is within I-interval frames of a cadence I-frame, the GOP is shrunk and/or stretched to the scene change I-frame. GOP stretch requires enabling lookahead as well as setting I-interval. The normal cadence resumes for the next GOP. Note: Maximum GOP stretch = GOP size + Min-I-interval - 1

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinQp`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MvOverPictureBoundaries`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MvTemporalPredictor`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParDenominator`

Pixel Aspect Ratio denominator.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParNumerator`

Pixel Aspect Ratio numerator.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Profile`

H.265 Profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QvbrQualityLevel`

Controls the target quality for the video encode. Applies only when the rate control mode is QVBR. Set values for the QVBR quality level field and Max bitrate field that suit your most important viewing devices. Recommended values are:
\- Primary screen: Quality level: 8 to 10. Max bitrate: 4M
\- PC or tablet: Quality level: 7. Max bitrate: 1.5M to 3M
\- Smartphone: Quality level: 6. Max bitrate: 1M to 1.5M

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RateControlMode`

Rate control mode.

QVBR: Quality will match the specified quality level except when it is constrained by the
maximum bitrate. Recommended if you or your viewers pay for bandwidth.

CBR: Quality varies, depending on the video complexity. Recommended only if you distribute
your assets to devices that cannot handle variable bitrates.

Multiplex: This rate control mode is only supported (and is required) when the video is being
delivered to a MediaLive Multiplex in which case the rate control configuration is controlled
by the properties within the Multiplex Program.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanType`

Sets the scan type of the output to progressive or top-field-first interlaced.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SceneChangeDetect`

Scene change detection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slices`

Number of slices per picture. Must be less than or equal to the number of macroblock rows for progressive pictures, and less than or equal to half the number of macroblock rows for interlaced pictures.
This field is optional; when no value is specified the encoder will choose the number of slices based on encode resolution.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubgopLength`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

H.265 Tier.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TileHeight`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TilePadding`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TileWidth`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeBurninSettings`

Property description not available.

_Required_: No

_Type_: [TimecodeBurninSettings](aws-properties-medialive-channel-timecodeburninsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeInsertion`

Determines how timecodes should be inserted into the video elementary stream.
\- 'disabled': Do not include timecodes
\- 'picTimingSei': Pass through picture timing SEI messages from the source specified in Timecode Config

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreeblockSize`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

H265FilterSettings

Hdr10Settings

All content copied from https://docs.aws.amazon.com/.
