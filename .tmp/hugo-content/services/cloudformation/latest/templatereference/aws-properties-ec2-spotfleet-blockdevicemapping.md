This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet BlockDeviceMapping

Specifies a block device mapping.

You can specify `Ebs` or `VirtualName`, but not both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceName" : String,
  "Ebs" : EbsBlockDevice,
  "NoDevice" : String,
  "VirtualName" : String
}

```

### YAML

```yaml

  DeviceName: String
  Ebs:
    EbsBlockDevice
  NoDevice: String
  VirtualName: String

```

## Properties

`DeviceName`

The device name. For available device names, see [Device names for volumes](../../../ec2/latest/userguide/device-naming.md).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ebs`

Parameters used to automatically set up EBS volumes when the instance is
launched.

_Required_: Conditional

_Type_: [EbsBlockDevice](aws-properties-ec2-spotfleet-ebsblockdevice.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NoDevice`

To omit the device from the block device mapping, specify an empty string. When this
property is specified, the device is removed from the block device mapping regardless of
the assigned value.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VirtualName`

The virtual device name ( `ephemeral` N). Instance store volumes are numbered
starting from 0. An instance type with 2 available instance store volumes can specify
mappings for `ephemeral0` and `ephemeral1`. The number of
available instance store volumes depends on the instance type. After you connect to the
instance, you must mount the volume.

NVMe instance store volumes are automatically enumerated and assigned a device name.
Including them in your block device mapping has no effect.

Constraints: For M3 instances, you must specify instance store volumes in the block
device mapping for the instance. When you launch an M3 instance, we ignore any instance
store volumes specified in the block device mapping for the AMI.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BaselinePerformanceFactorsRequest

ClassicLoadBalancer

All content copied from https://docs.aws.amazon.com/.
