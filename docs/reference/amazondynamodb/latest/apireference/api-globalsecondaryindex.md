# GlobalSecondaryIndex

Represents the properties of a global secondary index.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

The name of the global secondary index. The name must be unique among all other
indexes on this table.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**KeySchema**

The complete key schema for a global secondary index, which consists of one or more
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

Required: Yes

**Projection**

Represents attributes that are copied (projected) from the table into the global
secondary index. These are in addition to the primary key attributes and index key
attributes, which are automatically projected.

Type: [Projection](api-projection.md) object

Required: Yes

**OnDemandThroughput**

The maximum number of read and write units for the specified global secondary index.
If you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both. You must use either
`OnDemandThroughput` or `ProvisionedThroughput` based on your
table's capacity mode.

Type: [OnDemandThroughput](api-ondemandthroughput.md) object

Required: No

**ProvisionedThroughput**

Represents the provisioned throughput settings for the specified global secondary
index. You must use either `OnDemandThroughput` or
`ProvisionedThroughput` based on your table's capacity mode.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: [ProvisionedThroughput](api-provisionedthroughput.md) object

Required: No

**WarmThroughput**

Represents the warm throughput value (in read units per second and write units per
second) for the specified secondary index. If you use this parameter, you must specify
`ReadUnitsPerSecond`, `WriteUnitsPerSecond`, or both.

Type: [WarmThroughput](api-warmthroughput.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/globalsecondaryindex.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/globalsecondaryindex.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/globalsecondaryindex.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Get

GlobalSecondaryIndexAutoScalingUpdate

All content copied from https://docs.aws.amazon.com/.
