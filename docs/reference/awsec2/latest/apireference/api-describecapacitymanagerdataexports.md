# DescribeCapacityManagerDataExports

Describes one or more Capacity Manager data export configurations. Returns information about export settings, delivery status, and recent export activity.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityManagerDataExportId.N**

The IDs of the data export configurations to describe. If not specified, all export configurations are returned.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters to narrow the results. Supported filters include export status, creation date, and S3 bucket name.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return in a single call. If not specified, up to 1000 results are returned.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results. Use this value in a subsequent call to retrieve additional results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerDataExportSet**

Information about the data export configurations, including export settings, delivery status, and recent activity.

Type: Array of [CapacityManagerDataExportResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CapacityManagerDataExportResponse.html) objects

**nextToken**

The token to use to retrieve the next page of results. This value is null when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeCapacityManagerDataExports)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeCapacityManagerDataExports)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeCapacityBlockStatus

DescribeCapacityReservationBillingRequests
