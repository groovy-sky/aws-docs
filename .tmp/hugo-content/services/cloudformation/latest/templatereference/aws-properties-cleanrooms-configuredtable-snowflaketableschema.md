This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable SnowflakeTableSchema

The schema of a Snowflake table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "V1" : [ SnowflakeTableSchemaV1, ... ]
}

```

### YAML

```yaml

  V1:
    - SnowflakeTableSchemaV1

```

## Properties

`V1`

The schema of a Snowflake table.

_Required_: Yes

_Type_: Array of [SnowflakeTableSchemaV1](aws-properties-cleanrooms-configuredtable-snowflaketableschemav1.md)

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeTableReference

SnowflakeTableSchemaV1

All content copied from https://docs.aws.amazon.com/.
