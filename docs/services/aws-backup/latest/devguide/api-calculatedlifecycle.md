# CalculatedLifecycle

Contains `DeleteAt` and `MoveToColdStorageAt` timestamps, which
are used to specify a lifecycle for a recovery point.

The lifecycle defines when a protected resource is transitioned to cold storage and when
it expires. AWS Backup transitions and expires backups automatically according to
the lifecycle that you define.

Backups transitioned to cold storage must be stored in cold storage for a minimum of 90
days. Therefore, the “retention” setting must be 90 days greater than the “transition to
cold after days” setting. The “transition to cold after days” setting cannot be changed
after a backup has been transitioned to cold.

Resource types that can transition to cold storage are listed in the [Feature availability \
by resource](backup-feature-availability.md#features-by-resource) table. AWS Backup ignores this expression for other resource types.

## Contents

**DeleteAt**

A timestamp that specifies when to delete a recovery point.

Type: Timestamp

Required: No

**MoveToColdStorageAt**

A timestamp that specifies when to transition a recovery point to cold storage.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/calculatedlifecycle.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/calculatedlifecycle.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/calculatedlifecycle.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupVaultListMember

Condition

All content copied from https://docs.aws.amazon.com/.
