This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest NetworkInterfacesItems

A list of information about available network interfaces.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanNetworkInterface" : CanNetworkInterface,
  "CustomDecodingNetworkInterface" : CustomDecodingNetworkInterface,
  "ObdNetworkInterface" : ObdNetworkInterface
}

```

### YAML

```yaml

  CanNetworkInterface:
    CanNetworkInterface
  CustomDecodingNetworkInterface:
    CustomDecodingNetworkInterface
  ObdNetworkInterface:
    ObdNetworkInterface

```

## Properties

`CanNetworkInterface`

A single controller area network (CAN) device interface.

_Required_: No

_Type_: [CanNetworkInterface](aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDecodingNetworkInterface`

Represents a custom network interface as defined by the customer.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

_Required_: No

_Type_: [CustomDecodingNetworkInterface](aws-properties-iotfleetwise-decodermanifest-customdecodingnetworkinterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObdNetworkInterface`

A network interface that specifies the on-board diagnostic (OBD) II network
protocol.

_Required_: No

_Type_: [ObdNetworkInterface](aws-properties-iotfleetwise-decodermanifest-obdnetworkinterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomDecodingSignalDecoder

ObdInterface

All content copied from https://docs.aws.amazon.com/.
