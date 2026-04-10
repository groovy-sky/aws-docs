This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan CopyActionResourceType

Copies backups created by a backup rule to another vault.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationBackupVaultArn" : String,
  "Lifecycle" : LifecycleResourceType
}

```

### YAML

```yaml

  DestinationBackupVaultArn: String
  Lifecycle:
    LifecycleResourceType

```

## Properties

`DestinationBackupVaultArn`

An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for
the copied backup. For example,
`arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lifecycle`

Defines when a protected resource is transitioned to cold storage and when it expires.
AWS Backup transitions and expires backups automatically according to the
lifecycle that you define. If you do not specify a lifecycle, AWS Backup applies
the lifecycle policy of the source backup to the destination backup.

Backups transitioned to cold storage must be stored in cold storage for a minimum of 90
days.

_Required_: No

_Type_: [LifecycleResourceType](aws-properties-backup-backupplan-lifecycleresourcetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupRuleResourceType

IndexActionsResourceType

All content copied from https://docs.aws.amazon.com/.
