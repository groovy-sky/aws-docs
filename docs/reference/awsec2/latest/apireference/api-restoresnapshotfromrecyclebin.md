# RestoreSnapshotFromRecycleBin

Restores a snapshot from the Recycle Bin. For more information, see [Restore \
snapshots from the Recycle Bin](../../../../services/ebs/latest/userguide/recycle-bin-working-with-snaps.md#recycle-bin-restore-snaps) in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SnapshotId**

The ID of the snapshot to restore.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**description**

The description for the snapshot.

Type: String

**encrypted**

Indicates whether the snapshot is encrypted.

Type: Boolean

**outpostArn**

The ARN of the Outpost on which the snapshot is stored. For more information, see [Amazon EBS local snapshots on Outposts](../../../../services/ebs/latest/userguide/snapshots-outposts.md) in the
_Amazon EBS User Guide_.

Type: String

**ownerId**

The ID of the AWS account that owns the EBS snapshot.

Type: String

**progress**

The progress of the snapshot, as a percentage.

Type: String

**requestId**

The ID of the request.

Type: String

**snapshotId**

The ID of the snapshot.

Type: String

**sseType**

Reserved for future use.

Type: String

Valid Values: `sse-ebs | sse-kms | none`

**startTime**

The time stamp when the snapshot was initiated.

Type: Timestamp

**status**

The state of the snapshot.

Type: String

Valid Values: `pending | completed | error | recoverable | recovering`

**volumeId**

The ID of the volume that was used to create the snapshot.

Type: String

**volumeSize**

The size of the volume, in GiB.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/restoresnapshotfromrecyclebin.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RestoreManagedPrefixListVersion

RestoreSnapshotTier
