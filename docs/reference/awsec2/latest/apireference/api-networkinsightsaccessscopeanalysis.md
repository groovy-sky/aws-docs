# NetworkInsightsAccessScopeAnalysis

Describes a Network Access Scope analysis.

## Contents

**analyzedEniCount**

The number of network interfaces analyzed.

Type: Integer

Required: No

**endDate**

The analysis end date.

Type: Timestamp

Required: No

**findingsFound**

Indicates whether there are findings.

Type: String

Valid Values: `true | false | unknown`

Required: No

**networkInsightsAccessScopeAnalysisArn**

The Amazon Resource Name (ARN) of the Network Access Scope analysis.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**networkInsightsAccessScopeAnalysisId**

The ID of the Network Access Scope analysis.

Type: String

Required: No

**networkInsightsAccessScopeId**

The ID of the Network Access Scope.

Type: String

Required: No

**startDate**

The analysis start date.

Type: Timestamp

Required: No

**status**

The status.

Type: String

Valid Values: `running | succeeded | failed`

Required: No

**statusMessage**

The status message.

Type: String

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinsightsaccessscopeanalysis.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinsightsaccessscopeanalysis.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinsightsaccessscopeanalysis.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInsightsAccessScope

NetworkInsightsAccessScopeContent
