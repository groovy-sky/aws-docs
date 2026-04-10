This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate BlockDeviceMapping

Specifies a block device mapping for a launch template. You must specify `DeviceName`
plus exactly one of the following properties: `Ebs`, `NoDevice`, or `VirtualName`.

`BlockDeviceMapping` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceName" : String,
  "Ebs" : Ebs,
  "NoDevice" : String,
  "VirtualName" : String
}

```

### YAML

```yaml

  DeviceName: String
  Ebs:
    Ebs
  NoDevice: String
  VirtualName: String

```

## Properties

`DeviceName`

The device name (for example, /dev/sdh or xvdh).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ebs`

Parameters used to automatically set up EBS volumes when the instance is
launched.

_Required_: No

_Type_: [Ebs](aws-properties-ec2-launchtemplate-ebs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoDevice`

To omit the device from the block device mapping, specify an empty string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualName`

The virtual device name (ephemeralN). Instance store volumes are numbered starting
from 0. An instance type with 2 available instance store volumes can specify mappings
for ephemeral0 and ephemeral1. The number of available instance store volumes depends on
the instance type. After you connect to the instance, you must mount the volume.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Block device mappings](../../../ec2/latest/userguide/block-device-mapping-concepts.md) in the _Amazon EC2 User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BaselinePerformanceFactors

CapacityReservationSpecification

All content copied from https://docs.aws.amazon.com/.
