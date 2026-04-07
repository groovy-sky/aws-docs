# GlobalTableGlobalSecondaryIndexSettingsUpdate

Represents the settings of a global secondary index for a global table that will be
modified.

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

**ProvisionedWriteCapacityAutoScalingSettingsUpdate**

Auto scaling settings for managing a global secondary index's write capacity
units.

Type: [AutoScalingSettingsUpdate](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingSettingsUpdate.html) object

Required: No

**ProvisionedWriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException.`

Type: Long

Valid Range: Minimum value of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/GlobalTableGlobalSecondaryIndexSettingsUpdate)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/GlobalTableGlobalSecondaryIndexSettingsUpdate)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/GlobalTableGlobalSecondaryIndexSettingsUpdate)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GlobalTableDescription

GlobalTableWitnessDescription
