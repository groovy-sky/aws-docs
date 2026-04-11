This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::TaskDefinition LoRaWANUpdateGatewayTaskCreate

The signature used to verify the update firmware.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CurrentVersion" : LoRaWANGatewayVersion,
  "SigKeyCrc" : Integer,
  "UpdateSignature" : String,
  "UpdateVersion" : LoRaWANGatewayVersion
}

```

### YAML

```yaml

  CurrentVersion:
    LoRaWANGatewayVersion
  SigKeyCrc: Integer
  UpdateSignature: String
  UpdateVersion:
    LoRaWANGatewayVersion

```

## Properties

`CurrentVersion`

The version of the gateways that should receive the update.

_Required_: No

_Type_: [LoRaWANGatewayVersion](aws-properties-iotwireless-taskdefinition-lorawangatewayversion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigKeyCrc`

The CRC of the signature private key to check.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateSignature`

The signature used to verify the update firmware.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateVersion`

The firmware version to update the gateway to.

_Required_: No

_Type_: [LoRaWANGatewayVersion](aws-properties-iotwireless-taskdefinition-lorawangatewayversion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoRaWANGatewayVersion

LoRaWANUpdateGatewayTaskEntry

All content copied from https://docs.aws.amazon.com/.
