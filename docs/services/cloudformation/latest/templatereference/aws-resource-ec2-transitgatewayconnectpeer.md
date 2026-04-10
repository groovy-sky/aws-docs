This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayConnectPeer

Describes a transit gateway Connect peer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayConnectPeer",
  "Properties" : {
      "ConnectPeerConfiguration" : TransitGatewayConnectPeerConfiguration,
      "Tags" : [ Tag, ... ],
      "TransitGatewayAttachmentId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayConnectPeer
Properties:
  ConnectPeerConfiguration:
    TransitGatewayConnectPeerConfiguration
  Tags:
    - Tag
  TransitGatewayAttachmentId: String

```

## Properties

`ConnectPeerConfiguration`

The Connect peer details.

_Required_: Yes

_Type_: [TransitGatewayConnectPeerConfiguration](aws-properties-ec2-transitgatewayconnectpeer-transitgatewayconnectpeerconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the Connect peer.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-transitgatewayconnectpeer-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayAttachmentId`

The ID of the Connect attachment.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Connect peer.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`ConnectPeerConfiguration.BgpConfigurations`

The BGP configuration details.

`ConnectPeerConfiguration.Protocol`

The tunnel protocol.

`CreationTime`

The creation time.

`State`

The state of the Connect peer.

`TransitGatewayConnectPeerId`

The ID of the Connect peer.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransitGatewayConnectOptions

Tag

All content copied from https://docs.aws.amazon.com/.
