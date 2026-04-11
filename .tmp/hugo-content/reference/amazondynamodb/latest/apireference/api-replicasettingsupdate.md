# ReplicaSettingsUpdate

Represents the settings for a global table in a Region that will be modified.

## Contents

###### Note

In the following list, the required parameters are described first.

**RegionName**

The Region of the replica to be added.

Type: String

Required: Yes

**ReplicaGlobalSecondaryIndexSettingsUpdate**

Represents the settings of a global secondary index for a global table that will be
modified.

Type: Array of [ReplicaGlobalSecondaryIndexSettingsUpdate](api-replicaglobalsecondaryindexsettingsupdate.md) objects

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: No

**ReplicaProvisionedReadCapacityAutoScalingSettingsUpdate**

Auto scaling settings for managing a global table replica's read capacity
units.

Type: [AutoScalingSettingsUpdate](api-autoscalingsettingsupdate.md) object

Required: No

**ReplicaProvisionedReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. For more information, see [Specifying Read and Write Requirements](../../../../services/dynamodb/latest/developerguide/workingwithtables.md#ProvisionedThroughput) in the _Amazon DynamoDB_
_Developer Guide_.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**ReplicaTableClass**

Replica-specific table class. If not specified, uses the source table's table
class.

Type: String

Valid Values: `STANDARD | STANDARD_INFREQUENT_ACCESS`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicasettingsupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicasettingsupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicasettingsupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaSettingsDescription

ReplicationGroupUpdate

All content copied from https://docs.aws.amazon.com/.
