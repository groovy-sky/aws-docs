# DescribeTransitGatewayRouteTables

Describes one or more transit gateway route tables. By default, all transit gateway route tables are described.
Alternatively, you can filter the results.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters. The possible values are:

- `default-association-route-table` \- Indicates whether this is the default
association route table for the transit gateway ( `true` \| `false`).

- `default-propagation-route-table` \- Indicates whether this is the default
propagation route table for the transit gateway ( `true` \| `false`).

- `state` \- The state of the route table ( `available` \| `deleting` \| `deleted` \| `pending`).

- `transit-gateway-id` \- The ID of the transit gateway.

- `transit-gateway-route-table-id` \- The ID of the transit gateway route table.

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

**TransitGatewayRouteTableIds.N**

The IDs of the transit gateway route tables.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**transitGatewayRouteTables**

Information about the transit gateway route tables.

Type: Array of [TransitGatewayRouteTable](api-transitgatewayroutetable.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describetransitgatewayroutetables.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describetransitgatewayroutetables.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeTransitGatewayRouteTableAnnouncements

DescribeTransitGateways
