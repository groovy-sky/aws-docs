This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet RelationalTable

A physical table type for relational data sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Catalog" : String,
  "DataSourceArn" : String,
  "InputColumns" : [ InputColumn, ... ],
  "Name" : String,
  "Schema" : String
}

```

### YAML

```yaml

  Catalog: String
  DataSourceArn: String
  InputColumns:
    - InputColumn
  Name: String
  Schema: String

```

## Properties

`Catalog`

The catalog associated with a table.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceArn`

The Amazon Resource Name (ARN) for the data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputColumns`

The column schema of the table.

_Required_: Yes

_Type_: Array of [InputColumn](aws-properties-quicksight-dataset-inputcolumn.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the relational table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

The schema name. This name applies to certain relational database engines.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RefreshFailureEmailAlert

RenameColumnOperation

All content copied from https://docs.aws.amazon.com/.
