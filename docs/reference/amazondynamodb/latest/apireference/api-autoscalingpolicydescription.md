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

Type: [AutoScalingTargetTrackingScalingPolicyConfigurationDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AutoScalingTargetTrackingScalingPolicyConfigurationDescription.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/AutoScalingPolicyDescription)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/AutoScalingPolicyDescription)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/AutoScalingPolicyDescription)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttributeValueUpdate

AutoScalingPolicyUpdate
