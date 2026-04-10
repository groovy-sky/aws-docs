This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessGateway

Provisions a wireless gateway.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::WirelessGateway",
  "Properties" : {
      "Description" : String,
      "LastUplinkReceivedAt" : String,
      "LoRaWAN" : LoRaWANGateway,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "ThingArn" : String,
      "ThingName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::WirelessGateway
Properties:
  Description: String
  LastUplinkReceivedAt: String
  LoRaWAN:
    LoRaWANGateway
  Name: String
  Tags:
    - Tag
  ThingArn: String
  ThingName: String

```

## Properties

`Description`

The description of the new resource. The maximum length is 2048 characters.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUplinkReceivedAt`

The date and time when the most recent uplink was received.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRaWAN`

The gateway configuration information to use to create the wireless gateway.

_Required_: Yes

_Type_: [LoRaWANGateway](aws-properties-iotwireless-wirelessgateway-lorawangateway.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the new resource.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-wirelessgateway-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingArn`

The ARN of the thing to associate with the wireless gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingName`

The name of the thing associated with the wireless gateway. The value is empty if a
thing isn't associated with the gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the wireless gateway ID.

### Fn::GetAtt

`Arn`

The ARN of the wireless gateway created.

`Id`

The ID of the wireless gateway created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LoRaWANGateway

All content copied from https://docs.aws.amazon.com/.
