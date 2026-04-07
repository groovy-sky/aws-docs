This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::ConnectAttachment

Creates a core network Connect attachment from a specified core network attachment.

A core network Connect attachment is a GRE-based tunnel attachment that you can use to
establish a connection between a core network and an appliance. A core network Connect
attachment uses an existing VPC attachment as the underlying transport mechanism.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::ConnectAttachment",
  "Properties" : {
      "CoreNetworkId" : String,
      "EdgeLocation" : String,
      "NetworkFunctionGroupName" : String,
      "Options" : ConnectAttachmentOptions,
      "ProposedNetworkFunctionGroupChange" : ProposedNetworkFunctionGroupChange,
      "ProposedSegmentChange" : ProposedSegmentChange,
      "RoutingPolicyLabel" : String,
      "Tags" : [ Tag, ... ],
      "TransportAttachmentId" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::ConnectAttachment
Properties:
  CoreNetworkId: String
  EdgeLocation: String
  NetworkFunctionGroupName: String
  Options:
    ConnectAttachmentOptions
  ProposedNetworkFunctionGroupChange:
    ProposedNetworkFunctionGroupChange
  ProposedSegmentChange:
    ProposedSegmentChange
  RoutingPolicyLabel: String
  Tags:
    - Tag
  TransportAttachmentId: String

```

## Properties

`CoreNetworkId`

The ID of the core network where the Connect attachment is located.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EdgeLocation`

The Region where the edge is located.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkFunctionGroupName`

The name of the network function group.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

Options for connecting an attachment.

_Required_: Yes

_Type_: [ConnectAttachmentOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-connectattachment-connectattachmentoptions.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProposedNetworkFunctionGroupChange`

Describes proposed changes to a network function group.

_Required_: No

_Type_: [ProposedNetworkFunctionGroupChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-connectattachment-proposednetworkfunctiongroupchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedSegmentChange`

Describes a proposed segment change. In some cases, the segment change must first be evaluated and accepted.

_Required_: No

_Type_: [ProposedSegmentChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-connectattachment-proposedsegmentchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingPolicyLabel`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the Connect attachment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-connectattachment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransportAttachmentId`

The ID of the transport attachment.

_Required_: Yes

_Type_: String

_Pattern_: `^attachment-([0-9a-f]{8,17})$`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AttachmentId`. For example, `{ "Ref: "attachment-02767e74104a33690" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentId`

The ID of the Connect attachment.

`AttachmentPolicyRuleNumber`

The rule number associated with the attachment.

`AttachmentType`

The type of attachment. This will be `CONNECT`.

`CoreNetworkArn`

The ARN of the core network.

`CreatedAt`

The timestamp when the Connect attachment was created.

`LastModificationErrors`

Property description not available.

`OwnerAccountId`

The ID of the Connect attachment owner.

`ResourceArn`

The resource ARN for the Connect attachment.

`SegmentName`

The name of the Connect attachment's segment.

`State`

The state of the Connect attachment. This can be: `REJECTED` \| `PENDING_ATTACHMENT_ACCEPTANCE` \| `CREATING` \| `FAILED` \| `AVAILABLE` \| `UPDATING` \| ` PENDING_NETWORK_UPDATE` \| `PENDING_TAG_ACCEPTANCE` \| `DELETING`.

`UpdatedAt`

The timestamp when the Connect attachment was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Network Manager

ConnectAttachmentOptions
