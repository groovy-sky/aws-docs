This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow Fmtp

A set of parameters that define the media stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelOrder" : String,
  "Colorimetry" : String,
  "ExactFramerate" : String,
  "Par" : String,
  "Range" : String,
  "ScanMode" : String,
  "Tcs" : String
}

```

### YAML

```yaml

  ChannelOrder: String
  Colorimetry: String
  ExactFramerate: String
  Par: String
  Range: String
  ScanMode: String
  Tcs: String

```

## Properties

`ChannelOrder`

The format of the audio channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Colorimetry`

The format used for the representation of color.

_Required_: No

_Type_: String

_Allowed values_: `BT601 | BT709 | BT2020 | BT2100 | ST2065-1 | ST2065-3 | XYZ`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExactFramerate`

The frame rate for the video stream, in frames/second. For example: 60000/1001.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Par`

The pixel aspect ratio (PAR) of the video.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Range`

The encoding range of the video.

_Required_: No

_Type_: String

_Allowed values_: `NARROW | FULL | FULLPROTECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanMode`

The type of compression that was used to smooth the video’s appearance.

_Required_: No

_Type_: String

_Allowed values_: `progressive | interlace | progressive-segmented-frame`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tcs`

The transfer characteristic system (TCS) that is used in the video.

_Required_: No

_Type_: String

_Allowed values_: `SDR | PQ | HLG | LINEAR | BT2100LINPQ | BT2100LINHLG | ST2065-1 | ST428-1 | DENSITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowTransitEncryptionKeyConfiguration

FrozenFrames

All content copied from https://docs.aws.amazon.com/.
