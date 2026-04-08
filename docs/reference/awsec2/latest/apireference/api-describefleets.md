# DescribeFleets

Describes the specified EC2 Fleet or all of your EC2 Fleets.

###### Important

If a fleet is of type `instant`, you must specify the fleet ID in the
request, otherwise the fleet does not appear in the response.

For more information, see [Describe your\
EC2 Fleet](../../../../services/ec2/latest/userguide/manage-ec2-fleet.md#monitor-ec2-fleet) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `activity-status` \- The progress of the EC2 Fleet ( `error` \|
`pending-fulfillment` \| `pending-termination` \|
`fulfilled`).

- `excess-capacity-termination-policy` \- Indicates whether to terminate
running instances if the target capacity is decreased below the current EC2 Fleet size
( `true` \| `false`).

- `fleet-state` \- The state of the EC2 Fleet ( `submitted` \|
`active` \| `deleted` \| `failed` \|
`deleted-running` \| `deleted-terminating` \|
`modifying`).

- `replace-unhealthy-instances` \- Indicates whether EC2 Fleet should replace
unhealthy instances ( `true` \| `false`).

- `type` \- The type of request ( `instant` \|
`request` \| `maintain`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**FleetId.N**

The IDs of the EC2 Fleets.

###### Note

If a fleet is of type `instant`, you must specify the fleet ID, otherwise
it does not appear in the response.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**fleetSet**

Information about the EC2 Fleets.

Type: Array of [FleetData](api-fleetdata.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describefleets.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describefleets.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describefleets.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describefleets.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describefleets.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describefleets.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describefleets.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describefleets.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describefleets.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describefleets.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeFleetInstances

DescribeFlowLogs
