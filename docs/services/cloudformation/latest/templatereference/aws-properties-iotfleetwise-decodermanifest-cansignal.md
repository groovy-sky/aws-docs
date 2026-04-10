This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest CanSignal

Information about a single controller area network (CAN) signal and the messages it
receives and transmits.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Factor" : String,
  "IsBigEndian" : String,
  "IsSigned" : String,
  "Length" : String,
  "MessageId" : String,
  "Name" : String,
  "Offset" : String,
  "SignalValueType" : String,
  "StartBit" : String
}

```

### YAML

```yaml

  Factor: String
  IsBigEndian: String
  IsSigned: String
  Length: String
  MessageId: String
  Name: String
  Offset: String
  SignalValueType: String
  StartBit: String

```

## Properties

`Factor`

A multiplier used to decode the CAN message.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsBigEndian`

Whether the byte ordering of a CAN message is big-endian.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsSigned`

Whether the message data is specified as a signed value.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Length`

How many bytes of data are in the message.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageId`

The ID of the message.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the signal.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Offset`

The offset used to calculate the signal value. Combined with factor, the calculation is `value = raw_value * factor + offset`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignalValueType`

The value type of the signal. The default value is `INTEGER`.

_Required_: No

_Type_: String

_Allowed values_: `INTEGER | FLOATING_POINT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartBit`

Indicates the beginning of the CAN message.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CanNetworkInterface

CanSignalDecoder

All content copied from https://docs.aws.amazon.com/.
