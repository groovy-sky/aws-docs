# DescribeCapacityBlockExtensionHistory

Describes the events for the specified Capacity Block extension during the specified
time.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId.N**

The IDs of Capacity Block reservations that you want to display the history
for.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters

- `availability-zone` \- The Availability Zone of the
extension.

- `availability-zone-id` \- The Availability Zone ID of the
extension.

- `capacity-block-extension-offering-id` \- The ID of the extension
offering.

- `capacity-block-extension-status` \- The status of the extension
( `payment-pending` \| `payment-failed` \|
`payment-succeeded`).

- `capacity-reservation-id` \- The reservation ID of the
extension.

- `instance-type` \- The instance type of the extension.

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

**capacityBlockExtensionSet**

Describes one or more of your Capacity Block extensions. The results describe only the
Capacity Block extensions in the AWS Region that you're currently using.

Type: Array of [CapacityBlockExtension](api-capacityblockextension.md) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeCapacityBlockExtensionHistory)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeCapacityBlockExtensionHistory)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecapacityblockextensionhistory.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecapacityblockextensionhistory.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecapacityblockextensionhistory.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecapacityblockextensionhistory.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecapacityblockextensionhistory.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecapacityblockextensionhistory.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeCapacityBlockExtensionHistory)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecapacityblockextensionhistory.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeByoipCidrs

DescribeCapacityBlockExtensionOfferings
