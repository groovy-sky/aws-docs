# DescribeCapacityBlocks

Describes details about Capacity Blocks in the AWS Region that you're currently using.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityBlockId.N**

The IDs of the Capacity Blocks.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `capacity-block-id` \- The ID of the Capacity Block.

- `ultraserver-type` \- The Capacity Block type. The type can be
`instances` or `ultraservers`.

- `availability-zone` \- The Availability Zone of the Capacity
Block.

- `start-date` \- The date and time at which the Capacity Block was
started.

- `end-date` \- The date and time at which the Capacity Block expires.
When a Capacity Block expires, all instances in the Capacity Block are
terminated.

- `create-date` \- The date and time at which the Capacity Block was
created.

- `state` \- The state of the Capacity Block ( `active` \|
`expired` \| `unavailable` \| `cancelled` \|
`failed` \| `scheduled` \| `payment-pending` \|
`payment-failed`).

- `tags` \- The tags assigned to the Capacity Block.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request. To get the next page of items, make another request with the token returned in the output. For more information,
see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token to use to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**capacityBlockSet**

The Capacity Blocks.

Type: Array of [CapacityBlock](api-capacityblock.md) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeCapacityBlocks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeCapacityBlocks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecapacityblocks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecapacityblocks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecapacityblocks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecapacityblocks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecapacityblocks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecapacityblocks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeCapacityBlocks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecapacityblocks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeCapacityBlockOfferings

DescribeCapacityBlockStatus
