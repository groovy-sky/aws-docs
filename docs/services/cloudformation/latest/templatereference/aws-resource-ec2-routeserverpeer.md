This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteServerPeer

Specifies a BGP peer configuration for a route server endpoint.

A route server peer is a session between a route server endpoint and the device deployed in AWS (such as a firewall appliance or other network security function running on an EC2 instance).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::RouteServerPeer",
  "Properties" : {
      "BgpOptions" : BgpOptions,
      "PeerAddress" : String,
      "RouteServerEndpointId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::RouteServerPeer
Properties:
  BgpOptions:
    BgpOptions
  PeerAddress: String
  RouteServerEndpointId: String
  Tags:
    - Tag

```

## Properties

`BgpOptions`

The BGP configuration options for this peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.

_Required_: Yes

_Type_: [BgpOptions](aws-properties-ec2-routeserverpeer-bgpoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerAddress`

The IPv4 address of the peer device.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteServerEndpointId`

The ID of the route server endpoint associated with this peer.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the route server peer.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-routeserverpeer-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the peer ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the route server peer.

`EndpointEniAddress`

The IP address of the Elastic network interface for the route server endpoint.

`EndpointEniId`

The ID of the Elastic network interface for the route server endpoint.

`Id`

The ID of the route server peer.

`RouteServerId`

The ID of the route server associated with this peer.

`SubnetId`

The ID of the subnet containing the route server peer.

`VpcId`

The ID of the VPC containing the route server peer.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BgpOptions

All content copied from https://docs.aws.amazon.com/.
