This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::FuotaTask

A FUOTA task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::FuotaTask",
  "Properties" : {
      "AssociateMulticastGroup" : String,
      "AssociateWirelessDevice" : String,
      "Description" : String,
      "DisassociateMulticastGroup" : String,
      "DisassociateWirelessDevice" : String,
      "FirmwareUpdateImage" : String,
      "FirmwareUpdateRole" : String,
      "LoRaWAN" : LoRaWAN,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::FuotaTask
Properties:
  AssociateMulticastGroup: String
  AssociateWirelessDevice: String
  Description: String
  DisassociateMulticastGroup: String
  DisassociateWirelessDevice: String
  FirmwareUpdateImage: String
  FirmwareUpdateRole: String
  LoRaWAN:
    LoRaWAN
  Name: String
  Tags:
    - Tag

```

## Properties

`AssociateMulticastGroup`

The ID of the multicast group to associate with a FUOTA task.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociateWirelessDevice`

The ID of the wireless device to associate with a multicast group.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the new resource.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisassociateMulticastGroup`

The ID of the multicast group to disassociate from a FUOTA task.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisassociateWirelessDevice`

The ID of the wireless device to disassociate from a FUOTA task.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirmwareUpdateImage`

The S3 URI points to a firmware update image that is to be used with a FUOTA
task.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirmwareUpdateRole`

The firmware update role that is to be used with a FUOTA task.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRaWAN`

The LoRaWAN information used with a FUOTA task.

_Required_: Yes

_Type_: [LoRaWAN](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-fuotatask-lorawan.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a FUOTA task.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-fuotatask-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the FUOTA task.

### Fn::GetAtt

`Arn`

The ARN of a FUOTA task

`FuotaTaskStatus`

The status of a FUOTA task.

`Id`

The ID of a FUOTA task.

`LoRaWAN.StartTime`

Start time of a FUOTA task.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

LoRaWAN
