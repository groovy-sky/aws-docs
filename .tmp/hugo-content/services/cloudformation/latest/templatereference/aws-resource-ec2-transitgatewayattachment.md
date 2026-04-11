This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayAttachment

Attaches a VPC to a transit gateway.

If you attach a VPC with a CIDR range that overlaps the CIDR range of a VPC that is
already attached, the new VPC CIDR range is not propagated to the default propagation route
table.

To send VPC traffic to an attached transit gateway, add a route to the VPC route table
using [AWS::EC2::Route](../userguide/aws-resource-ec2-route.md).

To update tags for a VPC attachment after creation without replacing the attachment, use
[AWS::EC2::TransitGatewayVpcAttachment](../userguide/aws-resource-ec2-transitgatewayvpcattachment.md) instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayAttachment",
  "Properties" : {
      "Options" : Options,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayAttachment
Properties:
  Options:
    Options
  SubnetIds:
    - String
  Tags:
    - Tag
  TransitGatewayId: String
  VpcId: String

```

## Properties

`Options`

The VPC attachment options.

_Required_: No

_Type_: [Options](aws-properties-ec2-transitgatewayattachment-options.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The IDs of one or more subnets. You can specify only one subnet per Availability Zone.
You must specify at least one subnet, but we recommend that you specify two subnets for better availability.
The transit gateway uses one IP address from each specified subnet.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the attachment.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-transitgatewayattachment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the attachment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the attachment.

## See also

- [CreateTransitGatewayVpcAttachment](../../../../reference/awsec2/latest/apireference/api-createtransitgatewayvpcattachment.md) in the _Amazon EC2 API_
_Reference_.

- [ModifyTransitGatewayVpcAttachment](../../../../reference/awsec2/latest/apireference/api-modifytransitgatewayvpcattachment.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Options

All content copied from https://docs.aws.amazon.com/.
