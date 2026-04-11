This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice OtaaV10x

OTAA device object for v1.0.x

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppEui" : String,
  "AppKey" : String
}

```

### YAML

```yaml

  AppEui: String
  AppKey: String

```

## Properties

`AppEui`

The AppEUI value. You specify this value when using LoRaWAN versions v1.0.2 or
v1.0.3.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{16}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppKey`

The AppKey value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoRaWANDevice

OtaaV11

All content copied from https://docs.aws.amazon.com/.
