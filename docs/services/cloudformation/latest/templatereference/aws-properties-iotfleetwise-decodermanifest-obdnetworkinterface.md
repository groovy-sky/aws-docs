This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::DecoderManifest ObdNetworkInterface

Information about a network interface specified by the On-board diagnostic (OBD) II
protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InterfaceId" : String,
  "ObdInterface" : ObdInterface,
  "Type" : String
}

```

### YAML

```yaml

  InterfaceId: String
  ObdInterface:
    ObdInterface
  Type: String

```

## Properties

`InterfaceId`

The ID of the network interface.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObdInterface`

Information about a network interface specified by the On-board diagnostic (OBD) II
protocol.

_Required_: Yes

_Type_: [ObdInterface](aws-properties-iotfleetwise-decodermanifest-obdinterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The network protocol for the vehicle. For example, `CAN_SIGNAL` specifies
a protocol that defines how data is communicated between electronic control units
(ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic
data is communicated between ECUs.

_Required_: Yes

_Type_: String

_Allowed values_: `OBD_INTERFACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObdInterface

ObdSignal

All content copied from https://docs.aws.amazon.com/.
