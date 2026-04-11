This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint StreamSelection

Limitations for outputs from the endpoint, based on the video bitrate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxVideoBitsPerSecond" : Integer,
  "MinVideoBitsPerSecond" : Integer,
  "StreamOrder" : String
}

```

### YAML

```yaml

  MaxVideoBitsPerSecond: Integer
  MinVideoBitsPerSecond: Integer
  StreamOrder: String

```

## Properties

`MaxVideoBitsPerSecond`

The upper limit of the bitrates that this endpoint serves. If the video track exceeds this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 2147483647 bits per second.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinVideoBitsPerSecond`

The lower limit of the bitrates that this endpoint serves. If the video track is below this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 0 bits per second.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamOrder`

Order in which the different video bitrates are presented to the player.

Valid values: `ORIGINAL`, `VIDEO_BITRATE_ASCENDING`, `VIDEO_BITRATE_DESCENDING`.

_Required_: No

_Type_: String

_Allowed values_: `ORIGINAL | VIDEO_BITRATE_ASCENDING | VIDEO_BITRATE_DESCENDING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpekeKeyProvider

Tag

All content copied from https://docs.aws.amazon.com/.
