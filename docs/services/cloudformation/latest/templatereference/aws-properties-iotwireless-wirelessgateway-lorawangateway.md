This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessGateway LoRaWANGateway

LoRaWAN wireless gateway object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GatewayEui" : String,
  "RfRegion" : String
}

```

### YAML

```yaml

  GatewayEui: String
  RfRegion: String

```

## Properties

`GatewayEui`

The gateway's EUI value.

_Required_: Yes

_Type_: String

_Pattern_: `^(([0-9A-Fa-f]{2}-){7}|([0-9A-Fa-f]{2}:){7}|([0-9A-Fa-f]{2}\s){7}|([0-9A-Fa-f]{2}){7})([0-9A-Fa-f]{2})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RfRegion`

The frequency band (RFRegion) value.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::WirelessGateway

Tag

All content copied from https://docs.aws.amazon.com/.
