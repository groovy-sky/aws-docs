# DeleteClientVpnEndpoint

Deletes the specified Client VPN endpoint. You must disassociate all target networks before you
can delete a Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientVpnEndpointId**

The ID of the Client VPN to be deleted.

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

**status**

The current state of the Client VPN endpoint.

Type: [ClientVpnEndpointStatus](api-clientvpnendpointstatus.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example applies a security group to a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteClientVpnEndpoint
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DeleteClientVpnEndpointResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>f1e0fdfc-96a4-4d7d-bc78-22eb0EXAMPLE</requestId>
    <status>
        <code>deleting</code>
    </status>
</DeleteClientVpnEndpointResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deleteclientvpnendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deleteclientvpnendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteCarrierGateway

DeleteClientVpnRoute
