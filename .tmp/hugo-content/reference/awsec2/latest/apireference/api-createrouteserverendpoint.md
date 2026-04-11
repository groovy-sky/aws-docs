# CreateRouteServerEndpoint

Creates a new endpoint for a route server in a specified subnet.

A route server endpoint is an AWS-managed component inside a subnet that facilitates [BGP (Border Gateway Protocol)](https://en.wikipedia.org/wiki/Border_Gateway_Protocol) connections between your route server and your BGP peers.

For more information see [Dynamic routing in your VPC with VPC Route Server](../../../../services/vpc/latest/userguide/dynamic-routing-route-server.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier to ensure idempotency of the request.

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RouteServerId**

The ID of the route server for which to create an endpoint.

Type: String

Required: Yes

**SubnetId**

The ID of the subnet in which to create the route server endpoint.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the route server endpoint during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**routeServerEndpoint**

Information about the created route server endpoint.

Type: [RouteServerEndpoint](api-routeserverendpoint.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createrouteserverendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createrouteserverendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateRouteServer

CreateRouteServerPeer
