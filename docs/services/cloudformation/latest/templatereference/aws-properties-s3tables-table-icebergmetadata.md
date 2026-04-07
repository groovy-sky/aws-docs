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
  IcebergSortOrder:
    IcebergSortOrder
  TableProperties:
    Key: Value

```

## Properties

`IcebergPartitionSpec`

Property description not available.

_Required_: No

_Type_: [IcebergPartitionSpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-table-icebergpartitionspec.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IcebergSchema`

The schema for an Iceberg table. Use this property to define table schemas with primitive types only. For schemas that include nested or complex types such as `struct`, `list`, or `map`, use `schemaV2` instead.

_Required_: Yes

_Type_: [IcebergSchema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-table-icebergschema.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IcebergSortOrder`

Property description not available.

_Required_: No

_Type_: [IcebergSortOrder](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3tables-table-icebergsortorder.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableProperties`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Compaction

IcebergPartitionField
