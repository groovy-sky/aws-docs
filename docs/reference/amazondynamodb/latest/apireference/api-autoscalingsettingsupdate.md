# AutoScalingSettingsUpdate

Represents the auto scaling settings to be modified for a global table or global
secondary index.

## Contents

###### Note

In the following list, the required parameters are described first.

**AutoScalingDisabled**

Disabled auto scaling for this global table or global secondary index.

Type: Boolean

Required: No

**AutoScalingRoleArn**

Role ARN used for configuring auto scaling policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

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

**ScalingPolicyUpdate**

The scaling policy to apply for scaling target global table or global secondary index
capacity units.

Type: [AutoScalingPolicyUpdate](api-autoscalingpolicyupdate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/autoscalingsettingsupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/autoscalingsettingsupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/autoscalingsettingsupdate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AutoScalingSettingsDescription

AutoScalingTargetTrackingScalingPolicyConfigurationDescription
