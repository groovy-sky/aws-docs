# CopyVolumes

Creates a crash-consistent, point-in-time copy of an existing Amazon EBS volume within the same
Availability Zone. The volume copy can be attached to an Amazon EC2 instance once it reaches the
`available` state. For more information, see [Copy an Amazon EBS volume](../../../../services/ebs/latest/userguide/ebs-copying-volume.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. For more information, see [Ensure Idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Iops**

The number of I/O operations per second (IOPS) to provision for the volume copy.
Required for `io1` and `io2` volumes. Optional for `gp3`
volumes. Omit for all other volume types. Full provisioned IOPS performance can be
achieved only once the volume copy is fully initialized.

Valid ranges:

- gp3: `3,000 `( _default_) ` - 80,000` IOPS

- io1: `100 - 64,000` IOPS

- io2: `100 - 256,000` IOPS

###### Note

[Instances built on the Nitro System](../../../../services/ec2/latest/instancetypes/ec2-nitro-instances.md) can support up to 256,000 IOPS. Other instances can support up to 32,000
IOPS.

Type: Integer

Required: No

**MultiAttachEnabled**

Indicates whether to enable Amazon EBS Multi-Attach for the volume copy. If you enable
Multi-Attach, you can attach the volume to up to 16 Nitro instances in the same
Availability Zone simultaneously. Supported with `io1` and `io2` volumes only. For more
information, see [Amazon EBS Multi-Attach](../../../../services/ebs/latest/userguide/ebs-volumes-multi.md).

Type: Boolean

Required: No

**Size**

The size of the volume copy, in GiBs. The size must be equal to or greater than the
size of the source volume. If not specified, the size defaults to the size of the source
volume.

Maximum supported sizes:

- gp2: `16,384` GiB

- gp3: `65,536` GiB

- io1: `16,384` GiB

- io2: `65,536` GiB

- st1 and sc1: `16,384` GiB

- standard: `1024` GiB

Type: Integer

Required: No

**SourceVolumeId**

The ID of the source EBS volume to copy.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the volume copy during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**Throughput**

The throughput to provision for the volume copy, in MiB/s. Supported for `gp3`
volumes only. Omit for all other volume types. Full provisioned throughput performance can be
achieved only once the volume copy is fully initialized.

Valid Range: `125 - 2000` MiB/s

Type: Integer

Required: No

**VolumeType**

The volume type for the volume copy. If not specified, the volume type defaults to
`gp2`.

Type: String

Valid Values: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**volumeSet**

Information about the volume copy.

Type: Array of [Volume](api-volume.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CopyVolumes)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CopyVolumes)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/copyvolumes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/copyvolumes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/copyvolumes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/copyvolumes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/copyvolumes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/copyvolumes.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CopyVolumes)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/copyvolumes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CopySnapshot

CreateCapacityManagerDataExport
