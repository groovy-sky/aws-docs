This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayRouteTableAssociation

Associates the specified attachment with the specified transit gateway route table. You
can associate one route table with an attachment.

Before you can update the route table associated with an attachment, you must
disassociate the transit gateway route table that is currently associated with the
attachment. First update the stack to remove the associated transit gateway route table,
and then update the stack with the ID of the new transit gateway route table to
associate. In addition, the attachment must be in an `available` state; otherwise, the request will return an error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayRouteTableAssociation",
  "Properties" : {
      "TransitGatewayAttachmentId" : String,
      "TransitGatewayRouteTableId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayRouteTableAssociation
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

The ID of the route table for the transit gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the transit gateway route table
association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [AssociateTransitGatewayRouteTable](../../../../reference/awsec2/latest/apireference/api-associatetransitgatewayroutetable.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::TransitGatewayRouteTablePropagation
