This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Scte20SourceSettings

Information about the SCTE-20 captions to extract from the input.

The parent of this entity is CaptionSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Convert608To708" : String,
  "Source608ChannelNumber" : Integer
}

```

### YAML

```yaml

  Convert608To708: String
  Source608ChannelNumber: Integer

```

## Properties

`Convert608To708`

If upconvert, 608 data is both passed through the "608 compatibility bytes" fields of the 708 wrapper as well as
translated into 708. Any 708 data present in the source content is discarded.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source608ChannelNumber`

Specifies the 608/708 channel number within the video track from which to extract captions.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RtmpOutputSettings

Scte27SourceSettings

All content copied from https://docs.aws.amazon.com/.
