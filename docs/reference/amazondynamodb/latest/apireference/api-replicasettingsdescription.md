# ReplicaSettingsDescription

Represents the properties of a replica.

## Contents

###### Note

In the following list, the required parameters are described first.

**RegionName**

The Region name of the replica.

Type: String

Required: Yes

**ReplicaBillingModeSummary**

The read/write capacity mode of the replica.

Type: [BillingModeSummary](api-billingmodesummary.md) object

Required: No

**ReplicaGlobalSecondaryIndexSettings**

Replica global secondary index settings for the global table.

Type: Array of [ReplicaGlobalSecondaryIndexSettingsDescription](api-replicaglobalsecondaryindexsettingsdescription.md) objects

Required: No

**ReplicaProvisionedReadCapacityAutoScalingSettings**

Auto scaling settings for a global table replica's read capacity units.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ReplicaProvisionedReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. For more information, see [Specifying Read and Write Requirements](../../../../services/dynamodb/latest/developerguide/workingwithtables.md#ProvisionedThroughput) in the _Amazon DynamoDB_
_Developer Guide_.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ReplicaProvisionedWriteCapacityAutoScalingSettings**

Auto scaling settings for a global table replica's write capacity units.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ReplicaProvisionedWriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`. For more information, see [Specifying Read and Write Requirements](../../../../services/dynamodb/latest/developerguide/workingwithtables.md#ProvisionedThroughput) in the _Amazon DynamoDB_
_Developer Guide_.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ReplicaStatus**

The current state of the Region:

- `CREATING` \- The Region is being created.

- `UPDATING` \- The Region is being updated.

- `DELETING` \- The Region is being deleted.

- `ACTIVE` \- The Region is ready for use.

Type: String

Valid Values: `CREATING | CREATION_FAILED | UPDATING | DELETING | ACTIVE | REGION_DISABLED | INACCESSIBLE_ENCRYPTION_CREDENTIALS | ARCHIVING | ARCHIVED | REPLICATION_NOT_AUTHORIZED`

Required: No

**ReplicaTableClassSummary**

Contains details of the table class.

Type: [TableClassSummary](api-tableclasssummary.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicasettingsdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicasettingsdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicasettingsdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaGlobalSecondaryIndexSettingsUpdate

ReplicaSettingsUpdate

All content copied from https://docs.aws.amazon.com/.
