This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel WavSettings

The setup of WAV audio in the output.

The parent of this entity is AudioCodecSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BitDepth" : Number,
  "CodingMode" : String,
  "SampleRate" : Number
}

```

### YAML

```yaml

  BitDepth: Number
  CodingMode: String
  SampleRate: Number

```

## Properties

`BitDepth`

Bits per sample.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodingMode`

The audio coding mode for the WAV audio. The mode determines the number of channels in the audio.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleRate`

Sample rate in Hz.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcOutputSettings

WebvttDestinationSettings

All content copied from https://docs.aws.amazon.com/.
