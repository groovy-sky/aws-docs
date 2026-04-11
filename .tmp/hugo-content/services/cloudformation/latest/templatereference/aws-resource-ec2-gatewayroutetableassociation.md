This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::GatewayRouteTableAssociation

Associates a virtual private gateway or internet gateway with a route table. The gateway
and route table must be in the same VPC. This association causes the incoming traffic to
the gateway to be routed according to the routes in the route table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::GatewayRouteTableAssociation",
  "Properties" : {
      "GatewayId" : String,
      "RouteTableId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::GatewayRouteTableAssociation
Properties:
  GatewayId: String
  RouteTableId: String

```

## Properties

`GatewayId`

The ID of the gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteTableId`

The ID of the route table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the association.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

The ID of the route table association.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::Host

All content copied from https://docs.aws.amazon.com/.
