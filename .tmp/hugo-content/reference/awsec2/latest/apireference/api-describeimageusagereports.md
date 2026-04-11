# DescribeImageUsageReports

Describes the configuration and status of image usage reports, filtered by report IDs or
image IDs.

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

- `creation-time` \- The time when the report was created, in the ISO 8601
format in the UTC time zone (YYYY-MM-DDThh:mm:ss.sssZ), for example,
`2025-11-29T11:04:43.305Z`. You can use a wildcard ( `*`), for
example, `2025-11-29T*`, which matches an entire day.

- `state` \- The state of the report ( `available` \|
`pending` \| `error`).

- `tag:<key>` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

Type: Array of [Filter](api-filter.md) objects

Required: No

**ImageId.N**

The IDs of the images for filtering the reports. If specified, only reports containing
these images are returned.

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

The IDs of the image usage reports.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

## Response Elements

The following elements are returned by the service.

**imageUsageReportSet**

The image usage reports.

Type: Array of [ImageUsageReport](api-imageusagereport.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeimageusagereports.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeimageusagereports.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeImageUsageReportEntries

DescribeImportImageTasks
