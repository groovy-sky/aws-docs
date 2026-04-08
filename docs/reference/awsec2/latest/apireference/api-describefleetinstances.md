# DescribeFleetInstances

Describes the running instances for the specified EC2 Fleet.

###### Note

Currently, `DescribeFleetInstances` does not support fleets of type
`instant`. Instead, use `DescribeFleets`, specifying the
`instant` fleet ID in the request.

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

- `instance-type` \- The instance type.

Type: Array of [Filter](api-filter.md) objects

Required: No

**FleetId**

The ID of the EC2 Fleet.

Type: String

Required: Yes

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

**activeInstanceSet**

The running instances. This list is refreshed periodically and might be out of
date.

Type: Array of [ActiveInstance](api-activeinstance.md) objects

**fleetId**

The ID of the EC2 Fleet.

Type: String

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describefleetinstances.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describefleetinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeFleetHistory

DescribeFleets
