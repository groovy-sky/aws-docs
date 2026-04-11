This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergSortOrder

Defines the sort order for data within an Iceberg table. Sorting data can improve query performance by enabling more efficient data skipping.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Fields" : [ IcebergSortField, ... ],
  "OrderId" : Integer
}

```

### YAML

```yaml

  Fields:
    - IcebergSortField
  OrderId: Integer

```

## Properties

`Fields`

The list of sort fields that define how data is sorted within files. Each field specifies a source field, sort direction, and null ordering. This field is required if `writeOrder` is provided.

_Required_: Yes

_Type_: Array of [IcebergSortField](aws-properties-s3tables-table-icebergsortfield.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrderId`

The unique identifier for this sort order. If not specified, defaults to `1`. The order ID is used by Apache Iceberg to track sort order evolution.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergSortField

SchemaField

All content copied from https://docs.aws.amazon.com/.
