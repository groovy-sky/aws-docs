# NetworkInsightsAnalysis

Describes a network insights analysis.

## Contents

**AdditionalAccountSet.N**

The member accounts that contain resources that the path can traverse.

Type: Array of strings

Required: No

**AlternatePathHintSet.N**

Potential intermediate components.

Type: Array of [AlternatePathHint](api-alternatepathhint.md) objects

Required: No

**ExplanationSet.N**

The explanations. For more information, see [Reachability Analyzer explanation codes](../../../../services/vpc/latest/reachability/explanation-codes.md).

Type: Array of [Explanation](api-explanation.md) objects

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

Type: Array of [PathComponent](api-pathcomponent.md) objects

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

Type: Array of [PathComponent](api-pathcomponent.md) objects

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinsightsanalysis.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinsightsanalysis.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinsightsanalysis.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInsightsAccessScopeContent

NetworkInsightsPath
