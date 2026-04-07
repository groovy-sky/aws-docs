This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImageRecipe InstanceBlockDeviceMapping

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

_Type_: [EbsInstanceBlockDeviceSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagerecipe-ebsinstanceblockdevicespecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NoDevice`

Enter an empty string to remove a mapping from the parent image.

The following is an example of an empty string value in the `NoDevice` field.

`NoDevice:""`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VirtualName`

Manages the instance ephemeral devices.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EbsInstanceBlockDeviceSpecification

LatestVersion
