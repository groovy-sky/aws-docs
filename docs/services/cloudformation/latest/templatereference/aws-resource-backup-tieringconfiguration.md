This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::TieringConfiguration

This contains metadata about a tiering configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::TieringConfiguration",
  "Properties" : {
      "BackupVaultName" : String,
      "ResourceSelection" : [ ResourceSelection, ... ],
      "TieringConfigurationName" : String,
      "TieringConfigurationTags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Backup::TieringConfiguration
Properties:
  BackupVaultName: String
  ResourceSelection:
    - ResourceSelection
  TieringConfigurationName: String
  TieringConfigurationTags:
    Key: Value

```

## Properties

`BackupVaultName`

The name of the backup vault where the tiering configuration applies.
Use `*` to apply to all backup vaults.

_Required_: Yes

_Type_: String

_Pattern_: `^(\*|[a-zA-Z0-9\-\_]{2,50})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceSelection`

An array of resource selection objects that specify which resources
are included in the tiering configuration and their tiering settings.

_Required_: Yes

_Type_: [Array](aws-properties-backup-tieringconfiguration-resourceselection.md) of [ResourceSelection](aws-properties-backup-tieringconfiguration-resourceselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TieringConfigurationName`

The unique name of the tiering configuration. This cannot be changed
after creation, and it must consist of only alphanumeric characters and underscores.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]{1,200}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TieringConfigurationTags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

The date and time a tiering configuration was created, in Unix format
and Coordinated Universal Time (UTC). The value of `CreationTime`
is accurate to milliseconds. For example, the value 1516925490.087 represents
Friday, January 26, 2018 12:11:30.087AM.

`LastUpdatedTime`

The date and time a tiering configuration was updated, in Unix format
and Coordinated Universal Time (UTC). The value of `LastUpdatedTime`
is accurate to milliseconds. For example, the value 1516925490.087 represents
Friday, January 26, 2018 12:11:30.087AM.

`TieringConfigurationArn`

An Amazon Resource Name (ARN) that uniquely identifies the
tiering configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProtectedResourceConditions

ResourceSelection

All content copied from https://docs.aws.amazon.com/.
