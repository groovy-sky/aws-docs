This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::TaskDefinition UpdateWirelessGatewayTaskCreate

UpdateWirelessGatewayTaskCreate object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LoRaWAN" : LoRaWANUpdateGatewayTaskCreate,
  "UpdateDataRole" : String,
  "UpdateDataSource" : String
}

```

### YAML

```yaml

  LoRaWAN:
    LoRaWANUpdateGatewayTaskCreate
  UpdateDataRole: String
  UpdateDataSource: String

```

## Properties

`LoRaWAN`

The properties that relate to the LoRaWAN wireless gateway.

_Required_: No

_Type_: [LoRaWANUpdateGatewayTaskCreate](aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskcreate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateDataRole`

The IAM role used to read data from the S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateDataSource`

The link to the S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoTWireless::WirelessDevice

All content copied from https://docs.aws.amazon.com/.
