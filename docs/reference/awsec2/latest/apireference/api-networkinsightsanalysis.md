# NetworkInsightsAnalysis

Describes a network insights analysis.

## Contents

**AdditionalAccountSet.N**

The member accounts that contain resources that the path can traverse.

Type: Array of strings

Required: No

**AlternatePathHintSet.N**

Potential intermediate components.

Type: Array of [AlternatePathHint](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AlternatePathHint.html) objects

Required: No

**ExplanationSet.N**

The explanations. For more information, see [Reachability Analyzer explanation codes](https://docs.aws.amazon.com/vpc/latest/reachability/explanation-codes.html).

Type: Array of [Explanation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Explanation.html) objects

Required: No

**FilterInArnSet.N**

The Amazon Resource Names (ARN) of the resources that the path must traverse.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**FilterOutArnSet.N**

The Amazon Resource Names (ARN) of the resources that the path must ignore.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**ForwardPathComponentSet.N**

The components in the path from source to destination.

Type: Array of [PathComponent](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PathComponent.html) objects

Required: No

**networkInsightsAnalysisArn**

The Amazon Resource Name (ARN) of the network insights analysis.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**networkInsightsAnalysisId**

The ID of the network insights analysis.

Type: String

Required: No

**networkInsightsPathId**

The ID of the path.

Type: String

Required: No

**networkPathFound**

Indicates whether the destination is reachable from the source.

Type: Boolean

Required: No

**ReturnPathComponentSet.N**

The components in the path from destination to source.

Type: Array of [PathComponent](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PathComponent.html) objects

Required: No

**startDate**

The time the analysis started.

Type: Timestamp

Required: No

**status**

The status of the network insights analysis.

Type: String

Valid Values: `running | succeeded | failed`

Required: No

**statusMessage**

The status message, if the status is `failed`.

Type: String

Required: No

**SuggestedAccountSet.N**

Potential intermediate accounts.

Type: Array of strings

Required: No

**TagSet.N**

The tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

**warningMessage**

The warning message.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/NetworkInsightsAnalysis)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/NetworkInsightsAnalysis)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/NetworkInsightsAnalysis)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkInsightsAccessScopeContent

NetworkInsightsPath
