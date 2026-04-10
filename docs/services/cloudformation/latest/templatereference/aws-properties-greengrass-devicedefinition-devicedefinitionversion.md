This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::DeviceDefinition DeviceDefinitionVersion

A
device definition version contains a list of [devices](../userguide/aws-properties-greengrass-devicedefinition-device.md).

###### Note

After you create a device definition version that contains the devices you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In
an CloudFormation template, `DeviceDefinitionVersion` is the property type
of the `InitialVersion` property in the [`AWS::Greengrass::DeviceDefinition`](../userguide/aws-resource-greengrass-devicedefinition.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Devices" : [ Device, ... ]
}

```

### YAML

```yaml

  Devices:
    - Device

```

## Properties

`Devices`

The devices in this version.

_Required_: Yes

_Type_: Array of [Device](aws-properties-greengrass-devicedefinition-device.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [DeviceDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-devicedefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Device

AWS::Greengrass::DeviceDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
