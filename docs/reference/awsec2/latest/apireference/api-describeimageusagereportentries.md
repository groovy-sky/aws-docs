# DescribeImageUsageReportEntries

Describes the entries in image usage reports, showing how your images are used across
other AWS accounts.

For more information, see [View your AMI usage](../../../../services/ec2/latest/userguide/your-ec2-ami-usage.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `account-id` \- A 12-digit AWS account ID.

- `creation-time` \- The time when the report was created, in the ISO 8601
format in the UTC time zone (YYYY-MM-DDThh:mm:ss.sssZ), for example,
`2025-11-29T11:04:43.305Z`. You can use a wildcard ( `*`), for
example, `2025-11-29T*`, which matches an entire day.

- `resource-type` \- The resource type ( `ec2:Instance` \|
`ec2:LaunchTemplate`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**ImageId.N**

The IDs of the images for filtering the report entries. If specified, only report entries
containing these images are returned.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**ReportId.N**

The IDs of the usage reports.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

## Response Elements

The following elements are returned by the service.

**imageUsageReportEntrySet**

The content of the usage reports.

Type: Array of [ImageUsageReportEntry](api-imageusagereportentry.md) objects

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeimageusagereportentries.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeimageusagereportentries.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeImages

DescribeImageUsageReports
