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

Type: [BillingModeSummary](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BillingModeSummary.html) object

Required: No

**ReplicaGlobalSecondaryIndexSettings**

Replica global secondary index settings for the global table.

Type: Array of [ReplicaGlobalSecondaryIndexSettingsDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaGlobalSecondaryIndexSettingsDescription.html) objects

Required: No

**ReplicaProvisionedReadCapacityAutoScalingSettings**

Auto scaling settings for a global table replica's read capacity units.

Type: [AutoScalingSettingsDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingSettingsDescription.html) object

Required: No

**ReplicaProvisionedReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the _Amazon DynamoDB_
_Developer Guide_.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ReplicaProvisionedWriteCapacityAutoScalingSettings**

Auto scaling settings for a global table replica's write capacity units.

Type: [AutoScalingSettingsDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingSettingsDescription.html) object

Required: No

**ReplicaProvisionedWriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`. For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the _Amazon DynamoDB_
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

Type: [TableClassSummary](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableClassSummary.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ReplicaSettingsDescription)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ReplicaSettingsDescription)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ReplicaSettingsDescription)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicaGlobalSecondaryIndexSettingsUpdate

ReplicaSettingsUpdate
