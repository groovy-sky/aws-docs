This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest ObdSignal

Information about signal messages using the on-board diagnostics (OBD) II protocol in
a vehicle.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BitMaskLength" : String,
  "BitRightShift" : String,
  "ByteLength" : String,
  "IsSigned" : String,
  "Offset" : String,
  "Pid" : String,
  "PidResponseLength" : String,
  "Scaling" : String,
  "ServiceMode" : String,
  "SignalValueType" : String,
  "StartByte" : String
}

```

### YAML

```yaml

  BitMaskLength: String
  BitRightShift: String
  ByteLength: String
  IsSigned: String
  Offset: String
  Pid: String
  PidResponseLength: String
  Scaling: String
  ServiceMode: String
  SignalValueType: String
  StartByte: String

```

## Properties

`BitMaskLength`

The number of bits to mask in a message.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BitRightShift`

The number of positions to shift bits in the message.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ByteLength`

The length of a message.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsSigned`

Determines whether the message is signed ( `true`) or not ( `false`). If it's signed, the message can represent both positive and negative numbers. The `isSigned` parameter only applies to the `INTEGER` raw signal type, and it doesn't affect the `FLOATING_POINT` raw signal type. The default value is `false`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Offset`

The offset used to calculate the signal value. Combined with scaling, the calculation is `value = raw_value * scaling + offset`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pid`

The diagnostic code used to request data from a vehicle for this signal.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PidResponseLength`

The length of the requested data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scaling`

A multiplier used to decode the message.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceMode`

The mode of operation (diagnostic service) in a message.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignalValueType`

The value type of the signal. The default value is `INTEGER`.

_Required_: No

_Type_: String

_Allowed values_: `INTEGER | FLOATING_POINT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartByte`

Indicates the beginning of the message.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObdNetworkInterface

ObdSignalDecoder

All content copied from https://docs.aws.amazon.com/.
