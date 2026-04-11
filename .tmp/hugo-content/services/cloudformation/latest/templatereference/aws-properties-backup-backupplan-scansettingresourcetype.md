This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::BackupPlan ScanSettingResourceType

The `ScanSettingResourceType` property type specifies Property description not available. for an [AWS::Backup::BackupPlan](aws-resource-backup-backupplan.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MalwareScanner" : String,
  "ResourceTypes" : [ String, ... ],
  "ScannerRoleArn" : String
}

```

### YAML

```yaml

  MalwareScanner: String
  ResourceTypes:
    - String
  ScannerRoleArn: String

```

## Properties

`MalwareScanner`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `GUARDDUTY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTypes`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScannerRoleArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScanActionResourceType

AWS::Backup::BackupSelection

All content copied from https://docs.aws.amazon.com/.
