This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel RemixSettings

The settings for remixing audio in the output.

The parent of this entity is AudioDescription.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelMappings" : [ AudioChannelMapping, ... ],
  "ChannelsIn" : Integer,
  "ChannelsOut" : Integer
}

```

### YAML

```yaml

  ChannelMappings:
    - AudioChannelMapping
  ChannelsIn: Integer
  ChannelsOut: Integer

```

## Properties

`ChannelMappings`

A mapping of input channels to output channels, with appropriate gain adjustments.

_Required_: No

_Type_: Array of [AudioChannelMapping](aws-properties-medialive-channel-audiochannelmapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelsIn`

The number of input channels to be used.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelsOut`

The number of output channels to be produced. Valid values: 1, 2, 4, 6, 8.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrimaryChannelSettings

RtmpGroupSettings

All content copied from https://docs.aws.amazon.com/.
