# SearchLocalGatewayRoutes

Searches for routes in the specified local gateway route table.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `prefix-list-id` \- The ID of the prefix list.

- `route-search.exact-match` \- The exact match of the specified filter.

- `route-search.longest-prefix-match` \- The longest prefix that matches the route.

- `route-search.subnet-of-match` \- The routes with a subnet that match the specified CIDR filter.

- `route-search.supernet-of-match` \- The routes with a CIDR that encompass the CIDR filter.
For example, if you have 10.0.1.0/29 and 10.0.1.0/31 routes in your route table and you specify `supernet-of-match`
as 10.0.1.0/30, then the result returns 10.0.1.0/29.

- `state` \- The state of the route.

- `type` \- The route type.

Type: Array of [Filter](api-filter.md) objects

Required: No

**LocalGatewayRouteTableId**

The ID of the local gateway route table.

Type: String

Required: Yes

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**routeSet**

Information about the routes.

Type: Array of [LocalGatewayRoute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LocalGatewayRoute.html) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/SearchLocalGatewayRoutes)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/SearchLocalGatewayRoutes)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RunScheduledInstances

SearchTransitGatewayMulticastGroups
