# AnalysisRuleCriteria

The criteria for an analysis rule for an analyzer. The criteria determine which entities
will generate findings.

## Contents

**accountIds**

A list of AWS account IDs to apply to the analysis rule criteria. The accounts cannot
include the organization analyzer owner account. Account IDs can only be applied to the
analysis rule criteria for organization-level analyzers. The list cannot include more than
2,000 account IDs.

Type: Array of strings

Required: No

**resourceTags**

An array of key-value pairs to match for your resources. You can use the set of Unicode
letters, digits, whitespace, `_`, `.`, `/`,
`=`, `+`, and `-`.

For the tag key, you can specify a value that is 1 to 128 characters in length and
cannot be prefixed with `aws:`.

For the tag value, you can specify a value that is 0 to 256 characters in length. If the
specified tag value is 0 characters, the rule is applied to all principals with the
specified tag key.

Type: Array of string to string maps

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/analysisrulecriteria.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/analysisrulecriteria.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/analysisrulecriteria.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnalysisRule

AnalyzedResource
