This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig VolumeSpecification

`VolumeSpecification` is a subproperty of the `EbsBlockDeviceConfig` property type. `VolumeSecification` determines the volume type, IOPS, and size (GiB) for EBS volumes attached to EC2 instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iops" : Integer,
  "SizeInGB" : Integer,
  "Throughput" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  Iops: Integer
  SizeInGB: Integer
  Throughput: Integer
  VolumeType: String

```

## Properties

`Iops`

The number of I/O operations per second (IOPS) that the volume supports.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SizeInGB`

The volume size, in gibibytes (GiB). This can be a number from 1 - 1024. If the volume
type is EBS-optimized, the minimum value is 10.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Throughput`

The throughput, in mebibyte per second (MiB/s). This optional parameter can be a number
from 125 - 1000 and is valid only for gp3 volumes.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeType`

The volume type. Volume types supported are gp3, gp2, io1, st1, sc1, and
standard.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SimpleScalingPolicyConfiguration

AWS::EMR::SecurityConfiguration

All content copied from https://docs.aws.amazon.com/.
