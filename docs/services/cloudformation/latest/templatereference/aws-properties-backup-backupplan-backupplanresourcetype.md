This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan BackupPlanResourceType

Specifies an object containing properties used to create a backup plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvancedBackupSettings" : [ AdvancedBackupSettingResourceType, ... ],
  "BackupPlanName" : String,
  "BackupPlanRule" : [ BackupRuleResourceType, ... ],
  "ScanSettings" : [ ScanSettingResourceType, ... ]
}

```

### YAML

```yaml

  AdvancedBackupSettings:
    - AdvancedBackupSettingResourceType
  BackupPlanName: String
  BackupPlanRule:
    - BackupRuleResourceType
  ScanSettings:
    - ScanSettingResourceType

```

## Properties

`AdvancedBackupSettings`

A list of backup options for each resource type.

_Required_: No

_Type_: Array of [AdvancedBackupSettingResourceType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-backupplan-advancedbackupsettingresourcetype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackupPlanName`

The display name of a backup plan.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackupPlanRule`

An array of `BackupRule` objects, each of which specifies a scheduled task
that is used to back up a selection of resources.

_Required_: Yes

_Type_: Array of [BackupRuleResourceType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-backupplan-backupruleresourcetype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanSettings`

Property description not available.

_Required_: No

_Type_: Array of [ScanSettingResourceType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-backupplan-scansettingresourcetype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AdvancedBackupSettingResourceType

BackupRuleResourceType
