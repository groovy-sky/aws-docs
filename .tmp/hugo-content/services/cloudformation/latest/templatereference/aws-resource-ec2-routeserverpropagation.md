This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteServerPropagation

Specifies route propagation from a route server to a route table.

For more information see [Dynamic routing in your VPC with VPC Route Server](../../../vpc/latest/userguide/dynamic-routing-route-server.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::RouteServerPropagation",
  "Properties" : {
      "RouteServerId" : String,
      "RouteTableId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::RouteServerPropagation
Properties:
  RouteServerId: String
  RouteTableId: String

```

## Properties

`RouteServerId`

The ID of the route server configured for route propagation.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteTableId`

The ID of the route table configured for route server propagation.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the route server ID and route table ID separated by the pipe symbol (\|).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::RouteTable

All content copied from https://docs.aws.amazon.com/.
