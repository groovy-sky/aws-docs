# GlobalSecondaryIndexDescription

Represents the properties of a global secondary index.

## Contents

###### Note

In the following list, the required parameters are described first.

**Backfilling**

Indicates whether the index is currently backfilling. _Backfilling_
is the process of reading items from the table and determining whether they can be added
to the index. (Not all items will qualify: For example, a partition key cannot have any
duplicate values.) If an item can be added to the index, DynamoDB will do so. After all
items have been processed, the backfilling operation is complete and
`Backfilling` is false.

You can delete an index that is being created during the `Backfilling`
phase when `IndexStatus` is set to CREATING and `Backfilling` is
true. You can't delete the index that is being created when `IndexStatus` is
set to CREATING and `Backfilling` is false.

###### Note

For indexes that were created during a `CreateTable` operation, the
`Backfilling` attribute does not appear in the
`DescribeTable` output.

Type: Boolean

Required: No

**IndexArn**

The Amazon Resource Name (ARN) that uniquely identifies the index.

Type: String

Required: No

**IndexName**

The name of the global secondary index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**IndexSizeBytes**

The total size of the specified index, in bytes. DynamoDB updates this value
approximately every six hours. Recent changes might not be reflected in this
value.

Type: Long

Required: No

**IndexStatus**

The current state of the global secondary index:

- `CREATING` \- The index is being created.

- `UPDATING` \- The index is being updated.

- `DELETING` \- The index is being deleted.

- `ACTIVE` \- The index is ready for use.

Type: String

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE`

Required: No

**ItemCount**

The number of items in the specified index. DynamoDB updates this value approximately
every six hours. Recent changes might not be reflected in this value.

Type: Long

Required: No

**KeySchema**

The complete key schema for a global secondary index, which consists of one or more
pairs of attribute names and key types:

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

Type: Array of [KeySchemaElement](api-keyschemaelement.md) objects

Array Members: Minimum number of 1 item.

Required: No

**OnDemandThroughput**

The maximum number of read and write units for the specified global secondary index.
If you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

Type: [OnDemandThroughput](api-ondemandthroughput.md) object

Required: No

**Projection**

Represents attributes that are copied (projected) from the table into the global
secondary index. These are in addition to the primary key attributes and index key
attributes, which are automatically projected.

Type: [Projection](api-projection.md) object

Required: No

**ProvisionedThroughput**

Represents the provisioned throughput settings for the specified global secondary
index.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: [ProvisionedThroughputDescription](api-provisionedthroughputdescription.md) object

Required: No

**WarmThroughput**

Represents the warm throughput value (in read units per second and write units per
second) for the specified secondary index.

Type: [GlobalSecondaryIndexWarmThroughputDescription](api-globalsecondaryindexwarmthroughputdescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/globalsecondaryindexdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/globalsecondaryindexdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/globalsecondaryindexdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalSecondaryIndexAutoScalingUpdate

GlobalSecondaryIndexInfo

All content copied from https://docs.aws.amazon.com/.
