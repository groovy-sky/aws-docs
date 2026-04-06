# PolicyGeneration

Contains details about the policy generation status and properties.

## Contents

**jobId**

The `JobId` that is returned by the `StartPolicyGeneration`
operation. The `JobId` can be used with `GetGeneratedPolicy` to
retrieve the generated policies or used with `CancelPolicyGeneration` to cancel
the policy generation request.

Type: String

Required: Yes

**principalArn**

The ARN of the IAM entity (user or role) for which you are generating a policy.

Type: String

Pattern: `arn:[^:]*:iam::[^:]*:(role|user)/.{1,576}`

Required: Yes

**startedOn**

A timestamp of when the policy generation started.

Type: Timestamp

Required: Yes

**status**

The status of the policy generation request.

Type: String

Valid Values: `IN_PROGRESS | SUCCEEDED | FAILED | CANCELED`

Required: Yes

**completedOn**

A timestamp of when the policy generation was completed.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/PolicyGeneration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/PolicyGeneration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/PolicyGeneration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PathElement

PolicyGenerationDetails
