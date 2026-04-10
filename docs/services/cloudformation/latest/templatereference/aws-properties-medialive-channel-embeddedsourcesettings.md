This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel EmbeddedSourceSettings

Information about the embedded captions to extract from the input.

The parent of this entity is CaptionSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Convert608To708" : String,
  "Scte20Detection" : String,
  "Source608ChannelNumber" : Integer,
  "Source608TrackNumber" : Integer
}

```

### YAML

```yaml

  Convert608To708: String
  Scte20Detection: String
  Source608ChannelNumber: Integer
  Source608TrackNumber: Integer

```

## Properties

`Convert608To708`

If this is upconvert, 608 data is both passed through the "608 compatibility bytes" fields of the 708 wrapper as
well as translated into 708. If 708 data is present in the source content, it is discarded.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte20Detection`

Set to "auto" to handle streams with intermittent or non-aligned SCTE-20 and embedded captions.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source608ChannelNumber`

Specifies the 608/708 channel number within the video track from which to extract captions. This is unused for
passthrough.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source608TrackNumber`

This field is unused and deprecated.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EbuTtDDestinationSettings

EncoderSettings

All content copied from https://docs.aws.amazon.com/.
