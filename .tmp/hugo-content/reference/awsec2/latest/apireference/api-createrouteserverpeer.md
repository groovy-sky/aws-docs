# CreateRouteServerPeer

Creates a new BGP peer for a specified route server endpoint.

A route server peer is a session between a route server endpoint and the device deployed in AWS (such as a firewall appliance or other network security function running on an EC2 instance). The device must meet these requirements:

- Have an elastic network interface in the VPC

- Support BGP (Border Gateway Protocol)

- Can initiate BGP sessions

For more information see [Dynamic routing in your VPC with VPC Route Server](../../../../services/vpc/latest/userguide/dynamic-routing-route-server.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**BgpOptions**

The BGP options for the peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.

Type: [RouteServerBgpOptionsRequest](api-routeserverbgpoptionsrequest.md) object

Required: Yes

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PeerAddress**

The IPv4 address of the peer device.

Type: String

Required: Yes

**RouteServerEndpointId**

The ID of the route server endpoint for which to create a peer.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the route server peer during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**routeServerPeer**

Information about the created route server peer.

Type: [RouteServerPeer](api-routeserverpeer.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createrouteserverpeer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createrouteserverpeer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateRouteServerEndpoint

CreateRouteTable
