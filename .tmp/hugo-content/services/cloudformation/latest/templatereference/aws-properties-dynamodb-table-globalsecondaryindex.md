This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table GlobalSecondaryIndex

Represents the properties of a global secondary index.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContributorInsightsSpecification" : ContributorInsightsSpecification,
  "IndexName" : String,
  "KeySchema" : [ KeySchema, ... ],
  "OnDemandThroughput" : OnDemandThroughput,
  "Projection" : Projection,
  "ProvisionedThroughput" : ProvisionedThroughput,
  "WarmThroughput" : WarmThroughput
}

```

### YAML

```yaml

  ContributorInsightsSpecification:
    ContributorInsightsSpecification
  IndexName: String
  KeySchema:
    - KeySchema
  OnDemandThroughput:
    OnDemandThroughput
  Projection:
    Projection
  ProvisionedThroughput:
    ProvisionedThroughput
  WarmThroughput:
    WarmThroughput

```

## Properties

`ContributorInsightsSpecification`

The settings used to specify whether to enable CloudWatch Contributor Insights for the
global table and define which events to monitor.

_Required_: No

_Type_: [ContributorInsightsSpecification](aws-properties-dynamodb-table-contributorinsightsspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The name of the global secondary index. The name must be unique among all other
indexes on this table.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

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

`OnDemandThroughput`

The maximum number of read and write units for the specified global secondary index.
If you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both. You must use either
`OnDemandThroughput` or `ProvisionedThroughput` based on your
table's capacity mode.

_Required_: No

_Type_: [OnDemandThroughput](aws-properties-dynamodb-table-ondemandthroughput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Projection`

Represents attributes that are copied (projected) from the table into the global
secondary index. These are in addition to the primary key attributes and index key
attributes, which are automatically projected.

_Required_: Yes

_Type_: [Projection](aws-properties-dynamodb-table-projection.md)

_Update requires_: Updates are not supported.

`ProvisionedThroughput`

Represents the provisioned throughput settings for the specified global secondary
index. You must use either `OnDemandThroughput` or
`ProvisionedThroughput` based on your table's capacity mode.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

_Required_: No

_Type_: [ProvisionedThroughput](aws-properties-dynamodb-table-provisionedthroughput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmThroughput`

Represents the warm throughput value (in read units per second and write units per
second) for the specified secondary index. If you use this parameter, you must specify
`ReadUnitsPerSecond`, `WriteUnitsPerSecond`, or both.

_Required_: No

_Type_: [WarmThroughput](aws-properties-dynamodb-table-warmthroughput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Csv

ImportSourceSpecification

All content copied from https://docs.aws.amazon.com/.
