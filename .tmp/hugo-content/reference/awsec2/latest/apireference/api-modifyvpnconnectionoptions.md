# ModifyVpnConnectionOptions

Modifies the connection options for your Site-to-Site VPN connection.

When you modify the VPN connection options, the VPN endpoint IP addresses on the
AWS side do not change, and the tunnel options do not change. Your
VPN connection will be temporarily unavailable for a brief period while the VPN
connection is updated.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**LocalIpv4NetworkCidr**

The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.

Default: `0.0.0.0/0`

Type: String

Required: No

**LocalIpv6NetworkCidr**

The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.

Default: `::/0`

Type: String

Required: No

**RemoteIpv4NetworkCidr**

The IPv4 CIDR on the AWS side of the VPN connection.

Default: `0.0.0.0/0`

Type: String

Required: No

**RemoteIpv6NetworkCidr**

The IPv6 CIDR on the AWS side of the VPN connection.

Default: `::/0`

Type: String

Required: No

**VpnConnectionId**

The ID of the Site-to-Site VPN connection.

Type: String

Required: Yes

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyvpnconnectionoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpnconnectionoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpnConnection

ModifyVpnTunnelCertificate
