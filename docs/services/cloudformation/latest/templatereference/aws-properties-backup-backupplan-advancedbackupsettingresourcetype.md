This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan AdvancedBackupSettingResourceType

Specifies an object containing resource type and backup options. This is only supported
for Windows VSS backups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackupOptions" : Json,
  "ResourceType" : String
}

```

### YAML

```yaml

  BackupOptions: Json
  ResourceType: String

```

## Properties

`BackupOptions`

The backup option for the resource. Each option is a key-value pair. This option is only
available for Windows VSS backup jobs.

Valid values:

Set to `"WindowsVSS":"enabled"` to enable the `WindowsVSS` backup
option and create a Windows VSS backup.

Set to `"WindowsVSS":"disabled"` to create a regular backup. The
`WindowsVSS` option is not enabled by default.

If you specify an invalid option, you get an `InvalidParameterValueException`
exception.

For more information about Windows VSS backups, see [Creating a VSS-Enabled Windows\
Backup](../../../aws-backup/latest/devguide/windows-backups.md).

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The name of a resource type. The only supported resource type is EC2.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Backup::BackupPlan

BackupPlanResourceType

All content copied from https://docs.aws.amazon.com/.
