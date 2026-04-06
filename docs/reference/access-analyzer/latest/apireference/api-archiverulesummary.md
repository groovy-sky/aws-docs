# ArchiveRuleSummary

Contains information about an archive rule. Archive rules automatically archive new
findings that meet the criteria you define when you create the rule.

## Contents

**createdAt**

The time at which the archive rule was created.

Type: Timestamp

Required: Yes

**filter**

A filter used to define the archive rule.

Type: String to [Criterion](api-criterion.md) object map

Required: Yes

**ruleName**

The name of the archive rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

**updatedAt**

The time at which the archive rule was last updated.

Type: Timestamp

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/ArchiveRuleSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/ArchiveRuleSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/ArchiveRuleSummary)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalyzerSummary

CloudTrailDetails
