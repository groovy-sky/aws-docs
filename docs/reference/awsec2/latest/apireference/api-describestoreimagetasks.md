# DescribeStoreImageTasks

Describes the progress of the AMI store tasks. You can describe the store tasks for
specified AMIs. If you don't specify the AMIs, you get a paginated list of store tasks from
the last 31 days.

For each AMI task, the response indicates if the task is `InProgress`,
`Completed`, or `Failed`. For tasks `InProgress`, the
response shows the estimated progress as a percentage.

Tasks are listed in reverse chronological order. Currently, only tasks from the past 31
days can be viewed.

To use this API, you must have the required permissions. For more information, see [Permissions for storing and restoring AMIs using S3](../../../../services/ec2/latest/userguide/work-with-ami-store-restore.md#ami-s3-permissions) in the
_Amazon EC2 User Guide_.

For more information, see [Store and restore an AMI using\
S3](../../../../services/ec2/latest/userguide/ami-store-restore.md) in the _Amazon EC2 User Guide_.

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

- `task-state` \- Returns tasks in a certain state ( `InProgress` \|
`Completed` \| `Failed`)

- `bucket` \- Returns task information for tasks that targeted a specific
bucket. For the filter value, specify the bucket name.

###### Note

When you specify the `ImageIds` parameter, any filters that you specify are
ignored. To use the filters, you must remove the `ImageIds` parameter.

Type: Array of [Filter](api-filter.md) objects

Required: No

**ImageId.N**

The AMI IDs for which to show progress. Up to 20 AMI IDs can be included in a
request.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

You cannot specify this parameter and the `ImageIds` parameter in the same
call.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 200.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**storeImageTaskResultSet**

The information about the AMI store tasks.

Type: Array of [StoreImageTaskResult](api-storeimagetaskresult.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeStoreImageTasks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeStoreImageTasks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describestoreimagetasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describestoreimagetasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describestoreimagetasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describestoreimagetasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describestoreimagetasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describestoreimagetasks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeStoreImageTasks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describestoreimagetasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeStaleSecurityGroups

DescribeSubnets
