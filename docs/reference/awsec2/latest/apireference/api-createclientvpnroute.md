# CreateClientVpnRoute

Adds a route to a network to a Client VPN endpoint. Each Client VPN endpoint has a route table that describes the
available destination network routes. Each route in the route table specifies the path for traﬃc to speciﬁc resources or networks.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**ClientVpnEndpointId**

The ID of the Client VPN endpoint to which to add the route.

Type: String

Required: Yes

**Description**

A brief description of the route.

Type: String

Required: No

**DestinationCidrBlock**

The IPv4 address range, in CIDR notation, of the route destination. For example:

- To add a route for Internet access, enter `0.0.0.0/0`

- To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range

- To add a route for an on-premises network, enter the AWS Site-to-Site VPN connection's IPv4 CIDR range

- To add a route for the local network, enter the client CIDR range

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TargetVpcSubnetId**

The ID of the subnet through which you want to route traffic. The specified subnet must be
an existing target network of the Client VPN endpoint.

Alternatively, if you're adding a route for the local network, specify `local`.

Type: String

Required: Yes

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

This example adds a route for Internet access to the Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateClientVpnRoute
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&DestinationCidrBlock=0.0.0.0/0
&TargetVpcSubnetId=subnet-057fa0918fEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<CreateClientVpnRouteResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>5b301186-e6d3-436b-87d6-7c400EXAMPLE</requestId>
    <status>
        <code>creating</code>
    </status>
</CreateClientVpnRouteResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateClientVpnRoute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateClientVpnRoute)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createclientvpnroute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createclientvpnroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createclientvpnroute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createclientvpnroute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createclientvpnroute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createclientvpnroute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateClientVpnRoute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createclientvpnroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateClientVpnEndpoint

CreateCoipCidr
