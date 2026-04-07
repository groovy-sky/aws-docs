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

Type: [VpnConnection](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnConnection.html) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVpnConnectionOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ModifyVpnConnectionOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyVpnConnection

ModifyVpnTunnelCertificate
