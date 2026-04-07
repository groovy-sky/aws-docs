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

Type: [AutoScalingSettingsDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingSettingsDescription.html) object

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

Type: [AutoScalingSettingsDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingSettingsDescription.html) object

Required: No

**ProvisionedWriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`.

Type: Long

Valid Range: Minimum value of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ReplicaGlobalSecondaryIndexSettingsDescription)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ReplicaGlobalSecondaryIndexSettingsDescription)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ReplicaGlobalSecondaryIndexSettingsDescription)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicaGlobalSecondaryIndexDescription

ReplicaGlobalSecondaryIndexSettingsUpdate
