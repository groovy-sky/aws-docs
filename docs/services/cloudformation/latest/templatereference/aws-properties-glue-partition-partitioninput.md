This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Partition PartitionInput

The structure used to create and update a partition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : Json,
  "StorageDescriptor" : StorageDescriptor,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Parameters: Json
  StorageDescriptor:
    StorageDescriptor
  Values:
    - String

```

## Properties

`Parameters`

These key-value pairs define partition parameters.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageDescriptor`

Provides information about the physical
location where the partition is stored.

_Required_: No

_Type_: [StorageDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-partition-storagedescriptor.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values of the partition. Although this parameter is not required by the SDK, you must specify this parameter for a valid input.

The values for the keys for the new partition must be passed as an array of String objects that must be ordered in the same order as the partition keys appearing in the Amazon S3 prefix. Otherwise AWS Glue will add the values to the wrong keys.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PartitionInput](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-partitions.html#aws-glue-api-catalog-partitions-PartitionInput) in the _AWS Glue Developer_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Order

SchemaId
