This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayRouteTable

Specifies a route table for a transit gateway.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayRouteTable",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayRouteTable
Properties:
  Tags:
    - Tag
  TransitGatewayId: String

```

## Properties

`Tags`

Any tags assigned to the route table.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-transitgatewayroutetable-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the transit gateway route table.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TransitGatewayRouteTableId`

The ID of the transit gateway route table.

## See also

- [CreateTransitGatewayRouteTable](../../../../reference/awsec2/latest/apireference/api-createtransitgatewayroutetable.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::TransitGatewayRoute

Tag

All content copied from https://docs.aws.amazon.com/.
