# AnalysisRule

Contains information about analysis rules for the analyzer. Analysis rules determine
which entities will generate findings based on the criteria you define when you create the
rule.

## Contents

**exclusions**

A list of rules for the analyzer containing criteria to exclude from analysis. Entities
that meet the rule criteria will not generate findings.

Type: Array of [AnalysisRuleCriteria](api-analysisrulecriteria.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/AnalysisRule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/AnalysisRule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/AnalysisRule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AclGrantee

AnalysisRuleCriteria
