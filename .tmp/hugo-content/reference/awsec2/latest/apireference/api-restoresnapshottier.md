# RestoreSnapshotTier

Restores an archived Amazon EBS snapshot for use temporarily or permanently, or modifies the restore
period or restore type for a snapshot that was previously temporarily restored.

For more information see [Restore an archived snapshot](../../../../services/ebs/latest/userguide/working-with-snapshot-archiving.md#restore-archived-snapshot) and [modify the restore period or restore type for a temporarily restored snapshot](../../../../services/ebs/latest/userguide/working-with-snapshot-archiving.md#modify-temp-restore-period) in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PermanentRestore**

Indicates whether to permanently restore an archived snapshot. To permanently restore
an archived snapshot, specify `true` and omit the
**RestoreSnapshotTierRequest$TemporaryRestoreDays** parameter.

Type: Boolean

Required: No

**SnapshotId**

The ID of the snapshot to restore.

Type: String

Required: Yes

**TemporaryRestoreDays**

Specifies the number of days for which to temporarily restore an archived snapshot.
Required for temporary restores only. The snapshot will be automatically re-archived
after this period.

To temporarily restore an archived snapshot, specify the number of days and omit
the **PermanentRestore** parameter or set it to
`false`.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**isPermanentRestore**

Indicates whether the snapshot is permanently restored. `true` indicates a permanent
restore. `false` indicates a temporary restore.

Type: Boolean

**requestId**

The ID of the request.

Type: String

**restoreDuration**

For temporary restores only. The number of days for which the archived snapshot
is temporarily restored.

Type: Integer

**restoreStartTime**

The date and time when the snapshot restore process started.

Type: Timestamp

**snapshotId**

The ID of the snapshot.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/restoresnapshottier.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/restoresnapshottier.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RestoreSnapshotFromRecycleBin

RestoreVolumeFromRecycleBin
