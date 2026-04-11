This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel EncoderSettings

The settings for the encoding of outputs.

This entity is at the top level in the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioDescriptions" : [ AudioDescription, ... ],
  "AvailBlanking" : AvailBlanking,
  "AvailConfiguration" : AvailConfiguration,
  "BlackoutSlate" : BlackoutSlate,
  "CaptionDescriptions" : [ CaptionDescription, ... ],
  "ColorCorrectionSettings" : ColorCorrectionSettings,
  "FeatureActivations" : FeatureActivations,
  "GlobalConfiguration" : GlobalConfiguration,
  "MotionGraphicsConfiguration" : MotionGraphicsConfiguration,
  "NielsenConfiguration" : NielsenConfiguration,
  "OutputGroups" : [ OutputGroup, ... ],
  "ThumbnailConfiguration" : ThumbnailConfiguration,
  "TimecodeConfig" : TimecodeConfig,
  "VideoDescriptions" : [ VideoDescription, ... ]
}

```

### YAML

```yaml

  AudioDescriptions:
    - AudioDescription
  AvailBlanking:
    AvailBlanking
  AvailConfiguration:
    AvailConfiguration
  BlackoutSlate:
    BlackoutSlate
  CaptionDescriptions:
    - CaptionDescription
  ColorCorrectionSettings:
    ColorCorrectionSettings
  FeatureActivations:
    FeatureActivations
  GlobalConfiguration:
    GlobalConfiguration
  MotionGraphicsConfiguration:
    MotionGraphicsConfiguration
  NielsenConfiguration:
    NielsenConfiguration
  OutputGroups:
    - OutputGroup
  ThumbnailConfiguration:
    ThumbnailConfiguration
  TimecodeConfig:
    TimecodeConfig
  VideoDescriptions:
    - VideoDescription

```

## Properties

`AudioDescriptions`

The encoding information for output audio.

_Required_: No

_Type_: Array of [AudioDescription](aws-properties-medialive-channel-audiodescription.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailBlanking`

The settings for ad avail blanking.

_Required_: No

_Type_: [AvailBlanking](aws-properties-medialive-channel-availblanking.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailConfiguration`

The configuration settings for the ad avail handling.

_Required_: No

_Type_: [AvailConfiguration](aws-properties-medialive-channel-availconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlackoutSlate`

The settings for the blackout slate.

_Required_: No

_Type_: [BlackoutSlate](aws-properties-medialive-channel-blackoutslate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionDescriptions`

The encoding information for output captions.

_Required_: No

_Type_: Array of [CaptionDescription](aws-properties-medialive-channel-captiondescription.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorCorrectionSettings`

Property description not available.

_Required_: No

_Type_: [ColorCorrectionSettings](aws-properties-medialive-channel-colorcorrectionsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FeatureActivations`

Settings to enable specific features.

_Required_: No

_Type_: [FeatureActivations](aws-properties-medialive-channel-featureactivations.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalConfiguration`

The configuration settings that apply to the entire channel.

_Required_: No

_Type_: [GlobalConfiguration](aws-properties-medialive-channel-globalconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MotionGraphicsConfiguration`

Settings to enable and configure the motion graphics overlay feature in the channel.

_Required_: No

_Type_: [MotionGraphicsConfiguration](aws-properties-medialive-channel-motiongraphicsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenConfiguration`

The settings to configure Nielsen watermarks.

_Required_: No

_Type_: [NielsenConfiguration](aws-properties-medialive-channel-nielsenconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputGroups`

The settings for the output groups in the channel.

_Required_: No

_Type_: Array of [OutputGroup](aws-properties-medialive-channel-outputgroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThumbnailConfiguration`

Property description not available.

_Required_: No

_Type_: [ThumbnailConfiguration](aws-properties-medialive-channel-thumbnailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeConfig`

Contains settings used to acquire and adjust timecode information from the inputs.

_Required_: No

_Type_: [TimecodeConfig](aws-properties-medialive-channel-timecodeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoDescriptions`

The encoding information for output videos.

_Required_: No

_Type_: Array of [VideoDescription](aws-properties-medialive-channel-videodescription.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmbeddedSourceSettings

EpochLockingSettings

All content copied from https://docs.aws.amazon.com/.
