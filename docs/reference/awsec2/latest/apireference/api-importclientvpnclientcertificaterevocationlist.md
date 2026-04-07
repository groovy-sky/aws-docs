# ImportClientVpnClientCertificateRevocationList

Uploads a client certificate revocation list to the specified Client VPN endpoint. Uploading a client certificate revocation list overwrites the existing client certificate revocation list.

Uploading a client certificate revocation list resets existing client connections.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CertificateRevocationList**

The client certificate revocation list file. For more information, see [Generate a Client Certificate Revocation List](../../../../services/vpn/latest/clientvpn-admin/cvpn-working-certificates.md#cvpn-working-certificates-generate) in the
_AWS Client VPN Administrator Guide_.

Type: String

Required: Yes

**ClientVpnEndpointId**

The ID of the Client VPN endpoint to which the client certificate revocation list applies.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ImportClientVpnClientCertificateRevocationList)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ImportClientVpnClientCertificateRevocationList)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ImportClientVpnClientCertificateRevocationList)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/importclientvpnclientcertificaterevocationlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetVpnTunnelReplacementStatus

ImportImage
