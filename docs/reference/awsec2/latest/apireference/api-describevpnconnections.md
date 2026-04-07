# DescribeVpnConnections

Describes one or more of your VPN connections.

For more information, see [AWS Site-to-Site VPN](../../../../services/vpn/latest/s2svpn/vpc-vpn.md) in the _AWS Site-to-Site VPN_
_User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `customer-gateway-configuration` \- The configuration information
for the customer gateway.

- `customer-gateway-id` \- The ID of a customer gateway associated
with the VPN connection.

- `state` \- The state of the VPN connection ( `pending` \|
`available` \| `deleting` \|
`deleted`).

- `option.static-routes-only` \- Indicates whether the connection has
static routes only. Used for devices that do not support Border Gateway Protocol
(BGP).

- `route.destination-cidr-block` \- The destination CIDR block. This
corresponds to the subnet used in a customer data center.

- `bgp-asn` \- The BGP Autonomous System Number (ASN) associated with
a BGP device.

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `type` \- The type of VPN connection. Currently the only supported
type is `ipsec.1`.

- `vpn-connection-id` \- The ID of the VPN connection.

- `vpn-gateway-id` \- The ID of a virtual private gateway associated
with the VPN connection.

- `transit-gateway-id` \- The ID of a transit gateway associated with
the VPN connection.

Type: Array of [Filter](api-filter.md) objects

Required: No

**VpnConnectionId.N**

One or more VPN connection IDs.

Default: Describes your VPN connections.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpnConnectionSet**

Information about one or more VPN connections.

Type: Array of [VpnConnection](api-vpnconnection.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the specified VPN connection. The response includes the
customer gateway device configuration information. Because it's a long set of
information, we haven't displayed it here. To see an example of the
configuration information, see the [Your customer gateway\
device](../../../../services/vpn/latest/s2svpn/your-cgw.md).

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpnConnections
&VpnConnectionId.1=vpn-1122334455aabbccd
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVpnConnectionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>6791f4b8-5717-4272-aed2-faa18example</requestId>
    <vpnConnectionSet>
        <item>
            <vpnConnectionId>vpn-1122334455aabbccd</vpnConnectionId>
            <state>available</state>
            <customerGatewayConfiguration>..Customer gateway configuration data in escaped XML format...</customerGatewayConfiguration>
            <type>ipsec.1</type>
            <customerGatewayId>cgw-01234567abcde1234</customerGatewayId>
            <tagSet>
                <item>
                    <key>Name</key>
                    <value>CanadaVPN</value>
                </item>
            </tagSet>
            <vgwTelemetry>
                <item>
                    <outsideIpAddress>203.0.113.3</outsideIpAddress>
                    <status>DOWN</status>
                    <lastStatusChange>2020-07-29T10:35:11.000Z</lastStatusChange>
                    <statusMessage></statusMessage>
                    <acceptedRouteCount>0</acceptedRouteCount>
                </item>
                <item>
                    <outsideIpAddress>203.0.113.5</outsideIpAddress>
                    <status>UP</status>
                    <lastStatusChange>2020-09-02T09:09:33.000Z</lastStatusChange>
                    <statusMessage></statusMessage>
                    <acceptedRouteCount>0</acceptedRouteCount>
                </item>
            </vgwTelemetry>
            <transitGatewayId>tgw-00112233445566aab</transitGatewayId>
            <options>
                <enableAcceleration>false</enableAcceleration>
                <staticRoutesOnly>true</staticRoutesOnly>
                <localIpv4NetworkCidr>0.0.0.0/0</localIpv4NetworkCidr>
                <remoteIpv4NetworkCidr>0.0.0.0/0</remoteIpv4NetworkCidr>
                <tunnelInsideIpVersion>ipv4</tunnelInsideIpVersion>
            </options>
            <routes/>
            <category>VPN</category>
        </item>
    </vpnConnectionSet>
</DescribeVpnConnectionsResponse>
```

### Example 2

This example describes any VPN connection you own that is associated with the
customer gateway with ID cgw-b4dc3961, and whose state is either
`pending` or `available`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpnConnections
&Filter.1.Name=customer-gateway-id
&Filter.1.Value.1=cgw-b4dc3961
&Filter.2.Name=state
&Filter.2.Value.1=pending
&Filter.2.Value.2=available
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeVpnConnections)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeVpnConnections)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpnconnections.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpnconnections.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpnconnections.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpnconnections.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpnconnections.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpnconnections.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeVpnConnections)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpnconnections.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpnConcentrators

DescribeVpnGateways
