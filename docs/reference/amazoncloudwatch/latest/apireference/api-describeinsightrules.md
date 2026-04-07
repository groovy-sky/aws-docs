# DescribeInsightRules

Returns a list of all the Contributor Insights rules in your account.

For more information about Contributor Insights, see [Using Contributor\
Insights to Analyze High-Cardinality Data](../../../../services/amazoncloudwatch/latest/monitoring/contributorinsights.md).

## Request Parameters

**MaxResults**

The maximum number of results to return in one operation. If you omit this parameter,
the default of 500 is used.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 500.

Required: No

**NextToken**

Include this value, if it was returned by the previous operation, to get the next set
of rules.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**InsightRules**

The rules returned by the operation.

Type: Array of [InsightRule](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_InsightRule.html) objects

**NextToken**

If this parameter is present, it is a token that marks the start of the next batch of
returned results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidNextToken**

The next token specified is invalid.

**message**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/monitoring-2010-08-01/DescribeInsightRules)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/DescribeInsightRules)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeAnomalyDetectors

DisableAlarmActions
