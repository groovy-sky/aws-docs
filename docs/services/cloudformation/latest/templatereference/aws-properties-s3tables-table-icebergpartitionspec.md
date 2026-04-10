This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergPartitionSpec

Defines how data in an Iceberg table is partitioned. Partitioning helps optimize query performance by organizing data into separate files based on field values. Each partition field specifies a transform to apply to a source field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Fields" : [ IcebergPartitionField, ... ],
  "SpecId" : Integer
}

```

### YAML

```yaml

  Fields:
    - IcebergPartitionField
  SpecId: Integer

```

## Properties

`Fields`

The list of partition fields that define how the table data is partitioned. Each field specifies a source field and a transform to apply. This field is required if `partitionSpec` is provided.

_Required_: Yes

_Type_: Array of [IcebergPartitionField](aws-properties-s3tables-table-icebergpartitionfield.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpecId`

The unique identifier for this partition specification. If not specified, defaults to `0`.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergPartitionField

IcebergSchema

All content copied from https://docs.aws.amazon.com/.
