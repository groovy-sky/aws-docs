# ExportClientVpnClientConfiguration

Downloads the contents of the Client VPN endpoint configuration file for the specified Client VPN endpoint. The Client VPN endpoint configuration
file includes the Client VPN endpoint and certificate information clients need to establish a connection
with the Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**clientConfiguration**

The contents of the Client VPN endpoint configuration file.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example downloads a Client VPN endpoint configuration file.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ExportClientVpnClientConfiguration
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<ExportClientVpnClientConfigurationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>44e88bf8-d460-4c43-80b8-a27e4EXAMPLE</requestId>
    <clientConfiguration>client
dev tun
proto udp
remote cvpn-endpoint-00c5d11fc4EXAMPLE.clientvpn.us-east-1.amazonaws.com 443
remote-random-hostname
resolv-retry infinite
nobind
persist-key
persist-tun
remote-cert-tls server
cipher AES-256-CBC
verb 3
<ca>
-----BEGIN CERTIFICATE-----
EXAMPLECAgmgAwIBAgIJAOjnW3hL6o+7MA0GCSqGSIb3DQEBCwUAMBAxDEXAMPLE
-----END CERTIFICATE-----

</ca></clientConfiguration>
</ExportClientVpnClientConfigurationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ExportClientVpnClientConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ExportClientVpnClientConfiguration)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ExportClientVpnClientConfiguration)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/exportclientvpnclientconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportClientVpnClientCertificateRevocationList

ExportImage
