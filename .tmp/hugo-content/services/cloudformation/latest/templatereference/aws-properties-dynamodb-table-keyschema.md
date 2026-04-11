This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table KeySchema

Represents _a single element_ of a key schema. A key schema
specifies the attributes that make up the primary key of a table, or the key attributes
of an index.

A `KeySchemaElement` represents exactly one attribute of the primary key.
For example, a simple primary key would be represented by one
`KeySchemaElement` (for the partition key). A composite primary key would
require one `KeySchemaElement` for the partition key, and another
`KeySchemaElement` for the sort key.

A `KeySchemaElement` must be a scalar, top-level attribute (not a nested
attribute). The data type must be one of String, Number, or Binary. The attribute cannot
be nested within a List or a Map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "KeyType" : String
}

```

### YAML

```yaml

  AttributeName: String
  KeyType: String

```

## Properties

`AttributeName`

The name of a key attribute.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`KeyType`

The role that this key attribute will assume:

- `HASH` \- partition key

- `RANGE` \- sort key

###### Note

The partition key of an item is also known as its _hash_
_attribute_. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across
partitions, based on their partition key values.

The sort key of an item is also known as its _range attribute_.
The term "range attribute" derives from the way DynamoDB stores items with
the same partition key physically close together, in sorted order by the sort key
value.

_Required_: Yes

_Type_: String

_Allowed values_: `HASH | RANGE`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## See also

For an example of a declared key schema, see [AWS::DynamoDB::Table](../userguide/aws-resource-dynamodb-table.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputFormatOptions

KinesisStreamSpecification

All content copied from https://docs.aws.amazon.com/.
