# GeneratedPolicyProperties

Contains the generated policy details.

## Contents

**principalArn**

The ARN of the IAM entity (user or role) for which you are generating a policy.

Type: String

Pattern: `arn:[^:]*:iam::[^:]*:(role|user)/.{1,576}`

Required: Yes

**cloudTrailProperties**

Lists details about the `Trail` used to generated policy.

Type: [CloudTrailProperties](api-cloudtrailproperties.md) object

Required: No

**isComplete**

This value is set to `true` if the generated policy contains all possible
actions for a service that IAM Access Analyzer identified from the CloudTrail trail that you specified,
and `false` otherwise.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/GeneratedPolicyProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/GeneratedPolicyProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/GeneratedPolicyProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GeneratedPolicy

GeneratedPolicyResult
