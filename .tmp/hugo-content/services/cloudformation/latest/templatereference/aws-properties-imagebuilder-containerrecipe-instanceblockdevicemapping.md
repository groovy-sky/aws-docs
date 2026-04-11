This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ContainerRecipe InstanceBlockDeviceMapping

Defines block device mappings for the instance used to configure your image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceName" : String,
  "Ebs" : EbsInstanceBlockDeviceSpecification,
  "NoDevice" : String,
  "VirtualName" : String
}

```

### YAML

```yaml

  DeviceName: String
  Ebs:
    EbsInstanceBlockDeviceSpecification
  NoDevice: String
  VirtualName: String

```

## Properties

`DeviceName`

The device to which these mappings apply.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ebs`

Use to manage Amazon EBS-specific configuration for this mapping.

_Required_: No

_Type_: [EbsInstanceBlockDeviceSpecification](aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NoDevice`

Use to remove a mapping from the base image.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VirtualName`

Use to manage instance ephemeral devices.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EbsInstanceBlockDeviceSpecification

InstanceConfiguration

All content copied from https://docs.aws.amazon.com/.
