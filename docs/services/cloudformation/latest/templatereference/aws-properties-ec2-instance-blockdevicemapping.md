This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance BlockDeviceMapping

Specifies a block device mapping for an instance. You must specify exactly one of the
following properties: `VirtualName`, `Ebs`, or
`NoDevice`.

`BlockDeviceMapping` is a property of the [AWS::EC2::Instance](../userguide/aws-properties-ec2-instance.md) resource.

###### Important

After the instance is running, you can modify only the
`DeleteOnTermination` parameter for the attached volumes without
interrupting the instance. Modifying any other parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceName" : String,
  "Ebs" : Ebs,
  "NoDevice" : Json,
  "VirtualName" : String
}

```

### YAML

```yaml

  DeviceName: String
  Ebs:
    Ebs
  NoDevice: Json
  VirtualName: String

```

## Properties

`DeviceName`

The device name. For available device names, see [Device names for volumes](../../../ec2/latest/userguide/device-naming.md).

###### Important

After the instance is running, this parameter is used to specify the device name of
the block device mapping to update.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Ebs`

Parameters used to automatically set up EBS volumes when the instance is
launched.

###### Important

After the instance is running, you can modify only the
`DeleteOnTermination` parameter for the attached volumes without
interrupting the instance. Modifying any other parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt).

_Required_: Conditional

_Type_: [Ebs](aws-properties-ec2-instance-ebs.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`NoDevice`

To omit the device from the block device mapping, specify an empty string.

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Conditional

_Type_: Json

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VirtualName`

The virtual device name ( `ephemeral` N). The name must be in the form
`ephemeral` _X_ where _X_ is a number
starting from zero (0). For example, an instance type with 2 available instance store
volumes can specify mappings for `ephemeral0` and `ephemeral1`. The
number of available instance store volumes depends on the instance type. After you connect
to the instance, you must mount the volume.

NVMe instance store volumes are automatically enumerated and assigned a device name.
Including them in your block device mapping has no effect.

_Constraints_: For M3 instances, you must specify instance store volumes
in the block device mapping for the instance. When you launch an M3 instance, we ignore any
instance store volumes specified in the block device mapping for the AMI.

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Conditional

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Examples

- [Block device mapping with two EBS volumes](#aws-properties-ec2-instance-blockdevicemapping--examples--Block_device_mapping_with_two_EBS_volumes)

- [Unmap an AMI-defined device](#aws-properties-ec2-instance-blockdevicemapping--examples--Unmap_an_AMI-defined_device)

### Block device mapping with two EBS volumes

This example sets the size of an EBS-backed root device to 50 GiB and
adds another EBS volume that is 100 GiB in size. The supported device
names depend on the AMI for the instance. The volume names in this
example are supported for Nitro-based Windows instances.

#### JSON

```json

"BlockDeviceMappings" : [
   {
      "DeviceName" : "/dev/sda1",
      "Ebs" : { "VolumeSize" : "50" }
   },
   {
      "DeviceName" : "xvdbm",
      "Ebs" : { "VolumeSize" : "100" }
   }
]
```

#### YAML

```yaml

BlockDeviceMappings:
  - DeviceName: /dev/sda1
    Ebs:
      VolumeSize: 50
  - DeviceName: xvdbm
    Ebs:
      VolumeSize: 100
```

### Unmap an AMI-defined device

To unmap a device defined in the AMI, set the `NoDevice` property to an
empty map, as shown here:

#### JSON

```json

"BlockDeviceMappings" : [
   {
      "DeviceName":"/dev/sde",
      "NoDevice": {}
   }
]
```

#### YAML

```yaml

BlockDeviceMappings:
  - DeviceName: /dev/sde
    NoDevice: {}
```

## See also

- [Block device mappings](../../../ec2/latest/userguide/block-device-mapping-concepts.md) in the _Amazon EC2 User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociationParameter

CpuOptions

All content copied from https://docs.aws.amazon.com/.
