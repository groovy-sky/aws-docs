This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::VpcAttachment

Creates a VPC attachment on an edge location of a core network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::VpcAttachment",
  "Properties" : {
      "CoreNetworkId" : String,
      "Options" : VpcOptions,
      "ProposedNetworkFunctionGroupChange" : ProposedNetworkFunctionGroupChange,
      "ProposedSegmentChange" : ProposedSegmentChange,
      "RoutingPolicyLabel" : String,
      "SubnetArns" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::VpcAttachment
Properties:
  CoreNetworkId: String
  Options:
    VpcOptions
  ProposedNetworkFunctionGroupChange:
    ProposedNetworkFunctionGroupChange
  ProposedSegmentChange:
    ProposedSegmentChange
  RoutingPolicyLabel: String
  SubnetArns:
    - String
  Tags:
    - Tag
  VpcArn: String

```

## Properties

`CoreNetworkId`

The core network ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Options`

Options for creating the VPC attachment.

_Required_: No

_Type_: [VpcOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-vpcattachment-vpcoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedNetworkFunctionGroupChange`

Describes proposed changes to a network function group.

_Required_: No

_Type_: [ProposedNetworkFunctionGroupChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-vpcattachment-proposednetworkfunctiongroupchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedSegmentChange`

Describes a proposed segment change. In some cases, the segment change must first be evaluated and accepted.

_Required_: No

_Type_: [ProposedSegmentChange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-vpcattachment-proposedsegmentchange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingPolicyLabel`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetArns`

The subnet ARNs.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the VPC attachment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-vpcattachment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcArn`

The ARN of the VPC attachment.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AttachmentId`. For example, `{ "Ref: "attachment-00067e74104d33769" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentId`

The ID of the VPC attachment.

`AttachmentPolicyRuleNumber`

The policy rule number associated with the attachment.

`AttachmentType`

The type of attachment. This will be `VPC`.

`CoreNetworkArn`

The ARN of the core network.

`CreatedAt`

The timestamp when the VPC attachment was created.

`EdgeLocation`

The Region where the core network edge is located.

`LastModificationErrors`

Property description not available.

`NetworkFunctionGroupName`

The name of the network function group.

`OwnerAccountId`

The ID of the VPC attachment owner.

`ResourceArn`

The resource ARN for the VPC attachment.

`SegmentName`

The name of the attachment's segment.

`State`

The state of the attachment. This can be: `REJECTED` \| `PENDING_ATTACHMENT_ACCEPTANCE` \| `CREATING` \| `FAILED` \| `AVAILABLE` \| `UPDATING` \| ` PENDING_NETWORK_UPDATE` \| `PENDING_TAG_ACCEPTANCE` \| `DELETING`.

`UpdatedAt`

The timestamp when the VPC attachment was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ProposedNetworkFunctionGroupChange
