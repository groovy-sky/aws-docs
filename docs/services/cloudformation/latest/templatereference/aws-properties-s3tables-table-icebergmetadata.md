This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergMetadata

Contains details about the metadata for an Iceberg table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IcebergPartitionSpec" : IcebergPartitionSpec,
  "IcebergSchema" : IcebergSchema,
  "IcebergSchemaV2" : IcebergSchemaV2,
  "IcebergSortOrder" : IcebergSortOrder,
  "TableProperties" : {Key: Value, ...}
}

```

### YAML

```yaml

  IcebergPartitionSpec:
    IcebergPartitionSpec
  IcebergSchema:
    IcebergSchema
  IcebergSchemaV2:
    IcebergSchemaV2
  IcebergSortOrder:
    IcebergSortOrder
  TableProperties:
    Key: Value

```

## Properties

`IcebergPartitionSpec`

Property description not available.

_Required_: No

_Type_: [IcebergPartitionSpec](aws-properties-s3tables-table-icebergpartitionspec.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IcebergSchema`

The schema for an Iceberg table. Use this property to define table schemas with primitive types only. For schemas that include nested or complex types such as `struct`, `list`, or `map`, use `schemaV2` instead.

_Required_: No

_Type_: [IcebergSchema](aws-properties-s3tables-table-icebergschema.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IcebergSchemaV2`

The schema for an Iceberg table using the V2 format. Use this property to define table schemas that include nested or complex data types such as `struct`, `list`, or `map`, in addition to primitive types. For schemas with only primitive types, you can use either `schema` or `schemaV2`.

_Required_: No

_Type_: [IcebergSchemaV2](aws-properties-s3tables-table-icebergschemav2.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IcebergSortOrder`

Property description not available.

_Required_: No

_Type_: [IcebergSortOrder](aws-properties-s3tables-table-icebergsortorder.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableProperties`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Compaction

IcebergPartitionField

All content copied from https://docs.aws.amazon.com/.
