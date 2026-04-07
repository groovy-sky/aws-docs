This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataMigration DataMigrationSettings

Options for configuring a data migration, including whether to enable CloudWatch logs,
and the selection rules to use to include or exclude database objects from the
migration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudwatchLogsEnabled" : Boolean,
  "NumberOfJobs" : Integer,
  "SelectionRules" : String
}

```

### YAML

```yaml

  CloudwatchLogsEnabled: Boolean
  NumberOfJobs: Integer
  SelectionRules: String

```

## Properties

`CloudwatchLogsEnabled`

Whether to enable CloudWatch logging for the data migration.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfJobs`

The number of parallel jobs that trigger parallel threads to unload the tables from the
source, and then load them to the target.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectionRules`

A JSON-formatted string that defines what objects to include and exclude from the
migration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DMS::DataMigration

SourceDataSettings
