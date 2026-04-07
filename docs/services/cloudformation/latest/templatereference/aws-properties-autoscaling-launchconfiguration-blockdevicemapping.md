This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::LaunchConfiguration BlockDeviceMapping

`BlockDeviceMapping` specifies a block device mapping for the
`BlockDeviceMappings` property of the [AWS::AutoScaling::LaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html) resource.

Each instance that is launched has an associated root device volume, either an Amazon EBS
volume or an instance store volume. You can use block device mappings to specify additional
EBS volumes or instance store volumes to attach to an instance when it is launched.

For more information, see [Example block device mapping](../../../ec2/latest/userguide/block-device-mapping-concepts.md#block-device-mapping-ex) in the _Amazon EC2 User Guide for Linux_
_Instances_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceName" : String,
  "Ebs" : BlockDevice,
  "NoDevice" : Boolean,
  "VirtualName" : String
}

```

### YAML

```yaml

  DeviceName: String
  Ebs:
    BlockDevice
  NoDevice: Boolean
  VirtualName: String

```

## Properties

`DeviceName`

The device name assigned to the volume (for example, `/dev/sdh` or
`xvdh`). For more information, see [Device naming on Linux\
instances](../../../ec2/latest/userguide/device-naming.md) in the _Amazon EC2 User Guide_.

###### Note

To define a block device mapping, set the device name and exactly one of the
following properties: `Ebs`, `NoDevice`, or
`VirtualName`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ebs`

Information to attach an EBS volume to an instance at launch.

_Required_: No

_Type_: [BlockDevice](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-launchconfiguration-blockdevice.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NoDevice`

Setting this value to `true` prevents a volume that is included in the
block device mapping of the AMI from being mapped to the specified device name at
launch.

If `NoDevice` is `true` for the root device, instances might
fail the EC2 health check. In that case, Amazon EC2 Auto Scaling launches replacement instances.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VirtualName`

The name of the instance store volume (virtual device) to attach to an instance at
launch. The name must be in the form ephemeral _X_ where
_X_ is a number starting from zero (0), for example,
`ephemeral0`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Amazon\
EC2 instance store](../../../ec2/latest/userguide/instancestorage.md) in the _Amazon EC2 User Guide for Linux_
_Instances_

- [Amazon Elastic\
Block Store (Amazon EBS)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonEBS.html) in the _Amazon EC2 User Guide for Linux_
_Instances_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BlockDevice

MetadataOptions
