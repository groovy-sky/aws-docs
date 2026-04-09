# Choosing your controls

The following table lists the AWS Backup Audit Manager controls, their customizable
parameters, and their AWS Config recording resource types.

Available controlsControl nameControl descriptionCustomizable parametersAWS Config recording resource typeBackup resources are included in at least one backup planEvaluates if resources are included in at least one backup plan.None`AWS Backup: backup selection`Backup plan has minimum frequency and minimum retentionEvaluates if backup frequency is at least \[1 day\] and retention period is at
least \[35 days\].Backup frequency; retention period`AWS Backup: backup plans`Vaults prevent manual deletion of recovery pointsEvaluates if backup vaults do not allow manual deletion of recovery points
except by certain AWS Identity and Access Management (IAM) roles. By default, there are no IAM role
exceptions. There are also no IAM role exceptions when you deploy this control
with the AWS Backup framework.Up to 5 IAM roles that allow manual deletion of recovery points`AWS Backup: backup vaults`Recovery points are encryptedEvaluates if the recovery points are encrypted.None`AWS Backup: recovery points`Minimum retention established for recovery pointEvaluates if the recovery point retention period is at least \[35 days\].Recovery point retention period`AWS Backup: recovery points`Cross-Region backup copy is scheduledEvaluates if a resource is configured to create copies of its backups to another
AWS Region.AWS Region`AWS Backup: backup selection`Cross-account backup copy is scheduledEvaluates if a resource has a cross-account backup copy configured.AWS account ID`AWS Backup: backup selection`Resources are in a backup plan with an AWS Backup Vault LockEvaluates if a resource has a backup plan configured to store backups
in a locked backup vault.Min Retention Days; Max Retention Days`AWS Backup: backup selection`Last recovery point was createdEvaluates if a recovery point was created within specified time frame.Value in hours \[ `1` to `744`\] or days \[ `1`
to `31`\].`AWS Backup recovery points`Restore time for resources meet targetEvaluates if restore testing job completed within target restore timeValue in minutesNoneResources are inside a logically air-gapped vaultEvaluates if resources have at least one recovery point copied to
a logically air-gapped vault within the specified value and timeframe.Value in minutes, hours, or days`AWS Backup: recovery points`

For detailed information about these controls, see [Controls and remediation](controls-and-remediation.md).

For a list of AWS Backup-supported resources that don't support all controls, see the AWS Backup
Audit Manager section of the [Feature availability by resource](backup-feature-availability.md#features-by-resource) table.

###### Note

If you don't want to use any of the preceding controls, you can still use AWS Backup Audit
Manager to create daily reports of your backup, copy, and restore jobs. See [Working with audit reports](working-with-audit-reports.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Working with audit frameworks

Turning on resource tracking

All content copied from https://docs.aws.amazon.com/.
