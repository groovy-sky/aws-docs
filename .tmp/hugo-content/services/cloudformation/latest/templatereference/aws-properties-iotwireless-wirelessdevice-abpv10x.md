This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice AbpV10x

ABP device object for LoRaWAN specification v1.0.x

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DevAddr" : String,
  "SessionKeys" : SessionKeysAbpV10x
}

```

### YAML

```yaml

  DevAddr: String
  SessionKeys:
    SessionKeysAbpV10x

```

## Properties

`DevAddr`

The DevAddr value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{8}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionKeys`

Session keys for ABP v1.0.x.

_Required_: Yes

_Type_: [SessionKeysAbpV10x](aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::WirelessDevice

AbpV11

All content copied from https://docs.aws.amazon.com/.
