# ReplicaAutoScalingUpdate

Represents the auto scaling settings of a replica that will be modified.

## Contents

###### Note

In the following list, the required parameters are described first.

**RegionName**

The Region where the replica exists.

Type: String

Required: Yes

**ReplicaGlobalSecondaryIndexUpdates**

Represents the auto scaling settings of global secondary indexes that will be
modified.

Type: Array of [ReplicaGlobalSecondaryIndexAutoScalingUpdate](api-replicaglobalsecondaryindexautoscalingupdate.md) objects

Required: No

**ReplicaProvisionedReadCapacityAutoScalingUpdate**

Represents the auto scaling settings to be modified for a global table or global
secondary index.

Type: [AutoScalingSettingsUpdate](api-autoscalingsettingsupdate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicaautoscalingupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicaautoscalingupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicaautoscalingupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaAutoScalingDescription

ReplicaDescription

All content copied from https://docs.aws.amazon.com/.
