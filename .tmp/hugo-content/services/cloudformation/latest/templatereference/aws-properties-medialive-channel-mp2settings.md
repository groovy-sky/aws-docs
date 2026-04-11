This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Mp2Settings

The configuration for this MP2 audio.

The parent of this entity is AudioCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bitrate" : Number,
  "CodingMode" : String,
  "SampleRate" : Number
}

```

### YAML

```yaml

  Bitrate: Number
  CodingMode: String
  SampleRate: Number

```

## Properties

`Bitrate`

The average bitrate in bits/second.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodingMode`

The MPEG2 Audio coding mode. Valid values are codingMode10 (for mono) or codingMode20 (for stereo).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleRate`

The sample rate in Hz.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MotionGraphicsSettings

Mpeg2FilterSettings

All content copied from https://docs.aws.amazon.com/.
