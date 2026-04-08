# CreateVpnConnection

Creates a VPN connection between an existing virtual private gateway or transit
gateway and a customer gateway. The supported connection type is
`ipsec.1`.

The response includes information that you need to give to your network administrator
to configure your customer gateway.

###### Important

We strongly recommend that you use HTTPS when calling this operation because the
response contains sensitive cryptographic information for configuring your customer
gateway device.

If you decide to shut down your VPN connection for any reason and later create a new
VPN connection, you must reconfigure your customer gateway with the new information
returned from this call.

This is an idempotent operation. If you perform the operation more than once, Amazon
EC2 doesn't return an error.

For more information, see [AWS Site-to-Site VPN](../../../../services/vpn/latest/s2svpn/vpc-vpn.md) in the _AWS Site-to-Site VPN_
_User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CustomerGatewayId**

The ID of the customer gateway.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Options**

The options for the VPN connection.

Type: [VpnConnectionOptionsSpecification](api-vpnconnectionoptionsspecification.md) object

Required: No

**PreSharedKeyStorage**

Specifies the storage mode for the pre-shared key (PSK). Valid values are `Standard`" (stored in the Site-to-Site VPN service) or `SecretsManager` (stored in AWS Secrets Manager).

Type: String

Required: No

**TagSpecification.N**

The tags to apply to the VPN connection.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TransitGatewayId**

The ID of the transit gateway. If you specify a transit gateway, you cannot specify a virtual private
gateway.

Type: String

Required: No

**Type**

The type of VPN connection ( `ipsec.1`).

Type: String

Required: Yes

**VpnConcentratorId**

The ID of the VPN concentrator to associate with the VPN connection.

Type: String

Required: No

**VpnGatewayId**

The ID of the virtual private gateway. If you specify a virtual private gateway, you
cannot specify a transit gateway.

Type: String

Required: No

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

## Examples

### Example 1

This example creates a VPN connection between the specified virtual private
gateway and the specified customer gateway. The response includes configuration
information for configuring the customer gateway device. Because it's a long set
of information, we haven't included the complete response here. To see an
example of the configuration information, see the [Your customer gateway\
device](../../../../services/vpn/latest/s2svpn/your-cgw.md).

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnConnection
&Type=ipsec.1
&CustomerGatewayId=cgw-112233445566aabbc
&VpnGatewayId=vgw-aabbccddee1234567
&AUTHPARAMS
```

#### Sample Response

```

<CreateVpnConnectionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>22896b9b-e2fe-4574-9a20-example</requestId>
    <vpnConnection>
        <vpnConnectionId>vpn-01234abcabc123456</vpnConnectionId>
        <state>pending</state>
        <customerGatewayConfiguration>...Customer gateway configuration data in escaped XML format...</customerGatewayConfiguration>
        <customerGatewayId>cgw-112233445566aabbc</customerGatewayId>
        <vpnGatewayId>vgw-aabbccddee1234567</vpnGatewayId>
        <tagSet/>
        <options>
            <enableAcceleration>false</enableAcceleration>
            <staticRoutesOnly>false</staticRoutesOnly>
            <localIpv4NetworkCidr>0.0.0.0/0</localIpv4NetworkCidr>
            <remoteIpv4NetworkCidr>0.0.0.0/0</remoteIpv4NetworkCidr>
            <tunnelInsideIpVersion>ipv4</tunnelInsideIpVersion>
            <tunnelOptionSet>
                <item/>
                <item/>
            </tunnelOptionSet>
        </options>
        <routes/>
        <category>VPN</category>
    </vpnConnection>
</CreateVpnConnectionResponse>
```

### Example 2

This example creates a VPN connection with the static routes option between
the virtual private gateway with the ID `vgw-8db04f81`, and the
customer gateway with the ID `cgw-b4dc3961`, for a device that does
not support the Border Gateway Protocol (BGP).

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnConnection
&Type=ipsec.1
&CustomerGatewayId=cgw-b4dc3961
&VpnGatewayId=vgw-8db04f81
&Options.StaticRoutesOnly=true
&AUTHPARAMS
```

### Example 3

This example creates a VPN connection between the virtual private gateway with
the ID `vgw-8db04f81` and the customer gateway with the ID
`cgw-b4dc3961` and specifies the inside IP address CIDR block and
a custom pre-shared key for each tunnel.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnConnection
&Type=ipsec.1
&CustomerGatewayId=cgw-b4dc3961
&VpnGatewayId=vgw-8db04f81
&Options.TunnelOptions.1.PreSharedKey=wMp_IGfO1d1o9AT4lF6tJLFN4EXAMPLE
&Options.TunnelOptions.1.TunnelInsideCidr=169.254.44.110/30
&Options.TunnelOptions.2.PreSharedKey=HAM8lcnFYEvfl6gUrOatJLFN4EXAMPLE
&Options.TunnelOptions.2.TunnelInsideCidr=169.254.44.240/30
&AUTHPARAMS
```

### Example 4

This example creates a VPN connection between the specified transit gateway
and the specified customer gateway. The VPN connection processes IPv6 traffic
inside the tunnels, and the tunnel options for both tunnels specify that AWS must initiate the IKE negotiation. A tag with a key of
`Location` and a value of `NewYorkVPN` is applied to
the VPN connection.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateVpnConnection
&Type=ipsec.1
&CustomerGatewayId=cgw-112233445566aabbc
&TransitGatewayId=tgw-0123f96e7b3f5babc
&Options.StaticRoutesOnly=false
&Options.TunnelInsideIpVersion=ipv6
&Options.TunnelOptions.1.StartupAction=start
&Options.TunnelOptions.2.StartupAction=start
&TagSpecification.1.ResourceType=vpn-connection
&TagSpecification.1.Tag.1.Key=Location
&TagSpecification.1.Tag.1.Value=NewYorkVPN
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createvpnconnection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createvpnconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateVpnConcentrator

CreateVpnConnectionRoute
