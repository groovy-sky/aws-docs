This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputChannelLevel

The setting to remix the audio.

The parent of this entity is AudioChannelMappings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gain" : Integer,
  "InputChannel" : Integer
}

```

### YAML

```yaml

  Gain: Integer
  InputChannel: Integer

```

## Properties

`Gain`

The remixing value. Units are in dB, and acceptable values are within the range from -60 (mute) to 6 dB.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputChannel`

The index of the input channel that is used as a source.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputAttachment

InputLocation

All content copied from https://docs.aws.amazon.com/.
