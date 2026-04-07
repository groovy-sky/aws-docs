# LaunchTemplateEbsBlockDevice

Describes a block device for an EBS volume.

## Contents

**deleteOnTermination**

Indicates whether the EBS volume is deleted on instance termination.

Type: Boolean

Required: No

**ebsCardIndex**

The index of the EBS card. Some instance types support multiple EBS cards. The default EBS card index is 0.

Type: Integer

Required: No

**encrypted**

Indicates whether the EBS volume is encrypted.

Type: Boolean

Required: No

**iops**

The number of I/O operations per second (IOPS) that the volume supports.

Type: Integer

Required: No

**kmsKeyId**

Identifier (key ID, key alias, key ARN, or alias ARN) of the customer managed KMS key to use for EBS encryption.

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

**volumeInitializationRate**

The Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate) specified for the volume, in MiB/s.
If no volume initialization rate was specified, the value is `null`.

Type: Integer

Required: No

**volumeSize**

The size of the volume, in GiB.

Type: Integer

Required: No

**volumeType**

The volume type.

Type: String

Valid Values: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateebsblockdevice.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateebsblockdevice.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateebsblockdevice.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateCpuOptionsRequest

LaunchTemplateEbsBlockDeviceRequest
