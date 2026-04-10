This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table Schema

A Schema specifies the expected data model of the table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CompositePartitionKey" : [ PartitionKey, ... ]
}

```

### YAML

```yaml

  CompositePartitionKey:
    - PartitionKey

```

## Properties

`CompositePartitionKey`

A non-empty list of partition keys defining the attributes used to partition the table data. The order of the
list determines the partition hierarchy. The name and type of each partition key as well as the partition key order
cannot be changed after the table is created. However, the enforcement level of each partition key can be changed.

_Required_: No

_Type_: Array of [PartitionKey](aws-properties-timestream-table-partitionkey.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Configuration

Tag

All content copied from https://docs.aws.amazon.com/.
