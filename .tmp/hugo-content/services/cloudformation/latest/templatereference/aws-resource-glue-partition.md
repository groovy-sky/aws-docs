This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Partition

The `AWS::Glue::Partition` resource creates an AWS Glue partition, which
represents a slice of table data. For more information, see [CreatePartition Action](../../../glue/latest/dg/aws-glue-api-catalog-partitions.md#aws-glue-api-catalog-partitions-CreatePartition) and [Partition Structure](../../../glue/latest/dg/aws-glue-api-catalog-partitions.md#aws-glue-api-catalog-partitions-Partition) in the _AWS Glue Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Partition",
  "Properties" : {
      "CatalogId" : String,
      "DatabaseName" : String,
      "PartitionInput" : PartitionInput,
      "TableName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Partition
Properties:
  CatalogId: String
  DatabaseName: String
  PartitionInput:
    PartitionInput
  TableName: String

```

## Properties

`CatalogId`

The AWS account ID of the catalog in which the partion is to be created.

###### Note

To specify the account ID, you can use the `Ref` intrinsic function
with the `AWS::AccountId` pseudo parameter. For example: `!Ref
                    AWS::AccountId`

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the catalog database in which to create the partition.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PartitionInput`

The structure used to create and update a partition.

_Required_: Yes

_Type_: [PartitionInput](aws-properties-glue-partition-partitioninput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the metadata table in which the partition is to be created.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the partition name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## See also

- [CreatePartition Action](../../../glue/latest/dg/aws-glue-api-catalog-partitions.md#aws-glue-api-catalog-partitions-CreatePartition) in the _AWS Glue Developer_
_Guide_

- [Partition Structure](../../../glue/latest/dg/aws-glue-api-catalog-partitions.md#aws-glue-api-catalog-partitions-Partition) in the _AWS Glue Developer_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformParameters

Column

All content copied from https://docs.aws.amazon.com/.
