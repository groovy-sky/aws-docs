# CopyAction

The details of the copy operation.

## Contents

**DestinationBackupVaultArn**

An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for
the copied backup. For example,
`arn:aws:backup:us-east-1:123456789012:backup-vault:aBackupVault`.

Type: String

Required: Yes

**Lifecycle**

Specifies the time period, in days, before a recovery point transitions to cold storage
or is deleted.

Backups transitioned to cold storage must be stored in cold storage for a minimum of 90
days. Therefore, on the console, the retention setting must be 90 days greater than the
transition to cold after days setting. The transition to cold after days setting can't
be changed after a backup has been transitioned to cold.

Resource types that can transition to cold storage are listed in the [Feature \
availability by resource](backup-feature-availability.md#features-by-resource) table. AWS Backup ignores this expression for
other resource types.

To remove the existing lifecycle and retention periods and keep your recovery points indefinitely,
specify -1 for `MoveToColdStorageAfterDays` and `DeleteAfterDays`.

Type: [Lifecycle](api-lifecycle.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/copyaction.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/copyaction.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/copyaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlScope

CopyJob

All content copied from https://docs.aws.amazon.com/.
