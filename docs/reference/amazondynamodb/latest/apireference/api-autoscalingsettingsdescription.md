# AutoScalingSettingsDescription

Represents the auto scaling settings for a global table or global secondary
index.

## Contents

###### Note

In the following list, the required parameters are described first.

**AutoScalingDisabled**

Disabled auto scaling for this global table or global secondary index.

Type: Boolean

Required: No

**AutoScalingRoleArn**

Role ARN used for configuring the auto scaling policy.

Type: String

Required: No

**MaximumUnits**

The maximum capacity units that a global table or global secondary index should be
scaled up to.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**MinimumUnits**

The minimum capacity units that a global table or global secondary index should be
scaled down to.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**ScalingPolicies**

Information about the scaling policies.

Type: Array of [AutoScalingPolicyDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingPolicyDescription.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/AutoScalingSettingsDescription)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/AutoScalingSettingsDescription)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/AutoScalingSettingsDescription)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AutoScalingPolicyUpdate

AutoScalingSettingsUpdate
