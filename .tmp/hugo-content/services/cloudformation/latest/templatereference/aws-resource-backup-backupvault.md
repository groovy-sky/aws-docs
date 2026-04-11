This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupVault

Creates a logical container where backups are stored. A `CreateBackupVault`
request includes a name, optionally one or more resource tags, an encryption key, and a
request ID.

Do not include sensitive data, such as passport numbers, in the name of a backup
vault.

For a sample CloudFormation template, see the [AWS Backup Developer Guide](../../../aws-backup/latest/devguide/assigning-resources.md#assigning-resources-cfn).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::BackupVault",
  "Properties" : {
      "AccessPolicy" : Json,
      "BackupVaultName" : String,
      "BackupVaultTags" : {Key: Value, ...},
      "EncryptionKeyArn" : String,
      "LockConfiguration" : LockConfigurationType,
      "Notifications" : NotificationObjectType
    }
}

```

### YAML

```yaml

Type: AWS::Backup::BackupVault
Properties:
  AccessPolicy: Json
  BackupVaultName: String
  BackupVaultTags:
    Key: Value
  EncryptionKeyArn: String
  LockConfiguration:
    LockConfigurationType
  Notifications:
    NotificationObjectType

```

## Properties

`AccessPolicy`

A resource-based policy that is used to manage access permissions on the target backup
vault.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackupVaultName`

The name of a logical container where backups are stored. Backup vaults are identified
by names that are unique to the account used to create them and the AWS
Region where they are created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_]{2,50}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupVaultTags`

The tags to assign to the backup vault.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

A server-side encryption key you can specify to encrypt your backups from services
that support full AWS Backup management; for example,
`arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`.
If you specify a key, you must specify its ARN, not its alias. If you do not specify a key,
AWS Backup creates a KMS key for you by default.

To learn which AWS Backup services support full AWS Backup management
and how AWS Backup handles encryption for backups from services that do not yet support
full AWS Backup, see [Encryption for backups in AWS Backup](../../../aws-backup/latest/devguide/encryption.md)

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LockConfiguration`

Configuration for [AWS Backup Vault\
Lock](../../../aws-backup/latest/devguide/vault-lock.md).

_Required_: No

_Type_: [LockConfigurationType](aws-properties-backup-backupvault-lockconfigurationtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Notifications`

The SNS event notifications for the specified backup vault.

_Required_: No

_Type_: [NotificationObjectType](aws-properties-backup-backupvault-notificationobjecttype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `BackupVaultName`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BackupVaultArn`

An Amazon Resource Name (ARN) that uniquely identifies a
backup vault; for example,
`arn:aws:backup:us-east-1:123456789012:backup-vault:aBackupVault`.

`BackupVaultName`

The name of a logical container where backups are
stored. Backup vaults are identified by names that are unique to the account used to create
them and the Region where they are created. They consist of lowercase and uppercase
letters, numbers, and hyphens.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Conditions

LockConfigurationType

All content copied from https://docs.aws.amazon.com/.
