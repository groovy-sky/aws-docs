This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplex MultiplexSettings

Contains configuration for a Multiplex event

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumVideoBufferDelayMilliseconds" : Integer,
  "TransportStreamBitrate" : Integer,
  "TransportStreamId" : Integer,
  "TransportStreamReservedBitrate" : Integer
}

```

### YAML

```yaml

  MaximumVideoBufferDelayMilliseconds: Integer
  TransportStreamBitrate: Integer
  TransportStreamId: Integer
  TransportStreamReservedBitrate: Integer

```

## Properties

`MaximumVideoBufferDelayMilliseconds`

Maximum video buffer delay in milliseconds.

_Required_: No

_Type_: Integer

_Minimum_: `800`

_Maximum_: `3000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportStreamBitrate`

Transport stream bit rate.

_Required_: Yes

_Type_: Integer

_Minimum_: `1000000`

_Maximum_: `100000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportStreamId`

Transport stream ID.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportStreamReservedBitrate`

Transport stream reserved bit rate.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexOutputDestination

Tags

All content copied from https://docs.aws.amazon.com/.
