This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::RouteServerAssociation

Specifies the association between a route server and a VPC.

A route server association is the connection established between a route server and a VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::RouteServerAssociation",
  "Properties" : {
      "RouteServerId" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::RouteServerAssociation
Properties:
  RouteServerId: String
  VpcId: String

```

## Properties

`RouteServerId`

The ID of the associated route server.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the associated VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the route server ID and VPC ID separated by the pipe symbol (\|).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::RouteServerEndpoint
