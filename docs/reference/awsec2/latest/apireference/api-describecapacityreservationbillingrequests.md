# DescribeCapacityReservationBillingRequests

Describes a request to assign the billing of the unused capacity of a Capacity
Reservation. For more information, see [Billing assignment for shared\
Amazon EC2 Capacity Reservations](../../../../services/ec2/latest/userguide/transfer-billing.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId.N**

The ID of the Capacity Reservation.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `status` \- The state of the request ( `pending` \|
`accepted` \| `rejected` \| `cancelled` \|
`revoked` \| `expired`).

- `requested-by` \- The account ID of the Capacity Reservation owner
that initiated the request. Not supported if you specify
`requested-by` for **Role**.

- `unused-reservation-billing-owner` \- The ID of the consumer account
to which the request was sent. Not supported if you specify
`unused-reservation-billing-owner` for **Role**.

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

**Role**

Specify one of the following:

- `odcr-owner` \- If you are the Capacity Reservation owner, specify
this value to view requests that you have initiated. Not supported with the
`requested-by` filter.

- `unused-reservation-billing-owner` \- If you are the consumer
account, specify this value to view requests that have been sent to you. Not
supported with the `unused-reservation-billing-owner` filter.

Type: String

Valid Values: `odcr-owner | unused-reservation-billing-owner`

Required: Yes

## Response Elements

The following elements are returned by the service.

**capacityReservationBillingRequestSet**

Information about the request.

Type: Array of [CapacityReservationBillingRequest](api-capacityreservationbillingrequest.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecapacityreservationbillingrequests.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeCapacityManagerDataExports

DescribeCapacityReservationFleets
