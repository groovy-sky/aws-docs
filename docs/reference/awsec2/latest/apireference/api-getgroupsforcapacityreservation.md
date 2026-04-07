# GetGroupsForCapacityReservation

Lists the resource groups to which a Capacity Reservation has been added.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId**

The ID of the Capacity Reservation. If you specify a Capacity Reservation that is
shared with you, the operation returns only Capacity Reservation groups that you
own.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

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

**capacityReservationGroupSet**

Information about the resource groups to which the Capacity Reservation has been
added.

Type: Array of [CapacityReservationGroup](api-capacityreservationgroup.md) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetGroupsForCapacityReservation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetGroupsForCapacityReservation)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getgroupsforcapacityreservation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getgroupsforcapacityreservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getgroupsforcapacityreservation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getgroupsforcapacityreservation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getgroupsforcapacityreservation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getgroupsforcapacityreservation.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetGroupsForCapacityReservation)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getgroupsforcapacityreservation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetFlowLogsIntegrationTemplate

GetHostReservationPurchasePreview
