# DescribeExportImageTasks

Describes the specified export image tasks or all of your export image tasks.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ExportImageTaskId.N**

The IDs of the export image tasks.

Type: Array of strings

Required: No

**Filter.N**

Filter tasks using the `task-state` filter and one of the following values: `active`,
`completed`, `deleting`, or `deleted`.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return in a single call.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 500.

Required: No

**NextToken**

A token that indicates the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**exportImageTaskSet**

Information about the export image tasks.

Type: Array of [ExportImageTask](api-exportimagetask.md) objects

**nextToken**

The token to use to get the next page of results. This value is `null` when there are no more results
to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeExportImageTasks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeExportImageTasks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeexportimagetasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeexportimagetasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeexportimagetasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeexportimagetasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeexportimagetasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeexportimagetasks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeExportImageTasks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeexportimagetasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeElasticGpus

DescribeExportTasks
