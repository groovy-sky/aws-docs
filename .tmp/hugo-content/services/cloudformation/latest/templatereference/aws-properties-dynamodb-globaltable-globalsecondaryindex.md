This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable GlobalSecondaryIndex

Allows you to specify a global secondary index for the global table. The index will be
defined on all replicas.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IndexName" : String,
  "KeySchema" : [ KeySchema, ... ],
  "Projection" : Projection,
  "ReadOnDemandThroughputSettings" : ReadOnDemandThroughputSettings,
  "ReadProvisionedThroughputSettings" : GlobalReadProvisionedThroughputSettings,
  "WarmThroughput" : WarmThroughput,
  "WriteOnDemandThroughputSettings" : WriteOnDemandThroughputSettings,
  "WriteProvisionedThroughputSettings" : WriteProvisionedThroughputSettings
}

```

### YAML

```yaml

  IndexName: String
  KeySchema:
    - KeySchema
  Projection:
    Projection
  ReadOnDemandThroughputSettings:
    ReadOnDemandThroughputSettings
  ReadProvisionedThroughputSettings:
    GlobalReadProvisionedThroughputSettings
  WarmThroughput:
    WarmThroughput
  WriteOnDemandThroughputSettings:
    WriteOnDemandThroughputSettings
  WriteProvisionedThroughputSettings:
    WriteProvisionedThroughputSettings

```

## Properties

`IndexName`

The name of the global secondary index. The name must be unique among all other indexes
on this table.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`KeySchema`

The complete key schema for a global secondary index, which consists of one or more
pairs of attribute names and key types:

- `HASH` \- partition key

- `RANGE` \- sort key

###### Note

The partition key of an item is also known as its _hash_
_attribute_. The term "hash attribute" derives from DynamoDB's usage of an
internal hash function to evenly distribute data items across partitions, based on their
partition key values.

The sort key of an item is also known as its _range attribute_.
The term "range attribute" derives from the way DynamoDB stores items with the same
partition key physically close together, in sorted order by the sort key value.

_Required_: Yes

_Type_: [Array](aws-properties-dynamodb-globaltable-keyschema.md) of [KeySchema](aws-properties-dynamodb-globaltable-keyschema.md)

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: Updates are not supported.

`Projection`

Represents attributes that are copied (projected) from the table into the global
secondary index. These are in addition to the primary key attributes and index key
attributes, which are automatically projected.

_Required_: Yes

_Type_: [Projection](aws-properties-dynamodb-globaltable-projection.md)

_Update requires_: Updates are not supported.

`ReadOnDemandThroughputSettings`

Sets read capacity settings for the global secondary index of the global table, which applies to all
replicas. This can only be applied for multi-account global tables.

_Required_: No

_Type_: [ReadOnDemandThroughputSettings](aws-properties-dynamodb-globaltable-readondemandthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadProvisionedThroughputSettings`

Sets the read request settings for the global secondary index of the global table, which applies to all
replicas. This can only be applied for multi-account global tables.

_Required_: No

_Type_: [GlobalReadProvisionedThroughputSettings](aws-properties-dynamodb-globaltable-globalreadprovisionedthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmThroughput`

Represents the warm throughput value (in read units per second and write units per
second) for the specified secondary index. If you use this parameter, you must specify
`ReadUnitsPerSecond`, `WriteUnitsPerSecond`, or both.

_Required_: No

_Type_: [WarmThroughput](aws-properties-dynamodb-globaltable-warmthroughput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteOnDemandThroughputSettings`

Sets the write request settings for a global table or a global secondary index. You can
only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode`.

_Required_: No

_Type_: [WriteOnDemandThroughputSettings](aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteProvisionedThroughputSettings`

Defines write capacity settings for the global secondary index. You must specify a value
for this property if the table's `BillingMode` is `PROVISIONED`. All
replicas will have the same write capacity settings for this global secondary index.

_Required_: No

_Type_: [WriteProvisionedThroughputSettings](aws-properties-dynamodb-globaltable-writeprovisionedthroughputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalReadProvisionedThroughputSettings

GlobalTableWitness

All content copied from https://docs.aws.amazon.com/.
