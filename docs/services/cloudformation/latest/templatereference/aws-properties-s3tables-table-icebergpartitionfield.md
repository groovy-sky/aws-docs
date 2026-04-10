This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table IcebergPartitionField

Defines a single partition field in an Iceberg partition specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : Integer,
  "Name" : String,
  "SourceId" : Integer,
  "Transform" : String
}

```

### YAML

```yaml

  FieldId: Integer
  Name: String
  SourceId: Integer
  Transform: String

```

## Properties

`FieldId`

An optional unique identifier for this partition field. If not specified, S3 Tables automatically assigns a field ID.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name for this partition field. This name is used in the partitioned file paths.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceId`

The ID of the source schema field to partition by. This must reference a valid field ID from the table schema.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Transform`

The partition transform to apply to the source field. Supported transforms include `identity`, `year`, `month`, `day`, `hour`, `bucket`, and `truncate`. For more information, see the [Apache Iceberg partition transforms documentation](https://iceberg.apache.org/spec).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergMetadata

IcebergPartitionSpec

All content copied from https://docs.aws.amazon.com/.
