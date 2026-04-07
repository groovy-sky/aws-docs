# ModifyVpnConnection

Modifies the customer gateway or the target gateway of an AWS Site-to-Site VPN connection. To
modify the target gateway, the following migration options are available:

- An existing virtual private gateway to a new virtual private gateway

- An existing virtual private gateway to a transit gateway

- An existing transit gateway to a new transit gateway

- An existing transit gateway to a virtual private gateway

Before you perform the migration to the new gateway, you must configure the new
gateway. Use [CreateVpnGateway](api-createvpngateway.md) to create a virtual private gateway, or
[CreateTransitGateway](api-createtransitgateway.md) to create a transit gateway.

This step is required when you migrate from a virtual private gateway with static
routes to a transit gateway.

You must delete the static routes before you migrate to the new gateway.

Keep a copy of the static route before you delete it. You will need to add back these
routes to the transit gateway after the VPN connection migration is complete.

After you migrate to the new gateway, you might need to modify your VPC route table.
Use [CreateRoute](api-createroute.md) and [DeleteRoute](api-deleteroute.md) to make the changes
described in [Update VPC route\
tables](../../../../services/vpn/latest/s2svpn/modify-vpn-target.md#step-update-routing) in the _AWS Site-to-Site VPN User Guide_.

When the new gateway is a transit gateway, modify the transit gateway route table to
allow traffic between the VPC and the AWS Site-to-Site VPN connection.
Use [CreateTransitGatewayRoute](api-createtransitgatewayroute.md) to add the routes.

If you deleted VPN static routes, you must add the static routes to the transit
gateway route table.

After you perform this operation, the VPN endpoint's IP addresses on the AWS side and the tunnel options remain intact. Your AWS Site-to-Site VPN connection will
be temporarily unavailable for a brief period while we provision the new
endpoints.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CustomerGatewayId**

The ID of the customer gateway at your end of the VPN connection.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**TransitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**VpnConnectionId**

The ID of the VPN connection.

Type: String

Required: Yes

**VpnGatewayId**

The ID of the virtual private gateway at the AWS side of the VPN
connection.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpnConnection**

Information about the VPN connection.

Type: [VpnConnection](api-vpnconnection.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVpnConnection)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVpnConnection)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpnconnection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpnconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpnconnection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpnconnection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpnconnection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpnconnection.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVpnConnection)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpnconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpcTenancy

ModifyVpnConnectionOptions
