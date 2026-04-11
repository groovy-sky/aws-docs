This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet CustomSql

A physical table type built from the results of the custom SQL query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ InputColumn, ... ],
  "DataSourceArn" : String,
  "Name" : String,
  "SqlQuery" : String
}

```

### YAML

```yaml

  Columns:
    - InputColumn
  DataSourceArn: String
  Name: String
  SqlQuery: String

```

## Properties

`Columns`

The column schema from the SQL query result set.

_Required_: Yes

_Type_: Array of [InputColumn](aws-properties-quicksight-dataset-inputcolumn.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceArn`

The Amazon Resource Name (ARN) of the data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A display name for the SQL query result.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlQuery`

The SQL query.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `168000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateColumnsOperation

DataPrepAggregationFunction

All content copied from https://docs.aws.amazon.com/.
