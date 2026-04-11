This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet S3Source

A physical table type for an S3 data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceArn" : String,
  "InputColumns" : [ InputColumn, ... ],
  "UploadSettings" : UploadSettings
}

```

### YAML

```yaml

  DataSourceArn: String
  InputColumns:
    - InputColumn
  UploadSettings:
    UploadSettings

```

## Properties

`DataSourceArn`

The Amazon Resource Name (ARN) for the data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputColumns`

A physical table type for an S3 data source.

###### Note

For files that aren't JSON, only `STRING` data types are supported in input columns.

_Required_: Yes

_Type_: Array of [InputColumn](aws-properties-quicksight-dataset-inputcolumn.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UploadSettings`

Information about the format for the S3 source file or files.

_Required_: No

_Type_: [UploadSettings](aws-properties-quicksight-dataset-uploadsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RowLevelPermissionTagRule

SaaSTable

All content copied from https://docs.aws.amazon.com/.
