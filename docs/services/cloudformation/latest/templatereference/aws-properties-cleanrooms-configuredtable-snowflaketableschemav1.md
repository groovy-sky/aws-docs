This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable SnowflakeTableSchemaV1

The Snowflake table schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "ColumnType" : String
}

```

### YAML

```yaml

  ColumnName: String
  ColumnType: String

```

## Properties

`ColumnName`

The column name.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnType`

The column's data type. Supported data types: `ARRAY`, `BIGINT`,
`BOOLEAN`, `CHAR`, `DATE`,
`DECIMAL`, `DOUBLE`, `DOUBLE PRECISION`,
`FLOAT`, `FLOAT4`, `INT`, `INTEGER`,
`MAP`, `NUMERIC`, `NUMBER`, `REAL`,
`SMALLINT`, `STRING`, `TIMESTAMP`,
`TIMESTAMP_LTZ`, `TIMESTAMP_NTZ`, `DATETIME`,
`TINYINT`, `VARCHAR`, `TEXT`, `CHARACTER`.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeTableSchema

TableReference

All content copied from https://docs.aws.amazon.com/.
