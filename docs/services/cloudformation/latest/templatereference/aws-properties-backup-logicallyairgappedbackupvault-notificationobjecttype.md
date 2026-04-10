This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::LogicallyAirGappedBackupVault NotificationObjectType

The `NotificationObjectType` property type specifies Property description not available. for an [AWS::Backup::LogicallyAirGappedBackupVault](aws-resource-backup-logicallyairgappedbackupvault.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackupVaultEvents" : [ String, ... ],
  "SNSTopicArn" : String
}

```

### YAML

```yaml

  BackupVaultEvents:
    - String
  SNSTopicArn: String

```

## Properties

`BackupVaultEvents`

An array of events that indicate the status of jobs to back up resources to the backup
vault.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SNSTopicArn`

The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events; for
example, `arn:aws:sns:us-west-2:111122223333:MyVaultTopic`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Backup::LogicallyAirGappedBackupVault

AWS::Backup::ReportPlan

All content copied from https://docs.aws.amazon.com/.
