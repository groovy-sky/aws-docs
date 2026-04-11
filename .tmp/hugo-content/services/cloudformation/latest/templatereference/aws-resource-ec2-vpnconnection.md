This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection

Specifies a VPN connection between a virtual private gateway and a VPN customer gateway
or a transit gateway and a VPN customer gateway.

To specify a VPN connection between a transit gateway and customer gateway, use the
`TransitGatewayId` and `CustomerGatewayId` properties.

To specify a VPN connection between a virtual private gateway and customer gateway, use
the `VpnGatewayId` and `CustomerGatewayId` properties.

For more information, see [AWS Site-to-Site VPN](../../../vpn/latest/s2svpn/vpc-vpn.md) in the
_AWS Site-to-Site VPN User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPNConnection",
  "Properties" : {
      "CustomerGatewayId" : String,
      "EnableAcceleration" : Boolean,
      "LocalIpv4NetworkCidr" : String,
      "LocalIpv6NetworkCidr" : String,
      "OutsideIpAddressType" : String,
      "PreSharedKeyStorage" : String,
      "RemoteIpv4NetworkCidr" : String,
      "RemoteIpv6NetworkCidr" : String,
      "StaticRoutesOnly" : Boolean,
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String,
      "TransportTransitGatewayAttachmentId" : String,
      "TunnelBandwidth" : String,
      "TunnelInsideIpVersion" : String,
      "Type" : String,
      "VpnConcentratorId" : String,
      "VpnGatewayId" : String,
      "VpnTunnelOptionsSpecifications" : [ VpnTunnelOptionsSpecification, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPNConnection
Properties:
  CustomerGatewayId: String
  EnableAcceleration: Boolean
  LocalIpv4NetworkCidr: String
  LocalIpv6NetworkCidr: String
  OutsideIpAddressType: String
  PreSharedKeyStorage: String
  RemoteIpv4NetworkCidr: String
  RemoteIpv6NetworkCidr: String
  StaticRoutesOnly: Boolean
  Tags:
    - Tag
  TransitGatewayId: String
  TransportTransitGatewayAttachmentId: String
  TunnelBandwidth: String
  TunnelInsideIpVersion: String
  Type: String
  VpnConcentratorId: String
  VpnGatewayId: String
  VpnTunnelOptionsSpecifications:
    - VpnTunnelOptionsSpecification

```

## Properties

`CustomerGatewayId`

The ID of the customer gateway at your end of the VPN connection.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAcceleration`

Indicate whether to enable acceleration for the VPN connection.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalIpv4NetworkCidr`

The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.

Default: `0.0.0.0/0`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalIpv6NetworkCidr`

The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.

Default: `::/0`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutsideIpAddressType`

The type of IP address assigned to the outside interface of the customer gateway device.

Valid values: `PrivateIpv4` \| `PublicIpv4` \| `Ipv6`

Default: `PublicIpv4`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreSharedKeyStorage`

Describes the storage location for an instance store-backed AMI.

_Required_: No

_Type_: String

_Allowed values_: `Standard | SecretsManager`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RemoteIpv4NetworkCidr`

The IPv4 CIDR on the AWS side of the VPN connection.

Default: `0.0.0.0/0`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RemoteIpv6NetworkCidr`

The IPv6 CIDR on the AWS side of the VPN connection.

Default: `::/0`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StaticRoutesOnly`

Indicates whether the VPN connection uses static routes only. Static routes must be used
for devices that don't support BGP.

If you are creating a VPN connection for a device that does not support Border Gateway
Protocol (BGP), you must specify `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the VPN connection.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-vpnconnection-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway associated with the VPN connection.

You must specify either `TransitGatewayId` or `VpnGatewayId`, but
not both.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportTransitGatewayAttachmentId`

The transit gateway attachment ID to use for the VPN tunnel.

Required if `OutsideIpAddressType` is set to `PrivateIpv4`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TunnelBandwidth`

The desired bandwidth specification for the VPN tunnel, used when creating or modifying VPN connection options to set the tunnel's throughput
capacity. `standard` supports up to 1.25 Gbps per tunnel, while `large` supports up to 5 Gbps per tunnel. The default value is `standard`. Existing
VPN connections without a bandwidth setting will automatically default to `standard`.

_Required_: No

_Type_: String

_Allowed values_: `standard | large`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TunnelInsideIpVersion`

Indicate whether the VPN tunnels process IPv4 or IPv6 traffic.

Default: `ipv4`

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of VPN connection.

_Required_: Yes

_Type_: String

_Allowed values_: `ipsec.1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpnConcentratorId`

The ID of the VPN concentrator to associate with the VPN connection.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpnGatewayId`

The ID of the virtual private gateway at the AWS side of the VPN
connection.

You must specify either `TransitGatewayId` or `VpnGatewayId`, but
not both.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpnTunnelOptionsSpecifications`

The tunnel options for the VPN connection.

_Required_: No

_Type_: Array of [VpnTunnelOptionsSpecification](aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPN connection.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VpnConnectionId`

The ID of the VPN connection.

## Examples

### VPN connection

The following example specifies a VPN connection between myVPNGateway and
MyCustomerGateway.

#### JSON

```json

"myVPNConnection" : {
   "Type" : "AWS::EC2::VPNConnection",
   "Properties" : {
      "Type" : "ipsec.1",
      "StaticRoutesOnly" : "true",
      "CustomerGatewayId" : {"Ref" : "myCustomerGateway"},
      "VpnGatewayId" : {"Ref" : "myVPNGateway"}
   }
}
```

#### YAML

```yaml

   myVPNConnection:
      Type: AWS::EC2::VPNConnection
      Properties:
        Type: ipsec.1
        StaticRoutesOnly: true
        CustomerGatewayId:
          !Ref myCustomerGateway
        VpnGatewayId:
          !Ref myVPNGateway
```

## See also

- [VPNConnection](../../../../reference/awsec2/latest/apireference/api-vpnconnection.md)
in the _Amazon EC2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CloudwatchLogOptionsSpecification

All content copied from https://docs.aws.amazon.com/.
