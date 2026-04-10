This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::TableOptimizer IcebergConfiguration

IcebergConfiguration is a property type within the
`AWS::Glue::TableOptimizer` resource in AWS CloudFormation. This configuration is
used when setting up table optimization for Iceberg tables in AWS Glue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Location" : String,
  "OrphanFileRetentionPeriodInDays" : Integer
}

```

### YAML

```yaml

  Location: String
  OrphanFileRetentionPeriodInDays: Integer

```

## Properties

`Location`

Specifies a directory in which to look for orphan files (defaults to the table's location). You may choose a sub-directory rather than the top-level table location.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrphanFileRetentionPeriodInDays`

The specific number of days you want to keep the orphan files.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::TableOptimizer

IcebergRetentionConfiguration

All content copied from https://docs.aws.amazon.com/.
