# AdvancedBackupSetting

The backup options for each resource type.

## Contents

**BackupOptions**

Specifies the backup option for a selected resource. This option is available for
Windows VSS backup jobs and S3 backups.

Valid values:

Set to `"WindowsVSS":"enabled"` to enable the `WindowsVSS` backup
option and create a Windows VSS backup.

Set to `"WindowsVSS":"disabled"` to create a regular backup. The
`WindowsVSS` option is not enabled by default.

For S3 backups, set to `"BackupACLs":"disabled"` to exclude ACLs from the backup,
or `"BackupObjectTags":"disabled"` to exclude object tags from the backup.
By default, both ACLs and object tags are included in S3 backups.

If you specify an invalid option, you get an `InvalidParameterValueException`
exception.

For more information about Windows VSS backups, see [Creating a VSS-Enabled Windows\
Backup](windows-backups.md).

Type: String to string map

Key Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

Value Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

Required: No

**ResourceType**

Specifies an object containing resource type and backup options. The only supported
resource type is Amazon EC2 instances with Windows Volume Shadow Copy Service
(VSS). For a CloudFormation example, see the [sample\
CloudFormation template to enable Windows VSS](integrate-cloudformation-with-aws-backup.md) in the _AWS Backup User Guide_.

Valid values: `EC2`.

Type: String

Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/advancedbackupsetting.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/advancedbackupsetting.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/advancedbackupsetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Backup

AggregatedScanResult

All content copied from https://docs.aws.amazon.com/.
