This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan BackupRuleResourceType

Specifies an object containing properties used to schedule a task to back up a selection
of resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CompletionWindowMinutes" : Number,
  "CopyActions" : [ CopyActionResourceType, ... ],
  "EnableContinuousBackup" : Boolean,
  "IndexActions" : [ IndexActionsResourceType, ... ],
  "Lifecycle" : LifecycleResourceType,
  "RecoveryPointTags" : {Key: Value, ...},
  "RuleName" : String,
  "ScanActions" : [ ScanActionResourceType, ... ],
  "ScheduleExpression" : String,
  "ScheduleExpressionTimezone" : String,
  "StartWindowMinutes" : Number,
  "TargetBackupVault" : String,
  "TargetLogicallyAirGappedBackupVaultArn" : String
}

```

### YAML

```yaml

  CompletionWindowMinutes: Number
  CopyActions:
    - CopyActionResourceType
  EnableContinuousBackup: Boolean
  IndexActions:
    - IndexActionsResourceType
  Lifecycle:
    LifecycleResourceType
  RecoveryPointTags:
    Key: Value
  RuleName: String
  ScanActions:
    - ScanActionResourceType
  ScheduleExpression: String
  ScheduleExpressionTimezone: String
  StartWindowMinutes: Number
  TargetBackupVault: String
  TargetLogicallyAirGappedBackupVaultArn: String

```

## Properties

`CompletionWindowMinutes`

A value in minutes after a backup job is successfully started before it must be
completed or it is canceled by AWS Backup.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyActions`

An array of CopyAction objects, which contains the details of the copy operation.

_Required_: No

_Type_: Array of [CopyActionResourceType](aws-properties-backup-backupplan-copyactionresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableContinuousBackup`

Enables continuous backup and point-in-time restores (PITR).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexActions`

There can up to one IndexAction in each BackupRule, as each backup
can have 0 or 1 backup index associated with it.

Within the array is ResourceTypes. Only 1 resource type will
be accepted for each BackupRule. Valid values:

- `EBS` for Amazon Elastic Block Store

- `S3` for Amazon Simple Storage Service (Amazon S3)

_Required_: No

_Type_: Array of [IndexActionsResourceType](aws-properties-backup-backupplan-indexactionsresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lifecycle`

The lifecycle defines when a protected resource is transitioned to cold storage and when
it expires. AWS Backup transitions and expires backups automatically according to
the lifecycle that you define.

_Required_: No

_Type_: [LifecycleResourceType](aws-properties-backup-backupplan-lifecycleresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryPointTags`

The tags to assign to the resources.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

A display name for a backup rule.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanActions`

Property description not available.

_Required_: No

_Type_: Array of [ScanActionResourceType](aws-properties-backup-backupplan-scanactionresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

A CRON expression specifying when AWS Backup initiates a backup job.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpressionTimezone`

This is the timezone in which the schedule expression is set. By default,
ScheduleExpressions are in UTC. You can modify this to a specified timezone.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartWindowMinutes`

An optional value that specifies a period of time in minutes after a backup is scheduled
before a job is canceled if it doesn't start successfully.

If this value is included, it must be at least 60 minutes to avoid errors.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetBackupVault`

The name of a logical container where backups are stored. Backup vaults are identified
by names that are unique to the account used to create them and the AWS Region where they are created. They consist of letters, numbers, and
hyphens.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetLogicallyAirGappedBackupVaultArn`

The ARN of a logically air-gapped vault. ARN must be in the same account and Region. If provided, supported fully managed resources back up directly to logically air-gapped vault, while other supported resources create a temporary (billable) snapshot in backup vault, then copy it to logically air-gapped vault. Unsupported resources only back up to the specified backup vault.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupPlanResourceType

CopyActionResourceType

All content copied from https://docs.aws.amazon.com/.
