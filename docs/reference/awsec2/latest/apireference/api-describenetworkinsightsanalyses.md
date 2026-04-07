# DescribeNetworkInsightsAnalyses

Describes one or more of your network insights analyses.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AnalysisEndTime**

The time when the network insights analyses ended.

Type: Timestamp

Required: No

**AnalysisStartTime**

The time when the network insights analyses started.

Type: Timestamp

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters. The following are the possible values:

- path-found - A Boolean value that indicates whether a feasible path is found.

- status - The status of the analysis (running \| succeeded \| failed).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NetworkInsightsAnalysisId.N**

The ID of the network insights analyses. You must specify either analysis IDs or a path ID.

Type: Array of strings

Required: No

**NetworkInsightsPathId**

The ID of the path. You must specify either a path ID or analysis IDs.

Type: String

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**networkInsightsAnalysisSet**

Information about the network insights analyses.

Type: Array of [NetworkInsightsAnalysis](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInsightsAnalysis.html) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeNetworkInsightsAnalyses)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeNetworkInsightsAccessScopes

DescribeNetworkInsightsPaths
