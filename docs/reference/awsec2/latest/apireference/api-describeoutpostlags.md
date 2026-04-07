# DescribeOutpostLags

Describes the Outposts link aggregation groups (LAGs).

###### Note

LAGs are only available for second-generation Outposts racks at this time.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters to use for narrowing down the request. The following filters are
supported:

- `service-link-virtual-interface-id` \- The ID of the service link virtual interface.

- `service-link-virtual-interface-arn` \- The ARN of the service link virtual interface.

- `outpost-id` \- The Outpost ID.

- `outpost-arn` \- The Outpost ARN.

- `owner-id` \- The ID of the AWS account that owns the service link virtual interface.

- `vlan` \- The ID of the address pool.

- `local-address` \- The local address.

- `peer-address` \- The peer address.

- `peer-bgp-asn` \- The peer BGP ASN.

- `outpost-lag-id` \- The Outpost LAG ID.

- `configuration-state` \- The configuration state of the service link virtual interface.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**OutpostLagId.N**

The IDs of the Outpost LAGs.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**outpostLagSet**

The Outpost LAGs.

Type: Array of [OutpostLag](api-outpostlag.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeOutpostLags)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeOutpostLags)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeoutpostlags.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeoutpostlags.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeoutpostlags.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeoutpostlags.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeoutpostlags.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeoutpostlags.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeOutpostLags)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeoutpostlags.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeNetworkInterfaces

DescribePlacementGroups
