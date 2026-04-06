# FindingAggregationAccountDetails

Contains information about the findings for an AWS account in an organization unused
access analyzer.

## Contents

**account**

The ID of the AWS account for which unused access finding details are provided.

Type: String

Required: No

**details**

Provides the number of active findings for each type of unused access for the specified
AWS account.

Type: String to integer map

Required: No

**numberOfActiveFindings**

The number of active unused access findings for the specified AWS account.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/FindingAggregationAccountDetails)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/FindingAggregationAccountDetails)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/FindingAggregationAccountDetails)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Finding

FindingDetails
