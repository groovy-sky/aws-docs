This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayRouteTablePropagation

Enables the specified attachment to propagate routes to the specified propagation route
table.

For more information about enabling transit gateway route propagation, see [EnableTransitGatewayRouteTablePropagation](../../../../reference/awsec2/latest/apireference/api-enabletransitgatewayroutetablepropagation.md) in the _Amazon EC2 API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayRouteTablePropagation",
  "Properties" : {
      "TransitGatewayAttachmentId" : String,
      "TransitGatewayRouteTableId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayRouteTablePropagation
Properties:
  TransitGatewayAttachmentId: String
  TransitGatewayRouteTableId: String

```

## Properties

`TransitGatewayAttachmentId`

The ID of the attachment.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayRouteTableId`

The ID of the propagation route table.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the transit gateway route table that is
propagated.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [EnableTransitGatewayRouteTablePropagation](../../../../reference/awsec2/latest/apireference/api-enabletransitgatewayroutetablepropagation.md) in the _Amazon EC2_
_API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::TransitGatewayRouteTableAssociation

AWS::EC2::TransitGatewayVpcAttachment
