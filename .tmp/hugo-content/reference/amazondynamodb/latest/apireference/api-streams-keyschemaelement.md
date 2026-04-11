# KeySchemaElement

Represents _a single element_ of a key schema. A key schema specifies
the attributes that make up the primary key of a table, or the key attributes of an
index.

A `KeySchemaElement` represents exactly one attribute of the primary key. For
example, a simple primary key would be represented by one `KeySchemaElement`
(for the partition key). A composite primary key would require one
`KeySchemaElement` for the partition key, and another
`KeySchemaElement` for the sort key.

A `KeySchemaElement` must be a scalar, top-level attribute (not a nested
attribute). The data type must be one of String, Number, or Binary. The attribute cannot
be nested within a List or a Map.

## Contents

###### Note

In the following list, the required parameters are described first.

**AttributeName**

The name of a key attribute.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**KeyType**

The role that this key attribute will assume:

- `HASH` \- partition key

- `RANGE` \- sort key

###### Note

The partition key of an item is also known as its _hash_
_attribute_. The term "hash attribute" derives from DynamoDB's usage of
an internal hash function to evenly distribute data items across partitions, based
on their partition key values.

The sort key of an item is also known as its _range_
_attribute_. The term "range attribute" derives from the way DynamoDB
stores items with the same partition key physically close together, in sorted order
by the sort key value.

Type: String

Valid Values: `HASH | RANGE`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/keyschemaelement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/keyschemaelement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/keyschemaelement.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Identity

Record

All content copied from https://docs.aws.amazon.com/.
