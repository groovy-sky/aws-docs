This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::LaunchConfiguration BlockDevice

`BlockDevice` is a property of the `EBS` property of the [AWS::AutoScaling::LaunchConfiguration BlockDeviceMapping](../userguide/aws-properties-autoscaling-launchconfiguration-blockdevicemapping.md) property type that
describes an Amazon EBS volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteOnTermination" : Boolean,
  "Encrypted" : Boolean,
  "Iops" : Integer,
  "SnapshotId" : String,
  "Throughput" : Integer,
  "VolumeSize" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  DeleteOnTermination: Boolean
  Encrypted: Boolean
  Iops: Integer
  SnapshotId: String
  Throughput: Integer
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`DeleteOnTermination`

Indicates whether the volume is deleted on instance termination. For Amazon EC2 Auto Scaling, the
default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Encrypted`

Specifies whether the volume should be encrypted. Encrypted EBS volumes can only be
attached to instances that support Amazon EBS encryption. For more information, see [Requirements for Amazon EBS encryption](../../../ebs/latest/userguide/ebs-encryption-requirements.md) in the _Amazon EBS User Guide_. If your AMI uses encrypted volumes, you
can also only launch it on supported instance types.

###### Note

If you are creating a volume from a snapshot, you cannot create an unencrypted
volume from an encrypted snapshot. Also, you cannot specify a KMS key ID when using
a launch configuration.

If you enable encryption by default, the EBS volumes that you create are always
encrypted, either using the AWS managed KMS key or a customer-managed KMS key,
regardless of whether the snapshot was encrypted.

For more information, see [Use AWS KMS keys to encrypt Amazon EBS volumes](../../../autoscaling/ec2/userguide/ec2-auto-scaling-data-protection.md#encryption) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Iops`

The number of input/output (I/O) operations per second (IOPS) to provision for the volume.
For `gp3` and `io1` volumes, this represents the number of IOPS that are
provisioned for the volume. For `gp2` volumes, this represents the baseline
performance of the volume and the rate at which the volume accumulates I/O credits for
bursting.

The following are the supported values for each volume type:

- `gp3`: 3,000-16,000 IOPS

- `io1`: 100-64,000 IOPS

For `io1` volumes, we guarantee 64,000 IOPS only for [Instances built on the Nitro System](../../../ec2/latest/userguide/instance-types.md#ec2-nitro-instances). Other instance families guarantee performance
up to 32,000 IOPS.

`Iops` is supported when the volume type is `gp3` or `io1`
and required only when the volume type is `io1`. (Not used with
`standard`, `gp2`, `st1`, or `sc1` volumes.)

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotId`

The snapshot ID of the volume to use.

You must specify either a `VolumeSize` or a `SnapshotId`.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Throughput`

The throughput (MiBps) to provision for a `gp3` volume.

_Required_: No

_Type_: Integer

_Minimum_: `125`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeSize`

The volume size, in GiBs. The following are the supported volumes sizes for each
volume type:

- `gp2` and `gp3`: 1-16,384

- `io1`: 4-16,384

- `st1` and `sc1`: 125-16,384

- `standard`: 1-1,024

You must specify either a `SnapshotId` or a `VolumeSize`. If you
specify both `SnapshotId` and `VolumeSize`, the volume size must
be equal or greater than the size of the snapshot.

_Required_: Conditional

_Type_: Integer

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeType`

The volume type. For more information, see [Amazon EBS volume types](../../../ebs/latest/userguide/ebs-volume-types.md) in the
_Amazon EBS User Guide_.

Valid values: `standard` \| `io1` \| `gp2` \|
`st1` \| `sc1` \| `gp3`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Required\
KMS key policy for use with encrypted volumes](../../../autoscaling/ec2/userguide/key-policy-requirements-ebs-encryption.md) in the _Amazon EC2 Auto_
_Scaling User Guide_

- [Use\
encryption with EBS-backed AMIs](../../../ec2/latest/userguide/amiencryption.md) in the _Amazon EC2 User Guide for_
_Linux Instances_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AutoScaling::LaunchConfiguration

BlockDeviceMapping

All content copied from https://docs.aws.amazon.com/.
