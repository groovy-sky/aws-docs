This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::DeviceDefinition Device

A device is an
AWS IoT device (thing) that's added to a Greengrass group. Greengrass
devices can communicate with the Greengrass core in the same group. For more information,
see [What\
Is AWS IoT Greengrass?](../../../greengrass/v1/developerguide/what-is-gg.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, the `Devices` property of the [`DeviceDefinitionVersion`](../userguide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.md) property type contains a list of `Device` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "Id" : String,
  "SyncShadow" : Boolean,
  "ThingArn" : String
}

```

### YAML

```yaml

  CertificateArn: String
  Id: String
  SyncShadow: Boolean
  ThingArn: String

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) of the device certificate for the device. This X.509
certificate is used to authenticate the device with AWS IoT and AWS IoT Greengrass services.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Id`

A descriptive or arbitrary ID for the device. This value must be unique within the
device definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncShadow`

Indicates whether the device's local shadow is synced with the cloud
automatically.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThingArn`

The ARN of the device, which is an AWS IoT device (thing).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Device](../../../../reference/greengrass/v1/apireference/definitions-device.md) in
the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Greengrass::DeviceDefinition

DeviceDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
