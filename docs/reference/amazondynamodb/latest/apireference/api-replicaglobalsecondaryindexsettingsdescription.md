# ReplicaGlobalSecondaryIndexSettingsDescription

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

**IndexStatus**

The current status of the global secondary index:

- `CREATING` \- The global secondary index is being created.

- `UPDATING` \- The global secondary index is being updated.

- `DELETING` \- The global secondary index is being deleted.

- `ACTIVE` \- The global secondary index is ready for use.

Type: String

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE`

Required: No

**ProvisionedReadCapacityAutoScalingSettings**

Auto scaling settings for a global secondary index replica's read capacity
units.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ProvisionedReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**ProvisionedWriteCapacityAutoScalingSettings**

Auto scaling settings for a global secondary index replica's write capacity
units.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ProvisionedWriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`.

Type: Long

Valid Range: Minimum value of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicaglobalsecondaryindexsettingsdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicaglobalsecondaryindexsettingsdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicaglobalsecondaryindexsettingsdescription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ReplicaGlobalSecondaryIndexDescription

ReplicaGlobalSecondaryIndexSettingsUpdate
