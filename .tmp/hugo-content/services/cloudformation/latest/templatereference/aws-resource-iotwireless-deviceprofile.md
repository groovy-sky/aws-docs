This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::DeviceProfile

Creates a new device profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::DeviceProfile",
  "Properties" : {
      "LoRaWAN" : LoRaWANDeviceProfile,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::DeviceProfile
Properties:
  LoRaWAN:
    LoRaWANDeviceProfile
  Name: String
  Tags:
    - Tag

```

## Properties

`LoRaWAN`

LoRaWAN device profile object.

_Required_: No

_Type_: [LoRaWANDeviceProfile](aws-properties-iotwireless-deviceprofile-lorawandeviceprofile.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the new resource.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-deviceprofile-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the device profile ID.

### Fn::GetAtt

`Arn`

The ARN of the device profile created.

`Id`

The ID of the device profile created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LoRaWANDeviceProfile

All content copied from https://docs.aws.amazon.com/.
