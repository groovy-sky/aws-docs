This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteServerEndpoint

Creates a new endpoint for a route server in a specified subnet.

A route server endpoint is an AWS-managed component inside a subnet that facilitates [BGP (Border Gateway Protocol)](https://en.wikipedia.org/wiki/Border_Gateway_Protocol) connections between your route server and your BGP peers.

For more information see [Dynamic routing in your VPC with VPC Route Server](../../../vpc/latest/userguide/dynamic-routing-route-server.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::RouteServerEndpoint",
  "Properties" : {
      "RouteServerId" : String,
      "SubnetId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::RouteServerEndpoint
Properties:
  RouteServerId: String
  SubnetId: String
  Tags:
    - Tag

```

## Properties

`RouteServerId`

The ID of the route server associated with this endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet to place the route server endpoint into.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the route server endpoint.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-routeserverendpoint-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the endpoint ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN for the endpoint.

`EniAddress`

The IP address of the Elastic network interface for the endpoint.

`EniId`

The ID of the Elastic network interface for the endpoint.

`Id`

The unique identifier of the route server endpoint.

`VpcId`

The ID of the VPC containing the endpoint.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::RouteServerAssociation

Tag

All content copied from https://docs.aws.amazon.com/.
