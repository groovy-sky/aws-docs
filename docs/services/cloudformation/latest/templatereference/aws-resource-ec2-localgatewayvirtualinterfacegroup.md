This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LocalGatewayVirtualInterfaceGroup

Describes a local gateway virtual interface group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::LocalGatewayVirtualInterfaceGroup",
  "Properties" : {
      "LocalBgpAsn" : Integer,
      "LocalBgpAsnExtended" : Integer,
      "LocalGatewayId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::LocalGatewayVirtualInterfaceGroup
Properties:
  LocalBgpAsn: Integer
  LocalBgpAsnExtended: Integer
  LocalGatewayId: String
  Tags:
    - Tag

```

## Properties

`LocalBgpAsn`

The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP).

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalBgpAsnExtended`

The extended 32-bit ASN for the local BGP configuration.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalGatewayId`

The ID of the local gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the virtual interface group.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-localgatewayvirtualinterfacegroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the local gateway virtual interface group. For example:

`{ "Ref": "lgw-vif-grp-07145b276bEXAMPLE" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConfigurationState`

The current state of the local gateway virtual interface group.

`LocalGatewayVirtualInterfaceGroupArn`

The Amazon Resource Number (ARN) of the local gateway virtual interface group.

`LocalGatewayVirtualInterfaceGroupId`

The ID of the virtual interface group.

`LocalGatewayVirtualInterfaceIds`

The IDs of the virtual interfaces.

`OwnerId`

The ID of the AWS account that owns the local gateway virtual interface group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
