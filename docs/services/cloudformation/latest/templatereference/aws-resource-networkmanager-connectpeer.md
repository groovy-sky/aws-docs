This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::ConnectPeer

Creates a core network Connect peer for a specified core network connect attachment between a core network and an appliance.
The peer address and transit gateway address must be the same IP address family (IPv4 or IPv6).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::ConnectPeer",
  "Properties" : {
      "BgpOptions" : BgpOptions,
      "ConnectAttachmentId" : String,
      "CoreNetworkAddress" : String,
      "InsideCidrBlocks" : [ String, ... ],
      "PeerAddress" : String,
      "SubnetArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::ConnectPeer
Properties:
  BgpOptions:
    BgpOptions
  ConnectAttachmentId: String
  CoreNetworkAddress: String
  InsideCidrBlocks:
    - String
  PeerAddress: String
  SubnetArn: String
  Tags:
    - Tag

```

## Properties

`BgpOptions`

Describes the BGP options.

_Required_: No

_Type_: [BgpOptions](aws-properties-networkmanager-connectpeer-bgpoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectAttachmentId`

The ID of the attachment to connect.

_Required_: Yes

_Type_: String

_Pattern_: `^attachment-([0-9a-f]{8,17})$`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CoreNetworkAddress`

The IP address of a core network.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InsideCidrBlocks`

The inside IP addresses used for a Connect peer configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerAddress`

The IP address of the Connect peer.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetArn`

The subnet ARN of the Connect peer.

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:ec2:[^:]{0,63}:[^:]{0,63}:subnet\/subnet-[0-9a-f]{8,17}$|^$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of key-value tags associated with the Connect peer.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-connectpeer-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ConnectPeerId`. For example, `{ "Ref: "connect-peer--041e25dbc928d7e61" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectPeerId`

The ID of the Connect peer.

`CoreNetworkId`

The core network ID.

`CreatedAt`

The timestamp when the Connect peer was created.

`EdgeLocation`

The Connect peer Regions where edges are located.

`LastModificationErrors`

Describes the error associated with the attachment request.

`State`

The state of the Connect peer. This will be: `REJECTED` \| `PENDING_ATTACHMENT_ACCEPTANCE` \| `CREATING` \| `FAILED` \| `AVAILABLE` \| `UPDATING` \| ` PENDING_NETWORK_UPDATE` \| `PENDING_TAG_ACCEPTANCE` \| `DELETING`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BgpOptions

All content copied from https://docs.aws.amazon.com/.
