This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplexprogram MultiplexStatmuxVideoSettings

Statmux rate control settings

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumBitrate" : Integer,
  "MinimumBitrate" : Integer,
  "Priority" : Integer
}

```

### YAML

```yaml

  MaximumBitrate: Integer
  MinimumBitrate: Integer
  Priority: Integer

```

## Properties

`MaximumBitrate`

Maximum statmux bitrate.

_Required_: No

_Type_: Integer

_Minimum_: `100000`

_Maximum_: `100000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumBitrate`

Minimum statmux bitrate.

_Required_: No

_Type_: Integer

_Minimum_: `100000`

_Maximum_: `100000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The purpose of the priority is to use a combination of the\\nmultiplex rate control algorithm and the QVBR capability of the\\nencoder to prioritize the video quality of some channels in a\\nmultiplex over others. Channels that have a higher priority will\\nget higher video quality at the expense of the video quality of\\nother channels in the multiplex with lower priority.

_Required_: No

_Type_: Integer

_Minimum_: `-5`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexProgramSettings

MultiplexVideoSettings

All content copied from https://docs.aws.amazon.com/.
