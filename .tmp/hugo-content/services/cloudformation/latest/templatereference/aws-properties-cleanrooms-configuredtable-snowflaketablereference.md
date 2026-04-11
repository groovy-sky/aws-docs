This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable SnowflakeTableReference

A reference to a table within Snowflake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountIdentifier" : String,
  "DatabaseName" : String,
  "SchemaName" : String,
  "SecretArn" : String,
  "TableName" : String,
  "TableSchema" : SnowflakeTableSchema
}

```

### YAML

```yaml

  AccountIdentifier: String
  DatabaseName: String
  SchemaName: String
  SecretArn: String
  TableName: String
  TableSchema:
    SnowflakeTableSchema

```

## Properties

`AccountIdentifier`

The account identifier for the Snowflake table reference.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of the database the Snowflake table belongs to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaName`

The schema name of the Snowflake table reference.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The secret ARN of the Snowflake table reference.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the Snowflake table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableSchema`

The schema of the Snowflake table.

_Required_: Yes

_Type_: [SnowflakeTableSchema](aws-properties-cleanrooms-configuredtable-snowflaketableschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlueTableReference

SnowflakeTableSchema

All content copied from https://docs.aws.amazon.com/.
