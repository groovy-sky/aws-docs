# AttributeSummary

A summary report for the attribute across all Regions.

## Contents

**attributeName**

The name of the attribute.

Type: String

Required: No

**mostFrequentValue**

The configuration value that is most frequently observed for the attribute.

Type: String

Required: No

**numberOfMatchedAccounts**

The number of accounts with the same configuration value for the attribute that is
most frequently observed.

Type: Integer

Required: No

**numberOfUnmatchedAccounts**

The number of accounts with a configuration value different from the most frequently
observed value for the attribute.

Type: Integer

Required: No

**RegionalSummarySet.N**

The summary report for each Region for the attribute.

Type: Array of [RegionalSummary](api-regionalsummary.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/attributesummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/attributesummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/attributesummary.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttributeBooleanValue

AttributeValue
