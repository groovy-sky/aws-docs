This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet EbsBlockDevice

Describes a block device for an EBS volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteOnTermination" : Boolean,
  "Encrypted" : Boolean,
  "Iops" : Integer,
  "KmsKeyId" : String,
  "SnapshotId" : String,
  "VolumeSize" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  DeleteOnTermination: Boolean
  Encrypted: Boolean
  Iops: Integer
  KmsKeyId: String
  SnapshotId: String
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`DeleteOnTermination`

Indicates whether the EBS volume is deleted on instance termination. For more
information, see [Preserving Amazon EBS volumes on instance termination](../../../ec2/latest/userguide/terminating-instances.md#preserving-volumes-on-termination) in the
_Amazon EC2 User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Encrypted`

Indicates whether the encryption state of an EBS volume is changed while being
restored from a backing snapshot. The effect of setting the encryption state to `true` depends on
the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by
default is enabled. For more information, see [Amazon EBS encryption](../../../ebs/latest/userguide/ebs-encryption.md#encryption-parameters)
in the _Amazon EBS User Guide_.

In no case can you remove encryption from an encrypted volume.

Encrypted volumes can only be attached to instances that support Amazon EBS encryption. For
more information, see [Supported instance types](../../../ebs/latest/userguide/ebs-encryption-requirements.md#ebs-encryption_supported_instances).

- If you are creating a block device mapping for a **new (empty)**
**volume**, you can include this parameter, and specify either `true`
for an encrypted volume, or `false` for an unencrypted volume. If you omit
this parameter, it defaults to `false` (unencrypted).

- If you are creating a block device mapping from an **existing**
**encrypted or unencrypted snapshot**, you must omit this parameter. If you
include this parameter, the request will fail, regardless of the value that you
specify.

- If you are creating a block device mapping from an **existing**
**unencrypted volume**, you can include this parameter, but you must specify
`false`. If you specify `true`, the request will fail. In this
case, we recommend that you omit the parameter.

- If you are creating a block device mapping from an **existing**
**encrypted volume**, you can include this parameter, and specify either
`true` or `false`. However, if you specify `false`,
the parameter is ignored and the block device mapping is always encrypted. In this
case, we recommend that you omit the parameter.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Iops`

The number of I/O operations per second (IOPS). For `gp3`, `io1`, and `io2` volumes,
this represents the number of IOPS that are provisioned for the volume. For `gp2`
volumes, this represents the baseline performance of the volume and the rate at which
the volume accumulates I/O credits for bursting.

The following are the supported values for each volume type:

- `gp3`: 3,000 - 80,000 IOPS

- `io1`: 100 - 64,000 IOPS

- `io2`: 100 - 256,000 IOPS

For `io2` volumes, you can achieve up to 256,000 IOPS on
[instances \
built on the Nitro System](../../../ec2/latest/userguide/instance-types.md#ec2-nitro-instances). On other instances, you can achieve performance up to 32,000 IOPS.

This parameter is required for `io1` and `io2` volumes. The default for `gp3` volumes
is 3,000 IOPS.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

Identifier (key ID, key alias, key ARN, or alias ARN) of the customer managed KMS key
to use for EBS encryption.

This parameter is only supported on `BlockDeviceMapping` objects called by
[RunInstances](../../../../reference/awsec2/latest/apireference/api-runinstances.md), [RequestSpotFleet](../../../../reference/awsec2/latest/apireference/api-requestspotfleet.md),
and [RequestSpotInstances](../../../../reference/awsec2/latest/apireference/api-requestspotinstances.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotId`

The ID of the snapshot.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeSize`

The size of the volume, in GiBs. You must specify either a snapshot ID or a volume
size. If you specify a snapshot, the default is the snapshot size. You can specify a
volume size that is equal to or larger than the snapshot size.

The following are the supported sizes for each volume type:

- `gp2`: 1 - 16,384 GiB

- `gp3`: 1 - 65,536 GiB

- `io1`: 4 - 16,384 GiB

- `io2`: 4 - 65,536 GiB

- `st1` and `sc1`: 125 - 16,384 GiB

- `standard`: 1 - 1024 GiB

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeType`

The volume type. For more information, see [Amazon EBS volume types](../../../ebs/latest/userguide/ebs-volume-types.md) in the
_Amazon EBS User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `gp2 | gp3 | io1 | io2 | sc1 | st1 | standard`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CpuPerformanceFactorRequest

FleetLaunchTemplateConfigRequest

All content copied from https://docs.aws.amazon.com/.
