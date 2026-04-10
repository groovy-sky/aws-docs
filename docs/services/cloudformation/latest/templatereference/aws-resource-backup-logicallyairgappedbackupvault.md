This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::LogicallyAirGappedBackupVault

Creates a logical container to where backups may be copied.

This request includes a name, the Region, the maximum number of retention days, the
minimum number of retention days, and optionally can include tags and a creator request
ID.

###### Note

Do not include sensitive data, such as passport numbers, in the name of a backup
vault.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::LogicallyAirGappedBackupVault",
  "Properties" : {
      "AccessPolicy" : Json,
      "BackupVaultName" : String,
      "BackupVaultTags" : {Key: Value, ...},
      "EncryptionKeyArn" : String,
      "MaxRetentionDays" : Integer,
      "MinRetentionDays" : Integer,
      "MpaApprovalTeamArn" : String,
      "Notifications" : NotificationObjectType
    }
}

```

### YAML

```yaml

Type: AWS::Backup::LogicallyAirGappedBackupVault
Properties:
  AccessPolicy: Json
  BackupVaultName: String
  BackupVaultTags:
    Key: Value
  EncryptionKeyArn: String
  MaxRetentionDays: Integer
  MinRetentionDays: Integer
  MpaApprovalTeamArn: String
  Notifications:
    NotificationObjectType

```

## Properties

`AccessPolicy`

The backup vault access policy document in JSON format.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackupVaultName`

The name of a logical container where backups are stored. Logically air-gapped
backup vaults are identified by names that are unique to the account used to create
them and the Region where they are created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_]{2,50}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupVaultTags`

The tags to assign to the vault.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

The server-side encryption key that is used to protect your backups; for example,
`arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`.

If this field is left blank, AWS Backup will create an AWS owned key to be used to encrypt the content of the logically air-gapped vault. The ARN of this created key will be available as `Fn::GetAtt` output.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxRetentionDays`

The maximum retention period that the vault retains its recovery points.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinRetentionDays`

This setting specifies the minimum retention period
that the vault retains its recovery points.

The minimum value accepted is 7 days.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MpaApprovalTeamArn`

The Amazon Resource Name (ARN) of the MPA approval team to associate with the backup vault. This cannot be changed after it is set from the CloudFormation template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Notifications`

Returns event notifications for the specified backup vault.

_Required_: No

_Type_: [NotificationObjectType](aws-properties-backup-logicallyairgappedbackupvault-notificationobjecttype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the backup vault.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BackupVaultArn`

The ARN of the backup vault.

`VaultState`

The vault state. The possible values are `CREATING`, `AVAILABLE`,
and `FAILED`.

`VaultType`

The vault type. The possible values are `BACKUP_VAULT` and `LOGICALLY_AIR_GAPPED_BACKUP_VAULT`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

NotificationObjectType

All content copied from https://docs.aws.amazon.com/.
