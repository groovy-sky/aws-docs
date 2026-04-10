This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::FileSystem BackupPolicy

The backup policy turns automatic backups for the file system on or off.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String
}

```

### YAML

```yaml

  Status: String

```

## Properties

`Status`

Set the backup policy status for the file system.

- **`ENABLED`** \- Turns automatic backups on for the file system.

- **`DISABLED`** \- Turns automatic backups off for the file system.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EFS::FileSystem

ElasticFileSystemTag

All content copied from https://docs.aws.amazon.com/.
