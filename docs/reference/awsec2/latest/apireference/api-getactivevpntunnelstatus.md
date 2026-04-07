# GetActiveVpnTunnelStatus

Returns the currently negotiated security parameters for an active VPN tunnel, including IKE version, DH groups, encryption algorithms, and integrity algorithms.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request.

Type: Boolean

Required: No

**VpnConnectionId**

The ID of the VPN connection for which to retrieve the active tunnel status.

Type: String

Required: Yes

**VpnTunnelOutsideIpAddress**

The external IP address of the VPN tunnel for which to retrieve the active status.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**activeVpnTunnelStatus**

Information about the current security configuration of the VPN tunnel.

Type: [ActiveVpnTunnelStatus](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ActiveVpnTunnelStatus.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetActiveVpnTunnelStatus)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetActiveVpnTunnelStatus)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExportVerifiedAccessInstanceClientConfiguration

GetAllowedImagesSettings
