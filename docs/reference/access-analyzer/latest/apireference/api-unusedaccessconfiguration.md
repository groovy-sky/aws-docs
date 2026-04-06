# UnusedAccessConfiguration

Contains information about an unused access analyzer.

## Contents

**analysisRule**

Contains information about analysis rules for the analyzer. Analysis rules determine
which entities will generate findings based on the criteria you define when you create the
rule.

Type: [AnalysisRule](api-analysisrule.md) object

Required: No

**unusedAccessAge**

The specified access age in days for which to generate findings for unused access. For
example, if you specify 90 days, the analyzer will generate findings for IAM entities
within the accounts of the selected organization for any access that hasn't been used in 90
or more days since the analyzer's last scan. You can choose a value between 1 and 365
days.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/unusedaccessconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/unusedaccessconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/unusedaccessconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TrailProperties

UnusedAccessFindingsStatistics
