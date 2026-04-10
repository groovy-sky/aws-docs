This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergSchema

Contains details about the schema for an Iceberg table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SchemaFieldList" : [ SchemaField, ... ]
}

```

### YAML

```yaml

  SchemaFieldList:
    - SchemaField

```

## Properties

`SchemaFieldList`

The schema fields for the table

_Required_: Yes

_Type_: Array of [SchemaField](aws-properties-s3tables-table-schemafield.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergPartitionSpec

IcebergSchemaV2

All content copied from https://docs.aws.amazon.com/.
