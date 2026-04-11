# BackupPlan

Contains an optional backup plan display name and an array of `BackupRule`
objects, each of which specifies a backup rule. Each rule in a backup plan is a separate
scheduled task and can back up a different selection of AWS
resources.

## Contents

**BackupPlanName**

The display name of a backup plan. Must contain only alphanumeric or '-\_.'
special characters.

If this is set in the console, it can contain 1 to 50 characters; if this
is set through CLI or API, it can contain 1 to 200 characters.

Type: String

Required: Yes

**Rules**

An array of `BackupRule` objects, each of which specifies a scheduled task
that is used to back up a selection of resources.

Type: Array of [BackupRule](api-backuprule.md) objects

Required: Yes

**AdvancedBackupSettings**

Contains a list of `BackupOptions` for each resource type.

Type: Array of [AdvancedBackupSetting](api-advancedbackupsetting.md) objects

Required: No

**ScanSettings**

Contains your scanning configuration for the backup plan and includes the Malware scanner, your selected resources, and scanner role.

Type: Array of [ScanSetting](api-scansetting.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/backupplan.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/backupplan.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/backupplan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupJobSummary

BackupPlanInput

All content copied from https://docs.aws.amazon.com/.
