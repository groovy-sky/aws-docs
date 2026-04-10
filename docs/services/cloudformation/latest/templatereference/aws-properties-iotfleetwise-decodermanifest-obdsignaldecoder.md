This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest ObdSignalDecoder

A list of information about signal decoders.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FullyQualifiedName" : String,
  "InterfaceId" : String,
  "ObdSignal" : ObdSignal,
  "Type" : String
}

```

### YAML

```yaml

  FullyQualifiedName: String
  InterfaceId: String
  ObdSignal:
    ObdSignal
  Type: String

```

## Properties

`FullyQualifiedName`

The fully qualified name of a signal decoder as defined in a vehicle model.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterfaceId`

The ID of a network interface that specifies what network protocol a vehicle
follows.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObdSignal`

Information about signal messages using the on-board diagnostics (OBD) II protocol in
a vehicle.

_Required_: Yes

_Type_: [ObdSignal](aws-properties-iotfleetwise-decodermanifest-obdsignal.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The network protocol for the vehicle. For example, `CAN_SIGNAL` specifies a
protocol that defines how data is communicated between electronic control units (ECUs).
`OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data
is communicated between ECUs.

_Required_: Yes

_Type_: String

_Allowed values_: `OBD_SIGNAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObdSignal

SignalDecodersItems

All content copied from https://docs.aws.amazon.com/.
