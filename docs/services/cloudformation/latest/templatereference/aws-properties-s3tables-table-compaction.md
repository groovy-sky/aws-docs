This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table Compaction

Contains details about the compaction settings for an Iceberg table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String,
  "TargetFileSizeMB" : Integer
}

```

### YAML

```yaml

  Status: String
  TargetFileSizeMB: Integer

```

## Properties

`Status`

The status of the maintenance configuration.

_Required_: No

_Type_: String

_Allowed values_: `enabled | disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetFileSizeMB`

The target file size for the table in MB.

_Required_: No

_Type_: Integer

_Minimum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::S3Tables::Table

IcebergMetadata
