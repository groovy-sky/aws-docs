# RecoveryPointCreator

Contains information about the backup plan and rule that AWS Backup used to
initiate the recovery point backup.

## Contents

**BackupPlanArn**

An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for example,
`arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50`.

Type: String

Required: No

**BackupPlanId**

Uniquely identifies a backup plan.

Type: String

Required: No

**BackupPlanName**

The name of the backup plan that created this recovery point. This provides human-readable context about which backup plan was responsible for the backup job.

Type: String

Required: No

**BackupPlanVersion**

Version IDs are unique, randomly generated, Unicode, UTF-8 encoded strings that are at
most 1,024 bytes long. They cannot be edited.

Type: String

Required: No

**BackupRuleCron**

The cron expression that defines the schedule for the backup rule. This shows the frequency and timing of when backups are automatically triggered.

Type: String

Required: No

**BackupRuleId**

Uniquely identifies a rule used to schedule the backup of a selection of
resources.

Type: String

Required: No

**BackupRuleName**

The name of the backup rule within the backup plan that created this recovery point. This helps identify which specific rule triggered the backup job.

Type: String

Required: No

**BackupRuleTimezone**

The timezone used for the backup rule schedule. This provides context for when backups are scheduled to run in the specified timezone.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/recoverypointcreator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/recoverypointcreator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/recoverypointcreator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecoveryPointByResource

RecoveryPointMember

All content copied from https://docs.aws.amazon.com/.
