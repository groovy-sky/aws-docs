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

Type: Array of [AutoScalingPolicyDescription](api-autoscalingpolicydescription.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/autoscalingsettingsdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/autoscalingsettingsdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/autoscalingsettingsdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingPolicyUpdate

AutoScalingSettingsUpdate

All content copied from https://docs.aws.amazon.com/.
