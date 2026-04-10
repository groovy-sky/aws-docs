This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice LoRaWANDevice

LoRaWAN object for create functions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AbpV10x" : AbpV10x,
  "AbpV11" : AbpV11,
  "DevEui" : String,
  "DeviceProfileId" : String,
  "FPorts" : FPorts,
  "OtaaV10x" : OtaaV10x,
  "OtaaV11" : OtaaV11,
  "ServiceProfileId" : String
}

```

### YAML

```yaml

  AbpV10x:
    AbpV10x
  AbpV11:
    AbpV11
  DevEui: String
  DeviceProfileId: String
  FPorts:
    FPorts
  OtaaV10x:
    OtaaV10x
  OtaaV11:
    OtaaV11
  ServiceProfileId: String

```

## Properties

`AbpV10x`

ABP device object for LoRaWAN specification v1.0.x.

_Required_: No

_Type_: [AbpV10x](aws-properties-iotwireless-wirelessdevice-abpv10x.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AbpV11`

ABP device object for create APIs for v1.1.

_Required_: No

_Type_: [AbpV11](aws-properties-iotwireless-wirelessdevice-abpv11.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DevEui`

The DevEUI value.

_Required_: No

_Type_: String

_Pattern_: `[a-f0-9]{16}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceProfileId`

The ID of the device profile for the new wireless device.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FPorts`

List of FPort assigned for different LoRaWAN application packages to use.

_Required_: No

_Type_: [FPorts](aws-properties-iotwireless-wirelessdevice-fports.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OtaaV10x`

OTAA device object for create APIs for v1.0.x

_Required_: No

_Type_: [OtaaV10x](aws-properties-iotwireless-wirelessdevice-otaav10x.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OtaaV11`

OTAA device object for v1.1 for create APIs.

_Required_: No

_Type_: [OtaaV11](aws-properties-iotwireless-wirelessdevice-otaav11.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceProfileId`

The ID of the service profile.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FPorts

OtaaV10x

All content copied from https://docs.aws.amazon.com/.
