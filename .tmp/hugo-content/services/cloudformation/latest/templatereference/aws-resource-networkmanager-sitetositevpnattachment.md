This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::SiteToSiteVpnAttachment

Creates an Amazon Web Services site-to-site VPN attachment on an edge location of a core network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::SiteToSiteVpnAttachment",
  "Properties" : {
      "CoreNetworkId" : String,
      "NetworkFunctionGroupName" : String,
      "ProposedNetworkFunctionGroupChange" : ProposedNetworkFunctionGroupChange,
      "ProposedSegmentChange" : ProposedSegmentChange,
      "RoutingPolicyLabel" : String,
      "Tags" : [ Tag, ... ],
      "VpnConnectionArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::SiteToSiteVpnAttachment
Properties:
  CoreNetworkId: String
  NetworkFunctionGroupName: String
  ProposedNetworkFunctionGroupChange:
    ProposedNetworkFunctionGroupChange
  ProposedSegmentChange:
    ProposedSegmentChange
  RoutingPolicyLabel: String
  Tags:
    - Tag
  VpnConnectionArn: String

```

## Properties

`CoreNetworkId`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkFunctionGroupName`

The name of the network function group.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedNetworkFunctionGroupChange`

Describes proposed changes to a network function group.

_Required_: No

_Type_: [ProposedNetworkFunctionGroupChange](aws-properties-networkmanager-sitetositevpnattachment-proposednetworkfunctiongroupchange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProposedSegmentChange`

Describes a proposed segment change. In some cases, the segment change must first be evaluated and accepted.

_Required_: No

_Type_: [ProposedSegmentChange](aws-properties-networkmanager-sitetositevpnattachment-proposedsegmentchange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingPolicyLabel`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the Site-to-Site VPN attachment.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-sitetositevpnattachment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpnConnectionArn`

The ARN of the site-to-site VPN attachment.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:ec2:[^:]{0,63}:[^:]{0,63}:vpn-connection\/vpn-[0-9a-f]{8,17}$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AttachmentId`. For example, `{ "Ref: "attachment-05467e74104d33861" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentId`

The ID of the site-to-site VPN attachment.

`AttachmentPolicyRuleNumber`

The policy rule number associated with the attachment.

`AttachmentType`

The type of attachment. This will be `SITE_TO_SITE_VPN`.

`CoreNetworkArn`

The ARN of the core network.

`CreatedAt`

The timestamp when the site-to-site VPN attachment was created.

`EdgeLocation`

The Region where the core network edge is located.

`LastModificationErrors`

Property description not available.

`OwnerAccountId`

The ID of the site-to-site VPN attachment owner.

`ResourceArn`

The resource ARN for the site-to-site VPN attachment.

`SegmentName`

The name of the site-to-site VPN attachment's segment.

`State`

The state of the site-to-site VPN attachment. This can be: `REJECTED` \| `PENDING_ATTACHMENT_ACCEPTANCE` \| `CREATING` \| `FAILED` \| `AVAILABLE` \| `UPDATING` \| ` PENDING_NETWORK_UPDATE` \| `PENDING_TAG_ACCEPTANCE` \| `DELETING`.

`UpdatedAt`

The timestamp when the site-to-site VPN attachment was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ProposedNetworkFunctionGroupChange

All content copied from https://docs.aws.amazon.com/.
