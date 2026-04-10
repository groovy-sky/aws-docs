This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable TableReference

A pointer to the dataset that underlies this table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Athena" : AthenaTableReference,
  "Glue" : GlueTableReference,
  "Snowflake" : SnowflakeTableReference
}

```

### YAML

```yaml

  Athena:
    AthenaTableReference
  Glue:
    GlueTableReference
  Snowflake:
    SnowflakeTableReference

```

## Properties

`Athena`

If present, a reference to the Athena table referred to by this table reference.

_Required_: No

_Type_: [AthenaTableReference](aws-properties-cleanrooms-configuredtable-athenatablereference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Glue`

If present, a reference to the AWS Glue table referred to by this table
reference.

_Required_: No

_Type_: [GlueTableReference](aws-properties-cleanrooms-configuredtable-gluetablereference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

If present, a reference to the Snowflake table referred to by this table reference.

_Required_: No

_Type_: [SnowflakeTableReference](aws-properties-cleanrooms-configuredtable-snowflaketablereference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeTableSchemaV1

Tag

All content copied from https://docs.aws.amazon.com/.
