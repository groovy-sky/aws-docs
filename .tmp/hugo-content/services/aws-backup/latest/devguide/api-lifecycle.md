# Lifecycle

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

## Contents

**DeleteAfterDays**

The number of days after creation that a recovery point is deleted. This value must be
at least 90 days after the number of days specified in `MoveToColdStorageAfterDays`.

Type: Long

Required: No

**DeleteAfterEvent**

The event after which a recovery point is deleted. A recovery point with both
`DeleteAfterDays` and `DeleteAfterEvent` will delete after
whichever condition is satisfied first. Not valid as an input.

Type: String

Valid Values: `DELETE_AFTER_COPY`

Required: No

**MoveToColdStorageAfterDays**

The number of days after creation that a recovery point is moved to cold
storage.

Type: Long

Required: No

**OptInToArchiveForSupportedResources**

If the value is true, your backup plan transitions supported resources to
archive (cold) storage tier in accordance with your lifecycle settings.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/lifecycle.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/lifecycle.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/lifecycle.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LegalHold

ProtectedResource

All content copied from https://docs.aws.amazon.com/.
