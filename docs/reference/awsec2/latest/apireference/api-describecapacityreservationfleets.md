# DescribeCapacityReservationFleets

Describes one or more Capacity Reservation Fleets.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationFleetId.N**

The IDs of the Capacity Reservation Fleets to describe.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `state` \- The state of the Fleet ( `submitted` \|
`modifying` \| `active` \|
`partially_fulfilled` \| `expiring` \|
`expired` \| `cancelling` \| `cancelled` \|
`failed`).

- `instance-match-criteria` \- The instance matching criteria for the
Fleet. Only `open` is supported.

- `tenancy` \- The tenancy of the Fleet ( `default` \|
`dedicated`).

- `allocation-strategy` \- The allocation strategy used by the Fleet.
Only `prioritized` is supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request. To get the next page of items, make another request with the token returned in the output. For more information,
see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token to use to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**capacityReservationFleetSet**

Information about the Capacity Reservation Fleets.

Type: Array of [CapacityReservationFleet](api-capacityreservationfleet.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describecapacityreservationfleets.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecapacityreservationfleets.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeCapacityReservationBillingRequests

DescribeCapacityReservations
