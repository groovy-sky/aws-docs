# EbsBlockDeviceResponse

Describes a block device for an EBS volume.

## Contents

**deleteOnTermination**

Indicates whether the volume is deleted on instance termination.

Type: Boolean

Required: No

**encrypted**

Indicates whether the volume is encrypted.

Type: Boolean

Required: No

**iops**

The number of I/O operations per second (IOPS). For `gp3`, `io1`, and `io2` volumes,
this represents the number of IOPS that are provisioned for the volume. For `gp2`
volumes, this represents the baseline performance of the volume and the rate at which
the volume accumulates I/O credits for bursting.

Type: Integer

Required: No

**kmsKeyId**

Identifier (key ID, key alias, key ARN, or alias ARN) of the customer managed KMS key
to use for EBS encryption.

Type: String

Required: No

**snapshotId**

The ID of the snapshot.

Type: String

Required: No

**throughput**

The throughput that the volume supports, in MiB/s.

Type: Integer

Required: No

**volumeSize**

The size of the volume, in GiBs.

Type: Integer

Required: No

**volumeType**

The volume type. For more information, see [Amazon EBS volume types](../../../../services/ebs/latest/userguide/ebs-volume-types.md) in the
_Amazon EBS User Guide_.

Type: String

Valid Values: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ebsblockdeviceresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ebsblockdeviceresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ebsblockdeviceresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EbsBlockDevice

EbsCardInfo
