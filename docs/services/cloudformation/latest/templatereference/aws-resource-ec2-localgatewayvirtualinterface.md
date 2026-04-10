This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LocalGatewayVirtualInterface

Describes a local gateway virtual interface.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::LocalGatewayVirtualInterface",
  "Properties" : {
      "LocalAddress" : String,
      "LocalGatewayVirtualInterfaceGroupId" : String,
      "OutpostLagId" : String,
      "PeerAddress" : String,
      "PeerBgpAsn" : Integer,
      "PeerBgpAsnExtended" : Integer,
      "Tags" : [ Tag, ... ],
      "Vlan" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::LocalGatewayVirtualInterface
Properties:
  LocalAddress: String
  LocalGatewayVirtualInterfaceGroupId: String
  OutpostLagId: String
  PeerAddress: String
  PeerBgpAsn: Integer
  PeerBgpAsnExtended: Integer
  Tags:
    - Tag
  Vlan: Integer

```

## Properties

`LocalAddress`

The local address.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalGatewayVirtualInterfaceGroupId`

The ID of the local gateway virtual interface group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutpostLagId`

The Outpost LAG ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerAddress`

The peer address.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerBgpAsn`

The peer BGP ASN.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerBgpAsnExtended`

The extended 32-bit ASN of the BGP peer for use with larger ASN values.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the virtual interface.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-localgatewayvirtualinterface-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vlan`

The ID of the VLAN.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the local gateway virtual interface. For example:

`{ "Ref": "lgw-vif-07145b276bEXAMPLE" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConfigurationState`

The current state of the local gateway virtual interface.

`LocalBgpAsn`

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the local gateway.

`LocalGatewayId`

The ID of the local gateway.

`LocalGatewayVirtualInterfaceId`

The ID of the virtual interface.

`OwnerId`

The ID of the AWS account that owns the local gateway virtual interface.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
