This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table LocalSecondaryIndex

Represents the properties of a local secondary index. A local secondary index can only
be created when its parent table is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IndexName" : String,
  "KeySchema" : [ KeySchema, ... ],
  "Projection" : Projection
}

```

### YAML

```yaml

  IndexName: String
  KeySchema:
    - KeySchema
  Projection:
    Projection

```

## Properties

`IndexName`

The name of the local secondary index. The name must be unique among all other indexes
on this table.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`KeySchema`

The complete key schema for the local secondary index, consisting of one or more pairs
of attribute names and key types:

- `HASH` \- partition key

- `RANGE` \- sort key

###### Note

The partition key of an item is also known as its _hash_
_attribute_. The term "hash attribute" derives from DynamoDB's usage of
an internal hash function to evenly distribute data items across partitions, based
on their partition key values.

The sort key of an item is also known as its _range attribute_.
The term "range attribute" derives from the way DynamoDB stores items with the same
partition key physically close together, in sorted order by the sort key
value.

_Required_: Yes

_Type_: [Array](aws-properties-dynamodb-table-keyschema.md) of [KeySchema](aws-properties-dynamodb-table-keyschema.md)

_Minimum_: `1`

_Update requires_: Updates are not supported.

`Projection`

Represents attributes that are copied (projected) from the table into the local
secondary index. These are in addition to the primary key attributes and index key
attributes, which are automatically projected.

_Required_: Yes

_Type_: [Projection](aws-properties-dynamodb-table-projection.md)

_Update requires_: Updates are not supported.

## See also

For an example of a declared local secondary index, see [AWS::DynamoDB::Table](../userguide/aws-resource-dynamodb-table.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisStreamSpecification

OnDemandThroughput

All content copied from https://docs.aws.amazon.com/.
