This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VideoBlackFailoverSettings

MediaLive will perform a failover if content is considered black for the specified period.

The parent of this entity is FailoverConditionSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlackDetectThreshold" : Number,
  "VideoBlackThresholdMsec" : Integer
}

```

### YAML

```yaml

  BlackDetectThreshold: Number
  VideoBlackThresholdMsec: Integer

```

## Properties

`BlackDetectThreshold`

A value used in calculating the threshold below which MediaLive considers a pixel to be 'black'. For the input to be considered black, every pixel in a frame must be below this threshold. The threshold is calculated as a percentage (expressed as a decimal) of white. Therefore .1 means 10% white (or 90% black). Note how the formula works for any color depth. For example, if you set this field to 0.1 in 10-bit color depth: (1023\*0.1=102.3), which means a pixel value of 102 or less is 'black'. If you set this field to .1 in an 8-bit color depth: (255\*0.1=25.5), which means a pixel value of 25 or less is 'black'. The range is 0.0 to 1.0, with any number of decimal places.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoBlackThresholdMsec`

The amount of time (in milliseconds) that the active input must be black before automatic input failover occurs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UdpOutputSettings

VideoCodecSettings

All content copied from https://docs.aws.amazon.com/.
