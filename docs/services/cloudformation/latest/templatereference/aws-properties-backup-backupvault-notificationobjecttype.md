This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupVault NotificationObjectType

Specifies an object containing SNS event notification properties for the target backup
vault.

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
vault. For valid events, see [BackupVaultEvents](../../../aws-backup/latest/devguide/api-putbackupvaultnotifications.md#API_PutBackupVaultNotifications_RequestSyntax) in the _AWS Backup API_
_Guide_.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SNSTopicArn`

An ARN that uniquely identifies an Amazon Simple Notification Service (Amazon SNS)
topic; for example, `arn:aws:sns:us-west-2:111122223333:MyTopic`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LockConfigurationType

AWS::Backup::Framework

All content copied from https://docs.aws.amazon.com/.
