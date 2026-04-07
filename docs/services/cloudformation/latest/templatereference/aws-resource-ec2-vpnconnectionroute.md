This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnectionRoute

Specifies a static route for a VPN connection between an existing virtual private
gateway and a VPN customer gateway. The static route allows traffic to be routed from the
virtual private gateway to the VPN customer gateway.

For more information, see [AWS Site-to-Site VPN](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the
_AWS Site-to-Site VPN User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPNConnectionRoute",
  "Properties" : {
      "DestinationCidrBlock" : String,
      "VpnConnectionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPNConnectionRoute
Properties:
  DestinationCidrBlock: String
  VpnConnectionId: String

```

## Properties

`DestinationCidrBlock`

The CIDR block associated with the local subnet of the customer network.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpnConnectionId`

The ID of the VPN connection.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPN connection route.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### VPN connection route

The following example specifies a VPN connection route.

#### JSON

```json

"MyConnectionRoute0" : {
   "Type" : "AWS::EC2::VPNConnectionRoute",
    "Properties" : {
       "DestinationCidrBlock" : "10.0.0.0/16",
       "VpnConnectionId" : {"Ref" : "Connection0"}
    }
}
```

#### YAML

```yaml

MyConnectionRoute0:
  Type: AWS::EC2::VPNConnectionRoute
  Properties:
     DestinationCidrBlock: 10.0.0.0/16
     VpnConnectionId:
     !Ref Connection0
```

## See also

- [CreateVpnConnectionRoute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpnConnectionRoute.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpnTunnelOptionsSpecification

AWS::EC2::VPNGateway
