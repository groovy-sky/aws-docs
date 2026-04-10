This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice Application

A list of optional LoRaWAN application information, which can be used for
geolocation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationName" : String,
  "FPort" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  DestinationName: String
  FPort: Integer
  Type: String

```

## Properties

`DestinationName`

The name of the position data destination that describes the IoT rule that processes the
device's position data.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9-_]+`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FPort`

The name of the new destination for the device.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `223`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Application type, which can be specified to obtain real-time position information of
your LoRaWAN device.

_Required_: No

_Type_: String

_Allowed values_: `SemtechGeolocation | SemtechGNSS | SemtechGNSSNG | SemtechWiFi`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AbpV11

FPorts

All content copied from https://docs.aws.amazon.com/.
