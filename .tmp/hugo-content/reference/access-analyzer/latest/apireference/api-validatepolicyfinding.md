# ValidatePolicyFinding

A finding in a policy. Each finding is an actionable recommendation that can be used to
improve the policy.

## Contents

**findingDetails**

A localized message that explains the finding and provides guidance on how to address
it.

Type: String

Required: Yes

**findingType**

The impact of the finding.

Security warnings report when the policy allows access that we consider overly
permissive.

Errors report when a part of the policy is not functional.

Warnings report non-security issues when a policy does not conform to policy writing
best practices.

Suggestions recommend stylistic improvements in the policy that do not impact
access.

Type: String

Valid Values: `ERROR | SECURITY_WARNING | SUGGESTION | WARNING`

Required: Yes

**issueCode**

The issue code provides an identifier of the issue associated with this finding.

Type: String

Required: Yes

**learnMoreLink**

A link to additional documentation about the type of finding.

Type: String

Required: Yes

**locations**

The list of locations in the policy document that are related to the finding. The issue
code provides a summary of an issue identified by the finding.

Type: Array of [Location](api-location.md) objects

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/validatepolicyfinding.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/validatepolicyfinding.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/validatepolicyfinding.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnusedPermissionsRecommendedStep

ValidationExceptionField
