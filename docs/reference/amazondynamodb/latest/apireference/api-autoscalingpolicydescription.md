# AutoScalingPolicyDescription

Represents the properties of the scaling policy.

## Contents

###### Note

In the following list, the required parameters are described first.

**PolicyName**

The name of the scaling policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `\p{Print}+`

Required: No

**TargetTrackingScalingPolicyConfiguration**

Represents a target tracking scaling policy configuration.

Type: [AutoScalingTargetTrackingScalingPolicyConfigurationDescription](api-autoscalingtargettrackingscalingpolicyconfigurationdescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/autoscalingpolicydescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/autoscalingpolicydescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/autoscalingpolicydescription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttributeValueUpdate

AutoScalingPolicyUpdate
