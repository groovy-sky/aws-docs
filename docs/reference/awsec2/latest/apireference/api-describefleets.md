# DescribeFleets

Describes the specified EC2 Fleet or all of your EC2 Fleets.

###### Important

If a fleet is of type `instant`, you must specify the fleet ID in the
request, otherwise the fleet does not appear in the response.

For more information, see [Describe your\
EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#monitor-ec2-fleet) in the _Amazon EC2 User Guide_.

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

Type: Array of [FleetData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FleetData.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeFleets)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeFleets)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeFleetInstances

DescribeFlowLogs
