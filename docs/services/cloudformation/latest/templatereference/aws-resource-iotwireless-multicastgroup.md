This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::MulticastGroup

A multicast group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::MulticastGroup",
  "Properties" : {
      "AssociateWirelessDevice" : String,
      "Description" : String,
      "DisassociateWirelessDevice" : String,
      "LoRaWAN" : LoRaWAN,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::MulticastGroup
Properties:
  AssociateWirelessDevice: String
  Description: String
  DisassociateWirelessDevice: String
  LoRaWAN:
    LoRaWAN
  Name: String
  Tags:
    - Tag

```

## Properties

`AssociateWirelessDevice`

The ID of the wireless device to associate with a multicast group.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the multicast group.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisassociateWirelessDevice`

The ID of the wireless device to disassociate from a multicast group.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRaWAN`

The LoRaWAN information that is to be used with the multicast group.

_Required_: Yes

_Type_: [LoRaWAN](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-multicastgroup-lorawan.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the multicast group.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-multicastgroup-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the multicast group.

### Fn::GetAtt

`Arn`

The ARN of the multicast group.

`Id`

The ID of the multicast group.

`LoRaWAN.NumberOfDevicesInGroup`

The number of devices that are associated to the multicast group.

`LoRaWAN.NumberOfDevicesRequested`

The number of devices that are requested to be associated with the multicast
group.

`Status`

The status of a multicast group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

LoRaWAN
