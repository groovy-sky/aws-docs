# ReplicaAutoScalingDescription

Represents the auto scaling settings of the replica.

## Contents

###### Note

In the following list, the required parameters are described first.

**GlobalSecondaryIndexes**

Replica-specific global secondary index auto scaling settings.

Type: Array of [ReplicaGlobalSecondaryIndexAutoScalingDescription](api-replicaglobalsecondaryindexautoscalingdescription.md) objects

Required: No

**RegionName**

The Region where the replica exists.

Type: String

Required: No

**ReplicaProvisionedReadCapacityAutoScalingSettings**

Represents the auto scaling settings for a global table or global secondary
index.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ReplicaProvisionedWriteCapacityAutoScalingSettings**

Represents the auto scaling settings for a global table or global secondary
index.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ReplicaStatus**

The current state of the replica:

- `CREATING` \- The replica is being created.

- `UPDATING` \- The replica is being updated.

- `DELETING` \- The replica is being deleted.

- `ACTIVE` \- The replica is ready for use.

Type: String

Valid Values: `CREATING | CREATION_FAILED | UPDATING | DELETING | ACTIVE | REGION_DISABLED | INACCESSIBLE_ENCRYPTION_CREDENTIALS | ARCHIVING | ARCHIVED | REPLICATION_NOT_AUTHORIZED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicaautoscalingdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicaautoscalingdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicaautoscalingdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Replica

ReplicaAutoScalingUpdate

All content copied from https://docs.aws.amazon.com/.
