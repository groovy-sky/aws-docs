This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LocalGatewayRouteTableVPCAssociation

Associates the specified VPC with the specified local gateway route table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::LocalGatewayRouteTableVPCAssociation",
  "Properties" : {
      "LocalGatewayRouteTableId" : String,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::LocalGatewayRouteTableVPCAssociation
Properties:
  LocalGatewayRouteTableId: String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`LocalGatewayRouteTableId`

The ID of the local gateway route table.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the association.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-localgatewayroutetablevpcassociation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the local gateway route table VPC association. For example:

`{ "Ref": "lgw-vpc-assoc-0ee765bcc8EXAMPLE" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LocalGatewayId`

The ID of the local gateway.

`LocalGatewayRouteTableVpcAssociationId`

The ID of the association.

`State`

The state of the association.

## See also

- [Local\
gateway](../../../outposts/latest/userguide/outposts-local-gateways.md) in _AWS Outposts User_
_Guide_

- [CreateLocalGatewayRouteTableVpcAssociation](../../../../reference/awsec2/latest/apireference/api-createlocalgatewayroutetablevpcassociation.md) in the _Amazon EC2_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
