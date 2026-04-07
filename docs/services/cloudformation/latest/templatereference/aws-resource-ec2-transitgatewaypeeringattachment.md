This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayPeeringAttachment

Requests a transit gateway peering attachment between the specified transit gateway
(requester) and a peer transit gateway (accepter). The peer transit gateway can be in your
account or a different AWS account.

After you create the peering attachment, the owner of the accepter transit gateway must
accept the attachment request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayPeeringAttachment",
  "Properties" : {
      "PeerAccountId" : String,
      "PeerRegion" : String,
      "PeerTransitGatewayId" : String,
      "Tags" : [ Tag, ... ],
      "TransitGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayPeeringAttachment
Properties:
  PeerAccountId: String
  PeerRegion: String
  PeerTransitGatewayId: String
  Tags:
    - Tag
  TransitGatewayId: String

```

## Properties

`PeerAccountId`

The ID of the AWS account that owns the transit gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerRegion`

The Region where the transit gateway that you want to create the peer for is located.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerTransitGatewayId`

The ID of the transit gateway in the PeerRegion.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the transit gateway peering attachment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-transitgatewaypeeringattachment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitGatewayId`

The ID of the transit gateway peering attachment.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the transit gateway peering attachment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time the transit gateway peering attachment was created.

`State`

The state of the transit gateway peering attachment. Note that the
`initiating` state has been deprecated.

`TransitGatewayAttachmentId`

The ID of the transit gateway peering attachment.

## See also

- [CreateTransitGatewayPeeringAttachment](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTransitGatewayPeeringAttachment.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::TransitGatewayMulticastGroupSource

PeeringAttachmentStatus
