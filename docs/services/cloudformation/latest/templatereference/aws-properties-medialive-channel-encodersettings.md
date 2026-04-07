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

_Type_: Array of [AudioDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-audiodescription.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailBlanking`

The settings for ad avail blanking.

_Required_: No

_Type_: [AvailBlanking](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-availblanking.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailConfiguration`

The configuration settings for the ad avail handling.

_Required_: No

_Type_: [AvailConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-availconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlackoutSlate`

The settings for the blackout slate.

_Required_: No

_Type_: [BlackoutSlate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-blackoutslate.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionDescriptions`

The encoding information for output captions.

_Required_: No

_Type_: Array of [CaptionDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-captiondescription.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorCorrectionSettings`

Property description not available.

_Required_: No

_Type_: [ColorCorrectionSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-colorcorrectionsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FeatureActivations`

Settings to enable specific features.

_Required_: No

_Type_: [FeatureActivations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-featureactivations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalConfiguration`

The configuration settings that apply to the entire channel.

_Required_: No

_Type_: [GlobalConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-globalconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MotionGraphicsConfiguration`

Settings to enable and configure the motion graphics overlay feature in the channel.

_Required_: No

_Type_: [MotionGraphicsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-motiongraphicsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenConfiguration`

The settings to configure Nielsen watermarks.

_Required_: No

_Type_: [NielsenConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-nielsenconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputGroups`

The settings for the output groups in the channel.

_Required_: No

_Type_: Array of [OutputGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-outputgroup.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThumbnailConfiguration`

Property description not available.

_Required_: No

_Type_: [ThumbnailConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-thumbnailconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimecodeConfig`

Contains settings used to acquire and adjust timecode information from the inputs.

_Required_: No

_Type_: [TimecodeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-timecodeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoDescriptions`

The encoding information for output videos.

_Required_: No

_Type_: Array of [VideoDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-videodescription.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EmbeddedSourceSettings

EpochLockingSettings
