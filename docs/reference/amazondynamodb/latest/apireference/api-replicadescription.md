# ReplicaDescription

Contains the details of the replica.

## Contents

###### Note

In the following list, the required parameters are described first.

**GlobalSecondaryIndexes**

Replica-specific global secondary index settings.

Type: Array of [ReplicaGlobalSecondaryIndexDescription](api-replicaglobalsecondaryindexdescription.md) objects

Required: No

**GlobalTableSettingsReplicationMode**

Indicates one of the settings synchronization modes for the global table replica:

- `ENABLED`: Indicates that the settings synchronization mode for the global table
replica is enabled.

- `DISABLED`: Indicates that the settings synchronization mode for the global table
replica is disabled.

- `ENABLED_WITH_OVERRIDES`: This mode is set by default for a same account global table.
Indicates that certain global table settings can be overridden.

Type: String

Valid Values: `ENABLED | DISABLED | ENABLED_WITH_OVERRIDES`

Required: No

**KMSMasterKeyId**

The AWS KMS key of the replica that will be used for AWS KMS
encryption.

Type: String

Required: No

**OnDemandThroughputOverride**

Overrides the maximum on-demand throughput settings for the specified replica
table.

Type: [OnDemandThroughputOverride](api-ondemandthroughputoverride.md) object

Required: No

**ProvisionedThroughputOverride**

Replica-specific provisioned throughput. If not described, uses the source table's
provisioned throughput settings.

Type: [ProvisionedThroughputOverride](api-provisionedthroughputoverride.md) object

Required: No

**RegionName**

The name of the Region.

Type: String

Required: No

**ReplicaArn**

The Amazon Resource Name (ARN) of the global table replica.

Type: String

Required: No

**ReplicaInaccessibleDateTime**

The time at which the replica was first detected as inaccessible. To determine cause
of inaccessibility check the `ReplicaStatus` property.

Type: Timestamp

Required: No

**ReplicaStatus**

The current state of the replica:

- `CREATING` \- The replica is being created.

- `UPDATING` \- The replica is being updated.

- `DELETING` \- The replica is being deleted.

- `ACTIVE` \- The replica is ready for use.

- `REGION_DISABLED` \- The replica is inaccessible because the AWS Region has been disabled.

###### Note

If the AWS Region remains inaccessible for more than 20
hours, DynamoDB will remove this replica from the replication
group. The replica will not be deleted and replication will stop from and to
this region.

- `INACCESSIBLE_ENCRYPTION_CREDENTIALS ` \- The AWS KMS key
used to encrypt the table is inaccessible.

###### Note

If the AWS KMS key remains inaccessible for more than 20 hours,
DynamoDB will remove this replica from the replication group.
The replica will not be deleted and replication will stop from and to this
region.

Type: String

Valid Values: `CREATING | CREATION_FAILED | UPDATING | DELETING | ACTIVE | REGION_DISABLED | INACCESSIBLE_ENCRYPTION_CREDENTIALS | ARCHIVING | ARCHIVED | REPLICATION_NOT_AUTHORIZED`

Required: No

**ReplicaStatusDescription**

Detailed information about the replica status.

Type: String

Required: No

**ReplicaStatusPercentProgress**

Specifies the progress of a Create, Update, or Delete action on the replica as a
percentage.

Type: String

Required: No

**ReplicaTableClassSummary**

Contains details of the table class.

Type: [TableClassSummary](api-tableclasssummary.md) object

Required: No

**WarmThroughput**

Represents the warm throughput value for this replica.

Type: [TableWarmThroughputDescription](api-tablewarmthroughputdescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicadescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicadescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicadescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaAutoScalingUpdate

ReplicaGlobalSecondaryIndex

All content copied from https://docs.aws.amazon.com/.
