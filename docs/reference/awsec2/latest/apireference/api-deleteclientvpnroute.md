# DeleteClientVpnRoute

Deletes a route from a Client VPN endpoint. You can only delete routes that you manually added using
the **CreateClientVpnRoute** action. You cannot delete routes that were
automatically added when associating a subnet. To remove routes that have been automatically added,
disassociate the target subnet from the Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientVpnEndpointId**

The ID of the Client VPN endpoint from which the route is to be deleted.

Type: String

Required: Yes

**DestinationCidrBlock**

The IPv4 address range, in CIDR notation, of the route to be deleted.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TargetVpcSubnetId**

The ID of the target subnet used by the route.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**status**

The current state of the route.

Type: [ClientVpnRouteStatus](api-clientvpnroutestatus.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes a route from a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteClientVpnRoute
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&DestinationCidrBlock=0.0.0.0/0
&TargetVpcSubnetId=subnet-057fa0918fEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DeleteClientVpnRouteResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>00d80748-708d-40f7-8635-f34acEXAMPLE</requestId>
    <status>
        <code>deleting</code>
    </status>
</DeleteClientVpnRouteResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteClientVpnRoute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteClientVpnRoute)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deleteclientvpnroute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deleteclientvpnroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deleteclientvpnroute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deleteclientvpnroute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deleteclientvpnroute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deleteclientvpnroute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteClientVpnRoute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deleteclientvpnroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteClientVpnEndpoint

DeleteCoipCidr
