This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest CustomDecodingSignalDecoder

A list of information about signal decoders.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomDecodingSignal" : CustomDecodingSignal,
  "FullyQualifiedName" : String,
  "InterfaceId" : String,
  "Type" : String
}

```

### YAML

```yaml

  CustomDecodingSignal:
    CustomDecodingSignal
  FullyQualifiedName: String
  InterfaceId: String
  Type: String

```

## Properties

`CustomDecodingSignal`

Information about signals using a custom decoding protocol as defined by the customer.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

_Required_: Yes

_Type_: [CustomDecodingSignal](aws-properties-iotfleetwise-decodermanifest-customdecodingsignal.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

`Type`

The network protocol for the vehicle. For example, `CAN_SIGNAL` specifies a
protocol that defines how data is communicated between electronic control units (ECUs).
`OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data
is communicated between ECUs.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOM_DECODING_SIGNAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomDecodingSignal

NetworkInterfacesItems

All content copied from https://docs.aws.amazon.com/.
