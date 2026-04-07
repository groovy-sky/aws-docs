# CreateVpnConnectionRoute

Creates a static route associated with a VPN connection between an existing virtual
private gateway and a VPN customer gateway. The static route allows traffic to be routed
from the virtual private gateway to the VPN customer gateway.

For more information, see [AWS Site-to-Site VPN](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the _AWS Site-to-Site VPN_
_User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DestinationCidrBlock**

The CIDR block associated with the local subnet of the customer network.

Type: String

Required: Yes

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

This example creates a static route to the VPN connection for the VPN
connection with the ID `vpn-83ad48ea` to the destination CIDR block
`11.12.0.0/16`. Note that when using the Query API the "/" is
denoted as "%2F".

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnConnectionRoute
&DestinationCidrBlock=11.12.0.0%2F16
&VpnConnectionId=vpn-83ad48ea
&AUTHPARAMS
```

#### Sample Response

```

<CreateVpnConnectionRouteResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>4f35a1b2-c2c3-4093-b51f-abb9d7311990</requestId>
    <return>true</return>
</CreateVpnConnectionRouteResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateVpnConnectionRoute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateVpnConnectionRoute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateVpnConnection

CreateVpnGateway
