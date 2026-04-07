# ReplaceRoute

Replaces an existing route within a route table in a VPC.

You must specify either a destination CIDR block or a prefix list ID. You must also specify
exactly one of the resources from the parameter list, or reset the local route to its default
target.

For more information, see [Route tables](../../../../services/vpc/latest/userguide/vpc-route-tables.md) in the
_Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CarrierGatewayId**

\[IPv4 traffic only\] The ID of a carrier gateway.

Type: String

Required: No

**CoreNetworkArn**

The Amazon Resource Name (ARN) of the core network.

Type: String

Required: No

**DestinationCidrBlock**

The IPv4 CIDR address block used for the destination match. The value that you
provide must match the CIDR of an existing route in the table.

Type: String

Required: No

**DestinationIpv6CidrBlock**

The IPv6 CIDR address block used for the destination match. The value that you
provide must match the CIDR of an existing route in the table.

Type: String

Required: No

**DestinationPrefixListId**

The ID of the prefix list for the route.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EgressOnlyInternetGatewayId**

\[IPv6 traffic only\] The ID of an egress-only internet gateway.

Type: String

Required: No

**GatewayId**

The ID of an internet gateway or virtual private gateway.

Type: String

Required: No

**InstanceId**

The ID of a NAT instance in your VPC.

Type: String

Required: No

**LocalGatewayId**

The ID of the local gateway.

Type: String

Required: No

**LocalTarget**

Specifies whether to reset the local route to its default target ( `local`).

Type: Boolean

Required: No

**NatGatewayId**

\[IPv4 traffic only\] The ID of a NAT gateway.

Type: String

Required: No

**NetworkInterfaceId**

The ID of a network interface.

Type: String

Required: No

**OdbNetworkArn**

The Amazon Resource Name (ARN) of the ODB network.

Type: String

Required: No

**RouteTableId**

The ID of the route table.

Type: String

Required: Yes

**TransitGatewayId**

The ID of a transit gateway.

Type: String

Required: No

**VpcEndpointId**

The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.

Type: String

Required: No

**VpcPeeringConnectionId**

The ID of a VPC peering connection.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example replaces a route in the specified route table. The new route
matches the IPv4 CIDR `10.0.0.0/8` and sends the traffic to the
virtual private gateway with the ID `vgw-123456abcde123456`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ReplaceRoute
&RouteTableId=rtb-11223344556677889
&DestinationCidrBlock=10.0.0.0/8
&GatewayId=vgw-123456abcde123456
&AUTHPARAMS
```

#### Sample Response

```

<ReplaceRouteResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</ReplaceRouteResponse>
```

### Example 2

This examples resets the target for the default local route.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ReplaceRoute
&RouteTableId=rtb-11223344556677889
&DestinationCidrBlock=10.0.0.0/16
&LocalTarget=true
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ReplaceRoute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ReplaceRoute)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/replaceroute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/replaceroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/replaceroute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/replaceroute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/replaceroute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/replaceroute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ReplaceRoute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/replaceroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ReplaceNetworkAclEntry

ReplaceRouteTableAssociation
