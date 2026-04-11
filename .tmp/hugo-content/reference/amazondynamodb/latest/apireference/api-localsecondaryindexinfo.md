# LocalSecondaryIndexInfo

Represents the properties of a local secondary index for the table when the backup was
created.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

Represents the name of the local secondary index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**KeySchema**

The complete key schema for a local secondary index, which consists of one or more
pairs of attribute names and key types:

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

Type: Array of [KeySchemaElement](api-keyschemaelement.md) objects

Array Members: Minimum number of 1 item.

Required: No

**Projection**

Represents attributes that are copied (projected) from the table into the global
secondary index. These are in addition to the primary key attributes and index key
attributes, which are automatically projected.

Type: [Projection](api-projection.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/localsecondaryindexinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/localsecondaryindexinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/localsecondaryindexinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalSecondaryIndexDescription

OnDemandThroughput

All content copied from https://docs.aws.amazon.com/.
