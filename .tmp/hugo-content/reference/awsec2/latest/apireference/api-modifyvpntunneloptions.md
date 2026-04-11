# ModifyVpnTunnelOptions

Modifies the options for a VPN tunnel in an AWS Site-to-Site VPN connection. You can modify
multiple options for a tunnel in a single request, but you can only modify one tunnel at
a time. For more information, see [Site-to-Site VPN tunnel options for your Site-to-Site VPN\
connection](../../../../services/vpn/latest/s2svpn/vpntunnels.md) in the _AWS Site-to-Site VPN User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**PreSharedKeyStorage**

Specifies the storage mode for the pre-shared key (PSK). Valid values are `Standard` (stored in Site-to-Site VPN service) or `SecretsManager` (stored in AWS Secrets Manager).

Type: String

Required: No

**SkipTunnelReplacement**

Choose whether or not to trigger immediate tunnel replacement. This is only applicable when turning on or off `EnableTunnelLifecycleControl`.

Valid values: `True` \| `False`

Type: Boolean

Required: No

**TunnelOptions**

The tunnel options to modify.

Type: [ModifyVpnTunnelOptionsSpecification](api-modifyvpntunneloptionsspecification.md) object

Required: Yes

**VpnConnectionId**

The ID of the AWS Site-to-Site VPN connection.

Type: String

Required: Yes

**VpnTunnelOutsideIpAddress**

The external IP address of the VPN tunnel.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyvpntunneloptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpntunneloptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpnTunnelCertificate

MonitorInstances
