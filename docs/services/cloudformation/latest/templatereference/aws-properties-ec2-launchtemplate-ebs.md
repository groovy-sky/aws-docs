This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate Ebs

Parameters for a block device for an EBS volume in an Amazon EC2 launch template.

`Ebs` is a property of [AWS::EC2::LaunchTemplate BlockDeviceMapping](../userguide/aws-properties-ec2-launchtemplate-blockdevicemapping.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteOnTermination" : Boolean,
  "EbsCardIndex" : Integer,
  "Encrypted" : Boolean,
  "Iops" : Integer,
  "KmsKeyId" : String,
  "SnapshotId" : String,
  "Throughput" : Integer,
  "VolumeInitializationRate" : Integer,
  "VolumeSize" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  DeleteOnTermination: Boolean
  EbsCardIndex: Integer
  Encrypted: Boolean
  Iops: Integer
  KmsKeyId: String
  SnapshotId: String
  Throughput: Integer
  VolumeInitializationRate: Integer
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`DeleteOnTermination`

Indicates whether the EBS volume is deleted on instance termination.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EbsCardIndex`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encrypted`

Indicates whether the EBS volume is encrypted. Encrypted volumes can only be attached
to instances that support Amazon EBS encryption. If you are creating a volume from a
snapshot, you can't specify an encryption value.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Iops`

The number of I/O operations per second (IOPS). For `gp3`,
`io1`, and `io2` volumes, this represents the number of IOPS that
are provisioned for the volume. For `gp2` volumes, this represents the
baseline performance of the volume and the rate at which the volume accumulates I/O
credits for bursting.

The following are the supported values for each volume type:

- `gp3`: 3,000 - 80,000 IOPS

- `io1`: 100 - 64,000 IOPS

- `io2`: 100 - 256,000 IOPS

For `io2` volumes, you can achieve up to 256,000 IOPS on
[instances \
built on the Nitro System](../../../ec2/latest/instancetypes/ec2-nitro-instances.md). On other instances, you can achieve performance up to 32,000 IOPS.

This parameter is supported for `io1`, `io2`, and `gp3` volumes only.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

Identifier (key ID, key alias, key ARN, or alias ARN) of the customer managed KMS key to use for EBS encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotId`

The ID of the snapshot.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Throughput`

The throughput to provision for a `gp3` volume, with a maximum of 2,000
MiB/s.

Valid Range: Minimum value of 125. Maximum value of 2,000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeInitializationRate`

Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download
the snapshot blocks from Amazon S3 to the volume. This is also known as _volume_
_initialization_. Specifying a volume initialization rate ensures that the volume is
initialized at a predictable and consistent rate after creation.

This parameter is supported only for volumes created from snapshots. Omit this parameter
if:

- You want to create the volume using fast snapshot restore. You must specify a snapshot
that is enabled for fast snapshot restore. In this case, the volume is fully initialized
at creation.

###### Note

If you specify a snapshot that is enabled for fast snapshot restore and a volume initialization rate,
the volume will be initialized at the specified rate instead of fast snapshot restore.

- You want to create a volume that is initialized at the default rate.

For more information, see [Initialize Amazon EBS volumes](../../../ebs/latest/userguide/initalize-volume.md) in the _Amazon EC2 User Guide_.

Valid range: 100 - 300 MiB/s

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSize`

The size of the volume, in GiBs. You must specify either a snapshot ID or a volume
size. The following are the supported volumes sizes for each volume type:

- `gp2`: 1 - 16,384 GiB

- `gp3`: 1 - 65,536 GiB

- `io1`: 4 - 16,384 GiB

- `io2`: 4 - 65,536 GiB

- `st1` and `sc1`: 125 - 16,384 GiB

- `standard`: 1 - 1024 GiB

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeType`

The volume type. For more information, see [Amazon EBS volume types](../../../ebs/latest/userguide/ebs-volume-types.md) in the
_Amazon EBS User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplateEbsBlockDeviceRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplateebsblockdevicerequest.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreditSpecification

EnaSrdSpecification

All content copied from https://docs.aws.amazon.com/.
