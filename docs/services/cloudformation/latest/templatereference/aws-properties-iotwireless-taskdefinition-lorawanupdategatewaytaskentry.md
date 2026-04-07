This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::TaskDefinition LoRaWANUpdateGatewayTaskEntry

LoRaWANUpdateGatewayTaskEntry object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CurrentVersion" : LoRaWANGatewayVersion,
  "UpdateVersion" : LoRaWANGatewayVersion
}

```

### YAML

```yaml

  CurrentVersion:
    LoRaWANGatewayVersion
  UpdateVersion:
    LoRaWANGatewayVersion

```

## Properties

`CurrentVersion`

The version of the gateways that should receive the update.

_Required_: No

_Type_: [LoRaWANGatewayVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-taskdefinition-lorawangatewayversion.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateVersion`

The firmware version to update the gateway to.

_Required_: No

_Type_: [LoRaWANGatewayVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-taskdefinition-lorawangatewayversion.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LoRaWANUpdateGatewayTaskCreate

Tag
