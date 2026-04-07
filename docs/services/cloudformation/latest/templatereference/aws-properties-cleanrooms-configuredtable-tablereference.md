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

_Type_: [AthenaTableReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-configuredtable-athenatablereference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Glue`

If present, a reference to the AWS Glue table referred to by this table
reference.

_Required_: No

_Type_: [GlueTableReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-configuredtable-gluetablereference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

If present, a reference to the Snowflake table referred to by this table reference.

_Required_: No

_Type_: [SnowflakeTableReference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-configuredtable-snowflaketablereference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SnowflakeTableSchemaV1

Tag
