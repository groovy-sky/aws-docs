# VpnConnectionOptionsSpecification

Describes VPN connection options.

## Contents

**EnableAcceleration**

Indicate whether to enable acceleration for the VPN connection.

Default: `false`

Type: Boolean

Required: No

**LocalIpv4NetworkCidr**

The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.

Default: `0.0.0.0/0`

Type: String

Required: No

**LocalIpv6NetworkCidr**

The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.

Default: `::/0`

Type: String

Required: No

**OutsideIpAddressType**

The type of IP address assigned to the outside interface of the customer gateway device.

Valid values: `PrivateIpv4` \| `PublicIpv4` \| `Ipv6`

Default: `PublicIpv4`

Type: String

Required: No

**RemoteIpv4NetworkCidr**

The IPv4 CIDR on the AWS side of the VPN connection.

Default: `0.0.0.0/0`

Type: String

Required: No

**RemoteIpv6NetworkCidr**

The IPv6 CIDR on the AWS side of the VPN connection.

Default: `::/0`

Type: String

Required: No

**StaticRoutesOnly**

Indicate whether the VPN connection uses static routes only. If you are creating a VPN
connection for a device that does not support BGP, you must specify `true`.
Use [CreateVpnConnectionRoute](api-createvpnconnectionroute.md) to create a static route.

Default: `false`

Type: Boolean

Required: No

**TransportTransitGatewayAttachmentId**

The transit gateway attachment ID to use for the VPN tunnel.

Required if `OutsideIpAddressType` is set to `PrivateIpv4`.

Type: String

Required: No

**TunnelBandwidth**

The desired bandwidth specification for the VPN tunnel, used when creating or modifying VPN connection options to set the tunnel's throughput
capacity. `standard` supports up to 1.25 Gbps per tunnel, while `large` supports up to 5 Gbps per tunnel. The default value is `standard`. Existing
VPN connections without a bandwidth setting will automatically default to `standard`.

Type: String

Valid Values: `standard | large`

Required: No

**TunnelInsideIpVersion**

Indicate whether the VPN tunnels process IPv4 or IPv6 traffic.

Default: `ipv4`

Type: String

Valid Values: `ipv4 | ipv6`

Required: No

**TunnelOptions.N**

The tunnel options for the VPN connection.

Type: Array of [VpnTunnelOptionsSpecification](api-vpntunneloptionsspecification.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpnconnectionoptionsspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpnconnectionoptionsspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpnconnectionoptionsspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpnConnectionOptions

VpnGateway
