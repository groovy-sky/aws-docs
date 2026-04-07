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

Type: Array of [RegionalSummary](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RegionalSummary.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AttributeSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AttributeSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AttributeSummary)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttributeBooleanValue

AttributeValue
