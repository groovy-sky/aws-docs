This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation

Describes an association between a local gateway route table and a virtual interface group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation",
  "Properties" : {
      "LocalGatewayRouteTableId" : String,
      "LocalGatewayVirtualInterfaceGroupId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation
Properties:
  LocalGatewayRouteTableId: String
  LocalGatewayVirtualInterfaceGroupId: String
  Tags:
    - Tag

```

## Properties

`LocalGatewayRouteTableId`

The ID of the local gateway route table.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalGatewayVirtualInterfaceGroupId`

The ID of the virtual interface group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the association.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-localgatewayroutetablevirtualinterfacegroupassociation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the local gateway route table virtual interface group association. For example:

`{ "Ref": "lgw-vif-grp-assoc-07145b276bEXAMPLE" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LocalGatewayId`

The ID of the local gateway.

`LocalGatewayRouteTableArn`

The Amazon Resource Name (ARN) of the local gateway route table for the virtual interface group.

`LocalGatewayRouteTableVirtualInterfaceGroupAssociationId`

The ID of the association.

`OwnerId`

The ID of the AWS account that owns the local gateway virtual interface group association.

`State`

The state of the association.

## See also

- [Local\
gateway](../../../outposts/latest/userguide/outposts-local-gateways.md) in _AWS Outposts User_
_Guide_

- [CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation](../../../../reference/awsec2/latest/apireference/api-createlocalgatewayroutetablevirtualinterfacegroupassociation.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
