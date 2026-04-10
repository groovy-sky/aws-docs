This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::TransitGatewayPeering

Creates a transit gateway peering connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::TransitGatewayPeering",
  "Properties" : {
      "CoreNetworkId" : String,
      "Tags" : [ Tag, ... ],
      "TransitGatewayArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::TransitGatewayPeering
Properties:
  CoreNetworkId: String
  Tags:
    - Tag
  TransitGatewayArn: String

```

## Properties

`CoreNetworkId`

The ID of the core network.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of key-value tags associated with the peering.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-transitgatewaypeering-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayArn`

The ARN of the transit gateway.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `peeringId`. For example: `peering-01234ab1234a12a12`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`CoreNetworkArn`

The ARN of the core network.

`CreatedAt`

The timestamp when the core network peering was created.

`EdgeLocation`

The edge location for the peer.

`LastModificationErrors`

Property description not available.

`OwnerAccountId`

The ID of the account owner.

`PeeringId`

The ID of the peering.

`PeeringType`

The peering type. This will be `TRANSIT_GATEWAY`.

`ResourceArn`

The ARN of the resource peered to a core network.

`State`

The current state of the peer. This can be `CREATING` \| `FAILED` \| `AVAILABLE` \| `DELETING`.

`TransitGatewayPeeringAttachmentId`

The ID of the peering attachment.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
