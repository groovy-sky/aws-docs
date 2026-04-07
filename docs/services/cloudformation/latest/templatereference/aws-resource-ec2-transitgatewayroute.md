This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayRoute

Specifies a static route for a transit gateway route table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayRoute",
  "Properties" : {
      "Blackhole" : Boolean,
      "DestinationCidrBlock" : String,
      "TransitGatewayAttachmentId" : String,
      "TransitGatewayRouteTableId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayRoute
Properties:
  Blackhole: Boolean
  DestinationCidrBlock: String
  TransitGatewayAttachmentId: String
  TransitGatewayRouteTableId: String

```

## Properties

`Blackhole`

Indicates whether to drop traffic that matches this route.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationCidrBlock`

The CIDR block used for destination matches.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayAttachmentId`

The ID of the attachment.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayRouteTableId`

The ID of the transit gateway route table.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the transit gateway route.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [CreateTransitGatewayRoute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTransitGatewayRoute.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::TransitGatewayRouteTable
