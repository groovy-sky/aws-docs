# EbsBlockDevice

Describes a block device for an EBS volume.

## Contents

**AvailabilityZone** (request), **availabilityZone** (response)

The Availability Zone where the EBS volume will be created (for example,
`us-east-1a`).

Either `AvailabilityZone` or `AvailabilityZoneId` can be specified,
but not both. If neither is specified, Amazon EC2 automatically selects an Availability Zone within
the Region.

This parameter is not supported when using
[CreateFleet](api-createfleet.md),
[CreateImage](api-createimage.md),
[DescribeImages](api-describeimages.md),
[RequestSpotFleet](api-requestspotfleet.md),
[RequestSpotInstances](api-requestspotinstances.md), and
[RunInstances](api-runinstances.md).

Type: String

Required: No

**AvailabilityZoneId** (request), **AvailabilityZoneId** (response)

The ID of the Availability Zone where the EBS volume will be created (for example,
`use1-az1`).

Either `AvailabilityZone` or `AvailabilityZoneId` can be specified,
but not both. If neither is specified, Amazon EC2 automatically selects an Availability Zone within
the Region.

This parameter is not supported when using
[CreateFleet](api-createfleet.md),
[CreateImage](api-createimage.md),
[DescribeImages](api-describeimages.md),
[RequestSpotFleet](api-requestspotfleet.md),
[RequestSpotInstances](api-requestspotinstances.md), and
[RunInstances](api-runinstances.md).

Type: String

Required: No

**DeleteOnTermination** (request), **deleteOnTermination** (response)

Indicates whether the EBS volume is deleted on instance termination. For more
information, see [Preserving Amazon EBS volumes on instance termination](../../../../services/ec2/latest/userguide/terminating-instances.md#preserving-volumes-on-termination) in the
_Amazon EC2 User Guide_.

Type: Boolean

Required: No

**EbsCardIndex** (request), **EbsCardIndex** (response)

The index of the EBS card. Some instance types support multiple EBS cards. The default EBS card index is 0.

Type: Integer

Required: No

**Encrypted** (request), **encrypted** (response)

Indicates whether the encryption state of an EBS volume is changed while being
restored from a backing snapshot. The effect of setting the encryption state to `true` depends on
the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Amazon EBS encryption](../../../../services/ebs/latest/userguide/ebs-encryption.md#encryption-parameters) in the _Amazon EBS User Guide_.

In no case can you remove encryption from an encrypted volume.

Encrypted volumes can only be attached to instances that support Amazon EBS encryption. For
more information, see [Supported instance types](../../../../services/ebs/latest/userguide/ebs-encryption-requirements.md#ebs-encryption_supported_instances).

This parameter is not returned by [DescribeImageAttribute](api-describeimageattribute.md).

For [CreateImage](api-createimage.md) and [RegisterImage](api-registerimage.md), whether you can
include this parameter, and the allowed values differ depending on the type of block
device mapping you are creating.

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

Type: Boolean

Required: No

**Iops** (request), **iops** (response)

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
built on the Nitro System](../../../../services/ec2/latest/userguide/instance-types.md#ec2-nitro-instances). On other instances, you can achieve performance up to 32,000 IOPS.

This parameter is required for `io1` and `io2` volumes. The default for `gp3` volumes
is 3,000 IOPS.

Type: Integer

Required: No

**KmsKeyId** (request), **kmsKeyId** (response)

Identifier (key ID, key alias, key ARN, or alias ARN) of the customer managed KMS key
to use for EBS encryption.

This parameter is only supported on `BlockDeviceMapping` objects called by
[RunInstances](api-runinstances.md), [RequestSpotFleet](api-requestspotfleet.md),
and [RequestSpotInstances](api-requestspotinstances.md).

Type: String

Required: No

**OutpostArn** (request), **outpostArn** (response)

The ARN of the Outpost on which the snapshot is stored.

This parameter is not supported when using [CreateImage](api-createimage.md).

Type: String

Required: No

**SnapshotId** (request), **snapshotId** (response)

The ID of the snapshot.

Type: String

Required: No

**Throughput** (request), **throughput** (response)

The throughput that the volume supports, in MiB/s.

This parameter is valid only for `gp3` volumes.

Valid Range: Minimum value of 125. Maximum value of 2,000.

Type: Integer

Required: No

**VolumeInitializationRate** (request), **VolumeInitializationRate** (response)

Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download the snapshot
blocks from Amazon S3 to the volume. This is also known as _volume_
_initialization_. Specifying a volume initialization rate ensures that the volume is initialized
at a predictable and consistent rate after creation. For more information, see
[Initialize \
Amazon EBS volumes](../../../../services/ebs/latest/userguide/initalize-volume.md) in the _Amazon EC2 User Guide_.

This parameter is supported only for volumes created from snapshots. Omit this parameter
if:

- You want to create the volume using fast snapshot restore. You must specify a snapshot
that is enabled for fast snapshot restore. In this case, the volume is fully initialized at
creation.

###### Note

If you specify a snapshot that is enabled for fast snapshot restore and a volume initialization rate,
the volume will be initialized at the specified rate instead of fast snapshot restore.

- You want to create a volume that is initialized at the default rate.

This parameter is not supported when using [CreateImage](api-createimage.md)
and [DescribeImages](api-describeimages.md).

Valid range: 100 - 300 MiB/s

Type: Integer

Required: No

**VolumeSize** (request), **volumeSize** (response)

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

Type: Integer

Required: No

**VolumeType** (request), **volumeType** (response)

The volume type. For more information, see [Amazon EBS volume types](../../../../services/ebs/latest/userguide/ebs-volume-types.md) in the
_Amazon EBS User Guide_.

Type: String

Valid Values: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ebsblockdevice.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ebsblockdevice.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ebsblockdevice.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DnsServersOptionsModifyStructure

EbsBlockDeviceResponse
