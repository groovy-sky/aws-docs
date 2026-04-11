This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IpPoolRouteTableAssociation

A route server association is the connection established between a route server and a VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IpPoolRouteTableAssociation",
  "Properties" : {
      "PublicIpv4Pool" : String,
      "RouteTableId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IpPoolRouteTableAssociation
Properties:
  PublicIpv4Pool: String
  RouteTableId: String

```

## Properties

`PublicIpv4Pool`

The ID of a public IPv4 address pool.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteTableId`

The ID of a route table.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the route table association ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

The ID of a route table association.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::KeyPair

All content copied from https://docs.aws.amazon.com/.
