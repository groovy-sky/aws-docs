# EnableFastSnapshotRestores

Enables fast snapshot restores for the specified snapshots in the specified Availability Zones.

You get the full benefit of fast snapshot restores after they enter the `enabled` state.

To get the current state of fast snapshot restores, use [DescribeFastSnapshotRestores](api-describefastsnapshotrestores.md).
To disable fast snapshot restores, use [DisableFastSnapshotRestores](api-disablefastsnapshotrestores.md).

For more information, see [Amazon EBS fast snapshot\
restore](../../../../services/ebs/latest/userguide/ebs-fast-snapshot-restore.md) in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AvailabilityZone.N**

One or more Availability Zones. For example, `us-east-2a`.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified in the request, but not both.

Type: Array of strings

Required: No

**AvailabilityZoneId.N**

One or more Availability Zone IDs. For example, `use2-az1`.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified in the request, but not both.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SourceSnapshotId.N**

The IDs of one or more snapshots. For example, `snap-1234567890abcdef0`. You can specify
a snapshot that was shared with you from another AWS account.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**successful**

Information about the snapshots for which fast snapshot restores were successfully enabled.

Type: Array of [EnableFastSnapshotRestoreSuccessItem](api-enablefastsnapshotrestoresuccessitem.md) objects

**unsuccessful**

Information about the snapshots for which fast snapshot restores could not be enabled.

Type: Array of [EnableFastSnapshotRestoreErrorItem](api-enablefastsnapshotrestoreerroritem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/enablefastsnapshotrestores.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enablefastsnapshotrestores.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableFastLaunch

EnableImage
