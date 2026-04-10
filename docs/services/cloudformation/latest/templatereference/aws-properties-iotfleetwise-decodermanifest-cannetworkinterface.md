This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest CanNetworkInterface

Represents a node and its specifications in an in-vehicle communication network. All
signal decoders must be associated with a network node.

To return this information about all the network interfaces specified in a decoder
manifest, use the [ListDecoderManifestNetworkInterfaces](../../../../reference/iot-fleetwise/latest/apireference/api-listdecodermanifestnetworkinterfaces.md) in the _AWS IoT FleetWise API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanInterface" : CanInterface,
  "InterfaceId" : String,
  "Type" : String
}

```

### YAML

```yaml

  CanInterface:
    CanInterface
  InterfaceId: String
  Type: String

```

## Properties

`CanInterface`

Information about a network interface specified by the Controller Area Network (CAN)
protocol.

_Required_: Yes

_Type_: [CanInterface](aws-properties-iotfleetwise-decodermanifest-caninterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterfaceId`

The ID of the network interface.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The network protocol for the vehicle. For example, `CAN_SIGNAL` specifies
a protocol that defines how data is communicated between electronic control units
(ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic
data is communicated between ECUs.

_Required_: Yes

_Type_: String

_Allowed values_: `CAN_INTERFACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CanInterface

CanSignal

All content copied from https://docs.aws.amazon.com/.
