This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::TaskDefinition LoRaWANGatewayVersion

LoRaWANGatewayVersion object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Model" : String,
  "PackageVersion" : String,
  "Station" : String
}

```

### YAML

```yaml

  Model: String
  PackageVersion: String
  Station: String

```

## Properties

`Model`

The model number of the wireless gateway.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PackageVersion`

The version of the wireless gateway firmware.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Station`

The basic station version of the wireless gateway.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::TaskDefinition

LoRaWANUpdateGatewayTaskCreate

All content copied from https://docs.aws.amazon.com/.
