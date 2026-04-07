# DeleteVpnConnection

Deletes the specified VPN connection.

If you're deleting the VPC and its associated components, we recommend that you detach
the virtual private gateway from the VPC and delete the VPC before deleting the VPN
connection. If you believe that the tunnel credentials for your VPN connection have been
compromised, you can delete the VPN connection and create a new one that has new keys,
without needing to delete the VPC or virtual private gateway. If you create a new VPN
connection, you must reconfigure the customer gateway device using the new configuration
information returned with the new VPN connection ID.

For certificate-based authentication, delete all AWS Certificate Manager (ACM) private
certificates used for the AWS-side tunnel endpoints for the VPN
connection before deleting the VPN connection.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**VpnConnectionId**

The ID of the VPN connection.

Type: String

Required: Yes

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

### Example

This example deletes the specified VPN connection.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteVpnConnection
&vpnConnectionId=vpn-44a8938f
&AUTHPARAMS
```

#### Sample Response

```

<DeleteVpnConnectionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <return>true</return>
</DeleteVpnConnectionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteVpnConnection)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteVpnConnection)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletevpnconnection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletevpnconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletevpnconnection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletevpnconnection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletevpnconnection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletevpnconnection.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteVpnConnection)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletevpnconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteVpnConcentrator

DeleteVpnConnectionRoute
