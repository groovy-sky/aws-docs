# GetVpnConnectionDeviceSampleConfiguration

Download an AWS-provided sample configuration file to be used with the customer
gateway device specified for your Site-to-Site VPN connection.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**InternetKeyExchangeVersion**

The IKE version to be used in the sample configuration file for your customer gateway
device. You can specify one of the following versions: `ikev1` or
`ikev2`.

Type: String

Required: No

**SampleType**

The type of sample configuration to generate. Valid values are "compatibility" (includes IKEv1) or "recommended" (throws UnsupportedOperationException for IKEv1).

Type: String

Required: No

**VpnConnectionDeviceTypeId**

Device identifier provided by the `GetVpnConnectionDeviceTypes` API.

Type: String

Required: Yes

**VpnConnectionId**

The `VpnConnectionId` specifies the Site-to-Site VPN connection used for the sample
configuration.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpnConnectionDeviceSampleConfiguration**

Sample configuration file for the specified customer gateway device.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getvpnconnectiondevicesampleconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetVpcResourcesBlockingEncryptionEnforcement

GetVpnConnectionDeviceTypes
