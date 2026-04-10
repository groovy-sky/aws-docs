This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel TeletextSourceSettings

Information about the Teletext captions to extract from the input.

The parent of this entity is CaptionSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputRectangle" : CaptionRectangle,
  "PageNumber" : String
}

```

### YAML

```yaml

  OutputRectangle:
    CaptionRectangle
  PageNumber: String

```

## Properties

`OutputRectangle`

Settings to configure the caption rectangle for an output captions that will be created using this Teletext
source captions.

_Required_: No

_Type_: [CaptionRectangle](aws-properties-medialive-channel-captionrectangle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageNumber`

Specifies the Teletext page number within the data stream from which to extract captions. The range is 0x100
(256) to 0x8FF (2303). This is unused for passthrough. It should be specified as a hexadecimal string with no "0x"
prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StaticKeySettings

TemporalFilterSettings

All content copied from https://docs.aws.amazon.com/.
