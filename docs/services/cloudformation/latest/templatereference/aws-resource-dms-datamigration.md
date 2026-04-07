This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataMigration

This object provides information about a AWS DMS data migration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::DataMigration",
  "Properties" : {
      "DataMigrationIdentifier" : String,
      "DataMigrationName" : String,
      "DataMigrationSettings" : DataMigrationSettings,
      "DataMigrationType" : String,
      "MigrationProjectIdentifier" : String,
      "ServiceAccessRoleArn" : String,
      "SourceDataSettings" : [ SourceDataSettings, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DMS::DataMigration
Properties:
  DataMigrationIdentifier: String
  DataMigrationName: String
  DataMigrationSettings:
    DataMigrationSettings
  DataMigrationType: String
  MigrationProjectIdentifier: String
  ServiceAccessRoleArn: String
  SourceDataSettings:
    - SourceDataSettings
  Tags:
    - Tag

```

## Properties

`DataMigrationIdentifier`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataMigrationName`

The user-friendly name for the data migration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataMigrationSettings`

Specifies CloudWatch settings and selection rules for the data migration.

_Required_: No

_Type_: [DataMigrationSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-datamigration-datamigrationsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataMigrationType`

Specifies whether the data migration is full-load only, change data capture (CDC) only,
or full-load and CDC.

_Required_: Yes

_Type_: String

_Allowed values_: `full-load | cdc | full-load-and-cdc`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MigrationProjectIdentifier`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccessRoleArn`

The IAM role that the data migration uses to access AWS resources.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceDataSettings`

Specifies information about the data migration's source data provider.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-datamigration-sourcedatasettings.html) of [SourceDataSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-datamigration-sourcedatasettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-datamigration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DataMigrationArn`

The Amazon Resource Name (ARN) that identifies this replication.

`DataMigrationCreateTime`

The UTC time when DMS created the data migration.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DMS::Certificate

DataMigrationSettings
